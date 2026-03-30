package cms

import (
	"context"
	"fmt"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const urlRedirectsBasePath = "/cms/v3/url-redirects"

// UrlRedirectsService handles URL redirect (URL mapping) operations.
type UrlRedirectsService struct {
	requester api.Requester
}

// Create creates a new URL redirect.
func (s *UrlRedirectsService) Create(ctx context.Context, input *UrlMappingCreateRequest) (*UrlMapping, error) {
	var result UrlMapping
	if err := s.requester.Post(ctx, urlRedirectsBasePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a URL redirect by its ID.
func (s *UrlRedirectsService) GetByID(ctx context.Context, id string) (*UrlMapping, error) {
	path := fmt.Sprintf("%s/%s", urlRedirectsBasePath, id)
	var result UrlMapping
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing URL redirect.
func (s *UrlRedirectsService) Update(ctx context.Context, id string, input *UrlMapping) (*UrlMapping, error) {
	path := fmt.Sprintf("%s/%s", urlRedirectsBasePath, id)
	var result UrlMapping
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive deletes a URL redirect.
func (s *UrlRedirectsService) Archive(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s", urlRedirectsBasePath, id)
	return s.requester.Delete(ctx, path)
}

// List retrieves a page of URL redirects.
func (s *UrlRedirectsService) List(ctx context.Context, opts *ListOptions) (*UrlMappingListResult, error) {
	q := buildListQuery(opts)
	var result UrlMappingListResult
	if err := s.requester.Get(ctx, urlRedirectsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
