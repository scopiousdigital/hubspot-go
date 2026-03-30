package marketing_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/scopiousdigital/hubspot-go/marketing"
)

// mockRequester implements api.Requester by routing requests to a test server.
type mockRequester struct {
	server *httptest.Server
}

func (m *mockRequester) Get(ctx context.Context, path string, query url.Values, result any) error {
	u := m.server.URL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	return m.doRequest(req, result)
}

func (m *mockRequester) Post(ctx context.Context, path string, body, result any) error {
	var reader io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		reader = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, m.server.URL+path, reader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return m.doRequest(req, result)
}

func (m *mockRequester) Put(ctx context.Context, path string, body, result any) error {
	b, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, m.server.URL+path, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return m.doRequest(req, result)
}

func (m *mockRequester) Patch(ctx context.Context, path string, body, result any) error {
	b, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, m.server.URL+path, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return m.doRequest(req, result)
}

func (m *mockRequester) Delete(ctx context.Context, path string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, m.server.URL+path, nil)
	if err != nil {
		return err
	}
	return m.doRequest(req, nil)
}

func (m *mockRequester) doRequest(req *http.Request, result any) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if result != nil {
		body, _ := io.ReadAll(resp.Body)
		if len(body) > 0 {
			return json.Unmarshal(body, result)
		}
	}
	return nil
}

func setupMarketing(t *testing.T) (*marketing.Service, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	requester := &mockRequester{server: server}
	svc := marketing.NewService(requester)
	return svc, mux, server.Close
}

func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}
	return b
}

// --- Forms Tests ---

func TestForms_Create(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/forms", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input marketing.FormCreateRequest
		_ = json.Unmarshal(body, &input)
		if input.Name != "Test Form" {
			t.Errorf("name = %q, want 'Test Form'", input.Name)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(mustJSON(t, marketing.FormDefinition{
			ID:        "form-123",
			Name:      "Test Form",
			FormType:  "hubspot",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	form, err := svc.Forms.Create(context.Background(), &marketing.FormCreateRequest{
		Name:     "Test Form",
		FormType: "hubspot",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if form.ID != "form-123" {
		t.Errorf("ID = %q, want form-123", form.ID)
	}
	if form.Name != "Test Form" {
		t.Errorf("Name = %q, want 'Test Form'", form.Name)
	}
}

func TestForms_GetByID(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/forms/form-123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.FormDefinition{
			ID:        "form-123",
			Name:      "Test Form",
			FormType:  "hubspot",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	form, err := svc.Forms.GetByID(context.Background(), "form-123", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if form.ID != "form-123" {
		t.Errorf("ID = %q", form.ID)
	}
}

func TestForms_GetPage(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/forms", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if limit := r.URL.Query().Get("limit"); limit != "10" {
			t.Errorf("limit = %q, want 10", limit)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.FormListResult{
			Results: []marketing.FormDefinition{
				{ID: "f1", Name: "Form 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
				{ID: "f2", Name: "Form 2", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
			Paging: &marketing.ForwardPaging{Next: &marketing.NextPage{After: "f2"}},
		}))
	})

	result, err := svc.Forms.GetPage(context.Background(), &marketing.FormListOptions{Limit: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d, want 2", len(result.Results))
	}
	if result.Paging.Next.After != "f2" {
		t.Errorf("after = %q", result.Paging.Next.After)
	}
}

func TestForms_Update(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/forms/form-123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %s, want PATCH", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.FormDefinition{
			ID:        "form-123",
			Name:      "Updated Form",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	form, err := svc.Forms.Update(context.Background(), "form-123", &marketing.FormUpdateRequest{
		Name: "Updated Form",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if form.Name != "Updated Form" {
		t.Errorf("Name = %q", form.Name)
	}
}

func TestForms_Archive(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/forms/form-123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Forms.Archive(context.Background(), "form-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// --- Email Tests ---

func TestEmails_Create(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/emails", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input marketing.EmailCreateRequest
		_ = json.Unmarshal(body, &input)
		if input.Name != "Test Email" {
			t.Errorf("name = %q", input.Name)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(mustJSON(t, marketing.PublicEmail{
			ID:   "email-001",
			Name: "Test Email",
		}))
	})

	email, err := svc.Emails.Create(context.Background(), &marketing.EmailCreateRequest{
		Name: "Test Email",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if email.ID != "email-001" {
		t.Errorf("ID = %q", email.ID)
	}
}

func TestEmails_GetByID(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/emails/email-001", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.PublicEmail{
			ID:      "email-001",
			Name:    "Test Email",
			Subject: "Hello",
		}))
	})

	email, err := svc.Emails.GetByID(context.Background(), "email-001", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if email.Subject != "Hello" {
		t.Errorf("Subject = %q", email.Subject)
	}
}

func TestEmails_GetPage(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/emails", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if limit := r.URL.Query().Get("limit"); limit != "5" {
			t.Errorf("limit = %q, want 5", limit)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.EmailListResult{
			Total: 2,
			Results: []marketing.PublicEmail{
				{ID: "e1", Name: "Email 1"},
				{ID: "e2", Name: "Email 2"},
			},
		}))
	})

	result, err := svc.Emails.GetPage(context.Background(), &marketing.EmailListOptions{Limit: 5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 2 {
		t.Errorf("total = %d", result.Total)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d", len(result.Results))
	}
}

func TestEmails_Update(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/emails/email-001", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %s, want PATCH", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.PublicEmail{
			ID:      "email-001",
			Name:    "Updated Email",
			Subject: "Updated Subject",
		}))
	})

	email, err := svc.Emails.Update(context.Background(), "email-001", &marketing.EmailUpdateRequest{
		Name:    "Updated Email",
		Subject: "Updated Subject",
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if email.Name != "Updated Email" {
		t.Errorf("Name = %q", email.Name)
	}
}

func TestEmails_Archive(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/emails/email-001", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Emails.Archive(context.Background(), "email-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestEmails_Clone(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/emails/clone", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(mustJSON(t, marketing.PublicEmail{
			ID:         "email-002",
			Name:       "Cloned Email",
			ClonedFrom: "email-001",
		}))
	})

	email, err := svc.Emails.Clone(context.Background(), &marketing.ContentCloneRequest{
		ID:        "email-001",
		CloneName: "Cloned Email",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if email.ID != "email-002" {
		t.Errorf("ID = %q", email.ID)
	}
}

func TestEmails_PublishOrSend(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/emails/email-001/publish", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Emails.PublishOrSend(context.Background(), "email-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// --- Statistics Tests ---

func TestStatistics_GetEmailsList(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/emails/statistics/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.AggregateEmailStatistics{
			Emails: []int{1, 2, 3},
			Aggregate: &marketing.EmailStatisticsData{
				Counters: map[string]int{"sent": 100, "delivered": 95},
			},
		}))
	})

	stats, err := svc.Statistics.GetEmailsList(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stats.Emails) != 3 {
		t.Errorf("emails = %d", len(stats.Emails))
	}
	if stats.Aggregate.Counters["sent"] != 100 {
		t.Errorf("sent = %d", stats.Aggregate.Counters["sent"])
	}
}

