package files

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const filesV3Path = "/files/v3"

// Service provides access to the HubSpot File Manager APIs.
type Service struct {
	Files   *FilesService
	Folders *FoldersService
}

// NewService creates a new Files service. Called by the root hubspot package.
func NewService(r api.Requester) *Service {
	return &Service{
		Files:   &FilesService{requester: r},
		Folders: &FoldersService{requester: r},
	}
}

// --- FilesService ---

// FilesService handles file operations in the HubSpot File Manager.
type FilesService struct {
	requester api.Requester
}

// Upload uploads a file to the File Manager.
//
// Note: The standard Requester interface sends JSON bodies. File uploads to
// HubSpot require multipart/form-data encoding. Callers should ensure the
// underlying Requester implementation handles the io.Reader in FileUploadOptions
// appropriately, or use the raw HTTP client for file uploads.
func (s *FilesService) Upload(ctx context.Context, input *FileUploadOptions) (*File, error) {
	path := filesV3Path + "/files"
	var result File
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a file by its ID.
func (s *FilesService) GetByID(ctx context.Context, fileID string, properties []string) (*File, error) {
	path := fmt.Sprintf("%s/files/%s", filesV3Path, fileID)
	q := url.Values{}
	if len(properties) > 0 {
		q.Set("properties", strings.Join(properties, ","))
	}
	var result File
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Search searches for files using various filter parameters.
func (s *FilesService) Search(ctx context.Context, opts *FileSearchOptions) (*CollectionResponseFile, error) {
	path := filesV3Path + "/files/search"
	q := url.Values{}
	if opts != nil {
		if len(opts.Properties) > 0 {
			q.Set("properties", strings.Join(opts.Properties, ","))
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if opts.Before != "" {
			q.Set("before", opts.Before)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if len(opts.Sort) > 0 {
			q.Set("sort", strings.Join(opts.Sort, ","))
		}
		if opts.Name != "" {
			q.Set("name", opts.Name)
		}
		if opts.Path != "" {
			q.Set("path", opts.Path)
		}
		if opts.Type != "" {
			q.Set("type", opts.Type)
		}
		if opts.Extension != "" {
			q.Set("extension", opts.Extension)
		}
	}
	var result CollectionResponseFile
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive soft-deletes a file by its ID.
func (s *FilesService) Archive(ctx context.Context, fileID string) error {
	path := fmt.Sprintf("%s/files/%s", filesV3Path, fileID)
	return s.requester.Delete(ctx, path)
}

// Replace replaces the content of an existing file.
//
// Note: Like Upload, this requires multipart/form-data encoding. See Upload's
// documentation for details.
func (s *FilesService) Replace(ctx context.Context, fileID string, input *FileReplaceOptions) (*File, error) {
	path := fmt.Sprintf("%s/files/%s/replace", filesV3Path, fileID)
	var result File
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateProperties updates the metadata/properties of a file.
func (s *FilesService) UpdateProperties(ctx context.Context, fileID string, input *FileUpdateInput) (*File, error) {
	path := fmt.Sprintf("%s/files/%s", filesV3Path, fileID)
	var result File
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CheckImportStatus checks the status of a file import task.
func (s *FilesService) CheckImportStatus(ctx context.Context, taskID string) (*FileActionResponse, error) {
	path := fmt.Sprintf("%s/files/import-from-url/async/tasks/%s/status", filesV3Path, taskID)
	var result FileActionResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ImportFromURL imports a file from a URL.
func (s *FilesService) ImportFromURL(ctx context.Context, input *ImportFromURLInput) (*ImportFromURLTaskLocator, error) {
	path := filesV3Path + "/files/import-from-url/async"
	var result ImportFromURLTaskLocator
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSignedURL gets a temporarily signed URL for a file.
func (s *FilesService) GetSignedURL(ctx context.Context, fileID string, size string, expirationSeconds int, upscale bool) (*SignedURL, error) {
	path := fmt.Sprintf("%s/files/%s/signed-url", filesV3Path, fileID)
	q := url.Values{}
	if size != "" {
		q.Set("size", size)
	}
	if expirationSeconds > 0 {
		q.Set("expirationSeconds", strconv.Itoa(expirationSeconds))
	}
	if upscale {
		q.Set("upscale", "true")
	}
	var result SignedURL
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- FoldersService ---

// FoldersService handles folder operations in the HubSpot File Manager.
type FoldersService struct {
	requester api.Requester
}

// Create creates a new folder.
func (s *FoldersService) Create(ctx context.Context, input *FolderInput) (*Folder, error) {
	path := filesV3Path + "/folders"
	var result Folder
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Search searches for folders using various filter parameters.
func (s *FoldersService) Search(ctx context.Context, opts *FolderSearchOptions) (*CollectionResponseFolder, error) {
	path := filesV3Path + "/folders/search"
	q := url.Values{}
	if opts != nil {
		if len(opts.Properties) > 0 {
			q.Set("properties", strings.Join(opts.Properties, ","))
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if opts.Before != "" {
			q.Set("before", opts.Before)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if len(opts.Sort) > 0 {
			q.Set("sort", strings.Join(opts.Sort, ","))
		}
		if opts.Name != "" {
			q.Set("name", opts.Name)
		}
		if opts.Path != "" {
			q.Set("path", opts.Path)
		}
	}
	var result CollectionResponseFolder
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a folder by its ID.
func (s *FoldersService) GetByID(ctx context.Context, folderID string, properties []string) (*Folder, error) {
	path := fmt.Sprintf("%s/folders/%s", filesV3Path, folderID)
	q := url.Values{}
	if len(properties) > 0 {
		q.Set("properties", strings.Join(properties, ","))
	}
	var result Folder
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a folder's properties.
func (s *FoldersService) Update(ctx context.Context, folderID string, input *FolderUpdateInput) (*Folder, error) {
	path := fmt.Sprintf("%s/folders/%s", filesV3Path, folderID)
	var result Folder
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive soft-deletes a folder by its ID.
func (s *FoldersService) Archive(ctx context.Context, folderID string) error {
	path := fmt.Sprintf("%s/folders/%s", filesV3Path, folderID)
	return s.requester.Delete(ctx, path)
}

// UpdateProperties is an alias for Update, maintaining API parity with the
// Node.js client's updateProperties method.
func (s *FoldersService) UpdateProperties(ctx context.Context, folderID string, input *FolderUpdateInput) (*Folder, error) {
	return s.Update(ctx, folderID, input)
}
