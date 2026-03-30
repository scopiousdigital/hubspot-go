package automation_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/scopiousdigital/hubspot-go/automation"
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

func setupAutomation(t *testing.T) (*automation.Service, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	requester := &mockRequester{server: server}
	svc := automation.NewService(requester)
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

func TestCallbacks_Complete(t *testing.T) {
	svc, mux, teardown := setupAutomation(t)
	defer teardown()

	mux.HandleFunc("/automation/v4/actions/callbacks/cb1/complete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input automation.CallbackCompletionRequest
		json.Unmarshal(body, &input)
		if input.OutputFields["hs_status"] != "SUCCESS" {
			t.Errorf("hs_status = %q", input.OutputFields["hs_status"])
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Callbacks.Complete(context.Background(), "cb1", &automation.CallbackCompletionRequest{
		OutputFields: map[string]string{"hs_status": "SUCCESS"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCallbacks_CompleteBatch(t *testing.T) {
	svc, mux, teardown := setupAutomation(t)
	defer teardown()

	mux.HandleFunc("/automation/v4/actions/callbacks/complete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input automation.BatchInputCallbackCompletionBatchRequest
		json.Unmarshal(body, &input)
		if len(input.Inputs) != 2 {
			t.Errorf("inputs = %d, want 2", len(input.Inputs))
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Callbacks.CompleteBatch(context.Background(), &automation.BatchInputCallbackCompletionBatchRequest{
		Inputs: []automation.CallbackCompletionBatchRequest{
			{CallbackID: "cb1", OutputFields: map[string]string{"hs_status": "SUCCESS"}},
			{CallbackID: "cb2", OutputFields: map[string]string{"hs_status": "FAILURE"}},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDefinitions_Create(t *testing.T) {
	svc, mux, teardown := setupAutomation(t)
	defer teardown()

	mux.HandleFunc("/automation/v4/actions/12345", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, automation.PublicActionDefinition{
			ID:        "def1",
			ActionURL: "https://example.com/action",
			Published: true,
		}))
	})

	result, err := svc.Definitions.Create(context.Background(), 12345, &automation.PublicActionDefinitionEgg{
		ActionURL:   "https://example.com/action",
		Published:   true,
		ObjectTypes: []string{"CONTACT"},
		Labels:      map[string]automation.PublicActionLabels{"en": {ActionName: "Test"}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "def1" {
		t.Errorf("ID = %q", result.ID)
	}
}

func TestDefinitions_GetByID(t *testing.T) {
	svc, mux, teardown := setupAutomation(t)
	defer teardown()

	mux.HandleFunc("/automation/v4/actions/12345/def1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, automation.PublicActionDefinition{
			ID:        "def1",
			ActionURL: "https://example.com/action",
		}))
	})

	result, err := svc.Definitions.GetByID(context.Background(), "def1", 12345, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "def1" {
		t.Errorf("ID = %q", result.ID)
	}
}

func TestDefinitions_List(t *testing.T) {
	svc, mux, teardown := setupAutomation(t)
	defer teardown()

	mux.HandleFunc("/automation/v4/actions/12345", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, automation.DefinitionsListResult{
			Results: []*automation.PublicActionDefinition{
				{ID: "def1"},
				{ID: "def2"},
			},
		}))
	})

	result, err := svc.Definitions.List(context.Background(), 12345, &automation.DefinitionsListOptions{Limit: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d, want 2", len(result.Results))
	}
}

func TestDefinitions_Archive(t *testing.T) {
	svc, mux, teardown := setupAutomation(t)
	defer teardown()

	mux.HandleFunc("/automation/v4/actions/12345/def1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Definitions.Archive(context.Background(), "def1", 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFunctions_List(t *testing.T) {
	svc, mux, teardown := setupAutomation(t)
	defer teardown()

	mux.HandleFunc("/automation/v4/actions/12345/def1/functions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, automation.FunctionsResult{
			Results: []*automation.PublicActionFunctionIdentifier{
				{FunctionType: "PRE_ACTION_EXECUTION", ID: "fn1"},
			},
		}))
	})

	result, err := svc.Functions.List(context.Background(), "def1", 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Fatalf("results = %d, want 1", len(result.Results))
	}
	if result.Results[0].FunctionType != "PRE_ACTION_EXECUTION" {
		t.Errorf("functionType = %q", result.Results[0].FunctionType)
	}
}

func TestRevisions_GetByID(t *testing.T) {
	svc, mux, teardown := setupAutomation(t)
	defer teardown()

	mux.HandleFunc("/automation/v4/actions/12345/def1/revisions/rev1", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, automation.PublicActionRevision{
			ID:         "rev1",
			RevisionID: "rev1",
		}))
	})

	result, err := svc.Revisions.GetByID(context.Background(), "def1", "rev1", 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.RevisionID != "rev1" {
		t.Errorf("revisionId = %q", result.RevisionID)
	}
}
