package settings_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/scopiousdigital/hubspot-go/settings"
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

func setupSettings(t *testing.T) (*settings.Service, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	requester := &mockRequester{server: server}
	svc := settings.NewService(requester)
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

func TestBusinessUnits_GetByUserID(t *testing.T) {
	svc, mux, teardown := setupSettings(t)
	defer teardown()

	mux.HandleFunc("/settings/v3/business-units/users/12345", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, settings.BusinessUnitsResult{
			Results: []*settings.PublicBusinessUnit{
				{ID: "bu1", Name: "Main Unit"},
			},
		}))
	})

	result, err := svc.BusinessUnits.GetByUserID(context.Background(), "12345", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Fatalf("results = %d, want 1", len(result.Results))
	}
	if result.Results[0].Name != "Main Unit" {
		t.Errorf("name = %q", result.Results[0].Name)
	}
}

func TestUsers_Create(t *testing.T) {
	svc, mux, teardown := setupSettings(t)
	defer teardown()

	mux.HandleFunc("/settings/v3/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input settings.UserProvisionRequest
		json.Unmarshal(body, &input)
		if input.Email != "new@example.com" {
			t.Errorf("email = %q", input.Email)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, settings.PublicUser{
			ID:    "u1",
			Email: "new@example.com",
		}))
	})

	user, err := svc.Users.Create(context.Background(), &settings.UserProvisionRequest{
		Email: "new@example.com",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.ID != "u1" {
		t.Errorf("ID = %q", user.ID)
	}
}

func TestUsers_GetByID(t *testing.T) {
	svc, mux, teardown := setupSettings(t)
	defer teardown()

	mux.HandleFunc("/settings/v3/users/u1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, settings.PublicUser{ID: "u1", Email: "test@example.com"}))
	})

	user, err := svc.Users.GetByID(context.Background(), "u1", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if user.Email != "test@example.com" {
		t.Errorf("email = %q", user.Email)
	}
}

func TestUsers_List(t *testing.T) {
	svc, mux, teardown := setupSettings(t)
	defer teardown()

	mux.HandleFunc("/settings/v3/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, settings.UsersListResult{
			Results: []*settings.PublicUser{
				{ID: "u1", Email: "a@test.com"},
				{ID: "u2", Email: "b@test.com"},
			},
		}))
	})

	result, err := svc.Users.List(context.Background(), &settings.UsersListOptions{Limit: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d, want 2", len(result.Results))
	}
}

func TestUsers_Archive(t *testing.T) {
	svc, mux, teardown := setupSettings(t)
	defer teardown()

	mux.HandleFunc("/settings/v3/users/u1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Users.Archive(context.Background(), "u1", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRoles_GetAll(t *testing.T) {
	svc, mux, teardown := setupSettings(t)
	defer teardown()

	mux.HandleFunc("/settings/v3/users/roles", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, settings.RolesResult{
			Results: []*settings.PublicPermissionSet{
				{ID: "r1", Name: "Admin", RequiresBillingWrite: true},
			},
		}))
	})

	result, err := svc.Roles.GetAll(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Fatalf("results = %d, want 1", len(result.Results))
	}
	if result.Results[0].Name != "Admin" {
		t.Errorf("name = %q", result.Results[0].Name)
	}
}

func TestTeams_GetAll(t *testing.T) {
	svc, mux, teardown := setupSettings(t)
	defer teardown()

	mux.HandleFunc("/settings/v3/users/teams", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, settings.TeamsResult{
			Results: []*settings.PublicTeam{
				{ID: "t1", Name: "Sales", UserIDs: []string{"u1"}, SecondaryUserIDs: []string{}},
			},
		}))
	})

	result, err := svc.Teams.GetAll(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Fatalf("results = %d, want 1", len(result.Results))
	}
	if result.Results[0].Name != "Sales" {
		t.Errorf("name = %q", result.Results[0].Name)
	}
}
