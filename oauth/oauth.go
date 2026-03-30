package oauth

import (
	"context"
	"fmt"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const oauthV1Path = "/oauth/v1"

// Service provides access to the HubSpot OAuth APIs.
type Service struct {
	AccessTokens  *AccessTokensService
	RefreshTokens *RefreshTokensService
	Tokens        *TokensService
}

// NewService creates a new OAuth service. Called by the root hubspot package.
func NewService(r api.Requester) *Service {
	return &Service{
		AccessTokens:  &AccessTokensService{requester: r},
		RefreshTokens: &RefreshTokensService{requester: r},
		Tokens:        &TokensService{requester: r},
	}
}

// --- AccessTokensService ---

// AccessTokensService provides methods for retrieving access token metadata.
type AccessTokensService struct {
	requester api.Requester
}

// Get retrieves information about an OAuth access token.
func (s *AccessTokensService) Get(ctx context.Context, token string) (*AccessTokenInfo, error) {
	path := fmt.Sprintf("%s/access-tokens/%s", oauthV1Path, token)
	var result AccessTokenInfo
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- RefreshTokensService ---

// RefreshTokensService provides methods for retrieving and archiving refresh tokens.
type RefreshTokensService struct {
	requester api.Requester
}

// Get retrieves information about an OAuth refresh token.
func (s *RefreshTokensService) Get(ctx context.Context, token string) (*RefreshTokenInfo, error) {
	path := fmt.Sprintf("%s/refresh-tokens/%s", oauthV1Path, token)
	var result RefreshTokenInfo
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive deletes (archives) a refresh token.
func (s *RefreshTokensService) Archive(ctx context.Context, token string) error {
	path := fmt.Sprintf("%s/refresh-tokens/%s", oauthV1Path, token)
	return s.requester.Delete(ctx, path)
}

// --- TokensService ---

// TokensService handles OAuth token exchange (authorization code -> tokens,
// or refresh token -> new access token).
type TokensService struct {
	requester api.Requester
}

// Create exchanges an authorization code or refresh token for OAuth tokens.
func (s *TokensService) Create(ctx context.Context, input *TokenCreateRequest) (*TokenResponse, error) {
	path := oauthV1Path + "/token"
	var result TokenResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