func TestStatistics_GetHistogram(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/emails/statistics/histogram", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if interval := r.URL.Query().Get("interval"); interval != "DAY" {
			t.Errorf("interval = %q, want DAY", interval)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.EmailStatisticsHistogramResult{
			Total: 1,
			Results: []marketing.EmailStatisticInterval{
				{Aggregations: &marketing.EmailStatisticsData{
					Counters: map[string]int{"sent": 50},
				}},
			},
		}))
	})

	result, err := svc.Statistics.GetHistogram(context.Background(), &marketing.EmailStatisticsHistogramOptions{
		Interval: "DAY",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 1 {
		t.Errorf("total = %d", result.Total)
	}
}

// --- Events Tests ---

func TestEvents_Create(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/marketing-events/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input marketing.MarketingEventCreateRequest
		_ = json.Unmarshal(body, &input)
		if input.EventName != "My Event" {
			t.Errorf("eventName = %q", input.EventName)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(mustJSON(t, marketing.MarketingEventDefaultResponse{
			EventName:      "My Event",
			EventOrganizer: "Test Org",
		}))
	})

	event, err := svc.Events.Create(context.Background(), &marketing.MarketingEventCreateRequest{
		ExternalEventID:   "ext-001",
		ExternalAccountID: "acct-001",
		EventName:         "My Event",
		EventOrganizer:    "Test Org",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if event.EventName != "My Event" {
		t.Errorf("EventName = %q", event.EventName)
	}
}

func TestEvents_GetAll(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/marketing-events/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.MarketingEventListResult{
			Results: []marketing.MarketingEventReadResponseV2{
				{ObjectID: "obj-1", EventName: "Event 1", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
		}))
	})

	result, err := svc.Events.GetAll(context.Background(), "", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d", len(result.Results))
	}
}

func TestEvents_ArchiveByObjectID(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/marketing-events/events/obj-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Events.ArchiveByObjectID(context.Background(), "obj-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestEvents_Cancel(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/marketing-events/events/external/ext-001/cancel", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.MarketingEventDefaultResponse{
			EventName:      "Cancelled Event",
			EventOrganizer: "Test Org",
		}))
	})

	event, err := svc.Events.Cancel(context.Background(), "ext-001", "acct-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if event.EventName != "Cancelled Event" {
		t.Errorf("EventName = %q", event.EventName)
	}
}

func TestEvents_GetSettings(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/marketing-events/events/settings/12345", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.EventDetailSettings{
			AppID:           12345,
			EventDetailsURL: "https://example.com/event",
		}))
	})

	settings, err := svc.Events.GetSettings(context.Background(), 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if settings.AppID != 12345 {
		t.Errorf("AppID = %d", settings.AppID)
	}
	if settings.EventDetailsURL != "https://example.com/event" {
		t.Errorf("EventDetailsURL = %q", settings.EventDetailsURL)
	}
}

