package webhooks

import (
	"context"
	"fmt"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const webhooksV3Path = "/webhooks/v3"

// Service provides access to the HubSpot Webhooks APIs.
type Service struct {
	Settings      *SettingsService
	Subscriptions *SubscriptionsService
}

// NewService creates a new Webhooks service. Called by the root hubspot package.
func NewService(r api.Requester) *Service {
	return &Service{
		Settings:      &SettingsService{requester: r},
		Subscriptions: &SubscriptionsService{requester: r},
	}
}

// --- SettingsService ---

// SettingsService manages webhook settings for an app.
type SettingsService struct {
	requester api.Requester
}

// GetAll retrieves the current webhook settings for an app.
func (s *SettingsService) GetAll(ctx context.Context, appID int64) (*SettingsResponse, error) {
	path := fmt.Sprintf("%s/%d/settings", webhooksV3Path, appID)
	var result SettingsResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Configure updates the webhook settings for an app.
func (s *SettingsService) Configure(ctx context.Context, appID int64, input *SettingsChangeRequest) (*SettingsResponse, error) {
	path := fmt.Sprintf("%s/%d/settings", webhooksV3Path, appID)
	var result SettingsResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Clear removes the webhook settings for an app.
func (s *SettingsService) Clear(ctx context.Context, appID int64) error {
	path := fmt.Sprintf("%s/%d/settings", webhooksV3Path, appID)
	return s.requester.Delete(ctx, path)
}

// --- SubscriptionsService ---

// SubscriptionsService manages webhook event subscriptions for an app.
type SubscriptionsService struct {
	requester api.Requester
}

// GetAll retrieves all webhook subscriptions for an app.
func (s *SubscriptionsService) GetAll(ctx context.Context, appID int64) (*SubscriptionListResponse, error) {
	path := fmt.Sprintf("%s/%d/subscriptions", webhooksV3Path, appID)
	var result SubscriptionListResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new webhook subscription for an app.
func (s *SubscriptionsService) Create(ctx context.Context, appID int64, input *SubscriptionCreateRequest) (*SubscriptionResponse, error) {
	path := fmt.Sprintf("%s/%d/subscriptions", webhooksV3Path, appID)
	var result SubscriptionResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a single webhook subscription by ID.
func (s *SubscriptionsService) GetByID(ctx context.Context, subscriptionID int64, appID int64) (*SubscriptionResponse, error) {
	path := fmt.Sprintf("%s/%d/subscriptions/%d", webhooksV3Path, appID, subscriptionID)
	var result SubscriptionResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a webhook subscription (currently only the active flag).
func (s *SubscriptionsService) Update(ctx context.Context, subscriptionID int64, appID int64, input *SubscriptionPatchRequest) (*SubscriptionResponse, error) {
	path := fmt.Sprintf("%s/%d/subscriptions/%d", webhooksV3Path, appID, subscriptionID)
	var result SubscriptionResponse
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive deletes a webhook subscription.
func (s *SubscriptionsService) Archive(ctx context.Context, subscriptionID int64, appID int64) error {
	path := fmt.Sprintf("%s/%d/subscriptions/%d", webhooksV3Path, appID, subscriptionID)
	return s.requester.Delete(ctx, path)
}

// BatchUpdate updates multiple webhook subscriptions in a single request.
func (s *SubscriptionsService) BatchUpdate(ctx context.Context, appID int64, input *BatchInputSubscriptionBatchUpdateRequest) (*BatchResponseSubscriptionResponse, error) {
	path := fmt.Sprintf("%s/%d/subscriptions/batch/update", webhooksV3Path, appID)
	var result BatchResponseSubscriptionResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
