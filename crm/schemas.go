package crm

import (
	"context"
	"fmt"
	"net/url"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const schemasBasePath = "/crm/v3/schemas"

// SchemasService handles operations on HubSpot custom object schemas.
type SchemasService struct {
	requester api.Requester
}

// Create creates a new custom object schema.
func (s *SchemasService) Create(ctx context.Context, input *ObjectSchemaEgg) (*ObjectSchema, error) {
	var result ObjectSchema
	if err := s.requester.Post(ctx, schemasBasePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAll retrieves all custom object schemas.
func (s *SchemasService) GetAll(ctx context.Context, archived bool) (*ObjectSchemaListResult, error) {
	q := url.Values{}
	if archived {
		q.Set("archived", "true")
	}
	var result ObjectSchemaListResult
	if err := s.requester.Get(ctx, schemasBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a custom object schema by its ID or fully qualified name.
func (s *SchemasService) GetByID(ctx context.Context, objectType string) (*ObjectSchema, error) {
	path := fmt.Sprintf("%s/%s", schemasBasePath, objectType)
	var result ObjectSchema
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a custom object schema.
func (s *SchemasService) Update(ctx context.Context, objectType string, input *ObjectTypeDefinitionPatch) (*ObjectTypeDefinition, error) {
	path := fmt.Sprintf("%s/%s", schemasBasePath, objectType)
	var result ObjectTypeDefinition
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive deletes a custom object schema.
func (s *SchemasService) Archive(ctx context.Context, objectType string) error {
	path := fmt.Sprintf("%s/%s", schemasBasePath, objectType)
	return s.requester.Delete(ctx, path)
}

// CreateAssociation creates an association on a custom object schema.
func (s *SchemasService) CreateAssociation(ctx context.Context, objectType string, input *SchemaAssociationDefinitionEgg) (*SchemaAssociationDefinition, error) {
	path := fmt.Sprintf("%s/%s/associations", schemasBasePath, objectType)
	var result SchemaAssociationDefinition
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchiveAssociation removes an association from a custom object schema.
func (s *SchemasService) ArchiveAssociation(ctx context.Context, objectType, associationIdentifier string) error {
	path := fmt.Sprintf("%s/%s/associations/%s", schemasBasePath, objectType, associationIdentifier)
	return s.requester.Delete(ctx, path)
}
