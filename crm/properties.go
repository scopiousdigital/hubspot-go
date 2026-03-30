package crm

import (
	"context"
	"fmt"
	"net/url"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const propertiesBasePath = "/crm/v3/properties"

// PropertiesService handles operations on HubSpot CRM property definitions.
// It provides CRUD for properties and property groups, as well as batch operations.
type PropertiesService struct {
	requester api.Requester
}

// --- Core API ---

// Create creates a new property for the given object type.
func (s *PropertiesService) Create(ctx context.Context, objectType string, input *PropertyCreate) (*Property, error) {
	path := fmt.Sprintf("%s/%s", propertiesBasePath, objectType)
	var result Property
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAll retrieves all properties for the given object type.
func (s *PropertiesService) GetAll(ctx context.Context, objectType string, archived bool) (*PropertyListResult, error) {
	path := fmt.Sprintf("%s/%s", propertiesBasePath, objectType)
	q := url.Values{}
	if archived {
		q.Set("archived", "true")
	}
	var result PropertyListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByName retrieves a single property by name.
func (s *PropertiesService) GetByName(ctx context.Context, objectType, propertyName string, archived bool) (*Property, error) {
	path := fmt.Sprintf("%s/%s/%s", propertiesBasePath, objectType, propertyName)
	q := url.Values{}
	if archived {
		q.Set("archived", "true")
	}
	var result Property
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing property.
func (s *PropertiesService) Update(ctx context.Context, objectType, propertyName string, input *PropertyUpdate) (*Property, error) {
	path := fmt.Sprintf("%s/%s/%s", propertiesBasePath, objectType, propertyName)
	var result Property
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive deletes a property.
func (s *PropertiesService) Archive(ctx context.Context, objectType, propertyName string) error {
	path := fmt.Sprintf("%s/%s/%s", propertiesBasePath, objectType, propertyName)
	return s.requester.Delete(ctx, path)
}

// --- Batch API ---

// BatchCreate creates multiple properties in a single request.
func (s *PropertiesService) BatchCreate(ctx context.Context, objectType string, input *BatchPropertyCreateInput) (*BatchPropertyResult, error) {
	path := fmt.Sprintf("%s/%s/batch/create", propertiesBasePath, objectType)
	var result BatchPropertyResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchRead retrieves multiple properties by name.
func (s *PropertiesService) BatchRead(ctx context.Context, objectType string, input *BatchPropertyReadInput) (*BatchPropertyResult, error) {
	path := fmt.Sprintf("%s/%s/batch/read", propertiesBasePath, objectType)
	var result BatchPropertyResult
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Groups API ---

// CreateGroup creates a new property group.
func (s *PropertiesService) CreateGroup(ctx context.Context, objectType string, input *PropertyGroupCreate) (*PropertyGroup, error) {
	path := fmt.Sprintf("%s/%s/groups", propertiesBasePath, objectType)
	var result PropertyGroup
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAllGroups retrieves all property groups for the given object type.
func (s *PropertiesService) GetAllGroups(ctx context.Context, objectType string) (*PropertyGroupListResult, error) {
	path := fmt.Sprintf("%s/%s/groups", propertiesBasePath, objectType)
	var result PropertyGroupListResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetGroupByName retrieves a single property group by name.
func (s *PropertiesService) GetGroupByName(ctx context.Context, objectType, groupName string) (*PropertyGroup, error) {
	path := fmt.Sprintf("%s/%s/groups/%s", propertiesBasePath, objectType, groupName)
	var result PropertyGroup
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateGroup updates an existing property group.
func (s *PropertiesService) UpdateGroup(ctx context.Context, objectType, groupName string, input *PropertyGroupUpdate) (*PropertyGroup, error) {
	path := fmt.Sprintf("%s/%s/groups/%s", propertiesBasePath, objectType, groupName)
	var result PropertyGroup
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchiveGroup deletes a property group.
func (s *PropertiesService) ArchiveGroup(ctx context.Context, objectType, groupName string) error {
	path := fmt.Sprintf("%s/%s/groups/%s", propertiesBasePath, objectType, groupName)
	return s.requester.Delete(ctx, path)
}
