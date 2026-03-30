package hubspot

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/time/rate"
)

type httpClient struct {
	client          *http.Client
	baseURL         string
	accessToken     string
	apiKey          string
	developerAPIKey string
	userAgent       string
	headers         map[string]string
	retrier         *retrier
	limiter         *rate.Limiter
}

// requestConfig holds per-request configuration.
type requestConfig struct {
	method string
	path   string
	query  url.Values
	body   any
	result any
}

// newRequest builds an *http.Request from the config.
func (hc *httpClient) newRequest(ctx context.Context, rc requestConfig) (*http.Request, error) {
	u, err := url.Parse(hc.baseURL + rc.path)
	if err != nil {
		return nil, fmt.Errorf("hubspot: invalid URL: %w", err)
	}

	if rc.query != nil {
		u.RawQuery = rc.query.Encode()
	}

	var bodyReader io.Reader
	if rc.body != nil {
		b, err := json.Marshal(rc.body)
		if err != nil {
			return nil, fmt.Errorf("hubspot: failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, rc.method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	// Auth
	if hc.accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+hc.accessToken)
	} else if hc.apiKey != "" {
		q := req.URL.Query()
		q.Set("hapikey", hc.apiKey)
		req.URL.RawQuery = q.Encode()
	} else if hc.developerAPIKey != "" {
		q := req.URL.Query()
		q.Set("hapikey", hc.developerAPIKey)
		req.URL.RawQuery = q.Encode()
	}

	// Headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", hc.userAgent)
	for k, v := range hc.headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

// do executes an HTTP request, handling rate limiting, retries, and response parsing.
func (hc *httpClient) do(ctx context.Context, rc requestConfig) error {
	if hc.limiter != nil {
		if err := hc.limiter.Wait(ctx); err != nil {
			return fmt.Errorf("hubspot: rate limiter: %w", err)
		}
	}

	var resp *http.Response
	var err error
	if hc.retrier != nil {
		resp, err = hc.retrier.do(ctx, hc, rc)
	} else {
		req, reqErr := hc.newRequest(ctx, rc)
		if reqErr != nil {
			return reqErr
		}
		resp, err = hc.client.Do(req)
	}
	if err != nil {
		return fmt.Errorf("hubspot: request failed: %w", err)
	}
	defer resp.Body.Close()

	return hc.handleResponse(resp, rc.result)
}

// doRaw executes a single HTTP request without retries (used by the retrier).
func (hc *httpClient) doRaw(ctx context.Context, rc requestConfig) (*http.Response, error) {
	req, err := hc.newRequest(ctx, rc)
	if err != nil {
		return nil, err
	}
	return hc.client.Do(req)
}

// handleResponse reads the response body and either decodes it into result or returns an APIError.
func (hc *httpClient) handleResponse(resp *http.Response, result any) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("hubspot: failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		apiErr := &APIError{HTTPStatusCode: resp.StatusCode}
		if len(body) > 0 {
			_ = json.Unmarshal(body, apiErr)
		}
		if apiErr.Status == "" {
			apiErr.Status = http.StatusText(resp.StatusCode)
		}
		return apiErr
	}

	if result != nil && len(body) > 0 {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("hubspot: failed to decode response: %w", err)
		}
	}

	return nil
}

// Exported methods satisfying api.Requester interface.
// These are called by sub-packages (crm, cms, etc.) via the interface.

// Get performs a GET request.
func (hc *httpClient) Get(ctx context.Context, path string, query url.Values, result any) error {
	return hc.do(ctx, requestConfig{method: http.MethodGet, path: path, query: query, result: result})
}

// Post performs a POST request.
func (hc *httpClient) Post(ctx context.Context, path string, body, result any) error {
	return hc.do(ctx, requestConfig{method: http.MethodPost, path: path, body: body, result: result})
}

// Put performs a PUT request.
func (hc *httpClient) Put(ctx context.Context, path string, body, result any) error {
	return hc.do(ctx, requestConfig{method: http.MethodPut, path: path, body: body, result: result})
}

// Patch performs a PATCH request.
func (hc *httpClient) Patch(ctx context.Context, path string, body, result any) error {
	return hc.do(ctx, requestConfig{method: http.MethodPatch, path: path, body: body, result: result})
}

// Delete performs a DELETE request.
func (hc *httpClient) Delete(ctx context.Context, path string) error {
	return hc.do(ctx, requestConfig{method: http.MethodDelete, path: path})
}
