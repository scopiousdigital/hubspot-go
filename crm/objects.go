package crm

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const crmV3ObjectsPath = "/crm/v3/objects"

// ObjectService handles CRUD, batch, and search operations for a CRM object type.
// All standard CRM objects (contacts, companies, deals, tickets, etc.) use the same
// API shape at /crm/v3/objects/{objectType}.
type ObjectService struct {
	requester  api.Requester
	objectType string
	basePath   string
}

func newObjectService(r api.Requester, objectType string) *ObjectService {
	return &ObjectService{
		requester:  r,
		objectType: objectType,
		basePath:   crmV3ObjectsPath + "/" + objectType,
	}
}

// --- Basic API ---

// Create creates a new CRM object.
func (s *ObjectService) Create(ctx context.Context, input *SimplePublicObjectInputForCreate) (*SimplePublicObject, error) {
	var result SimplePublicObject
	if err := s.requester.Post(ctx, s.basePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a single CRM object by its ID.
func (s *ObjectService) GetByID(ctx context.Context, id string, opts *GetByIDOptions) (*SimplePublicObjectWithAssociations, error) {
	path := fmt.Sprintf("%s/%s", s.basePath, id)
	q := url.Values{}
	if opts != nil {
		if len(opts.Properties) > 0 {
			q.Set("properties", strings.Join(opts.Properties, ","))
		}
		if len(opts.PropertiesWithHistory) > 0 {
			q.Set("propertiesWithHistory", strings.Join(opts.PropertiesWithHistory, ","))
		}
		if len(opts.Associations) > 0 {
			q.Set("associations", strings.Join(opts.Associations, ","))
		}
		if opts.Archived {
			q.Set("archived", "true")
		}
		if opts.IDProperty != "" {
			q.Set("idProperty", opts.IDProperty)
		}
	}
	var result SimplePublicObjectWithAssociations
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing CRM object.
func (s *ObjectService) Update(ctx context.Context, id string, input *SimplePublicObjectInput, opts *UpdateOptions) (*SimplePublicObject, error) {
	path := fmt.Sprintf("%s/%s", s.basePath, id)
	if opts != nil && opts.IDProperty != "" {
		path += "?idProperty=" + url.QueryEscape(opts.IDProperty)
	}
	var result SimplePublicObject
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive soft-deletes a CRM object.
func (s *ObjectService) Archive(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s", s.basePath, id)
	return s.requester.Delete(ctx, path)
}

// List retrieves a page of CRM objects.
func (s *ObjectService) List(ctx context.Context, opts *ListOptions) (*ListResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if len(opts.Properties) > 0 {
			q.Set("properties", strings.Join(opts.Properties, ","))
		}
		if len(opts.PropertiesWithHistory) > 0 {
			q.Set("propertiesWithHistory", strings.Join(opts.PropertiesWithHistory, ","))
		}
		if len(opts.Associations) > 0 {
			q.Set("associations", strings.Join(opts.Associations, ","))
		}
		if opts.Archived {
			q.Set("archived", "true")
		}
	}
	var result ListResult
	if err := s.requester.Get(ctx, s.basePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Merge merges two objects of this type.
func (s *ObjectService) Merge(ctx context.Context, input *PublicMergeInput) (*SimplePublicObject, error) {
	path := s.basePath + "/merge"
	var result SimplePublicObject
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GdprDelete performs a GDPR-compliant deletion.
func (s *ObjectService) GdprDelete(ctx context.Context, input *PublicGdprDeleteInput) error {
	path := s.basePath + "/gdpr-delete"
	return s.requester.Post(ctx, path, input, nil)
}

// --- Batch API ---

// BatchCreate creates multiple objects in a single request.
func (s *ObjectService) BatchCreate(ctx context.Context, input *BatchCreateInput) (*BatchResult, error) {
	path := s.basePath + "/batch/create"
	var result BatchResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchRead retrieves multiple objects by ID.
func (s *ObjectService) BatchRead(ctx context.Context, input *BatchReadInput) (*BatchResult, error) {
	path := s.basePath + "/batch/read"
	var result BatchResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchUpdate updates multiple objects in a single request.
func (s *ObjectService) BatchUpdate(ctx context.Context, input *BatchUpdateInput) (*BatchResult, error) {
	path := s.basePath + "/batch/update"
	var result BatchResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchArchive archives multiple objects in a single request.
func (s *ObjectService) BatchArchive(ctx context.Context, input *BatchArchiveInput) error {
	path := s.basePath + "/batch/archive"
	return s.requester.Post(ctx, path, input, nil)
}

// BatchUpsert creates or updates multiple objects in a single request.
func (s *ObjectService) BatchUpsert(ctx context.Context, input *BatchUpsertInput) (*BatchUpsertResult, error) {
	path := s.basePath + "/batch/upsert"
	var result BatchUpsertResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Search API ---

// Search searches for objects using filters, query text, and sorting.
func (s *ObjectService) Search(ctx context.Context, input *PublicObjectSearchRequest) (*SearchResult, error) {
	path := s.basePath + "/search"
	var result SearchResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Helpers ---

// GetAll retrieves all objects by automatically paginating through all pages.
// Equivalent to the Node client's getAll helper.
func (s *ObjectService) GetAll(ctx context.Context, opts *GetAllOptions) ([]*SimplePublicObjectWithAssociations, error) {
	var allResults []*SimplePublicObjectWithAssociations
	after := ""

	limit := 100
	if opts != nil && opts.Limit > 0 {
		limit = opts.Limit
	}

	for {
		listOpts := &ListOptions{
			Limit: limit,
			After: after,
		}
		if opts != nil {
			listOpts.Properties = opts.Properties
			listOpts.PropertiesWithHistory = opts.PropertiesWithHistory
			listOpts.Associations = opts.Associations
			listOpts.Archived = opts.Archived
		}

		page, err := s.List(ctx, listOpts)
		if err != nil {
			return nil, err
		}

		allResults = append(allResults, page.Results...)

		if page.Paging == nil || page.Paging.Next == nil || page.Paging.Next.After == "" {
			break
		}
		after = page.Paging.Next.After
	}

	return allResults, nil
}
