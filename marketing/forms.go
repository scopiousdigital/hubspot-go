package marketing

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const formsBasePath = "/marketing/v3/forms"

// FormsService handles operations on HubSpot marketing forms.
type FormsService struct {
	requester api.Requester
}

// Create creates a new marketing form.
func (s *FormsService) Create(ctx context.Context, input *FormCreateRequest) (*FormDefinition, error) {
	var result FormDefinition
	if err := s.requester.Post(ctx, formsBasePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a single form by its ID.
func (s *FormsService) GetByID(ctx context.Context, formID string, archived bool) (*FormDefinition, error) {
	path := fmt.Sprintf("%s/%s", formsBasePath, formID)
	q := url.Values{}
	if archived {
		q.Set("archived", "true")
	}
	var result FormDefinition
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPage retrieves a page of form definitions.
func (s *FormsService) GetPage(ctx context.Context, opts *FormListOptions) (*FormListResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.Archived {
			q.Set("archived", "true")
		}
		if len(opts.FormTypes) > 0 {
			q.Set("formTypes", strings.Join(opts.FormTypes, ","))
		}
	}
	var result FormListResult
	if err := s.requester.Get(ctx, formsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update patches an existing form.
func (s *FormsService) Update(ctx context.Context, formID string, input *FormUpdateRequest) (*FormDefinition, error) {
	path := fmt.Sprintf("%s/%s", formsBasePath, formID)
	var result FormDefinition
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Replace fully replaces a form definition (PUT).
func (s *FormsService) Replace(ctx context.Context, formID string, input *FormReplaceRequest) (*FormDefinition, error) {
	path := fmt.Sprintf("%s/%s", formsBasePath, formID)
	var result FormDefinition
	if err := s.requester.Put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive soft-deletes a form.
func (s *FormsService) Archive(ctx context.Context, formID string) error {
	path := fmt.Sprintf("%s/%s", formsBasePath, formID)
	return s.requester.Delete(ctx, path)
}
