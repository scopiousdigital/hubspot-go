package hubspot

import (
	"context"
	"net/http"
	"strings"
	"time"
)

// retrier handles automatic retry of failed API requests.
// Matches the Node client's RetryDecorator behavior.
type retrier struct {
	maxRetries int
}

func newRetrier(maxRetries int) *retrier {
	return &retrier{maxRetries: maxRetries}
}

// do executes a request with automatic retries on transient failures.
func (r *retrier) do(ctx context.Context, hc *httpClient, rc requestConfig) (*http.Response, error) {
	var resp *http.Response
	var err error

	for attempt := 0; attempt <= r.maxRetries; attempt++ {
		resp, err = hc.doRaw(ctx, rc)
		if err != nil {
			return nil, err
		}

		if !r.shouldRetry(resp, attempt) {
			return resp, nil
		}

		delay := r.retryDelay(resp, attempt+1)

		// Close the body before retrying
		resp.Body.Close()

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(delay):
		}
	}

	return resp, nil
}

// shouldRetry determines if a response warrants a retry.
func (r *retrier) shouldRetry(resp *http.Response, attempt int) bool {
	if attempt >= r.maxRetries {
		return false
	}

	// Retry on 5xx server errors
	if resp.StatusCode >= 500 {
		return true
	}

	// Retry on 429 rate limit
	if resp.StatusCode == http.StatusTooManyRequests {
		return true
	}

	return false
}

// retryDelay calculates the wait time before the next retry.
// Matches the Node client's retry timing:
//   - 5xx errors: 200ms * retryNumber
//   - 429 with TEN_SECONDLY_ROLLING: 10s * retryNumber
//   - 429 with "secondly limit": 1s * retryNumber
//   - 429 other: 1s * retryNumber
func (r *retrier) retryDelay(resp *http.Response, retryNumber int) time.Duration {
	if resp.StatusCode >= 500 {
		return 200 * time.Millisecond * time.Duration(retryNumber)
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		// Check Retry-After header first
		if ra := resp.Header.Get("Retry-After"); ra != "" {
			if d, err := time.ParseDuration(ra + "s"); err == nil {
				return d
			}
		}

		// Check for HubSpot-specific rate limit policies via header/body heuristics
		policyName := resp.Header.Get("X-HubSpot-RateLimit-Policy")
		if strings.Contains(policyName, "TEN_SECONDLY") {
			return 10 * time.Second * time.Duration(retryNumber)
		}

		return time.Second * time.Duration(retryNumber)
	}

	return time.Second * time.Duration(retryNumber)
}
