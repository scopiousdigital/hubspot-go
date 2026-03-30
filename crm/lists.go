package crm

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const listsBasePath = "/crm/v3/lists"

// ListsService handles operations on HubSpot CRM lists, memberships, folders, and mapping.
type ListsService struct {
	requester api.Requester
}

// =============================================================================
// Lists API
// =============================================================================

// Create creates a new list.
func (s *ListsService) Create(ctx context.Context, input *ListCreateRequest) (*ListCreateResponse, error) {
	var result ListCreateResponse
	if err := s.requester.Post(ctx, listsBasePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a list by its ID.
func (s *ListsService) GetByID(ctx context.Context, listID string, includeFilters bool) (*ListFetchResponse, error) {
	path := fmt.Sprintf("%s/%s", listsBasePath, listID)
	q := url.Values{}
	if includeFilters {
		q.Set("includeFilters", "true")
	}
	var result ListFetchResponse
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAll retrieves multiple lists by their IDs.
func (s *ListsService) GetAll(ctx context.Context, listIDs []string, includeFilters bool) (*ListsByIDResponse, error) {
	q := url.Values{}
	if len(listIDs) > 0 {
		q.Set("listIds", strings.Join(listIDs, ","))
	}
	if includeFilters {
		q.Set("includeFilters", "true")
	}
	var result ListsByIDResponse
	if err := s.requester.Get(ctx, listsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByName retrieves a list by name and object type.
func (s *ListsService) GetByName(ctx context.Context, listName, objectTypeID string, includeFilters bool) (*ListFetchResponse, error) {
	path := fmt.Sprintf("%s/object-type-id/%s/name/%s", listsBasePath, objectTypeID, listName)
	q := url.Values{}
	if includeFilters {
		q.Set("includeFilters", "true")
	}
	var result ListFetchResponse
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Search searches for lists.
func (s *ListsService) Search(ctx context.Context, input *ListSearchRequest) (*ListSearchResponse, error) {
	path := listsBasePath + "/search"
	var result ListSearchResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Remove deletes a list.
func (s *ListsService) Remove(ctx context.Context, listID string) error {
	path := fmt.Sprintf("%s/%s", listsBasePath, listID)
	return s.requester.Delete(ctx, path)
}

// Restore restores a previously deleted list.
func (s *ListsService) Restore(ctx context.Context, listID string) error {
	path := fmt.Sprintf("%s/%s/restore", listsBasePath, listID)
	return s.requester.Put(ctx, path, nil, nil)
}

// UpdateName updates a list's name.
func (s *ListsService) UpdateName(ctx context.Context, listID, listName string, includeFilters bool) (*ListUpdateResponse, error) {
	path := fmt.Sprintf("%s/%s/update-list-name", listsBasePath, listID)
	q := url.Values{}
	if listName != "" {
		q.Set("listName", listName)
	}
	if includeFilters {
		q.Set("includeFilters", "true")
	}
	var result ListUpdateResponse
	// The Node API uses PUT for this, but we approximate via PUT
	if err := s.requester.Put(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateFilters updates a list's filter definition.
func (s *ListsService) UpdateFilters(ctx context.Context, listID string, input *ListFilterUpdateRequest, enrollObjectsInWorkflows bool) (*ListUpdateResponse, error) {
	path := fmt.Sprintf("%s/%s/update-list-filters", listsBasePath, listID)
	q := url.Values{}
	if enrollObjectsInWorkflows {
		q.Set("enrollObjectsInWorkflows", "true")
	}
	var result ListUpdateResponse
	if err := s.requester.Put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============================================================================
// Memberships API
// =============================================================================

// AddMembers adds records to a list.
func (s *ListsService) AddMembers(ctx context.Context, listID string, recordIDs []string) (*MembershipsUpdateResponse, error) {
	path := fmt.Sprintf("%s/%s/memberships/add", listsBasePath, listID)
	var result MembershipsUpdateResponse
	if err := s.requester.Put(ctx, path, recordIDs, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RemoveMembers removes records from a list.
func (s *ListsService) RemoveMembers(ctx context.Context, listID string, recordIDs []string) (*MembershipsUpdateResponse, error) {
	path := fmt.Sprintf("%s/%s/memberships/remove", listsBasePath, listID)
	var result MembershipsUpdateResponse
	if err := s.requester.Put(ctx, path, recordIDs, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RemoveAllMembers removes all records from a list.
func (s *ListsService) RemoveAllMembers(ctx context.Context, listID string) error {
	path := fmt.Sprintf("%s/%s/memberships", listsBasePath, listID)
	return s.requester.Delete(ctx, path)
}

// AddAndRemoveMembers atomically adds and removes records from a list.
func (s *ListsService) AddAndRemoveMembers(ctx context.Context, listID string, input *MembershipChangeRequest) (*MembershipsUpdateResponse, error) {
	path := fmt.Sprintf("%s/%s/memberships/add-and-remove", listsBasePath, listID)
	var result MembershipsUpdateResponse
	if err := s.requester.Put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMembersPage retrieves a page of list members.
func (s *ListsService) GetMembersPage(ctx context.Context, listID string, opts *MembershipListOptions) (*MembershipPageResult, error) {
	path := fmt.Sprintf("%s/%s/memberships", listsBasePath, listID)
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
	var result MembershipPageResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRecordLists retrieves the lists a record belongs to.
func (s *ListsService) GetRecordLists(ctx context.Context, objectTypeID, recordID string) (*RecordListMembershipResult, error) {
	path := fmt.Sprintf("%s/records/%s/%s/memberships", listsBasePath, objectTypeID, recordID)
	var result RecordListMembershipResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============================================================================
// Folders API
// =============================================================================

// CreateFolder creates a new list folder.
func (s *ListsService) CreateFolder(ctx context.Context, input *ListFolderCreateRequest) (*ListFolderCreateResponse, error) {
	path := listsBasePath + "/folders"
	var result ListFolderCreateResponse
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFolder retrieves a folder and its children.
func (s *ListsService) GetFolder(ctx context.Context, folderID string) (*ListFolderFetchResponse, error) {
	path := fmt.Sprintf("%s/folders/%s", listsBasePath, folderID)
	var result ListFolderFetchResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RenameFolder renames a folder.
func (s *ListsService) RenameFolder(ctx context.Context, folderID, newName string) (*ListFolderFetchResponse, error) {
	path := fmt.Sprintf("%s/folders/%s/rename", listsBasePath, folderID)
	q := url.Values{}
	if newName != "" {
		q.Set("newFolderName", newName)
	}
	var result ListFolderFetchResponse
	if err := s.requester.Put(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// MoveFolder moves a folder to a new parent.
func (s *ListsService) MoveFolder(ctx context.Context, folderID, newParentFolderID string) (*ListFolderFetchResponse, error) {
	path := fmt.Sprintf("%s/folders/%s/move/%s", listsBasePath, folderID, newParentFolderID)
	var result ListFolderFetchResponse
	if err := s.requester.Put(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RemoveFolder deletes a folder.
func (s *ListsService) RemoveFolder(ctx context.Context, folderID string) error {
	path := fmt.Sprintf("%s/folders/%s", listsBasePath, folderID)
	return s.requester.Delete(ctx, path)
}

// MoveList moves a list to a folder.
func (s *ListsService) MoveList(ctx context.Context, input *ListMoveRequest) error {
	path := listsBasePath + "/folders/move-list"
	return s.requester.Put(ctx, path, input, nil)
}

// =============================================================================
// Mapping API
// =============================================================================

// TranslateLegacyListID maps a legacy (ILS) list ID to a v3 list ID.
func (s *ListsService) TranslateLegacyListID(ctx context.Context, legacyListID string) (*PublicMigrationMapping, error) {
	path := listsBasePath + "/idmapping"
	q := url.Values{}
	if legacyListID != "" {
		q.Set("legacyListId", legacyListID)
	}
	var result PublicMigrationMapping
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TranslateLegacyListIDBatch maps multiple legacy list IDs to v3 list IDs.
func (s *ListsService) TranslateLegacyListIDBatch(ctx context.Context, legacyListIDs []string) (*PublicBatchMigrationMapping, error) {
	path := listsBasePath + "/idmapping"
	var result PublicBatchMigrationMapping
	if err := s.requester.Post(ctx, path, legacyListIDs, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
