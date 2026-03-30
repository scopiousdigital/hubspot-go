package crm

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const ownersBasePath = "/crm/v3/owners"

// OwnersService handles operations on HubSpot owners.
type OwnersService struct {
	requester api.Requester
}

// GetByID retrieves an owner by their ID.
func (s *OwnersService) GetByID(ctx context.Context, ownerID string, opts *OwnerGetByIDOptions) (*PublicOwner, error) {
	path := fmt.Sprintf("%s/%s", ownersBasePath, ownerID)
	q := url.Values{}
	if opts != nil {
		if opts.IDProperty != "" {
			q.Set("idProperty", opts.IDProperty)
		}
		if opts.Archived {
			q.Set("archived", "true")
		}
	}
	var result PublicOwner
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List retrieves a page of owners.
func (s *OwnersService) List(ctx context.Context, opts *OwnerListOptions) (*OwnerListResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Email != "" {
			q.Set("email", opts.Email)
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.Archived {
			q.Set("archived", "true")
		}
	}
	var result OwnerListResult
	if err := s.requester.Get(ctx, ownersBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
