package cms

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const hubdbBasePath = "/cms/v3/hubdb/tables"

// --- TablesService ---

// TablesService handles HubDB table operations.
type TablesService struct {
	requester api.Requester
}

// Create creates a new HubDB table.
func (s *TablesService) Create(ctx context.Context, input *HubDbTableRequest) (*HubDbTable, error) {
	var result HubDbTable
	if err := s.requester.Post(ctx, hubdbBasePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDetails retrieves a published table by ID or name.
func (s *TablesService) GetDetails(ctx context.Context, tableIDOrName string, opts *TableDetailsOptions) (*HubDbTable, error) {
	path := fmt.Sprintf("%s/%s", hubdbBasePath, tableIDOrName)
	q := url.Values{}
	if opts != nil {
		if opts.Archived {
			q.Set("archived", "true")
		}
		if opts.IncludeForeignIDs {
			q.Set("includeForeignIds", "true")
		}
	}
	var result HubDbTable
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDraftDetails retrieves the draft version of a table.
func (s *TablesService) GetDraftDetails(ctx context.Context, tableIDOrName string, opts *TableDetailsOptions) (*HubDbTable, error) {
	path := fmt.Sprintf("%s/%s/draft", hubdbBasePath, tableIDOrName)
	q := url.Values{}
	if opts != nil {
		if opts.Archived {
			q.Set("archived", "true")
		}
		if opts.IncludeForeignIDs {
			q.Set("includeForeignIds", "true")
		}
	}
	var result HubDbTable
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List retrieves all published tables.
func (s *TablesService) List(ctx context.Context, opts *TableListOptions) (*HubDbTableListResult, error) {
	q := buildTableListQuery(opts)
	var result HubDbTableListResult
	if err := s.requester.Get(ctx, hubdbBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListDraft retrieves all draft tables.
func (s *TablesService) ListDraft(ctx context.Context, opts *TableListOptions) (*HubDbTableListResult, error) {
	path := hubdbBasePath + "/draft"
	q := buildTableListQuery(opts)
	var result HubDbTableListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDraft updates the draft of a table.
func (s *TablesService) UpdateDraft(ctx context.Context, tableIDOrName string, input *HubDbTableRequest) (*HubDbTable, error) {
	path := fmt.Sprintf("%s/%s/draft", hubdbBasePath, tableIDOrName)
	var result HubDbTable
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Clone clones a table.
func (s *TablesService) Clone(ctx context.Context, tableIDOrName string, input *HubDbTableCloneRequest) (*HubDbTable, error) {
	path := fmt.Sprintf("%s/%s/draft/clone", hubdbBasePath, tableIDOrName)
	var result HubDbTable
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive deletes a table.
func (s *TablesService) Archive(ctx context.Context, tableIDOrName string) error {
	path := fmt.Sprintf("%s/%s", hubdbBasePath, tableIDOrName)
	return s.requester.Delete(ctx, path)
}

// Publish publishes the draft version of a table.
func (s *TablesService) Publish(ctx context.Context, tableIDOrName string) (*HubDbTable, error) {
	path := fmt.Sprintf("%s/%s/draft/publish", hubdbBasePath, tableIDOrName)
	var result HubDbTable
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unpublish unpublishes a table, returning it to draft state.
func (s *TablesService) Unpublish(ctx context.Context, tableIDOrName string) (*HubDbTable, error) {
	path := fmt.Sprintf("%s/%s/unpublish", hubdbBasePath, tableIDOrName)
	var result HubDbTable
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ResetDraft resets the draft to match the published version.
func (s *TablesService) ResetDraft(ctx context.Context, tableIDOrName string) (*HubDbTable, error) {
	path := fmt.Sprintf("%s/%s/draft/reset", hubdbBasePath, tableIDOrName)
	var result HubDbTable
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- RowsService ---

// RowsService handles HubDB row operations.
type RowsService struct {
	requester api.Requester
}

// Create creates a new row in a table's draft.
func (s *RowsService) Create(ctx context.Context, tableIDOrName string, input *HubDbTableRowRequest) (*HubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows", hubdbBasePath, tableIDOrName)
	var result HubDbTableRow
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a published row by ID.
func (s *RowsService) Get(ctx context.Context, tableIDOrName, rowID string) (*HubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/%s", hubdbBasePath, tableIDOrName, rowID)
	var result HubDbTableRow
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDraft retrieves a draft row by ID.
func (s *RowsService) GetDraft(ctx context.Context, tableIDOrName, rowID string) (*HubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/%s/draft", hubdbBasePath, tableIDOrName, rowID)
	var result HubDbTableRow
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List retrieves published rows for a table.
func (s *RowsService) List(ctx context.Context, tableIDOrName string, opts *RowListOptions) (*HubDbRowListResult, error) {
	path := fmt.Sprintf("%s/%s/rows", hubdbBasePath, tableIDOrName)
	q := buildRowListQuery(opts)
	var result HubDbRowListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListDraft retrieves draft rows for a table.
func (s *RowsService) ListDraft(ctx context.Context, tableIDOrName string, opts *RowListOptions) (*HubDbRowListResult, error) {
	path := fmt.Sprintf("%s/%s/rows/draft", hubdbBasePath, tableIDOrName)
	q := buildRowListQuery(opts)
	var result HubDbRowListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a draft row (partial update).
func (s *RowsService) Update(ctx context.Context, tableIDOrName, rowID string, input *HubDbTableRowRequest) (*HubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/%s/draft", hubdbBasePath, tableIDOrName, rowID)
	var result HubDbTableRow
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Replace replaces a draft row entirely.
func (s *RowsService) Replace(ctx context.Context, tableIDOrName, rowID string, input *HubDbTableRowRequest) (*HubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/%s/draft", hubdbBasePath, tableIDOrName, rowID)
	var result HubDbTableRow
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Clone clones a draft row.
func (s *RowsService) Clone(ctx context.Context, tableIDOrName, rowID string) (*HubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/%s/draft/clone", hubdbBasePath, tableIDOrName, rowID)
	var result HubDbTableRow
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Purge permanently removes a draft row.
func (s *RowsService) Purge(ctx context.Context, tableIDOrName, rowID string) error {
	path := fmt.Sprintf("%s/%s/rows/%s/draft", hubdbBasePath, tableIDOrName, rowID)
	return s.requester.Delete(ctx, path)
}

// --- RowsBatchService ---

// RowsBatchService handles batch operations on HubDB rows.
type RowsBatchService struct {
	requester api.Requester
}

// CreateDraft creates multiple draft rows.
func (s *RowsBatchService) CreateDraft(ctx context.Context, tableIDOrName string, input *BatchInputHubDbTableRowRequest) (*BatchResponseHubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/draft/batch/create", hubdbBasePath, tableIDOrName)
	var result BatchResponseHubDbTableRow
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReadPublished reads multiple published rows by ID.
func (s *RowsBatchService) ReadPublished(ctx context.Context, tableIDOrName string, input *BatchInputString) (*BatchResponseHubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/batch/read", hubdbBasePath, tableIDOrName)
	var result BatchResponseHubDbTableRow
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReadDraft reads multiple draft rows by ID.
func (s *RowsBatchService) ReadDraft(ctx context.Context, tableIDOrName string, input *BatchInputString) (*BatchResponseHubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/draft/batch/read", hubdbBasePath, tableIDOrName)
	var result BatchResponseHubDbTableRow
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDraft updates multiple draft rows.
func (s *RowsBatchService) UpdateDraft(ctx context.Context, tableIDOrName string, input *BatchInputHubDbTableRowBatchUpdateRequest) (*BatchResponseHubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/draft/batch/update", hubdbBasePath, tableIDOrName)
	var result BatchResponseHubDbTableRow
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReplaceDraft replaces multiple draft rows.
func (s *RowsBatchService) ReplaceDraft(ctx context.Context, tableIDOrName string, input *BatchInputHubDbTableRowBatchUpdateRequest) (*BatchResponseHubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/draft/batch/replace", hubdbBasePath, tableIDOrName)
	var result BatchResponseHubDbTableRow
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CloneDraft clones multiple draft rows.
func (s *RowsBatchService) CloneDraft(ctx context.Context, tableIDOrName string, input *BatchInputHubDbTableRowBatchCloneRequest) (*BatchResponseHubDbTableRow, error) {
	path := fmt.Sprintf("%s/%s/rows/draft/batch/clone", hubdbBasePath, tableIDOrName)
	var result BatchResponseHubDbTableRow
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PurgeDraft permanently removes multiple draft rows.
func (s *RowsBatchService) PurgeDraft(ctx context.Context, tableIDOrName string, input *BatchInputString) error {
	path := fmt.Sprintf("%s/%s/rows/draft/batch/purge", hubdbBasePath, tableIDOrName)
	return s.requester.Post(ctx, path, input, nil)
}

// --- Option types ---

// TableDetailsOptions configures a GetDetails request.
type TableDetailsOptions struct {
	Archived          bool
	IncludeForeignIDs bool
}

// TableListOptions configures a table list request.
type TableListOptions struct {
	Limit       int
	After       string
	Sort        []string
	Archived    bool
	ContentType string
}

// RowListOptions configures a row list request.
type RowListOptions struct {
	Limit      int
	After      string
	Sort       []string
	Properties []string
	Archived   bool
}

func buildTableListQuery(opts *TableListOptions) url.Values {
	q := url.Values{}
	if opts == nil {
		return q
	}
	if opts.Limit > 0 {
		q.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.After != "" {
		q.Set("after", opts.After)
	}
	if len(opts.Sort) > 0 {
		q.Set("sort", strings.Join(opts.Sort, ","))
	}
	if opts.Archived {
		q.Set("archived", "true")
	}
	if opts.ContentType != "" {
		q.Set("contentType", opts.ContentType)
	}
	return q
}

func buildRowListQuery(opts *RowListOptions) url.Values {
	q := url.Values{}
	if opts == nil {
		return q
	}
	if opts.Limit > 0 {
		q.Set("limit", strconv.Itoa(opts.Limit))
	}
	if opts.After != "" {
		q.Set("after", opts.After)
	}
	if len(opts.Sort) > 0 {
		q.Set("sort", strings.Join(opts.Sort, ","))
	}
	if len(opts.Properties) > 0 {
		q.Set("properties", strings.Join(opts.Properties, ","))
	}
	if opts.Archived {
		q.Set("archived", "true")
	}
	return q
}
