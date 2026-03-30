package marketing

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const eventsBasePath = "/marketing/v3/marketing-events"

// EventsService handles operations on HubSpot marketing events.
// It covers the BasicApi, BatchApi, ChangePropertyApi, IdentifiersApi,
// ListAssociationsApi, RetrieveParticipantStateApi, SettingsApi,
// SubscriberStateChangesApi, and AddEventAttendeesApi.
type EventsService struct {
	requester api.Requester
}

// --- Basic API ---

// Create creates a new marketing event.
func (s *EventsService) Create(ctx context.Context, input *MarketingEventCreateRequest) (*MarketingEventDefaultResponse, error) {
	path := eventsBasePath + "/events"
	var result MarketingEventDefaultResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAll retrieves a paginated list of marketing events.
func (s *EventsService) GetAll(ctx context.Context, after string, limit int) (*MarketingEventListResult, error) {
	path := eventsBasePath + "/events"
	q := url.Values{}
	if after != "" {
		q.Set("after", after)
	}
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	var result MarketingEventListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByObjectID retrieves a marketing event by its HubSpot object ID.
func (s *EventsService) GetByObjectID(ctx context.Context, objectID string) (*MarketingEventReadResponseV2, error) {
	path := fmt.Sprintf("%s/events/%s", eventsBasePath, objectID)
	var result MarketingEventReadResponseV2
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDetails retrieves a marketing event by external event ID and account ID.
func (s *EventsService) GetDetails(ctx context.Context, externalEventID, externalAccountID string) (*MarketingEventReadResponse, error) {
	path := fmt.Sprintf("%s/events/external/%s", eventsBasePath, externalEventID)
	q := url.Values{}
	q.Set("externalAccountId", externalAccountID)
	var result MarketingEventReadResponse
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a marketing event by external IDs.
func (s *EventsService) Update(ctx context.Context, externalEventID, externalAccountID string, input *MarketingEventUpdateRequest) (*MarketingEventPublicDefaultResponse, error) {
	path := fmt.Sprintf("%s/events/external/%s", eventsBasePath, externalEventID)
	q := url.Values{}
	q.Set("externalAccountId", externalAccountID)
	// Patch doesn't support query params natively in our interface, so we append them to the path.
	if externalAccountID != "" {
		path += "?externalAccountId=" + url.QueryEscape(externalAccountID)
	}
	var result MarketingEventPublicDefaultResponse
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateByObjectID updates a marketing event by its HubSpot object ID (V2).
func (s *EventsService) UpdateByObjectID(ctx context.Context, objectID string, input *MarketingEventV2UpdateRequest) (*MarketingEventV2DefaultResponse, error) {
	path := fmt.Sprintf("%s/events/%s", eventsBasePath, objectID)
	var result MarketingEventV2DefaultResponse
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Upsert creates or updates a marketing event by external event ID.
func (s *EventsService) Upsert(ctx context.Context, externalEventID string, input *MarketingEventCreateRequest) (*MarketingEventPublicDefaultResponse, error) {
	path := fmt.Sprintf("%s/events/upsert/%s", eventsBasePath, externalEventID)
	var result MarketingEventPublicDefaultResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive archives a marketing event by external IDs.
func (s *EventsService) Archive(ctx context.Context, externalEventID, externalAccountID string) error {
	path := fmt.Sprintf("%s/events/external/%s?externalAccountId=%s", eventsBasePath, externalEventID, url.QueryEscape(externalAccountID))
	return s.requester.Delete(ctx, path)
}

// ArchiveByObjectID archives a marketing event by its HubSpot object ID.
func (s *EventsService) ArchiveByObjectID(ctx context.Context, objectID string) error {
	path := fmt.Sprintf("%s/events/%s", eventsBasePath, objectID)
	return s.requester.Delete(ctx, path)
}

// --- Batch API ---

// BatchUpsert creates or updates multiple marketing events.
func (s *EventsService) BatchUpsert(ctx context.Context, input *BatchInputMarketingEventCreateRequests) (*BatchResponseMarketingEventPublicDefault, error) {
	path := eventsBasePath + "/events/batch/upsert"
	var result BatchResponseMarketingEventPublicDefault
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchArchive archives multiple marketing events by external identifiers.
func (s *EventsService) BatchArchive(ctx context.Context, input *BatchInputMarketingEventExternalUniqueIdentifiers) error {
	path := eventsBasePath + "/events/batch/archive"
	return s.requester.Post(ctx, path, input, nil)
}

// --- Change Property API ---

// Cancel marks a marketing event as cancelled.
func (s *EventsService) Cancel(ctx context.Context, externalEventID, externalAccountID string) (*MarketingEventDefaultResponse, error) {
	path := fmt.Sprintf("%s/events/external/%s/cancel?externalAccountId=%s", eventsBasePath, externalEventID, url.QueryEscape(externalAccountID))
	var result MarketingEventDefaultResponse
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Complete marks a marketing event as completed.
func (s *EventsService) Complete(ctx context.Context, externalEventID, externalAccountID string, input *MarketingEventCompleteRequest) (*MarketingEventDefaultResponse, error) {
	path := fmt.Sprintf("%s/events/external/%s/complete?externalAccountId=%s", eventsBasePath, externalEventID, url.QueryEscape(externalAccountID))
	var result MarketingEventDefaultResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Identifiers API ---

// SearchEvents searches for marketing events by query string.
func (s *EventsService) SearchEvents(ctx context.Context, query string) (*EventSearchResult, error) {
	path := eventsBasePath + "/events/search"
	q := url.Values{}
	q.Set("q", query)
	var result EventSearchResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SearchPortalEvents searches portal events by external event ID.
func (s *EventsService) SearchPortalEvents(ctx context.Context, externalEventID string) (*MarketingEventIdentifiersResult, error) {
	path := fmt.Sprintf("%s/events/external/%s/identifiers", eventsBasePath, externalEventID)
	var result MarketingEventIdentifiersResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- List Associations API ---

// AssociateListByExternalIDs associates a list with a marketing event using external IDs.
func (s *EventsService) AssociateListByExternalIDs(ctx context.Context, externalAccountID, externalEventID, listID string) error {
	path := fmt.Sprintf("%s/events/external/%s/associations/%s?externalAccountId=%s",
		eventsBasePath, externalEventID, listID, url.QueryEscape(externalAccountID))
	return s.requester.Post(ctx, path, nil, nil)
}

// AssociateListByEventID associates a list with a marketing event by event ID.
func (s *EventsService) AssociateListByEventID(ctx context.Context, marketingEventID, listID string) error {
	path := fmt.Sprintf("%s/events/%s/associations/%s", eventsBasePath, marketingEventID, listID)
	return s.requester.Post(ctx, path, nil, nil)
}

// DisassociateListByExternalIDs removes a list association using external IDs.
func (s *EventsService) DisassociateListByExternalIDs(ctx context.Context, externalAccountID, externalEventID, listID string) error {
	path := fmt.Sprintf("%s/events/external/%s/associations/%s?externalAccountId=%s",
		eventsBasePath, externalEventID, listID, url.QueryEscape(externalAccountID))
	return s.requester.Delete(ctx, path)
}

// DisassociateListByEventID removes a list association by event ID.
func (s *EventsService) DisassociateListByEventID(ctx context.Context, marketingEventID, listID string) error {
	path := fmt.Sprintf("%s/events/%s/associations/%s", eventsBasePath, marketingEventID, listID)
	return s.requester.Delete(ctx, path)
}

// GetListsByExternalIDs retrieves all lists associated with a marketing event by external IDs.
func (s *EventsService) GetListsByExternalIDs(ctx context.Context, externalAccountID, externalEventID string) (*PublicListResult, error) {
	path := fmt.Sprintf("%s/events/external/%s/associations", eventsBasePath, externalEventID)
	q := url.Values{}
	q.Set("externalAccountId", externalAccountID)
	var result PublicListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetListsByEventID retrieves all lists associated with a marketing event by event ID.
func (s *EventsService) GetListsByEventID(ctx context.Context, marketingEventID string) (*PublicListResult, error) {
	path := fmt.Sprintf("%s/events/%s/associations", eventsBasePath, marketingEventID)
	var result PublicListResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Retrieve Participant State API ---

// GetParticipationsBreakdownByContactID retrieves participation breakdown for a contact.
func (s *EventsService) GetParticipationsBreakdownByContactID(ctx context.Context, contactIdentifier string, opts *ParticipationBreakdownOptions) (*ParticipationBreakdownResult, error) {
	path := fmt.Sprintf("%s/participations/contacts/%s/breakdown", eventsBasePath, contactIdentifier)
	q := url.Values{}
	if opts != nil {
		if opts.State != "" {
			q.Set("state", opts.State)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
	}
	var result ParticipationBreakdownResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetParticipationsBreakdownByExternalEventID retrieves participation breakdown by external event.
func (s *EventsService) GetParticipationsBreakdownByExternalEventID(ctx context.Context, externalAccountID, externalEventID string, opts *ParticipationBreakdownOptions) (*ParticipationBreakdownResult, error) {
	path := fmt.Sprintf("%s/events/external/%s/participations/breakdown", eventsBasePath, externalEventID)
	q := url.Values{}
	q.Set("externalAccountId", externalAccountID)
	if opts != nil {
		if opts.ContactIdentifier != "" {
			q.Set("contactIdentifier", opts.ContactIdentifier)
		}
		if opts.State != "" {
			q.Set("state", opts.State)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
	}
	var result ParticipationBreakdownResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetParticipationsBreakdownByMarketingEventID retrieves participation breakdown by marketing event ID.
func (s *EventsService) GetParticipationsBreakdownByMarketingEventID(ctx context.Context, marketingEventID int, opts *ParticipationBreakdownOptions) (*ParticipationBreakdownResult, error) {
	path := fmt.Sprintf("%s/events/%d/participations/breakdown", eventsBasePath, marketingEventID)
	q := url.Values{}
	if opts != nil {
		if opts.ContactIdentifier != "" {
			q.Set("contactIdentifier", opts.ContactIdentifier)
		}
		if opts.State != "" {
			q.Set("state", opts.State)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
	}
	var result ParticipationBreakdownResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetParticipationsCountersByExternalEventID retrieves attendance counters by external event.
func (s *EventsService) GetParticipationsCountersByExternalEventID(ctx context.Context, externalAccountID, externalEventID string) (*AttendanceCounters, error) {
	path := fmt.Sprintf("%s/events/external/%s/participations/counters", eventsBasePath, externalEventID)
	q := url.Values{}
	q.Set("externalAccountId", externalAccountID)
	var result AttendanceCounters
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetParticipationsCountersByMarketingEventID retrieves attendance counters by marketing event ID.
func (s *EventsService) GetParticipationsCountersByMarketingEventID(ctx context.Context, marketingEventID int) (*AttendanceCounters, error) {
	path := fmt.Sprintf("%s/events/%d/participations/counters", eventsBasePath, marketingEventID)
	var result AttendanceCounters
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Settings API ---

// GetSettings retrieves marketing event settings for an app.
func (s *EventsService) GetSettings(ctx context.Context, appID int) (*EventDetailSettings, error) {
	path := fmt.Sprintf("%s/events/settings/%d", eventsBasePath, appID)
	var result EventDetailSettings
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateSettings updates marketing event settings for an app.
func (s *EventsService) UpdateSettings(ctx context.Context, appID int, input *EventDetailSettingsURL) (*EventDetailSettings, error) {
	path := fmt.Sprintf("%s/events/settings/%d", eventsBasePath, appID)
	var result EventDetailSettings
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Subscriber State Changes API ---

// UpsertByContactEmail upserts subscriber state changes by contact email.
func (s *EventsService) UpsertByContactEmail(ctx context.Context, externalEventID, subscriberState, externalAccountID string, input *BatchInputEmailSubscribers) error {
	path := fmt.Sprintf("%s/events/external/%s/%s/upsert-by-email?externalAccountId=%s",
		eventsBasePath, externalEventID, subscriberState, url.QueryEscape(externalAccountID))
	return s.requester.Post(ctx, path, input, nil)
}

// UpsertByContactID upserts subscriber state changes by contact ID.
func (s *EventsService) UpsertByContactID(ctx context.Context, externalEventID, subscriberState, externalAccountID string, input *BatchInputSubscribers) error {
	path := fmt.Sprintf("%s/events/external/%s/%s/upsert?externalAccountId=%s",
		eventsBasePath, externalEventID, subscriberState, url.QueryEscape(externalAccountID))
	return s.requester.Post(ctx, path, input, nil)
}

// --- Add Event Attendees API ---

// RecordByContactEmails records attendance by contact emails (external event).
func (s *EventsService) RecordByContactEmails(ctx context.Context, externalEventID, subscriberState string, input *BatchInputEmailSubscribers, externalAccountID string) (*BatchResponseSubscriberEmail, error) {
	path := fmt.Sprintf("%s/events/external/%s/%s/record-by-email", eventsBasePath, externalEventID, subscriberState)
	if externalAccountID != "" {
		path += "?externalAccountId=" + url.QueryEscape(externalAccountID)
	}
	var result BatchResponseSubscriberEmail
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RecordByContactIDs records attendance by contact IDs (external event).
func (s *EventsService) RecordByContactIDs(ctx context.Context, externalEventID, subscriberState string, input *BatchInputSubscribers, externalAccountID string) (*BatchResponseSubscriberVid, error) {
	path := fmt.Sprintf("%s/events/external/%s/%s/record-by-id", eventsBasePath, externalEventID, subscriberState)
	if externalAccountID != "" {
		path += "?externalAccountId=" + url.QueryEscape(externalAccountID)
	}
	var result BatchResponseSubscriberVid
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RecordByContactID records attendance by contact ID (object-level).
func (s *EventsService) RecordByContactID(ctx context.Context, objectID, subscriberState string, input *BatchInputSubscribers) (*BatchResponseSubscriberVid, error) {
	path := fmt.Sprintf("%s/events/%s/%s/record-by-id", eventsBasePath, objectID, subscriberState)
	var result BatchResponseSubscriberVid
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RecordByEmail records attendance by email (object-level).
func (s *EventsService) RecordByEmail(ctx context.Context, objectID, subscriberState string, input *BatchInputEmailSubscribers) (*BatchResponseSubscriberEmail, error) {
	path := fmt.Sprintf("%s/events/%s/%s/record-by-email", eventsBasePath, objectID, subscriberState)
	var result BatchResponseSubscriberEmail
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