func TestEvents_GetParticipationsCounters(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/marketing-events/events/42/participations/counters", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.AttendanceCounters{
			Attended:   50,
			Registered: 100,
			Cancelled:  10,
			NoShows:    5,
		}))
	})

	counters, err := svc.Events.GetParticipationsCountersByMarketingEventID(context.Background(), 42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if counters.Attended != 50 {
		t.Errorf("Attended = %d", counters.Attended)
	}
	if counters.Registered != 100 {
		t.Errorf("Registered = %d", counters.Registered)
	}
}

// --- Transactional Tests ---

func TestTransactional_SendEmail(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/transactional/single-email/send", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input marketing.SingleSendRequest
		_ = json.Unmarshal(body, &input)
		if input.EmailID != 12345 {
			t.Errorf("emailId = %d", input.EmailID)
		}
		if input.Message.To != "user@example.com" {
			t.Errorf("to = %q", input.Message.To)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.EmailSendStatusView{
			StatusID: "status-001",
			Status:   "PENDING",
		}))
	})

	status, err := svc.Transactional.SendEmail(context.Background(), &marketing.SingleSendRequest{
		EmailID: 12345,
		Message: marketing.SingleSendEmail{To: "user@example.com"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if status.StatusID != "status-001" {
		t.Errorf("StatusID = %q", status.StatusID)
	}
	if status.Status != "PENDING" {
		t.Errorf("Status = %q", status.Status)
	}
}

func TestTransactional_CreateToken(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/transactional/smtp-tokens", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input marketing.SmtpApiTokenCreateRequest
		_ = json.Unmarshal(body, &input)
		if input.CampaignName != "My Campaign" {
			t.Errorf("campaignName = %q", input.CampaignName)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(mustJSON(t, marketing.SmtpApiTokenView{
			ID:              "token-001",
			CampaignName:    "My Campaign",
			EmailCampaignID: "ec-001",
			CreatedBy:       "user-1",
			CreateContact:   true,
			Password:        "secret-password",
			CreatedAt:       time.Now(),
		}))
	})

	token, err := svc.Transactional.CreateToken(context.Background(), &marketing.SmtpApiTokenCreateRequest{
		CampaignName:  "My Campaign",
		CreateContact: true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token.ID != "token-001" {
		t.Errorf("ID = %q", token.ID)
	}
	if token.Password != "secret-password" {
		t.Errorf("Password = %q", token.Password)
	}
}

func TestTransactional_GetTokenByID(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/transactional/smtp-tokens/token-001", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.SmtpApiTokenView{
			ID:              "token-001",
			CampaignName:    "My Campaign",
			EmailCampaignID: "ec-001",
			CreatedBy:       "user-1",
			CreatedAt:       time.Now(),
		}))
	})

	token, err := svc.Transactional.GetTokenByID(context.Background(), "token-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token.ID != "token-001" {
		t.Errorf("ID = %q", token.ID)
	}
}

func TestTransactional_GetTokensPage(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/transactional/smtp-tokens", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if limit := r.URL.Query().Get("limit"); limit != "10" {
			t.Errorf("limit = %q, want 10", limit)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.SmtpTokenListResult{
			Results: []marketing.SmtpApiTokenView{
				{ID: "t1", CampaignName: "C1", CreatedAt: time.Now()},
				{ID: "t2", CampaignName: "C2", CreatedAt: time.Now()},
			},
		}))
	})

	result, err := svc.Transactional.GetTokensPage(context.Background(), &marketing.SmtpTokenListOptions{Limit: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d", len(result.Results))
	}
}

func TestTransactional_ArchiveToken(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/transactional/smtp-tokens/token-001", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Transactional.ArchiveToken(context.Background(), "token-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTransactional_ResetPassword(t *testing.T) {
	svc, mux, teardown := setupMarketing(t)
	defer teardown()

	mux.HandleFunc("/marketing/v3/transactional/smtp-tokens/token-001/password-reset", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, marketing.SmtpApiTokenView{
			ID:       "token-001",
			Password: "new-password",
		}))
	})

	token, err := svc.Transactional.ResetPassword(context.Background(), "token-001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token.Password != "new-password" {
		t.Errorf("Password = %q", token.Password)
	}
}

// --- Service Container Test ---

func TestNewService(t *testing.T) {
	svc, _, teardown := setupMarketing(t)
	defer teardown()

	if svc.Forms == nil {
		t.Error("Forms is nil")
	}
	if svc.Emails == nil {
		t.Error("Emails is nil")
	}
	if svc.Statistics == nil {
		t.Error("Statistics is nil")
	}
	if svc.Events == nil {
		t.Error("Events is nil")
	}
	if svc.Transactional == nil {
		t.Error("Transactional is nil")
	}
}
