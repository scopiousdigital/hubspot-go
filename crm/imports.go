package crm

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const importsBasePath = "/crm/v3/imports"

// ImportsService handles operations on HubSpot CRM imports.
type ImportsService struct {
	requester api.Requester
}

// Create starts a new import. Note: the HubSpot API expects a multipart/form-data
// request with a file upload. This method accepts a JSON import request for cases
// where the import request is submitted via JSON. For file uploads, callers should
// use the raw HTTP client directly.
func (s *ImportsService) Create(ctx context.Context, importRequest any) (*PublicImportResponse, error) {
	var result PublicImportResponse
	if err := s.requester.Post(ctx, importsBasePath, importRequest, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List retrieves a page of imports.
func (s *ImportsService) List(ctx context.Context, opts *ImportListOptions) (*ImportListResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if opts.Before != "" {
			q.Set("before", opts.Before)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
	}
	var result ImportListResult
	if err := s.requester.Get(ctx, importsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves an import by its ID.
func (s *ImportsService) GetByID(ctx context.Context, importID string) (*PublicImportResponse, error) {
	path := fmt.Sprintf("%s/%s", importsBasePath, importID)
	var result PublicImportResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Cancel cancels an active import.
func (s *ImportsService) Cancel(ctx context.Context, importID string) (*ActionResponse, error) {
	path := fmt.Sprintf("%s/%s/cancel", importsBasePath, importID)
	var result ActionResponse
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetErrors retrieves errors for an import.
func (s *ImportsService) GetErrors(ctx context.Context, importID string, opts *ImportGetErrorsOptions) (*ImportErrorListResult, error) {
	path := fmt.Sprintf("%s/%s/errors", importsBasePath, importID)
	q := url.Values{}
	if opts != nil {
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.IncludeErrorMessage {
			q.Set("includeErrorMessage", "true")
		}
		if opts.IncludeRowData {
			q.Set("includeRowData", "true")
		}
	}
	var result ImportErrorListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
