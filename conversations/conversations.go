package conversations

import (
	"context"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const visitorIdentificationBasePath = "/conversations/v3/visitor-identification/tokens/create"

// Service provides access to HubSpot Conversations APIs.
type Service struct {
	requester api.Requester

	// VisitorIdentification provides methods for generating visitor identification tokens.
	VisitorIdentification *VisitorIdentificationService
}

// NewService creates a new Conversations service.
func NewService(r api.Requester) *Service {
	return &Service{
		requester:             r,
		VisitorIdentification: &VisitorIdentificationService{requester: r},
	}
}

// --- VisitorIdentificationService ---

// VisitorIdentificationService provides methods for the Visitor Identification API.
type VisitorIdentificationService struct {
	requester api.Requester
}

// GenerateToken generates an identification token for a visitor.
func (s *VisitorIdentificationService) GenerateToken(ctx context.Context, input *IdentificationTokenGenerationRequest) (*IdentificationTokenResponse, error) {
	var result IdentificationTokenResponse
	if err := s.requester.Post(ctx, visitorIdentificationBasePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
