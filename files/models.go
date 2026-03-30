package files

import (
	"io"
	"time"
)

// --- File models ---

// File represents a file object in the HubSpot File Manager.
type File struct {
	ID                string     `json:"id"`
	Name              string     `json:"name,omitempty"`
	Extension         string     `json:"extension,omitempty"`
	Access            string     `json:"access"`
	ParentFolderID    string     `json:"parentFolderId,omitempty"`
	SourceGroup       string     `json:"sourceGroup,omitempty"`
	FileMD5           string     `json:"fileMd5,omitempty"`
	Encoding          string     `json:"encoding,omitempty"`
	Type              string     `json:"type,omitempty"`
	IsUsableInContent *bool      `json:"isUsableInContent,omitempty"`
	URL               string     `json:"url,omitempty"`
	ExpiresAt         *int64     `json:"expiresAt,omitempty"`
	CreatedAt         time.Time  `json:"createdAt"`
	ArchivedAt        *time.Time `json:"archivedAt,omitempty"`
	Path              string     `json:"path,omitempty"`
	Archived          bool       `json:"archived"`
	Size              *int64     `json:"size,omitempty"`
	Width             *int       `json:"width,omitempty"`
	Height            *int       `json:"height,omitempty"`
	DefaultHostingURL string     `json:"defaultHostingUrl,omitempty"`
	UpdatedAt         time.Time  `json:"updatedAt"`
}

// FileUpdateInput is the input for updating file properties.
type FileUpdateInput struct {
	Access           string     `json:"access,omitempty"`
	ParentFolderID   string     `json:"parentFolderId,omitempty"`
	Name             string     `json:"name,omitempty"`
	ParentFolderPath string     `json:"parentFolderPath,omitempty"`
	ClearExpires     *bool      `json:"clearExpires,omitempty"`
	IsUsableInContent *bool     `json:"isUsableInContent,omitempty"`
	ExpiresAt        *time.Time `json:"expiresAt,omitempty"`
}

// FileSearchOptions contains query parameters for searching files.
type FileSearchOptions struct {
	Properties []string `json:"properties,omitempty"`
	After      string   `json:"after,omitempty"`
	Before     string   `json:"before,omitempty"`
	Limit      int      `json:"limit,omitempty"`
	Sort       []string `json:"sort,omitempty"`
	Name       string   `json:"name,omitempty"`
	Path       string   `json:"path,omitempty"`
	Type       string   `json:"type,omitempty"`
	Extension  string   `json:"extension,omitempty"`
}

// FileUploadOptions contains the metadata for a file upload.
type FileUploadOptions struct {
	File         io.Reader `json:"-"`
	FileName     string    `json:"fileName,omitempty"`
	FolderID     string    `json:"folderId,omitempty"`
	FolderPath   string    `json:"folderPath,omitempty"`
	CharsetHunch string    `json:"charsetHunch,omitempty"`
	Options      string    `json:"options,omitempty"`
}

// FileReplaceOptions contains the metadata for replacing a file.
type FileReplaceOptions struct {
	File         io.Reader `json:"-"`
	CharsetHunch string    `json:"charsetHunch,omitempty"`
	Options      string    `json:"options,omitempty"`
}

// --- Folder models ---

