package communicationpreferences

import (
	"context"
	"net/url"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const (
	definitionsBasePath = "/communication-preferences/v3/definitions"
	statusBasePath      = "/communication-preferences/v3/status"
	subscribePath       = "/communication-preferences/v3/subscribe"
	unsubscribePath     = "/communication-preferences/v3/unsubscribe"
)

// Service provides access to HubSpot Communication Preferences APIs.
type Service struct {
	requester api.Requester

	// Definitions provides methods for retrieving subscription definitions.
	Definitions *DefinitionsService
	// Status provides methods for managing subscription statuses.
	Status *StatusService
}

// NewService creates a new Communication Preferences service.
func NewService(r api.Requester) *Service {
	return &Service{
		requester:   r,
		Definitions: &DefinitionsService{requester: r},
		Status:      &StatusService{requester: r},
	}
}

// --- DefinitionsService ---

// DefinitionsService provides methods for retrieving communication subscription definitions.
type DefinitionsService struct {
	requester api.Requester
}

// GetAll retrieves all subscription definitions.
func (s *DefinitionsService) GetAll(ctx context.Context) (*SubscriptionDefinitionsResponse, error) {
	var result SubscriptionDefinitionsResponse
	if err := s.requester.Get(ctx, definitionsBasePath, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- StatusService ---

// StatusService provides methods for managing email subscription statuses.
type StatusService struct {
	requester api.Requester
}

// GetEmailStatus retrieves the subscription statuses for the given email address.
func (s *StatusService) GetEmailStatus(ctx context.Context, emailAddress string) (*PublicSubscriptionStatusesResponse, error) {
	q := url.Values{}
	q.Set("emailAddress", emailAddress)
	var result PublicSubscriptionStatusesResponse
	if err := s.requester.Get(ctx, statusBasePath+"/email/"+url.PathEscape(emailAddress), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Subscribe subscribes a contact to a subscription type.
func (s *StatusService) Subscribe(ctx context.Context, input *PublicUpdateSubscriptionStatusRequest) (*PublicSubscriptionStatus, error) {
	var result PublicSubscriptionStatus
	if err := s.requester.Post(ctx, subscribePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unsubscribe unsubscribes a contact from a subscription type.
func (s *StatusService) Unsubscribe(ctx context.Context, input *PublicUpdateSubscriptionStatusRequest) (*PublicSubscriptionStatus, error) {
	var result PublicSubscriptionStatus
	if err := s.requester.Post(ctx, unsubscribePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
