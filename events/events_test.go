package events_test

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

	"github.com/scopiousdigital/hubspot-go/events"
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

func setupEvents(t *testing.T) (*events.Service, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	requester := &mockRequester{server: server}
	svc := events.NewService(requester)
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

func TestEvents_List(t *testing.T) {
	svc, mux, teardown := setupEvents(t)
	defer teardown()

	mux.HandleFunc("/events/v3/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if ot := r.URL.Query().Get("objectType"); ot != "contact" {
			t.Errorf("objectType = %q, want contact", ot)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, events.EventsListResult{
			Results: []*events.ExternalUnifiedEvent{
				{ID: "evt1", EventType: "click", ObjectID: "123", ObjectType: "contact", OccurredAt: time.Now()},
			},
			Paging: &events.Paging{Next: &events.NextPage{After: "abc"}},
		}))
	})

	result, err := svc.List(context.Background(), &events.ListEventsOptions{ObjectType: "contact"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
	if result.Results[0].ID != "evt1" {
		t.Errorf("ID = %q, want evt1", result.Results[0].ID)
	}
}

func TestEvents_GetTypes(t *testing.T) {
	svc, mux, teardown := setupEvents(t)
	defer teardown()

	mux.HandleFunc("/events/v3/events/event-types", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, events.VisibleExternalEventTypeNames{
			EventTypes: []string{"click", "view"},
		}))
	})

	result, err := svc.GetTypes(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.EventTypes) != 2 {
		t.Errorf("eventTypes = %d, want 2", len(result.EventTypes))
	}
}

func TestEvents_Send(t *testing.T) {
	svc, mux, teardown := setupEvents(t)
	defer teardown()

	mux.HandleFunc("/events/v3/send", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input events.BehavioralEventHttpCompletionRequest
		json.Unmarshal(body, &input)
		if input.EventName != "pe12345_test_event" {
			t.Errorf("eventName = %q", input.EventName)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Send.Send(context.Background(), &events.BehavioralEventHttpCompletionRequest{
		EventName: "pe12345_test_event",
		Email:     "test@example.com",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestEvents_SendBatch(t *testing.T) {
	svc, mux, teardown := setupEvents(t)
	defer teardown()

	mux.HandleFunc("/events/v3/send/batch", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input events.BatchedBehavioralEventHttpCompletionRequest
		json.Unmarshal(body, &input)
		if len(input.Inputs) != 2 {
			t.Errorf("inputs = %d, want 2", len(input.Inputs))
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Send.SendBatch(context.Background(), &events.BatchedBehavioralEventHttpCompletionRequest{
		Inputs: []events.BehavioralEventHttpCompletionRequest{
			{EventName: "event1", Email: "a@test.com"},
			{EventName: "event2", Email: "b@test.com"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
