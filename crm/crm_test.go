package crm_test

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

	"github.com/scopiousdigital/hubspot-go/crm"
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

func setupCRM(t *testing.T) (*crm.Service, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	requester := &mockRequester{server: server}
	svc := crm.NewService(requester)
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

// --- Tests ---

func TestContacts_Create(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/contacts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input crm.SimplePublicObjectInputForCreate
		json.Unmarshal(body, &input)
		if input.Properties["email"] != "test@example.com" {
			t.Errorf("email = %q", input.Properties["email"])
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.SimplePublicObject{
			ID:         "501",
			Properties: crm.Properties{"email": "test@example.com"},
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}))
	})

	contact, err := svc.Contacts.Create(context.Background(), &crm.SimplePublicObjectInputForCreate{
		Properties: crm.Properties{"email": "test@example.com"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if contact.ID != "501" {
		t.Errorf("ID = %q, want 501", contact.ID)
	}
}

func TestContacts_GetByID(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/contacts/501", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if props := r.URL.Query().Get("properties"); props != "email,firstname" {
			t.Errorf("properties = %q", props)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.SimplePublicObjectWithAssociations{
			ID:         "501",
			Properties: crm.Properties{"email": "test@example.com"},
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}))
	})

	contact, err := svc.Contacts.GetByID(context.Background(), "501", &crm.GetByIDOptions{
		Properties: []string{"email", "firstname"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if contact.ID != "501" {
		t.Errorf("ID = %q", contact.ID)
	}
}

func TestContacts_Update(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/contacts/501", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %s, want PATCH", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.SimplePublicObject{
			ID:         "501",
			Properties: crm.Properties{"firstname": "Updated"},
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}))
	})

	contact, err := svc.Contacts.Update(context.Background(), "501", &crm.SimplePublicObjectInput{
		Properties: crm.Properties{"firstname": "Updated"},
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if contact.Properties["firstname"] != "Updated" {
		t.Errorf("firstname = %q", contact.Properties["firstname"])
	}
}

func TestContacts_Archive(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/contacts/501", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Contacts.Archive(context.Background(), "501")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestContacts_List(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/contacts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if limit := r.URL.Query().Get("limit"); limit != "10" {
			t.Errorf("limit = %q, want 10", limit)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ListResult{
			Results: []*crm.SimplePublicObjectWithAssociations{
				{ID: "1", Properties: crm.Properties{"email": "a@test.com"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
				{ID: "2", Properties: crm.Properties{"email": "b@test.com"}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
			Paging: &crm.ForwardPaging{Next: &crm.NextPage{After: "2"}},
		}))
	})

	result, err := svc.Contacts.List(context.Background(), &crm.ListOptions{Limit: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d, want 2", len(result.Results))
	}
	if result.Paging.Next.After != "2" {
		t.Errorf("after = %q", result.Paging.Next.After)
	}
}

func TestContacts_Search(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/contacts/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var req crm.PublicObjectSearchRequest
		json.Unmarshal(body, &req)
		if req.FilterGroups[0].Filters[0].Operator != crm.FilterOperatorEQ {
			t.Errorf("operator = %q", req.FilterGroups[0].Filters[0].Operator)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.SearchResult{
			Total:   1,
			Results: []*crm.SimplePublicObject{{ID: "501", Properties: crm.Properties{}, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}))
	})

	result, err := svc.Contacts.Search(context.Background(), &crm.PublicObjectSearchRequest{
		FilterGroups: []crm.FilterGroup{{
			Filters: []crm.Filter{{PropertyName: "email", Operator: crm.FilterOperatorEQ, Value: "test@test.com"}},
		}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 1 {
		t.Errorf("total = %d", result.Total)
	}
}

func TestContacts_BatchCreate(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/contacts/batch/create", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var input crm.BatchCreateInput
		json.Unmarshal(body, &input)
		if len(input.Inputs) != 2 {
			t.Errorf("inputs = %d, want 2", len(input.Inputs))
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.BatchResult{
			Status: "COMPLETE",
			Results: []*crm.SimplePublicObject{
				{ID: "1", Properties: crm.Properties{}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
				{ID: "2", Properties: crm.Properties{}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
			StartedAt: time.Now(), CompletedAt: time.Now(),
		}))
	})

	result, err := svc.Contacts.BatchCreate(context.Background(), &crm.BatchCreateInput{
		Inputs: []crm.SimplePublicObjectInputForCreate{
			{Properties: crm.Properties{"email": "a@test.com"}},
			{Properties: crm.Properties{"email": "b@test.com"}},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Status != "COMPLETE" {
		t.Errorf("status = %q", result.Status)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d", len(result.Results))
	}
}

func TestContacts_BatchArchive(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/contacts/batch/archive", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Contacts.BatchArchive(context.Background(), &crm.BatchArchiveInput{
		Inputs: []crm.ObjectID{{ID: "1"}, {ID: "2"}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestContacts_Merge(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/contacts/merge", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.SimplePublicObject{
			ID: "501", Properties: crm.Properties{}, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}))
	})

	result, err := svc.Contacts.Merge(context.Background(), &crm.PublicMergeInput{
		PrimaryObjectID: "501", ObjectIDToMerge: "502",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "501" {
		t.Errorf("ID = %q", result.ID)
	}
}

func TestContacts_GdprDelete(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/contacts/gdpr-delete", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Contacts.GdprDelete(context.Background(), &crm.PublicGdprDeleteInput{ObjectID: "501"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestContacts_GetAll(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	callCount := 0
	mux.HandleFunc("/crm/v3/objects/contacts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			return
		}
		callCount++
		var result crm.ListResult
		switch callCount {
		case 1:
			result = crm.ListResult{
				Results: []*crm.SimplePublicObjectWithAssociations{
					{ID: "1", Properties: crm.Properties{}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
					{ID: "2", Properties: crm.Properties{}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
				},
				Paging: &crm.ForwardPaging{Next: &crm.NextPage{After: "2"}},
			}
		case 2:
			if after := r.URL.Query().Get("after"); after != "2" {
				t.Errorf("page 2 after = %q, want 2", after)
			}
			result = crm.ListResult{
				Results: []*crm.SimplePublicObjectWithAssociations{
					{ID: "3", Properties: crm.Properties{}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
				},
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, result))
	})

	all, err := svc.Contacts.GetAll(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(all) != 3 {
		t.Errorf("total = %d, want 3", len(all))
	}
	if callCount != 2 {
		t.Errorf("API calls = %d, want 2", callCount)
	}
}

func TestDifferentObjectTypes(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	for _, objType := range []string{"contacts", "companies", "deals", "tickets"} {
		mux.HandleFunc("/crm/v3/objects/"+objType+"/123", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(mustJSON(t, crm.SimplePublicObjectWithAssociations{
				ID: "123", Properties: crm.Properties{}, CreatedAt: time.Now(), UpdatedAt: time.Now(),
			}))
		})
	}

	services := map[string]*crm.ObjectService{
		"contacts":  svc.Contacts,
		"companies": svc.Companies,
		"deals":     svc.Deals,
		"tickets":   svc.Tickets,
	}

	for name, objSvc := range services {
		t.Run(name, func(t *testing.T) {
			result, err := objSvc.GetByID(context.Background(), "123", nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.ID != "123" {
				t.Errorf("ID = %q", result.ID)
			}
		})
	}
}

func TestProperties_UnmarshalJSON_HandlesNulls(t *testing.T) {
	input := `{"firstname": "John", "lastname": null, "email": "john@test.com"}`
	var props crm.Properties
	if err := json.Unmarshal([]byte(input), &props); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if props["firstname"] != "John" {
		t.Errorf("firstname = %q", props["firstname"])
	}
	if _, exists := props["lastname"]; exists {
		t.Error("null lastname should be omitted")
	}
}

func TestSimplePublicObject_Deserialize(t *testing.T) {
	input := `{
		"id": "501",
		"properties": {"email": "test@example.com", "lastname": null},
		"createdAt": "2024-01-15T10:30:00.000Z",
		"updatedAt": "2024-01-16T11:00:00.000Z",
		"archived": false
	}`
	var obj crm.SimplePublicObject
	if err := json.Unmarshal([]byte(input), &obj); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if obj.ID != "501" {
		t.Errorf("ID = %q", obj.ID)
	}
	if obj.Properties["email"] != "test@example.com" {
		t.Errorf("email = %q", obj.Properties["email"])
	}
	if _, exists := obj.Properties["lastname"]; exists {
		t.Error("null should be omitted")
	}
	if obj.CreatedAt.IsZero() {
		t.Error("createdAt is zero")
	}
}
