package hubspot

import (
	"net/http"
	"testing"
	"time"
)

func TestNewClient_Defaults(t *testing.T) {
	c := NewClient(WithAccessToken("tok"))

	if c.httpClient.baseURL != defaultBaseURL {
		t.Errorf("baseURL = %q, want %q", c.httpClient.baseURL, defaultBaseURL)
	}
	if c.httpClient.accessToken != "tok" {
		t.Errorf("accessToken = %q, want %q", c.httpClient.accessToken, "tok")
	}
	if c.httpClient.userAgent != defaultUserAgent {
		t.Errorf("userAgent = %q, want %q", c.httpClient.userAgent, defaultUserAgent)
	}
	if c.CRM == nil {
		t.Fatal("CRM service is nil")
	}
	if c.CRM.Contacts == nil {
		t.Fatal("CRM.Contacts is nil")
	}
}

func TestNewClient_Options(t *testing.T) {
	customHTTP := &http.Client{Timeout: 60 * time.Second}
	c := NewClient(
		WithAccessToken("my-token"),
		WithBaseURL("https://custom.api.com"),
		WithHTTPClient(customHTTP),
		WithHeaders(map[string]string{"X-Custom": "value"}),
		WithRetries(3),
		WithRateLimiter(10, 5),
	)

	if c.httpClient.accessToken != "my-token" {
		t.Errorf("accessToken = %q", c.httpClient.accessToken)
	}
	if c.httpClient.baseURL != "https://custom.api.com" {
		t.Errorf("baseURL = %q", c.httpClient.baseURL)
	}
	if c.httpClient.client != customHTTP {
		t.Error("custom HTTP client not set")
	}
	if c.httpClient.headers["X-Custom"] != "value" {
		t.Error("custom header not set")
	}
	if c.httpClient.retrier == nil {
		t.Fatal("retrier is nil")
	}
	if c.httpClient.retrier.maxRetries != 3 {
		t.Errorf("maxRetries = %d, want 3", c.httpClient.retrier.maxRetries)
	}
	if c.httpClient.limiter == nil {
		t.Fatal("limiter is nil")
	}
}

func TestNewClient_RetriesClamp(t *testing.T) {
	c := NewClient(WithRetries(10))
	if c.httpClient.retrier.maxRetries != 6 {
		t.Errorf("maxRetries = %d, want 6 (clamped)", c.httpClient.retrier.maxRetries)
	}

	c = NewClient(WithRetries(-1))
	if c.httpClient.retrier.maxRetries != 0 {
		t.Errorf("maxRetries = %d, want 0 (clamped)", c.httpClient.retrier.maxRetries)
	}
}

func TestSetAccessToken(t *testing.T) {
	c := NewClient(WithAccessToken("old"))
	c.SetAccessToken("new")
	if c.httpClient.accessToken != "new" {
		t.Errorf("accessToken = %q, want %q", c.httpClient.accessToken, "new")
	}
}

func TestAllServicesInitialized(t *testing.T) {
	c := NewClient(WithAccessToken("tok"))

	services := map[string]any{
		"CRM":                      c.CRM,
		"CRM.Companies":           c.CRM.Companies,
		"CRM.Deals":               c.CRM.Deals,
		"CRM.Owners":              c.CRM.Owners,
		"CRM.Properties":          c.CRM.Properties,
		"CMS":                     c.CMS,
		"Marketing":               c.Marketing,
		"Automation":              c.Automation,
		"Conversations":           c.Conversations,
		"Events":                  c.Events,
		"Files":                   c.Files,
		"Settings":                c.Settings,
		"OAuth":                   c.OAuth,
		"Webhooks":                c.Webhooks,
		"CommunicationPreferences": c.CommunicationPreferences,
	}
	for name, svc := range services {
		if svc == nil {
			t.Errorf("%s is nil", name)
		}
	}
}
