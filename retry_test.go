package hubspot

import (
	"context"
	"net/http"
	"sync/atomic"
	"testing"
)

func TestRetrier_RetriesOn500(t *testing.T) {
	client, mux, teardown := setup(t)
	defer teardown()
	client.httpClient.retrier = newRetrier(2)

	var attempts int32
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n <= 2 {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{"status":"error","message":"server error"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})

	var result map[string]any
	err := client.httpClient.Get(context.Background(), "/test", nil, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := atomic.LoadInt32(&attempts); got != 3 {
		t.Errorf("attempts = %d, want 3", got)
	}
}

func TestRetrier_RetriesOn429(t *testing.T) {
	client, mux, teardown := setup(t)
	defer teardown()
	client.httpClient.retrier = newRetrier(1)

	var attempts int32
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n <= 1 {
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"status":"error","message":"rate limited"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})

	var result map[string]any
	err := client.httpClient.Get(context.Background(), "/test", nil, &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := atomic.LoadInt32(&attempts); got != 2 {
		t.Errorf("attempts = %d, want 2", got)
	}
}

func TestRetrier_GivesUpAfterMaxRetries(t *testing.T) {
	client, mux, teardown := setup(t)
	defer teardown()
	client.httpClient.retrier = newRetrier(2)

	var attempts int32
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"status":"error","message":"always fails"}`))
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
	if apiErr.HTTPStatusCode != http.StatusInternalServerError {
		t.Errorf("status = %d, want 500", apiErr.HTTPStatusCode)
	}
	if got := atomic.LoadInt32(&attempts); got != 3 {
		t.Errorf("attempts = %d, want 3", got)
	}
}

func TestRetrier_NoRetryOn4xx(t *testing.T) {
	client, mux, teardown := setup(t)
	defer teardown()
	client.httpClient.retrier = newRetrier(3)

	var attempts int32
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&attempts, 1)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"status":"error","message":"bad request"}`))
	})

	var result map[string]any
	err := client.httpClient.Get(context.Background(), "/test", nil, &result)
	if err == nil {
		t.Fatal("expected error")
	}
	if got := atomic.LoadInt32(&attempts); got != 1 {
		t.Errorf("attempts = %d, want 1 (no retries for 400)", got)
	}
}

func TestRetrier_ContextCancellation(t *testing.T) {
	client, mux, teardown := setup(t)
	defer teardown()
	client.httpClient.retrier = newRetrier(5)

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"status":"error"}`))
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	var result map[string]any
	err := client.httpClient.Get(ctx, "/test", nil, &result)
	if err == nil {
		t.Fatal("expected error from cancelled context")
	}
}
