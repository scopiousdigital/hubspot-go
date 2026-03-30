package crm

import (
	"context"
	"fmt"
	"strconv"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const timelineBasePath = "/crm/v3/timeline"

// TimelineService handles operations on HubSpot timeline events, templates, and tokens.
type TimelineService struct {
	requester api.Requester
}

// =============================================================================
// Events API
// =============================================================================

// CreateEvent creates a single timeline event.
func (s *TimelineService) CreateEvent(ctx context.Context, input *TimelineEvent) (*TimelineEventResponse, error) {
	path := timelineBasePath + "/events"
	var result TimelineEventResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateEventBatch creates multiple timeline events in batch.
func (s *TimelineService) CreateEventBatch(ctx context.Context, input *BatchTimelineEventInput) (*BatchTimelineEventResult, error) {
	path := timelineBasePath + "/events/batch/create"
	var result BatchTimelineEventResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEventByID retrieves a timeline event by its template and event ID.
func (s *TimelineService) GetEventByID(ctx context.Context, eventTemplateID, eventID string) (*TimelineEventResponse, error) {
	path := fmt.Sprintf("%s/events/%s/%s", timelineBasePath, eventTemplateID, eventID)
	var result TimelineEventResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetEventDetailByID retrieves the detail of a timeline event.
func (s *TimelineService) GetEventDetailByID(ctx context.Context, eventTemplateID, eventID string) (*EventDetail, error) {
	path := fmt.Sprintf("%s/events/%s/%s/detail", timelineBasePath, eventTemplateID, eventID)
	var result EventDetail
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============================================================================
// Templates API
// =============================================================================

// CreateTemplate creates a new timeline event template.
func (s *TimelineService) CreateTemplate(ctx context.Context, appID int, input *TimelineEventTemplateCreateRequest) (*TimelineEventTemplate, error) {
	path := fmt.Sprintf("%s/%d/event-templates", timelineBasePath, appID)
	var result TimelineEventTemplate
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAllTemplates retrieves all timeline event templates for an app.
func (s *TimelineService) GetAllTemplates(ctx context.Context, appID int) (*TimelineEventTemplateListResult, error) {
	path := fmt.Sprintf("%s/%d/event-templates", timelineBasePath, appID)
	var result TimelineEventTemplateListResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTemplateByID retrieves a timeline event template by its ID.
func (s *TimelineService) GetTemplateByID(ctx context.Context, eventTemplateID string, appID int) (*TimelineEventTemplate, error) {
	path := fmt.Sprintf("%s/%s/event-templates/%s", timelineBasePath, strconv.Itoa(appID), eventTemplateID)
	var result TimelineEventTemplate
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateTemplate updates a timeline event template.
func (s *TimelineService) UpdateTemplate(ctx context.Context, eventTemplateID string, appID int, input *TimelineEventTemplateUpdateRequest) (*TimelineEventTemplate, error) {
	path := fmt.Sprintf("%s/%d/event-templates/%s", timelineBasePath, appID, eventTemplateID)
	var result TimelineEventTemplate
	if err := s.requester.Put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchiveTemplate deletes a timeline event template.
func (s *TimelineService) ArchiveTemplate(ctx context.Context, eventTemplateID string, appID int) error {
	path := fmt.Sprintf("%s/%d/event-templates/%s", timelineBasePath, appID, eventTemplateID)
	return s.requester.Delete(ctx, path)
}

// =============================================================================
// Tokens API
// =============================================================================

// CreateToken creates a token on a timeline event template.
func (s *TimelineService) CreateToken(ctx context.Context, eventTemplateID string, appID int, input *TimelineEventTemplateToken) (*TimelineEventTemplateToken, error) {
	path := fmt.Sprintf("%s/%d/event-templates/%s/tokens", timelineBasePath, appID, eventTemplateID)
	var result TimelineEventTemplateToken
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateToken updates a token on a timeline event template.
func (s *TimelineService) UpdateToken(ctx context.Context, eventTemplateID, tokenName string, appID int, input *TimelineEventTemplateTokenUpdateRequest) (*TimelineEventTemplateToken, error) {
	path := fmt.Sprintf("%s/%d/event-templates/%s/tokens/%s", timelineBasePath, appID, eventTemplateID, tokenName)
	var result TimelineEventTemplateToken
	if err := s.requester.Put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchiveToken deletes a token from a timeline event template.
func (s *TimelineService) ArchiveToken(ctx context.Context, eventTemplateID, tokenName string, appID int) error {
	path := fmt.Sprintf("%s/%d/event-templates/%s/tokens/%s", timelineBasePath, appID, eventTemplateID, tokenName)
	return s.requester.Delete(ctx, path)
}
