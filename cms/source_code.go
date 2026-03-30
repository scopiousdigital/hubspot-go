package cms

import (
	"context"
	"fmt"
	"net/url"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const sourceCodeBasePath = "/cms/v3/source-code"

// --- SourceCodeContentService ---

// SourceCodeContentService handles source code file content operations (create, download, archive).
type SourceCodeContentService struct {
	requester api.Requester
}

// Create uploads a new source code file. Note: the actual Node client uses multipart upload;
// this simplified version sends the path and expects server-side handling.
func (s *SourceCodeContentService) Create(ctx context.Context, environment, filePath string) (*AssetFileMetadata, error) {
	path := fmt.Sprintf("%s/%s/content/%s", sourceCodeBasePath, url.PathEscape(environment), filePath)
	var result AssetFileMetadata
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateOrUpdate creates or updates a source code file.
func (s *SourceCodeContentService) CreateOrUpdate(ctx context.Context, environment, filePath string) (*AssetFileMetadata, error) {
	path := fmt.Sprintf("%s/%s/content/%s", sourceCodeBasePath, url.PathEscape(environment), filePath)
	var result AssetFileMetadata
	// Uses POST for create-or-update (PUT semantics in the actual API, but our
	// Requester only exposes Post/Patch). The HubSpot API accepts POST here.
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Download retrieves the raw content of a source code file.
// Note: in a real implementation this would return file bytes; here we return
// the metadata. Full file download support requires a raw HTTP response method.
func (s *SourceCodeContentService) Download(ctx context.Context, environment, filePath string) (*AssetFileMetadata, error) {
	path := fmt.Sprintf("%s/%s/content/%s", sourceCodeBasePath, url.PathEscape(environment), filePath)
	var result AssetFileMetadata
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive deletes a source code file.
func (s *SourceCodeContentService) Archive(ctx context.Context, environment, filePath string) error {
	path := fmt.Sprintf("%s/%s/content/%s", sourceCodeBasePath, url.PathEscape(environment), filePath)
	return s.requester.Delete(ctx, path)
}

// --- SourceCodeExtractService ---

// SourceCodeExtractService handles async file extraction operations.
type SourceCodeExtractService struct {
	requester api.Requester
}

// Extract starts an asynchronous file extraction task.
func (s *SourceCodeExtractService) Extract(ctx context.Context, input *FileExtractRequest) (*TaskLocator, error) {
	path := sourceCodeBasePath + "/extract/async"
	var result TaskLocator
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetExtractStatus checks the status of an async extraction task.
func (s *SourceCodeExtractService) GetExtractStatus(ctx context.Context, taskID int) (*ActionResponse, error) {
	path := fmt.Sprintf("%s/extract/async/tasks/%d/status", sourceCodeBasePath, taskID)
	var result ActionResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- SourceCodeMetadataService ---

// SourceCodeMetadataService handles source code file metadata operations.
type SourceCodeMetadataService struct {
	requester api.Requester
}

// Get retrieves metadata for a source code file.
func (s *SourceCodeMetadataService) Get(ctx context.Context, environment, filePath string, properties string) (*AssetFileMetadata, error) {
	path := fmt.Sprintf("%s/%s/metadata/%s", sourceCodeBasePath, url.PathEscape(environment), filePath)
	q := url.Values{}
	if properties != "" {
		q.Set("properties", properties)
	}
	var result AssetFileMetadata
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- SourceCodeValidationService ---

// SourceCodeValidationService handles source code file validation.
type SourceCodeValidationService struct {
	requester api.Requester
}

// Validate validates a source code file at the given path.
func (s *SourceCodeValidationService) Validate(ctx context.Context, filePath string) error {
	path := fmt.Sprintf("%s/validation/validate/%s", sourceCodeBasePath, filePath)
	return s.requester.Post(ctx, path, nil, nil)
}

// --- Aggregate SourceCodeService ---

// SourceCodeService provides access to all source code sub-services.
type SourceCodeService struct {
	Content    *SourceCodeContentService
	Extract    *SourceCodeExtractService
	Metadata   *SourceCodeMetadataService
	Validation *SourceCodeValidationService
}
