package hubspot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// setup creates a test HTTP server and a Client pointing at it.
// The handler receives all requests and the caller controls responses.
// Returns the client, server mux, and a teardown function.
func setup(t *testing.T) (*Client, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(
		WithAccessToken("test-token"),
		WithBaseURL(server.URL),
	)
	return client, mux, server.Close
}

// mustJSON marshals v to JSON or fails the test.
func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}
	return b
}

// assertMethod checks the HTTP method of the request.
func assertMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if r.Method != want {
		t.Errorf("method = %s, want %s", r.Method, want)
	}
}

// assertAuth checks the Authorization header.
func assertAuth(t *testing.T, r *http.Request, wantToken string) {
	t.Helper()
	got := r.Header.Get("Authorization")
	want := "Bearer " + wantToken
	if got != want {
		t.Errorf("Authorization = %q, want %q", got, want)
	}
}
