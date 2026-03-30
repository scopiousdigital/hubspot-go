package events

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const (
	eventsBasePath = "/events/v3/events"
	sendBasePath   = "/events/v3/send"
)

// Service provides access to HubSpot Events APIs.
type Service struct {
	requester api.Requester

	// Send provides methods for sending behavioral events.
	Send *SendService
}

// NewService creates a new Events service.
func NewService(r api.Requester) *Service {
	return &Service{
		requester: r,
		Send:      &SendService{requester: r},
	}
}

// --- EventsApi ---

// List retrieves a page of events.
func (s *Service) List(ctx context.Context, opts *ListEventsOptions) (*EventsListResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.ObjectType != "" {
			q.Set("objectType", opts.ObjectType)
		}
		if opts.EventType != "" {
			q.Set("eventType", opts.EventType)
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if opts.Before != "" {
			q.Set("before", opts.Before)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if len(opts.Sort) > 0 {
			q.Set("sort", strings.Join(opts.Sort, ","))
		}
		if opts.OccurredAfter != "" {
			q.Set("occurredAfter", opts.OccurredAfter)
		}
		if opts.OccurredBefore != "" {
			q.Set("occurredBefore", opts.OccurredBefore)
		}
		if opts.ObjectID != 0 {
			q.Set("objectId", strconv.FormatInt(opts.ObjectID, 10))
		}
		if len(opts.ID) > 0 {
			for _, id := range opts.ID {
				q.Add("id", id)
			}
		}
	}
	var result EventsListResult
	if err := s.requester.Get(ctx, eventsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTypes retrieves the list of available event types.
func (s *Service) GetTypes(ctx context.Context) (*VisibleExternalEventTypeNames, error) {
	var result VisibleExternalEventTypeNames
	if err := s.requester.Get(ctx, eventsBasePath+"/event-types", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- SendService ---

// SendService provides methods for sending behavioral events.
type SendService struct {
	requester api.Requester
}

// Send sends a single behavioral event.
func (s *SendService) Send(ctx context.Context, input *BehavioralEventHttpCompletionRequest) error {
	return s.requester.Post(ctx, sendBasePath, input, nil)
}

// SendBatch sends a batch of behavioral events.
func (s *SendService) SendBatch(ctx context.Context, input *BatchedBehavioralEventHttpCompletionRequest) error {
	return s.requester.Post(ctx, sendBasePath+"/batch", input, nil)
}
