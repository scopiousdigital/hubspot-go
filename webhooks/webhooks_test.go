package webhooks_test

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

	"github.com/scopiousdigital/hubspot-go/webhooks"
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
	b, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, m.server.URL+path, bytes.NewReader(b))
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

func setupWebhooks(t *testing.T) (*webhooks.Service, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	requester := &mockRequester{server: server}
	svc := webhooks.NewService(requester)
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

// --- Settings Tests ---

func TestSettings_GetAll(t *testing.T) {
	svc, mux, teardown := setupWebhooks(t)
	defer teardown()

	mux.HandleFunc("/webhooks/v3/12345/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, webhooks.SettingsResponse{
			CreatedAt: time.Now(),
			TargetURL: "https://example.com/webhook",
			Throttling: webhooks.ThrottlingSettings{
				MaxConcurrentRequests: 10,
			},
		}))
	})

	settings, err := svc.Settings.GetAll(context.Background(), 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if settings.TargetURL != "https://example.com/webhook" {
		t.Errorf("TargetURL = %q", settings.TargetURL)
	}
	if settings.Throttling.MaxConcurrentRequests != 10 {
		t.Errorf("MaxConcurrentRequests = %d", settings.Throttling.MaxConcurrentRequests)
	}
}

func TestSettings_Configure(t *testing.T) {
	svc, mux, teardown := setupWebhooks(t)
	defer teardown()

	mux.HandleFunc("/webhooks/v3/12345/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input webhooks.SettingsChangeRequest
		json.Unmarshal(body, &input)
		if input.TargetURL != "https://example.com/new-webhook" {
			t.Errorf("TargetURL = %q", input.TargetURL)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, webhooks.SettingsResponse{
			CreatedAt: time.Now(),
			TargetURL: input.TargetURL,
			Throttling: input.Throttling,
		}))
	})

	settings, err := svc.Settings.Configure(context.Background(), 12345, &webhooks.SettingsChangeRequest{
		TargetURL:  "https://example.com/new-webhook",
		Throttling: webhooks.ThrottlingSettings{MaxConcurrentRequests: 5},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if settings.TargetURL != "https://example.com/new-webhook" {
		t.Errorf("TargetURL = %q", settings.TargetURL)
	}
}

func TestSettings_Clear(t *testing.T) {
	svc, mux, teardown := setupWebhooks(t)
	defer teardown()

	mux.HandleFunc("/webhooks/v3/12345/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Settings.Clear(context.Background(), 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// --- Subscriptions Tests ---

func TestSubscriptions_GetAll(t *testing.T) {
	svc, mux, teardown := setupWebhooks(t)
	defer teardown()

	mux.HandleFunc("/webhooks/v3/12345/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, webhooks.SubscriptionListResponse{
			Results: []webhooks.SubscriptionResponse{
				{ID: "1", EventType: webhooks.EventTypeContactCreation, Active: true, CreatedAt: time.Now()},
				{ID: "2", EventType: webhooks.EventTypeDealCreation, Active: false, CreatedAt: time.Now()},
			},
		}))
	})

	list, err := svc.Subscriptions.GetAll(context.Background(), 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(list.Results) != 2 {
		t.Errorf("results = %d, want 2", len(list.Results))
	}
	if list.Results[0].EventType != webhooks.EventTypeContactCreation {
		t.Errorf("EventType = %q", list.Results[0].EventType)
	}
}

func TestSubscriptions_Create(t *testing.T) {
	svc, mux, teardown := setupWebhooks(t)
	defer teardown()

	mux.HandleFunc("/webhooks/v3/12345/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input webhooks.SubscriptionCreateRequest
		json.Unmarshal(body, &input)
		if input.EventType != webhooks.EventTypeContactCreation {
			t.Errorf("EventType = %q", input.EventType)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, webhooks.SubscriptionResponse{
			ID:        "100",
			EventType: input.EventType,
			Active:    true,
			CreatedAt: time.Now(),
		}))
	})

	sub, err := svc.Subscriptions.Create(context.Background(), 12345, &webhooks.SubscriptionCreateRequest{
		EventType: webhooks.EventTypeContactCreation,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sub.ID != "100" {
		t.Errorf("ID = %q, want 100", sub.ID)
	}
}

func TestSubscriptions_GetByID(t *testing.T) {
	svc, mux, teardown := setupWebhooks(t)
	defer teardown()

	mux.HandleFunc("/webhooks/v3/12345/subscriptions/100", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, webhooks.SubscriptionResponse{
			ID:        "100",
			EventType: webhooks.EventTypeContactCreation,
			Active:    true,
			CreatedAt: time.Now(),
		}))
	})

	sub, err := svc.Subscriptions.GetByID(context.Background(), 100, 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sub.ID != "100" {
		t.Errorf("ID = %q", sub.ID)
	}
}

func TestSubscriptions_Update(t *testing.T) {
	svc, mux, teardown := setupWebhooks(t)
	defer teardown()

	mux.HandleFunc("/webhooks/v3/12345/subscriptions/100", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %s, want PATCH", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input webhooks.SubscriptionPatchRequest
		json.Unmarshal(body, &input)
		if input.Active == nil || *input.Active != false {
			t.Errorf("Active should be false")
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, webhooks.SubscriptionResponse{
			ID:        "100",
			EventType: webhooks.EventTypeContactCreation,
			Active:    false,
			CreatedAt: time.Now(),
		}))
	})

	active := false
	sub, err := svc.Subscriptions.Update(context.Background(), 100, 12345, &webhooks.SubscriptionPatchRequest{
		Active: &active,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sub.Active {
		t.Error("Active should be false")
	}
}

func TestSubscriptions_Archive(t *testing.T) {
	svc, mux, teardown := setupWebhooks(t)
	defer teardown()

	mux.HandleFunc("/webhooks/v3/12345/subscriptions/100", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Subscriptions.Archive(context.Background(), 100, 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSubscriptions_BatchUpdate(t *testing.T) {
	svc, mux, teardown := setupWebhooks(t)
	defer teardown()

	mux.HandleFunc("/webhooks/v3/12345/subscriptions/batch/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input webhooks.BatchInputSubscriptionBatchUpdateRequest
		json.Unmarshal(body, &input)
		if len(input.Inputs) != 2 {
			t.Errorf("inputs = %d, want 2", len(input.Inputs))
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, webhooks.BatchResponseSubscriptionResponse{
			Status:      "COMPLETE",
			CompletedAt: time.Now(),
			StartedAt:   time.Now(),
			Results: []webhooks.SubscriptionResponse{
				{ID: "1", Active: true, CreatedAt: time.Now()},
				{ID: "2", Active: false, CreatedAt: time.Now()},
			},
		}))
	})

	result, err := svc.Subscriptions.BatchUpdate(context.Background(), 12345, &webhooks.BatchInputSubscriptionBatchUpdateRequest{
		Inputs: []webhooks.SubscriptionBatchUpdateRequest{
			{ID: 1, Active: true},
			{ID: 2, Active: false},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Status != "COMPLETE" {
		t.Errorf("Status = %q", result.Status)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d", len(result.Results))
	}
}
