package hubspot

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestHTTPClient_AuthHeader(t *testing.T) {
	client, mux, teardown := setup(t)
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assertAuth(t, r, "test-token")
		assertMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	})

	var result map[string]any
	err := client.httpClient.Get(context.Background(), "/test", nil, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestHTTPClient_APIKeyAuth(t *testing.T) {
	c := NewClient(WithAPIKey("my-api-key"), WithBaseURL("https://example.com"))
	req, err := c.httpClient.newRequest(context.Background(), requestConfig{
		method: http.MethodGet,
		path:   "/test",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := req.URL.Query().Get("hapikey"); got != "my-api-key" {
		t.Errorf("hapikey = %q, want %q", got, "my-api-key")
	}
}

func TestHTTPClient_PostBody(t *testing.T) {
	client, mux, teardown := setup(t)
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, r, http.MethodPost)

		body, _ := io.ReadAll(r.Body)
		var got map[string]any
		_ = json.Unmarshal(body, &got)

		if got["name"] != "test" {
			t.Errorf("body name = %v, want %q", got["name"], "test")
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, map[string]string{"id": "123"}))
	})

	var result map[string]string
	err := client.httpClient.Post(context.Background(), "/test", map[string]string{"name": "test"}, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["id"] != "123" {
		t.Errorf("id = %q, want %q", result["id"], "123")
	}
}

func TestHTTPClient_APIError(t *testing.T) {
	client, mux, teardown := setup(t)
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{
			"status": "error",
			"message": "Object not found",
			"correlationId": "abc-123",
			"category": "OBJECT_NOT_FOUND"
		}`))
	})

	var result map[string]any
	err := client.httpClient.Get(context.Background(), "/test", nil, &result)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.HTTPStatusCode != 404 {
		t.Errorf("status code = %d, want 404", apiErr.HTTPStatusCode)
	}
	if apiErr.Category != "OBJECT_NOT_FOUND" {
		t.Errorf("category = %q, want OBJECT_NOT_FOUND", apiErr.Category)
	}
	if !IsNotFound(err) {
		t.Error("IsNotFound should return true")
	}
}

func TestHTTPClient_ContentTypeHeaders(t *testing.T) {
	client, mux, teardown := setup(t)
	defer teardown()

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		if ct := r.Header.Get("Content-Type"); ct != "application/json" {
			t.Errorf("Content-Type = %q, want application/json", ct)
		}
		if accept := r.Header.Get("Accept"); accept != "application/json" {
			t.Errorf("Accept = %q, want application/json", accept)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	})

	var result map[string]any
	_ = client.httpClient.Get(context.Background(), "/test", nil, &result)
}
