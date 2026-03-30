package hubspot

import (
	"net/http"

	xrate "golang.org/x/time/rate"
)

// Option configures a Client.
type Option func(*Client)

// WithAccessToken sets the private app or OAuth access token.
func WithAccessToken(token string) Option {
	return func(c *Client) {
		c.httpClient.accessToken = token
	}
}

// WithAPIKey sets the API key (legacy).
func WithAPIKey(key string) Option {
	return func(c *Client) {
		c.httpClient.apiKey = key
	}
}

// WithDeveloperAPIKey sets the developer API key.
func WithDeveloperAPIKey(key string) Option {
	return func(c *Client) {
		c.httpClient.developerAPIKey = key
	}
}

// WithBaseURL overrides the default HubSpot API base URL.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.httpClient.baseURL = url
	}
}

// WithHTTPClient provides a custom *http.Client for requests.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient.client = client
	}
}

// WithHeaders sets default headers sent with every request.
func WithHeaders(headers map[string]string) Option {
	return func(c *Client) {
		for k, v := range headers {
			c.httpClient.headers[k] = v
		}
	}
}

// WithRetries sets the number of automatic retries for failed API calls (0-6).
func WithRetries(n int) Option {
	return func(c *Client) {
		if n > 6 {
			n = 6
		}
		if n < 0 {
			n = 0
		}
		c.httpClient.retrier = newRetrier(n)
	}
}

// WithRateLimiter enables client-side rate limiting.
// rate is requests per second, burst is the maximum burst size.
func WithRateLimiter(rate float64, burst int) Option {
	return func(c *Client) {
		c.httpClient.limiter = xrate.NewLimiter(xrate.Limit(rate), burst)
	}
}
