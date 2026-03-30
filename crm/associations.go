package crm

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const (
	associationsV3BasePath = "/crm/v3/associations"
	associationsV4BasePath = "/crm/v4/associations"
)

// AssociationsService handles operations on HubSpot CRM associations (v3 and v4).
type AssociationsService struct {
	requester api.Requester
}

// =============================================================================
// v3 Batch API
// =============================================================================

// BatchCreate creates associations in batch (v3).
func (s *AssociationsService) BatchCreate(ctx context.Context, fromObjectType, toObjectType string, input *BatchPublicAssociationInput) (*BatchPublicAssociationResult, error) {
	path := fmt.Sprintf("%s/%s/%s/batch/create", associationsV3BasePath, fromObjectType, toObjectType)
	var result BatchPublicAssociationResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchRead reads associations in batch (v3).
func (s *AssociationsService) BatchRead(ctx context.Context, fromObjectType, toObjectType string, input *BatchPublicObjectIDInput) (*BatchPublicAssociationMultiResult, error) {
	path := fmt.Sprintf("%s/%s/%s/batch/read", associationsV3BasePath, fromObjectType, toObjectType)
	var result BatchPublicAssociationMultiResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchArchive archives associations in batch (v3).
func (s *AssociationsService) BatchArchive(ctx context.Context, fromObjectType, toObjectType string, input *BatchPublicAssociationInput) error {
	path := fmt.Sprintf("%s/%s/%s/batch/archive", associationsV3BasePath, fromObjectType, toObjectType)
	return s.requester.Post(ctx, path, input, nil)
}

// =============================================================================
// v4 Basic API
// =============================================================================

// V4Create creates a labeled association between two objects (v4).
func (s *AssociationsService) V4Create(ctx context.Context, objectType, objectID, toObjectType, toObjectID string, specs []AssociationV4Spec) (*LabelsBetweenObjectPair, error) {
	path := fmt.Sprintf("%s/%s/%s/%s/%s", associationsV4BasePath, objectType, objectID, toObjectType, toObjectID)
	var result LabelsBetweenObjectPair
	if err := s.requester.Put(ctx, path, specs, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// V4Archive removes all associations between two objects (v4).
func (s *AssociationsService) V4Archive(ctx context.Context, objectType, objectID, toObjectType, toObjectID string) error {
	path := fmt.Sprintf("%s/%s/%s/%s/%s", associationsV4BasePath, objectType, objectID, toObjectType, toObjectID)
	return s.requester.Delete(ctx, path)
}

// V4CreateDefault creates a default association between two objects (v4).
func (s *AssociationsService) V4CreateDefault(ctx context.Context, fromObjectType, fromObjectID, toObjectType, toObjectID string) (*BatchPublicDefaultAssociationResult, error) {
	path := fmt.Sprintf("%s/%s/%s/%s/%s", associationsV4BasePath, fromObjectType, fromObjectID, toObjectType, toObjectID)
	var result BatchPublicDefaultAssociationResult
	if err := s.requester.Put(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// V4GetPage retrieves a page of associations for an object (v4).
func (s *AssociationsService) V4GetPage(ctx context.Context, objectType, objectID, toObjectType string, after string, limit int) (*CollectionMultiAssociatedObjectWithLabel, error) {
	path := fmt.Sprintf("%s/%s/%s/%s", associationsV4BasePath, objectType, objectID, toObjectType)
	q := url.Values{}
	if after != "" {
		q.Set("after", after)
	}
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	var result CollectionMultiAssociatedObjectWithLabel
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============================================================================
// v4 Batch API
// =============================================================================

// V4BatchCreate creates associations in batch (v4).
func (s *AssociationsService) V4BatchCreate(ctx context.Context, fromObjectType, toObjectType string, input *BatchAssociationV4MultiPostInput) (*BatchLabelsBetweenObjectPairResult, error) {
	path := fmt.Sprintf("%s/%s/%s/batch/create", associationsV4BasePath, fromObjectType, toObjectType)
	var result BatchLabelsBetweenObjectPairResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// V4BatchCreateDefault creates default associations in batch (v4).
func (s *AssociationsService) V4BatchCreateDefault(ctx context.Context, fromObjectType, toObjectType string, input *BatchAssociationV4DefaultMultiPostInput) (*BatchPublicDefaultAssociationResult, error) {
	path := fmt.Sprintf("%s/%s/%s/batch/create/default", associationsV4BasePath, fromObjectType, toObjectType)
	var result BatchPublicDefaultAssociationResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// V4BatchGetPage retrieves associations in batch (v4).
func (s *AssociationsService) V4BatchGetPage(ctx context.Context, fromObjectType, toObjectType string, input *BatchFetchAssociationsInput) (*BatchAssociationMultiWithLabelResult, error) {
	path := fmt.Sprintf("%s/%s/%s/batch/read", associationsV4BasePath, fromObjectType, toObjectType)
	var result BatchAssociationMultiWithLabelResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// V4BatchArchive archives associations in batch (v4).
func (s *AssociationsService) V4BatchArchive(ctx context.Context, fromObjectType, toObjectType string, input *BatchAssociationV4MultiArchiveInput) error {
	path := fmt.Sprintf("%s/%s/%s/batch/archive", associationsV4BasePath, fromObjectType, toObjectType)
	return s.requester.Post(ctx, path, input, nil)
}

// =============================================================================
// v4 Schema - Definitions API
// =============================================================================

// V4GetDefinitions retrieves all association definitions between two object types.
func (s *AssociationsService) V4GetDefinitions(ctx context.Context, fromObjectType, toObjectType string) (*AssociationDefinitionSpecWithLabelResult, error) {
	path := fmt.Sprintf("%s/%s/%s/labels", associationsV4BasePath, fromObjectType, toObjectType)
	var result AssociationDefinitionSpecWithLabelResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// V4CreateDefinition creates a custom association definition.
func (s *AssociationsService) V4CreateDefinition(ctx context.Context, fromObjectType, toObjectType string, input *PublicAssociationDefinitionCreateRequest) (*AssociationDefinitionSpecWithLabelResult, error) {
	path := fmt.Sprintf("%s/%s/%s/labels", associationsV4BasePath, fromObjectType, toObjectType)
	var result AssociationDefinitionSpecWithLabelResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// V4UpdateDefinition updates a custom association definition.
func (s *AssociationsService) V4UpdateDefinition(ctx context.Context, fromObjectType, toObjectType string, input *PublicAssociationDefinitionUpdateRequest) error {
	path := fmt.Sprintf("%s/%s/%s/labels", associationsV4BasePath, fromObjectType, toObjectType)
	return s.requester.Put(ctx, path, input, nil)
}

// V4RemoveDefinition removes a custom association definition.
func (s *AssociationsService) V4RemoveDefinition(ctx context.Context, fromObjectType, toObjectType string, associationTypeID int32) error {
	path := fmt.Sprintf("%s/%s/%s/labels/%d", associationsV4BasePath, fromObjectType, toObjectType, associationTypeID)
	return s.requester.Delete(ctx, path)
}
