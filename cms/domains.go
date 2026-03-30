package cms

import (
	"context"
	"fmt"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const domainsBasePath = "/cms/v3/domains"

// DomainsService handles HubSpot domain operations.
type DomainsService struct {
	requester api.Requester
}

// GetByID retrieves a domain by its ID.
func (s *DomainsService) GetByID(ctx context.Context, domainID string) (*Domain, error) {
	path := fmt.Sprintf("%s/%s", domainsBasePath, domainID)
	var result Domain
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List retrieves a page of domains.
func (s *DomainsService) List(ctx context.Context, opts *ListOptions) (*DomainListResult, error) {
	q := buildListQuery(opts)
	var result DomainListResult
	if err := s.requester.Get(ctx, domainsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
