package hubspot

import "fmt"

// APIError represents an error response from the HubSpot API.
type APIError struct {
	// HTTPStatusCode is the HTTP status code of the response.
	HTTPStatusCode int `json:"-"`

	Status        string            `json:"status"`
	Message       string            `json:"message"`
	CorrelationID string            `json:"correlationId"`
	Category      string            `json:"category"`
	SubCategory   string            `json:"subCategory,omitempty"`
	Errors        []ErrorDetail     `json:"errors,omitempty"`
	Links         map[string]string `json:"links,omitempty"`
	Context       map[string][]string `json:"context,omitempty"`
}

// ErrorDetail represents a single validation or processing error.
type ErrorDetail struct {
	Message     string              `json:"message"`
	In          string              `json:"in"`
	Code        string              `json:"code"`
	SubCategory string              `json:"subCategory"`
	Context     map[string][]string `json:"context,omitempty"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("hubspot: %d %s (correlation_id: %s)", e.HTTPStatusCode, e.Message, e.CorrelationID)
	}
	return fmt.Sprintf("hubspot: %d %s", e.HTTPStatusCode, e.Status)
}

// IsNotFound returns true if the error is a 404.
func IsNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.HTTPStatusCode == 404
	}
	return false
}

// IsRateLimited returns true if the error is a 429.
func IsRateLimited(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.HTTPStatusCode == 429
	}
	return false
}

// IsConflict returns true if the error is a 409.
func IsConflict(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.HTTPStatusCode == 409
	}
	return false
}
