package oauth_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/scopiousdigital/hubspot-go/oauth"
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

func setupOAuth(t *testing.T) (*oauth.Service, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	requester := &mockRequester{server: server}
	svc := oauth.NewService(requester)
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

func TestAccessTokens_Get(t *testing.T) {
	svc, mux, teardown := setupOAuth(t)
	defer teardown()

	mux.HandleFunc("/oauth/v1/access-tokens/test-token-123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, oauth.AccessTokenInfo{
			HubID:     12345,
			UserID:    67890,
			Scopes:    []string{"contacts", "content"},
			TokenType: "access",
			AppID:     111,
			ExpiresIn: 1800,
			Token:     "test-token-123",
		}))
	})

	info, err := svc.AccessTokens.Get(context.Background(), "test-token-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if info.HubID != 12345 {
		t.Errorf("HubID = %d, want 12345", info.HubID)
	}
	if info.Token != "test-token-123" {
		t.Errorf("Token = %q, want test-token-123", info.Token)
	}
	if len(info.Scopes) != 2 {
		t.Errorf("Scopes = %d, want 2", len(info.Scopes))
	}
}

func TestRefreshTokens_Get(t *testing.T) {
	svc, mux, teardown := setupOAuth(t)
	defer teardown()

	mux.HandleFunc("/oauth/v1/refresh-tokens/refresh-abc", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, oauth.RefreshTokenInfo{
			HubID:    12345,
			UserID:   67890,
			Scopes:   []string{"contacts"},
			ClientID: "client-id-abc",
			Token:    "refresh-abc",
		}))
	})

	info, err := svc.RefreshTokens.Get(context.Background(), "refresh-abc")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if info.ClientID != "client-id-abc" {
		t.Errorf("ClientID = %q, want client-id-abc", info.ClientID)
	}
	if info.Token != "refresh-abc" {
		t.Errorf("Token = %q, want refresh-abc", info.Token)
	}
}

func TestRefreshTokens_Archive(t *testing.T) {
	svc, mux, teardown := setupOAuth(t)
	defer teardown()

	mux.HandleFunc("/oauth/v1/refresh-tokens/refresh-abc", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.RefreshTokens.Archive(context.Background(), "refresh-abc")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTokens_Create(t *testing.T) {
	svc, mux, teardown := setupOAuth(t)
	defer teardown()

	mux.HandleFunc("/oauth/v1/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input oauth.TokenCreateRequest
		json.Unmarshal(body, &input)
		if input.GrantType != "authorization_code" {
			t.Errorf("grant_type = %q, want authorization_code", input.GrantType)
		}
		if input.ClientID != "my-client-id" {
			t.Errorf("client_id = %q, want my-client-id", input.ClientID)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, oauth.TokenResponse{
			AccessToken:  "new-access-token",
			RefreshToken: "new-refresh-token",
			TokenType:    "bearer",
			ExpiresIn:    1800,
		}))
	})

	resp, err := svc.Tokens.Create(context.Background(), &oauth.TokenCreateRequest{
		GrantType:    "authorization_code",
		Code:         "auth-code-xyz",
		RedirectURI:  "https://example.com/callback",
		ClientID:     "my-client-id",
		ClientSecret: "my-client-secret",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.AccessToken != "new-access-token" {
		t.Errorf("AccessToken = %q", resp.AccessToken)
	}
	if resp.RefreshToken != "new-refresh-token" {
		t.Errorf("RefreshToken = %q", resp.RefreshToken)
	}
	if resp.ExpiresIn != 1800 {
		t.Errorf("ExpiresIn = %d, want 1800", resp.ExpiresIn)
	}
}

func TestTokenResponse_JSONRoundTrip(t *testing.T) {
	input := `{"access_token":"at","refresh_token":"rt","token_type":"bearer","expires_in":3600}`
	var resp oauth.TokenResponse
	if err := json.Unmarshal([]byte(input), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if resp.AccessToken != "at" {
		t.Errorf("AccessToken = %q", resp.AccessToken)
	}
	if resp.ExpiresIn != 3600 {
		t.Errorf("ExpiresIn = %d", resp.ExpiresIn)
	}
}
