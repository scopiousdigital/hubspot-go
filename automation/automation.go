package automation

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const actionsBasePath = "/automation/v4/actions"

// Service provides access to HubSpot Automation Actions APIs.
type Service struct {
	requester api.Requester

	// Callbacks provides methods for completing automation callbacks.
	Callbacks *CallbacksService
	// Definitions provides methods for managing action definitions.
	Definitions *DefinitionsService
	// Functions provides methods for managing action functions.
	Functions *FunctionsService
	// Revisions provides methods for viewing action revisions.
	Revisions *RevisionsService
}

// NewService creates a new Automation service.
func NewService(r api.Requester) *Service {
	return &Service{
		requester:   r,
		Callbacks:   &CallbacksService{requester: r},
		Definitions: &DefinitionsService{requester: r},
		Functions:   &FunctionsService{requester: r},
		Revisions:   &RevisionsService{requester: r},
	}
}

// --- CallbacksService ---

// CallbacksService provides methods for completing automation action callbacks.
type CallbacksService struct {
	requester api.Requester
}

// Complete completes a single callback.
func (s *CallbacksService) Complete(ctx context.Context, callbackID string, input *CallbackCompletionRequest) error {
	path := fmt.Sprintf("%s/callbacks/%s/complete", actionsBasePath, callbackID)
	return s.requester.Post(ctx, path, input, nil)
}

// CompleteBatch completes multiple callbacks in a single request.
func (s *CallbacksService) CompleteBatch(ctx context.Context, input *BatchInputCallbackCompletionBatchRequest) error {
	path := actionsBasePath + "/callbacks/complete"
	return s.requester.Post(ctx, path, input, nil)
}

// --- DefinitionsService ---

// DefinitionsService provides methods for managing action definitions.
type DefinitionsService struct {
	requester api.Requester
}

// Create creates a new action definition for the given app.
func (s *DefinitionsService) Create(ctx context.Context, appID int64, input *PublicActionDefinitionEgg) (*PublicActionDefinition, error) {
	path := fmt.Sprintf("%s/%d", actionsBasePath, appID)
	var result PublicActionDefinition
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a single action definition.
func (s *DefinitionsService) GetByID(ctx context.Context, definitionID string, appID int64, archived bool) (*PublicActionDefinition, error) {
	path := fmt.Sprintf("%s/%d/%s", actionsBasePath, appID, definitionID)
	q := url.Values{}
	if archived {
		q.Set("archived", "true")
	}
	var result PublicActionDefinition
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List retrieves a page of action definitions for the given app.
func (s *DefinitionsService) List(ctx context.Context, appID int64, opts *DefinitionsListOptions) (*DefinitionsListResult, error) {
	path := fmt.Sprintf("%s/%d", actionsBasePath, appID)
	q := url.Values{}
	if opts != nil {
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if opts.Archived {
			q.Set("archived", "true")
		}
	}
	var result DefinitionsListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing action definition.
func (s *DefinitionsService) Update(ctx context.Context, definitionID string, appID int64, input *PublicActionDefinitionPatch) (*PublicActionDefinition, error) {
	path := fmt.Sprintf("%s/%d/%s", actionsBasePath, appID, definitionID)
	var result PublicActionDefinition
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive deletes an action definition.
func (s *DefinitionsService) Archive(ctx context.Context, definitionID string, appID int64) error {
	path := fmt.Sprintf("%s/%d/%s", actionsBasePath, appID, definitionID)
	return s.requester.Delete(ctx, path)
}

// --- FunctionsService ---

// FunctionsService provides methods for managing serverless functions attached to action definitions.
type FunctionsService struct {
	requester api.Requester
}

// List retrieves all function identifiers for an action definition.
func (s *FunctionsService) List(ctx context.Context, definitionID string, appID int64) (*FunctionsResult, error) {
	path := fmt.Sprintf("%s/%d/%s/functions", actionsBasePath, appID, definitionID)
	var result FunctionsResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a specific function by its ID and type.
func (s *FunctionsService) GetByID(ctx context.Context, definitionID string, functionType string, functionID string, appID int64) (*PublicActionFunction, error) {
	path := fmt.Sprintf("%s/%d/%s/functions/%s/%s", actionsBasePath, appID, definitionID, functionType, functionID)
	var result PublicActionFunction
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByFunctionType retrieves a function by its type.
func (s *FunctionsService) GetByFunctionType(ctx context.Context, definitionID string, functionType string, appID int64) (*PublicActionFunction, error) {
	path := fmt.Sprintf("%s/%d/%s/functions/%s", actionsBasePath, appID, definitionID, functionType)
	var result PublicActionFunction
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateOrReplace creates or replaces a function by ID and type.
func (s *FunctionsService) CreateOrReplace(ctx context.Context, definitionID string, functionType string, functionID string, appID int64, body string) (*PublicActionFunctionIdentifier, error) {
	path := fmt.Sprintf("%s/%d/%s/functions/%s/%s", actionsBasePath, appID, definitionID, functionType, functionID)
	var result PublicActionFunctionIdentifier
	if err := s.requester.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateOrReplaceByFunctionType creates or replaces a function by type.
func (s *FunctionsService) CreateOrReplaceByFunctionType(ctx context.Context, definitionID string, functionType string, appID int64, body string) (*PublicActionFunctionIdentifier, error) {
	path := fmt.Sprintf("%s/%d/%s/functions/%s", actionsBasePath, appID, definitionID, functionType)
	var result PublicActionFunctionIdentifier
	if err := s.requester.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive deletes a function by ID and type.
func (s *FunctionsService) Archive(ctx context.Context, definitionID string, functionType string, functionID string, appID int64) error {
	path := fmt.Sprintf("%s/%d/%s/functions/%s/%s", actionsBasePath, appID, definitionID, functionType, functionID)
	return s.requester.Delete(ctx, path)
}

// ArchiveByFunctionType deletes a function by type.
func (s *FunctionsService) ArchiveByFunctionType(ctx context.Context, definitionID string, functionType string, appID int64) error {
	path := fmt.Sprintf("%s/%d/%s/functions/%s", actionsBasePath, appID, definitionID, functionType)
	return s.requester.Delete(ctx, path)
}

// --- RevisionsService ---

// RevisionsService provides methods for viewing action definition revisions.
type RevisionsService struct {
	requester api.Requester
}

// GetByID retrieves a specific revision.
func (s *RevisionsService) GetByID(ctx context.Context, definitionID string, revisionID string, appID int64) (*PublicActionRevision, error) {
	path := fmt.Sprintf("%s/%d/%s/revisions/%s", actionsBasePath, appID, definitionID, revisionID)
	var result PublicActionRevision
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List retrieves a page of revisions for an action definition.
func (s *RevisionsService) List(ctx context.Context, definitionID string, appID int64, opts *RevisionsListOptions) (*RevisionsListResult, error) {
	path := fmt.Sprintf("%s/%d/%s/revisions", actionsBasePath, appID, definitionID)
	q := url.Values{}
	if opts != nil {
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
	}
	var result RevisionsListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