// Folder represents a folder in the HubSpot File Manager.
type Folder struct {
	ID             string     `json:"id"`
	Name           string     `json:"name,omitempty"`
	Path           string     `json:"path,omitempty"`
	ParentFolderID string     `json:"parentFolderId,omitempty"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	Archived       bool       `json:"archived"`
	ArchivedAt     *time.Time `json:"archivedAt,omitempty"`
}

// FolderInput is the input for creating a new folder.
type FolderInput struct {
	ParentFolderID string `json:"parentFolderId,omitempty"`
	ParentPath     string `json:"parentPath,omitempty"`
	Name           string `json:"name"`
}

// FolderUpdateInput is the input for updating a folder.
type FolderUpdateInput struct {
	ParentFolderID *int64 `json:"parentFolderId,omitempty"`
	Name           string `json:"name,omitempty"`
}

// FolderSearchOptions contains query parameters for searching folders.
type FolderSearchOptions struct {
	Properties []string `json:"properties,omitempty"`
	After      string   `json:"after,omitempty"`
	Before     string   `json:"before,omitempty"`
	Limit      int      `json:"limit,omitempty"`
	Sort       []string `json:"sort,omitempty"`
	Name       string   `json:"name,omitempty"`
	Path       string   `json:"path,omitempty"`
}

// --- Action response models ---

// FileActionResponse is the response for async file operations (e.g. import status).
type FileActionResponse struct {
	Result      any               `json:"result,omitempty"`
	CompletedAt time.Time         `json:"completedAt"`
	NumErrors   int               `json:"numErrors,omitempty"`
	RequestedAt *time.Time        `json:"requestedAt,omitempty"`
	StartedAt   time.Time         `json:"startedAt"`
	Links       map[string]string `json:"links,omitempty"`
	TaskID      string            `json:"taskId"`
	Status      string            `json:"status"`
}

// FolderActionResponse is the response for async folder operations.
type FolderActionResponse struct {
	Result      *Folder           `json:"result,omitempty"`
	CompletedAt time.Time         `json:"completedAt"`
	NumErrors   int               `json:"numErrors,omitempty"`
	RequestedAt *time.Time        `json:"requestedAt,omitempty"`
	StartedAt   time.Time         `json:"startedAt"`
	Links       map[string]string `json:"links,omitempty"`
	TaskID      string            `json:"taskId"`
	Status      string            `json:"status"`
}

// ImportFromURLInput is the input for importing a file from a URL.
type ImportFromURLInput struct {
	FolderPath                  string     `json:"folderPath,omitempty"`
	Access                      string     `json:"access"`
	DuplicateValidationScope    string     `json:"duplicateValidationScope,omitempty"`
	Name                        string     `json:"name,omitempty"`
	DuplicateValidationStrategy string     `json:"duplicateValidationStrategy,omitempty"`
	TTL                         string     `json:"ttl,omitempty"`
	Overwrite                   *bool      `json:"overwrite,omitempty"`
	ExpiresAt                   *time.Time `json:"expiresAt,omitempty"`
	URL                         string     `json:"url"`
	FolderID                    string     `json:"folderId,omitempty"`
}

// ImportFromURLTaskLocator is the response from importing a file via URL.
type ImportFromURLTaskLocator struct {
	Links map[string]string `json:"links"`
	ID    string            `json:"id"`
}

// SignedURL is a temporary signed URL for accessing a file.
type SignedURL struct {
	Extension string     `json:"extension"`
	Size      int64      `json:"size"`
	Name      string     `json:"name"`
	Width     *int       `json:"width,omitempty"`
	Type      string     `json:"type"`
	URL       string     `json:"url"`
	ExpiresAt time.Time  `json:"expiresAt"`
	Height    *int       `json:"height,omitempty"`
}

// --- Collection responses ---

// Paging holds pagination cursors.
type Paging struct {
	Next *NextPage `json:"next,omitempty"`
}

// NextPage holds the cursor for the next page.
type NextPage struct {
	Link  string `json:"link,omitempty"`
	After string `json:"after"`
}

// CollectionResponseFile is a paginated list of files.
type CollectionResponseFile struct {
	Paging  *Paging `json:"paging,omitempty"`
	Results []File  `json:"results"`
}

// CollectionResponseFolder is a paginated list of folders.
type CollectionResponseFolder struct {
	Paging  *Paging  `json:"paging,omitempty"`
	Results []Folder `json:"results"`
}

// --- Access level constants ---

const (
	AccessPublicIndexable    = "PUBLIC_INDEXABLE"
	AccessPublicNotIndexable = "PUBLIC_NOT_INDEXABLE"
	AccessHiddenIndexable    = "HIDDEN_INDEXABLE"
	AccessHiddenNotIndexable = "HIDDEN_NOT_INDEXABLE"
	AccessHiddenPrivate      = "HIDDEN_PRIVATE"
	AccessPrivate            = "PRIVATE"
	AccessHiddenSensitive    = "HIDDEN_SENSITIVE"
	AccessSensitive          = "SENSITIVE"
)

// --- Duplicate validation constants ---

const (
	DuplicateValidationScopeEntirePortal = "ENTIRE_PORTAL"
	DuplicateValidationScopeExactFolder  = "EXACT_FOLDER"
)

const (
	DuplicateValidationStrategyNone           = "NONE"
	DuplicateValidationStrategyReject         = "REJECT"
	DuplicateValidationStrategyReturnExisting = "RETURN_EXISTING"
)

// --- Action status constants ---

const (
	StatusPending    = "PENDING"
	StatusProcessing = "PROCESSING"
	StatusCanceled   = "CANCELED"
	StatusComplete   = "COMPLETE"
)
