package crm

import (
	"context"
	"fmt"
	"strconv"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

// ExtensionsService handles operations on HubSpot CRM extensions (calling, cards, video conferencing).
type ExtensionsService struct {
	requester api.Requester
}

// =============================================================================
// Calling Settings API
// =============================================================================

const callingSettingsBasePath = "/crm/v3/extensions/calling"

// CreateCallingSettings creates calling app settings.
func (s *ExtensionsService) CreateCallingSettings(ctx context.Context, appID int, input *CallingSettingsRequest) (*CallingSettingsResponse, error) {
	path := fmt.Sprintf("%s/%d/settings", callingSettingsBasePath, appID)
	var result CallingSettingsResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCallingSettings retrieves calling app settings.
func (s *ExtensionsService) GetCallingSettings(ctx context.Context, appID int) (*CallingSettingsResponse, error) {
	path := fmt.Sprintf("%s/%d/settings", callingSettingsBasePath, appID)
	var result CallingSettingsResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCallingSettings updates calling app settings.
func (s *ExtensionsService) UpdateCallingSettings(ctx context.Context, appID int, input *CallingSettingsPatchRequest) (*CallingSettingsResponse, error) {
	path := fmt.Sprintf("%s/%d/settings", callingSettingsBasePath, appID)
	var result CallingSettingsResponse
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchiveCallingSettings deletes calling app settings.
func (s *ExtensionsService) ArchiveCallingSettings(ctx context.Context, appID int) error {
	path := fmt.Sprintf("%s/%d/settings", callingSettingsBasePath, appID)
	return s.requester.Delete(ctx, path)
}

// --- Recording Settings ---

// GetRecordingSettings retrieves recording URL format settings.
func (s *ExtensionsService) GetRecordingSettings(ctx context.Context, appID int) (*RecordingSettingsResponse, error) {
	path := fmt.Sprintf("%s/%d/settings/recording", callingSettingsBasePath, appID)
	var result RecordingSettingsResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RegisterRecordingSettings registers a recording URL format.
func (s *ExtensionsService) RegisterRecordingSettings(ctx context.Context, appID int, input *RecordingSettingsRequest) (*RecordingSettingsResponse, error) {
	path := fmt.Sprintf("%s/%d/settings/recording", callingSettingsBasePath, appID)
	var result RecordingSettingsResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateRecordingSettings updates recording URL format settings.
func (s *ExtensionsService) UpdateRecordingSettings(ctx context.Context, appID int, input *RecordingSettingsPatchRequest) (*RecordingSettingsResponse, error) {
	path := fmt.Sprintf("%s/%d/settings/recording", callingSettingsBasePath, appID)
	var result RecordingSettingsResponse
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// MarkRecordingAsReady marks a recording as ready for playback.
func (s *ExtensionsService) MarkRecordingAsReady(ctx context.Context, input *MarkRecordingAsReadyRequest) error {
	path := callingSettingsBasePath + "/recordings/ready"
	return s.requester.Post(ctx, path, input, nil)
}

// =============================================================================
// Cards API
// =============================================================================

const cardsBasePath = "/crm/v3/extensions/cards"

// CreateCard creates a new CRM extension card.
func (s *ExtensionsService) CreateCard(ctx context.Context, appID int, input *CardCreateRequest) (*PublicCardResponse, error) {
	path := fmt.Sprintf("%s/%d", cardsBasePath, appID)
	var result PublicCardResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAllCards retrieves all CRM extension cards for an app.
func (s *ExtensionsService) GetAllCards(ctx context.Context, appID int) (*PublicCardListResponse, error) {
	path := fmt.Sprintf("%s/%d", cardsBasePath, appID)
	var result PublicCardListResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCardByID retrieves a CRM extension card by its ID.
func (s *ExtensionsService) GetCardByID(ctx context.Context, cardID string, appID int) (*PublicCardResponse, error) {
	path := fmt.Sprintf("%s/%d/%s", cardsBasePath, appID, cardID)
	var result PublicCardResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCard updates a CRM extension card.
func (s *ExtensionsService) UpdateCard(ctx context.Context, cardID string, appID int, input *CardPatchRequest) (*PublicCardResponse, error) {
	path := fmt.Sprintf("%s/%d/%s", cardsBasePath, appID, cardID)
	var result PublicCardResponse
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchiveCard deletes a CRM extension card.
func (s *ExtensionsService) ArchiveCard(ctx context.Context, cardID string, appID int) error {
	path := fmt.Sprintf("%s/%s/%s", cardsBasePath, strconv.Itoa(appID), cardID)
	return s.requester.Delete(ctx, path)
}

// =============================================================================
// Video Conferencing API
// =============================================================================

const videoConferencingBasePath = "/crm/v3/extensions/videoconferencing/settings"

// GetVideoConferencingSettings retrieves video conferencing settings.
func (s *ExtensionsService) GetVideoConferencingSettings(ctx context.Context, appID int) (*VideoConferencingExternalSettings, error) {
	path := fmt.Sprintf("%s/%d", videoConferencingBasePath, appID)
	var result VideoConferencingExternalSettings
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReplaceVideoConferencingSettings replaces video conferencing settings (PUT).
func (s *ExtensionsService) ReplaceVideoConferencingSettings(ctx context.Context, appID int, input *VideoConferencingExternalSettings) (*VideoConferencingExternalSettings, error) {
	path := fmt.Sprintf("%s/%d", videoConferencingBasePath, appID)
	var result VideoConferencingExternalSettings
	if err := s.requester.Put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchiveVideoConferencingSettings deletes video conferencing settings.
func (s *ExtensionsService) ArchiveVideoConferencingSettings(ctx context.Context, appID int) error {
	path := fmt.Sprintf("%s/%d", videoConferencingBasePath, appID)
	return s.requester.Delete(ctx, path)
}
