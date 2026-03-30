package communicationpreferences_test

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

	"github.com/scopiousdigital/hubspot-go/communicationpreferences"
)

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

func setupCommPrefs(t *testing.T) (*communicationpreferences.Service, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	requester := &mockRequester{server: server}
	svc := communicationpreferences.NewService(requester)
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

func TestDefinitions_GetAll(t *testing.T) {
	svc, mux, teardown := setupCommPrefs(t)
	defer teardown()

	mux.HandleFunc("/communication-preferences/v3/definitions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, communicationpreferences.SubscriptionDefinitionsResponse{
			SubscriptionDefinitions: []*communicationpreferences.SubscriptionDefinition{
				{
					ID:          "sub1",
					Name:        "Marketing",
					Description: "Marketing emails",
					IsActive:    true,
					IsInternal:  false,
					IsDefault:   false,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				},
			},
		}))
	})

	result, err := svc.Definitions.GetAll(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.SubscriptionDefinitions) != 1 {
		t.Fatalf("definitions = %d, want 1", len(result.SubscriptionDefinitions))
	}
	if result.SubscriptionDefinitions[0].Name != "Marketing" {
		t.Errorf("name = %q", result.SubscriptionDefinitions[0].Name)
	}
}

func TestStatus_GetEmailStatus(t *testing.T) {
	svc, mux, teardown := setupCommPrefs(t)
	defer teardown()

	mux.HandleFunc("/communication-preferences/v3/status/email/test@example.com", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, communicationpreferences.PublicSubscriptionStatusesResponse{
			Recipient: "test@example.com",
			SubscriptionStatuses: []*communicationpreferences.PublicSubscriptionStatus{
				{
					ID:             "sub1",
					Name:           "Marketing",
					Description:    "Marketing emails",
					Status:         communicationpreferences.StatusSubscribed,
					SourceOfStatus: communicationpreferences.SourceSubscriptionStatus,
				},
			},
		}))
	})

	result, err := svc.Status.GetEmailStatus(context.Background(), "test@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Recipient != "test@example.com" {
		t.Errorf("recipient = %q", result.Recipient)
	}
	if len(result.SubscriptionStatuses) != 1 {
		t.Fatalf("statuses = %d, want 1", len(result.SubscriptionStatuses))
	}
	if result.SubscriptionStatuses[0].Status != "SUBSCRIBED" {
		t.Errorf("status = %q", result.SubscriptionStatuses[0].Status)
	}
}

func TestStatus_Subscribe(t *testing.T) {
	svc, mux, teardown := setupCommPrefs(t)
	defer teardown()

	mux.HandleFunc("/communication-preferences/v3/subscribe", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input communicationpreferences.PublicUpdateSubscriptionStatusRequest
		json.Unmarshal(body, &input)
		if input.EmailAddress != "test@example.com" {
			t.Errorf("emailAddress = %q", input.EmailAddress)
		}
		if input.SubscriptionID != "sub1" {
			t.Errorf("subscriptionId = %q", input.SubscriptionID)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, communicationpreferences.PublicSubscriptionStatus{
			ID:     "sub1",
			Name:   "Marketing",
			Status: communicationpreferences.StatusSubscribed,
		}))
	})

	result, err := svc.Status.Subscribe(context.Background(), &communicationpreferences.PublicUpdateSubscriptionStatusRequest{
		EmailAddress:   "test@example.com",
		SubscriptionID: "sub1",
		LegalBasis:     communicationpreferences.LegalBasisConsentWithNotice,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Status != "SUBSCRIBED" {
		t.Errorf("status = %q", result.Status)
	}
}

func TestStatus_Unsubscribe(t *testing.T) {
	svc, mux, teardown := setupCommPrefs(t)
	defer teardown()

	mux.HandleFunc("/communication-preferences/v3/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, communicationpreferences.PublicSubscriptionStatus{
			ID:     "sub1",
			Name:   "Marketing",
			Status: communicationpreferences.StatusNotSubscribed,
		}))
	})

	result, err := svc.Status.Unsubscribe(context.Background(), &communicationpreferences.PublicUpdateSubscriptionStatusRequest{
		EmailAddress:   "test@example.com",
		SubscriptionID: "sub1",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Status != "NOT_SUBSCRIBED" {
		t.Errorf("status = %q", result.Status)
	}
}
