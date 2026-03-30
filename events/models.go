package events

import "time"

// --- Events API models ---

// ExternalUnifiedEvent represents a single event from HubSpot's events API.
type ExternalUnifiedEvent struct {
	OccurredAt time.Time         `json:"occurredAt"`
	EventType  string            `json:"eventType"`
	ID         string            `json:"id"`
	ObjectID   string            `json:"objectId"`
	Properties map[string]string `json:"properties,omitempty"`
	ObjectType string            `json:"objectType"`
}

// EventsListResult is the paginated response for listing events.
type EventsListResult struct {
	Results []*ExternalUnifiedEvent `json:"results"`
	Paging  *Paging                 `json:"paging,omitempty"`
}

// VisibleExternalEventTypeNames contains a list of event type names.
type VisibleExternalEventTypeNames struct {
	EventTypes []string `json:"eventTypes"`
}

// ListEventsOptions contains the query parameters for listing events.
type ListEventsOptions struct {
	ObjectType     string   `json:"-"`
	EventType      string   `json:"-"`
	After          string   `json:"-"`
	Before         string   `json:"-"`
	Limit          int      `json:"-"`
	Sort           []string `json:"-"`
	OccurredAfter  string   `json:"-"`
	OccurredBefore string   `json:"-"`
	ObjectID       int64    `json:"-"`
	ID             []string `json:"-"`
}

// --- Send API models ---

// BehavioralEventHttpCompletionRequest is the body for sending a single behavioral event.
type BehavioralEventHttpCompletionRequest struct {
	OccurredAt *time.Time        `json:"occurredAt,omitempty"`
	EventName  string            `json:"eventName"`
	Utk        string            `json:"utk,omitempty"`
	UUID       string            `json:"uuid,omitempty"`
	Email      string            `json:"email,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
	ObjectID   string            `json:"objectId,omitempty"`
}

// BatchedBehavioralEventHttpCompletionRequest wraps multiple events for a batch send.
type BatchedBehavioralEventHttpCompletionRequest struct {
	Inputs []BehavioralEventHttpCompletionRequest `json:"inputs"`
}

// --- Shared paging models ---

// Paging contains pagination cursors.
type Paging struct {
	Next *NextPage     `json:"next,omitempty"`
	Prev *PreviousPage `json:"prev,omitempty"`
}

// NextPage contains the cursor for the next page of results.
type NextPage struct {
	Link  string `json:"link,omitempty"`
	After string `json:"after"`
}

// PreviousPage contains the cursor for the previous page of results.
type PreviousPage struct {
	Before string `json:"before"`
	Link   string `json:"link,omitempty"`
}
