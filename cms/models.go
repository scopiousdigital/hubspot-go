package cms

import "time"

// --- Pagination ---

// ForwardPaging contains the cursor for the next page of results.
type ForwardPaging struct {
	Next *NextPage `json:"next,omitempty"`
}

// NextPage holds the cursor and optional link for pagination.
type NextPage struct {
	After string `json:"after"`
	Link  string `json:"link,omitempty"`
}

// Paging contains both next and previous page cursors.
type Paging struct {
	Next *NextPage     `json:"next,omitempty"`
	Prev *PreviousPage `json:"prev,omitempty"`
}

// PreviousPage holds the cursor for the previous page.
type PreviousPage struct {
	Before string `json:"before"`
	Link   string `json:"link,omitempty"`
}

// --- Blog models ---

// BlogPost represents a HubSpot CMS blog post.
type BlogPost struct {
	ID                   string     `json:"id,omitempty"`
	Name                 string     `json:"name,omitempty"`
	Slug                 string     `json:"slug,omitempty"`
	ContentGroupID       string     `json:"contentGroupId,omitempty"`
	Campaign             string     `json:"campaign,omitempty"`
	CategoryID           int64      `json:"categoryId,omitempty"`
	State                string     `json:"state,omitempty"`
	AuthorName           string     `json:"authorName,omitempty"`
	BlogAuthorID         string     `json:"blogAuthorId,omitempty"`
	TagIDs               []int64    `json:"tagIds,omitempty"`
	HtmlTitle            string     `json:"htmlTitle,omitempty"`
	PostBody             string     `json:"postBody,omitempty"`
	PostSummary          string     `json:"postSummary,omitempty"`
	RssBody              string     `json:"rssBody,omitempty"`
	RssSummary           string     `json:"rssSummary,omitempty"`
	MetaDescription      string     `json:"metaDescription,omitempty"`
	FeaturedImage        string     `json:"featuredImage,omitempty"`
	FeaturedImageAltText string     `json:"featuredImageAltText,omitempty"`
	CurrentState         string     `json:"currentState,omitempty"`
	Language             string     `json:"language,omitempty"`
	TranslatedFromID     string     `json:"translatedFromId,omitempty"`
	PublishDate          *time.Time `json:"publishDate,omitempty"`
	Created              *time.Time `json:"created,omitempty"`
	Updated              *time.Time `json:"updated,omitempty"`
	ArchivedAt           *time.Time `json:"archivedAt,omitempty"`
	Archived             bool       `json:"archived,omitempty"`
	URL                  string     `json:"url,omitempty"`
}

// BlogAuthor represents a blog author.
type BlogAuthor struct {
	ID               string     `json:"id,omitempty"`
	FullName         string     `json:"fullName,omitempty"`
	Email            string     `json:"email,omitempty"`
	Slug             string     `json:"slug,omitempty"`
	Language         string     `json:"language,omitempty"`
	TranslatedFromID string     `json:"translatedFromId,omitempty"`
	Bio              string     `json:"bio,omitempty"`
	Website          string     `json:"website,omitempty"`
	Twitter          string     `json:"twitter,omitempty"`
	Facebook         string     `json:"facebook,omitempty"`
	LinkedIn         string     `json:"linkedin,omitempty"`
	Avatar           string     `json:"avatar,omitempty"`
	DisplayName      string     `json:"displayName,omitempty"`
	Created          *time.Time `json:"created,omitempty"`
	Updated          *time.Time `json:"updated,omitempty"`
	ArchivedAt       *time.Time `json:"archivedAt,omitempty"`
	Archived         bool       `json:"archived,omitempty"`
}

// BlogTag represents a blog tag.
type BlogTag struct {
	ID               string     `json:"id,omitempty"`
	Name             string     `json:"name,omitempty"`
	Slug             string     `json:"slug,omitempty"`
	Language         string     `json:"language,omitempty"`
	TranslatedFromID string     `json:"translatedFromId,omitempty"`
	Created          *time.Time `json:"created,omitempty"`
	Updated          *time.Time `json:"updated,omitempty"`
	ArchivedAt       *time.Time `json:"archivedAt,omitempty"`
	Archived         bool       `json:"archived,omitempty"`
}

// --- Blog batch types ---

// BatchInputBlogPost is the input for batch blog post operations.
type BatchInputBlogPost struct {
	Inputs []BlogPost `json:"inputs"`
}

// BatchResponseBlogPost is the response for batch blog post operations.
type BatchResponseBlogPost struct {
	Status      string          `json:"status"`
	Results     []*BlogPost     `json:"results"`
	StartedAt   time.Time       `json:"startedAt"`
	CompletedAt time.Time       `json:"completedAt"`
	NumErrors   int             `json:"numErrors,omitempty"`
	Errors      []StandardError `json:"errors,omitempty"`
}

// BatchInputBlogAuthor is the input for batch blog author operations.
type BatchInputBlogAuthor struct {
	Inputs []BlogAuthor `json:"inputs"`
}

// BatchResponseBlogAuthor is the response for batch blog author operations.
type BatchResponseBlogAuthor struct {
	Status      string          `json:"status"`
	Results     []*BlogAuthor   `json:"results"`
	StartedAt   time.Time       `json:"startedAt"`
	CompletedAt time.Time       `json:"completedAt"`
	NumErrors   int             `json:"numErrors,omitempty"`
	Errors      []StandardError `json:"errors,omitempty"`
}

// BatchInputBlogTag is the input for batch blog tag operations.
type BatchInputBlogTag struct {
	Inputs []BlogTag `json:"inputs"`
}

// BatchResponseBlogTag is the response for batch blog tag operations.
type BatchResponseBlogTag struct {
	Status      string          `json:"status"`
	Results     []*BlogTag      `json:"results"`
	StartedAt   time.Time       `json:"startedAt"`
	CompletedAt time.Time       `json:"completedAt"`
	NumErrors   int             `json:"numErrors,omitempty"`
	Errors      []StandardError `json:"errors,omitempty"`
}

// --- Blog list results ---

// BlogPostListResult is the paginated response for blog posts.
type BlogPostListResult struct {
	Total   int            `json:"total"`
	Results []*BlogPost    `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// BlogAuthorListResult is the paginated response for blog authors.
type BlogAuthorListResult struct {
	Total   int            `json:"total"`
	Results []*BlogAuthor  `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// BlogTagListResult is the paginated response for blog tags.
type BlogTagListResult struct {
	Total   int            `json:"total"`
	Results []*BlogTag     `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// --- Blog versioning ---

// VersionBlogPost represents a previous version of a blog post.
type VersionBlogPost struct {
	Object *BlogPost `json:"object,omitempty"`
	User   *User     `json:"user,omitempty"`
	ID     string    `json:"id"`
}

// VersionListResult is the paginated response for versions.
type VersionListResult struct {
	Total   int                `json:"total"`
	Results []*VersionBlogPost `json:"results"`
	Paging  *ForwardPaging     `json:"paging,omitempty"`
}

// User represents a HubSpot user reference in version history.
type User struct {
	ID    string `json:"id"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// --- Blog multi-language types ---

// AttachToLangPrimaryRequest is used to attach content to a language group.
type AttachToLangPrimaryRequest struct {
	ID        string `json:"id"`
	Language  string `json:"language"`
	PrimaryID string `json:"primaryId"`
}

// DetachFromLangGroupRequest is used to detach content from a language group.
type DetachFromLangGroupRequest struct {
	ID string `json:"id"`
}

// SetNewLanguagePrimaryRequest sets a new primary language for a group.
type SetNewLanguagePrimaryRequest struct {
	ID string `json:"id"`
}

// UpdateLanguagesRequest updates language settings for content.
type UpdateLanguagesRequest struct {
	Languages map[string]string `json:"languages"`
	PrimaryID string            `json:"primaryId"`
}

// LanguageCloneRequest is used to create a language variation.
type LanguageCloneRequest struct {
	ID       string `json:"id"`
	Language string `json:"language,omitempty"`
}

// --- Content shared types ---

// ContentCloneRequest is used to clone content items.
type ContentCloneRequest struct {
	ID        string `json:"id"`
	CloneName string `json:"cloneName,omitempty"`
}

// ContentScheduleRequest is used to schedule content for publishing.
type ContentScheduleRequest struct {
	ID          string     `json:"id"`
	PublishDate *time.Time `json:"publishDate"`
}

// --- Batch input types ---

// BatchInputString is a batch input of string IDs.
type BatchInputString struct {
	Inputs []string `json:"inputs"`
}

// BatchInputJsonNode is a generic batch input for updates (maps).
type BatchInputJsonNode struct {
	Inputs []map[string]any `json:"inputs"`
}

// --- Common error types ---

// StandardError represents an error within a batch response.
type StandardError struct {
	Status      string              `json:"status"`
	ID          string              `json:"id,omitempty"`
	Category    string              `json:"category"`
	SubCategory string              `json:"subCategory,omitempty"`
	Message     string              `json:"message"`
	Errors      []ErrorDetail       `json:"errors,omitempty"`
	Context     map[string][]string `json:"context,omitempty"`
	Links       map[string]string   `json:"links,omitempty"`
}

// ErrorDetail represents a single validation or processing error.
type ErrorDetail struct {
	Message     string              `json:"message"`
	In          string              `json:"in"`
	Code        string              `json:"code"`
	SubCategory string              `json:"subCategory"`
	Context     map[string][]string `json:"context,omitempty"`
}

// --- HubDB models ---

// HubDbTable represents a HubDB table.
type HubDbTable struct {
	ID                    string         `json:"id,omitempty"`
	Name                  string         `json:"name,omitempty"`
	Label                 string         `json:"label,omitempty"`
	Columns               []*HubDbColumn `json:"columns,omitempty"`
	Published             bool           `json:"published,omitempty"`
	RowCount              int            `json:"rowCount,omitempty"`
	CreatedBy             *User          `json:"createdBy,omitempty"`
	UpdatedBy             *User          `json:"updatedBy,omitempty"`
	PublishedAt           *time.Time     `json:"publishedAt,omitempty"`
	AllowPublicAPIAccess  bool           `json:"allowPublicApiAccess,omitempty"`
	UseForPages           bool           `json:"useForPages,omitempty"`
	EnableChildTablePages bool           `json:"enableChildTablePages,omitempty"`
	DynamicMetaTags       map[string]int `json:"dynamicMetaTags,omitempty"`
	AllowChildTables      bool           `json:"allowChildTables,omitempty"`
	CreatedAt             *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt             *time.Time     `json:"updatedAt,omitempty"`
	Archived              bool           `json:"archived,omitempty"`
}

// HubDbTableRequest is the input for creating or updating a HubDB table.
type HubDbTableRequest struct {
	Name                  string         `json:"name,omitempty"`
	Label                 string         `json:"label,omitempty"`
	Columns               []*HubDbColumn `json:"columns,omitempty"`
	AllowPublicAPIAccess  bool           `json:"allowPublicApiAccess,omitempty"`
	UseForPages           bool           `json:"useForPages,omitempty"`
	EnableChildTablePages bool           `json:"enableChildTablePages,omitempty"`
	DynamicMetaTags       map[string]int `json:"dynamicMetaTags,omitempty"`
	AllowChildTables      bool           `json:"allowChildTables,omitempty"`
}

// HubDbColumn represents a column in a HubDB table.
type HubDbColumn struct {
	Name     string              `json:"name,omitempty"`
	Label    string              `json:"label,omitempty"`
	ID       int                 `json:"id,omitempty"`
	Type     string              `json:"type,omitempty"`
	Options  []HubDbColumnOption `json:"options,omitempty"`
	Archived bool                `json:"archived,omitempty"`
}

// HubDbColumnOption represents an option for a select-type column.
type HubDbColumnOption struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Label string `json:"label,omitempty"`
	Order int    `json:"order,omitempty"`
}

// HubDbTableRow represents a row in a HubDB table.
type HubDbTableRow struct {
	ID           string         `json:"id,omitempty"`
	Path         string         `json:"path,omitempty"`
	Name         string         `json:"name,omitempty"`
	Values       map[string]any `json:"values,omitempty"`
	ChildTableID string         `json:"childTableId,omitempty"`
	CreatedAt    *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time     `json:"updatedAt,omitempty"`
}

// HubDbTableRowRequest is the input for creating or updating a row.
type HubDbTableRowRequest struct {
	Path         string         `json:"path,omitempty"`
	Name         string         `json:"name,omitempty"`
	Values       map[string]any `json:"values,omitempty"`
	ChildTableID string         `json:"childTableId,omitempty"`
}

// HubDbTableRowBatchUpdateRequest is the input for batch updating a row.
type HubDbTableRowBatchUpdateRequest struct {
	ID     string         `json:"id"`
	Path   string         `json:"path,omitempty"`
	Name   string         `json:"name,omitempty"`
	Values map[string]any `json:"values,omitempty"`
}

// HubDbTableRowBatchCloneRequest is the input for batch cloning a row.
type HubDbTableRowBatchCloneRequest struct {
	ID string `json:"id"`
}

// HubDbTableCloneRequest is the input for cloning a table.
type HubDbTableCloneRequest struct {
	NewName  string `json:"newName,omitempty"`
	NewLabel string `json:"newLabel,omitempty"`
	CopyRows bool   `json:"copyRows,omitempty"`
}

// HubDbImportResult is the result of importing rows into a HubDB table.
type HubDbImportResult struct {
	RowsImported     int             `json:"rowsImported"`
	DuplicateRows    int             `json:"duplicateRows"`
	RowLimitExceeded bool            `json:"rowLimitExceeded"`
	Errors           []StandardError `json:"errors,omitempty"`
}

// --- HubDB list results ---

// HubDbTableListResult is the paginated response for HubDB tables.
type HubDbTableListResult struct {
	Total   int            `json:"total"`
	Results []*HubDbTable  `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// HubDbRowListResult is the paginated response for HubDB rows.
type HubDbRowListResult struct {
	Total   int              `json:"total"`
	Results []*HubDbTableRow `json:"results"`
	Paging  *ForwardPaging   `json:"paging,omitempty"`
}

// --- HubDB batch types ---

// BatchInputHubDbTableRowRequest is the input for batch row creation.
type BatchInputHubDbTableRowRequest struct {
	Inputs []HubDbTableRowRequest `json:"inputs"`
}

// BatchInputHubDbTableRowBatchUpdateRequest is the input for batch row updates.
type BatchInputHubDbTableRowBatchUpdateRequest struct {
	Inputs []HubDbTableRowBatchUpdateRequest `json:"inputs"`
}

// BatchInputHubDbTableRowBatchCloneRequest is the input for batch row cloning.
type BatchInputHubDbTableRowBatchCloneRequest struct {
	Inputs []HubDbTableRowBatchCloneRequest `json:"inputs"`
}

// BatchResponseHubDbTableRow is the response for batch row operations.
type BatchResponseHubDbTableRow struct {
	Status      string           `json:"status"`
	Results     []*HubDbTableRow `json:"results"`
	StartedAt   time.Time        `json:"startedAt"`
	CompletedAt time.Time        `json:"completedAt"`
	NumErrors   int              `json:"numErrors,omitempty"`
	Errors      []StandardError  `json:"errors,omitempty"`
}

// --- Domain models ---

// Domain represents a HubSpot domain.
type Domain struct {
	ID                        string     `json:"id"`
	Domain                    string     `json:"domain"`
	IsUsedForLandingPage      bool       `json:"isUsedForLandingPage"`
	IsUsedForBlogPost         bool       `json:"isUsedForBlogPost"`
	IsUsedForSitePage         bool       `json:"isUsedForSitePage"`
	IsUsedForEmail            bool       `json:"isUsedForEmail"`
	IsUsedForKnowledge        bool       `json:"isUsedForKnowledge"`
	IsResolving               bool       `json:"isResolving"`
	IsSslEnabled              *bool      `json:"isSslEnabled,omitempty"`
	IsSslOnly                 *bool      `json:"isSslOnly,omitempty"`
	PrimaryBlogPost           *bool      `json:"primaryBlogPost,omitempty"`
	PrimaryLandingPage        *bool      `json:"primaryLandingPage,omitempty"`
	PrimarySitePage           *bool      `json:"primarySitePage,omitempty"`
	PrimaryEmail              *bool      `json:"primaryEmail,omitempty"`
	PrimaryKnowledge          *bool      `json:"primaryKnowledge,omitempty"`
	SecondaryToDomain         string     `json:"secondaryToDomain,omitempty"`
	ManuallyMarkedAsResolving bool       `json:"manuallyMarkedAsResolving,omitempty"`
	CorrectCname              string     `json:"correctCname,omitempty"`
	Created                   *time.Time `json:"created,omitempty"`
	Updated                   *time.Time `json:"updated,omitempty"`
}

// DomainListResult is the paginated response for domains.
type DomainListResult struct {
	Total   int            `json:"total"`
	Results []*Domain      `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// --- Page models ---

// Page represents a HubSpot CMS page (landing page or site page).
type Page struct {
	ID                   string     `json:"id,omitempty"`
	Name                 string     `json:"name,omitempty"`
	Slug                 string     `json:"slug,omitempty"`
	State                string     `json:"state,omitempty"`
	Domain               string     `json:"domain,omitempty"`
	Subcategory          string     `json:"subcategory,omitempty"`
	ContentGroupID       string     `json:"contentGroupId,omitempty"`
	Campaign             string     `json:"campaign,omitempty"`
	HtmlTitle            string     `json:"htmlTitle,omitempty"`
	MetaDescription      string     `json:"metaDescription,omitempty"`
	Language             string     `json:"language,omitempty"`
	TranslatedFromID     string     `json:"translatedFromId,omitempty"`
	FeaturedImage        string     `json:"featuredImage,omitempty"`
	FeaturedImageAltText string     `json:"featuredImageAltText,omitempty"`
	TemplateID           string     `json:"templateId,omitempty"`
	TemplatePath         string     `json:"templatePath,omitempty"`
	CurrentState         string     `json:"currentState,omitempty"`
	PublishDate          *time.Time `json:"publishDate,omitempty"`
	Created              *time.Time `json:"created,omitempty"`
	Updated              *time.Time `json:"updated,omitempty"`
	ArchivedAt           *time.Time `json:"archivedAt,omitempty"`
	Archived             bool       `json:"archived,omitempty"`
	URL                  string     `json:"url,omitempty"`
}

// ContentFolder represents a content folder in the CMS.
type ContentFolder struct {
	ID             string     `json:"id,omitempty"`
	Name           string     `json:"name,omitempty"`
	ParentFolderID string     `json:"parentFolderId,omitempty"`
	Created        *time.Time `json:"created,omitempty"`
	Updated        *time.Time `json:"updated,omitempty"`
	Archived       bool       `json:"archived,omitempty"`
}

// VersionPage represents a previous version of a page.
type VersionPage struct {
	Object *Page  `json:"object,omitempty"`
	User   *User  `json:"user,omitempty"`
	ID     string `json:"id"`
}

// VersionPageListResult is the paginated response for page versions.
type VersionPageListResult struct {
	Total   int            `json:"total"`
	Results []*VersionPage `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// VersionContentFolder represents a previous version of a content folder.
type VersionContentFolder struct {
	Object *ContentFolder `json:"object,omitempty"`
	User   *User          `json:"user,omitempty"`
	ID     string         `json:"id"`
}

// VersionContentFolderListResult is the paginated response for folder versions.
type VersionContentFolderListResult struct {
	Total   int                     `json:"total"`
	Results []*VersionContentFolder `json:"results"`
	Paging  *ForwardPaging          `json:"paging,omitempty"`
}

// --- Page A/B test types ---

// ABTestCreateRequest is used to create an A/B test variation.
type ABTestCreateRequest struct {
	VariationID string `json:"variationId"`
}

// ABTestEndRequest is used to end an active A/B test.
type ABTestEndRequest struct {
	WinningVariationID string `json:"winningVariationId"`
}

// ABTestRerunRequest is used to rerun a previous A/B test.
type ABTestRerunRequest struct {
	ABTestID string `json:"abTestId"`
}

// --- Page batch types ---

// BatchInputPage is the input for batch page operations.
type BatchInputPage struct {
	Inputs []Page `json:"inputs"`
}

// BatchResponsePage is the response for batch page operations.
type BatchResponsePage struct {
	Status      string          `json:"status"`
	Results     []*Page         `json:"results"`
	StartedAt   time.Time       `json:"startedAt"`
	CompletedAt time.Time       `json:"completedAt"`
	NumErrors   int             `json:"numErrors,omitempty"`
	Errors      []StandardError `json:"errors,omitempty"`
}

// BatchInputContentFolder is the input for batch folder operations.
type BatchInputContentFolder struct {
	Inputs []ContentFolder `json:"inputs"`
}

// BatchResponseContentFolder is the response for batch folder operations.
type BatchResponseContentFolder struct {
	Status      string           `json:"status"`
	Results     []*ContentFolder `json:"results"`
	StartedAt   time.Time        `json:"startedAt"`
	CompletedAt time.Time        `json:"completedAt"`
	NumErrors   int              `json:"numErrors,omitempty"`
	Errors      []StandardError  `json:"errors,omitempty"`
}

// --- Page list results ---

// PageListResult is the paginated response for pages.
type PageListResult struct {
	Total   int            `json:"total"`
	Results []*Page        `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// ContentFolderListResult is the paginated response for content folders.
type ContentFolderListResult struct {
	Total   int              `json:"total"`
	Results []*ContentFolder `json:"results"`
	Paging  *ForwardPaging   `json:"paging,omitempty"`
}

// --- Audit log models ---

// PublicAuditLog represents a CMS audit log entry.
type PublicAuditLog struct {
	ObjectName string    `json:"objectName"`
	FullName   string    `json:"fullName"`
	Event      string    `json:"event"`
	UserID     string    `json:"userId"`
	ObjectID   string    `json:"objectId"`
	ObjectType string    `json:"objectType"`
	Timestamp  time.Time `json:"timestamp"`
	Meta       any       `json:"meta,omitempty"`
}

// AuditLogListResult is the response for audit log queries.
type AuditLogListResult struct {
	Results []*PublicAuditLog `json:"results"`
	Paging  *Paging           `json:"paging,omitempty"`
}

// Audit log event type constants.
const (
	AuditLogEventCreated     = "CREATED"
	AuditLogEventUpdated     = "UPDATED"
	AuditLogEventPublished   = "PUBLISHED"
	AuditLogEventDeleted     = "DELETED"
	AuditLogEventUnpublished = "UNPUBLISHED"
	AuditLogEventRestore     = "RESTORE"
)

// Audit log object type constants.
const (
	AuditLogObjectTypeBlog               = "BLOG"
	AuditLogObjectTypeBlogPost           = "BLOG_POST"
	AuditLogObjectTypeLandingPage        = "LANDING_PAGE"
	AuditLogObjectTypeWebsitePage        = "WEBSITE_PAGE"
	AuditLogObjectTypeTemplate           = "TEMPLATE"
	AuditLogObjectTypeModule             = "MODULE"
	AuditLogObjectTypeGlobalModule       = "GLOBAL_MODULE"
	AuditLogObjectTypeServerlessFunction = "SERVERLESS_FUNCTION"
	AuditLogObjectTypeDomain             = "DOMAIN"
	AuditLogObjectTypeURLMapping         = "URL_MAPPING"
	AuditLogObjectTypeEmail              = "EMAIL"
	AuditLogObjectTypeContentSettings    = "CONTENT_SETTINGS"
	AuditLogObjectTypeHubDBTable         = "HUBDB_TABLE"
	AuditLogObjectTypeKnowledgeArticle   = "KNOWLEDGE_BASE_ARTICLE"
	AuditLogObjectTypeKnowledgeBase      = "KNOWLEDGE_BASE"
	AuditLogObjectTypeTheme              = "THEME"
	AuditLogObjectTypeCSS                = "CSS"
	AuditLogObjectTypeJS                 = "JS"
	AuditLogObjectTypeCTA                = "CTA"
	AuditLogObjectTypeFile               = "FILE"
)

// --- URL redirect models ---

// UrlMapping represents a URL redirect (URL mapping).
type UrlMapping struct {
	ID                      string     `json:"id"`
	RoutePrefix             string     `json:"routePrefix"`
	Destination             string     `json:"destination"`
	RedirectStyle           int        `json:"redirectStyle"`
	IsTrailingSlashOptional bool       `json:"isTrailingSlashOptional"`
	IsMatchQueryString      bool       `json:"isMatchQueryString"`
	IsMatchFullUrl          bool       `json:"isMatchFullUrl"`
	IsOnlyAfterNotFound     bool       `json:"isOnlyAfterNotFound"`
	IsPattern               bool       `json:"isPattern"`
	IsProtocolAgnostic      bool       `json:"isProtocolAgnostic"`
	Precedence              int        `json:"precedence"`
	Created                 *time.Time `json:"created,omitempty"`
	Updated                 *time.Time `json:"updated,omitempty"`
}

// UrlMappingCreateRequest is the input for creating a URL redirect.
type UrlMappingCreateRequest struct {
	RoutePrefix             string `json:"routePrefix"`
	Destination             string `json:"destination"`
	RedirectStyle           int    `json:"redirectStyle"`
	IsTrailingSlashOptional *bool  `json:"isTrailingSlashOptional,omitempty"`
	IsMatchQueryString      *bool  `json:"isMatchQueryString,omitempty"`
	IsMatchFullUrl          *bool  `json:"isMatchFullUrl,omitempty"`
	IsOnlyAfterNotFound     *bool  `json:"isOnlyAfterNotFound,omitempty"`
	IsPattern               *bool  `json:"isPattern,omitempty"`
	IsProtocolAgnostic      *bool  `json:"isProtocolAgnostic,omitempty"`
	Precedence              *int   `json:"precedence,omitempty"`
}

// UrlMappingListResult is the paginated response for URL redirects.
type UrlMappingListResult struct {
	Total   int            `json:"total"`
	Results []*UrlMapping  `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// --- Source code models ---

// AssetFileMetadata represents metadata for a source code file.
type AssetFileMetadata struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Folder     bool     `json:"folder"`
	Children   []string `json:"children,omitempty"`
	Hash       string   `json:"hash,omitempty"`
	CreatedAt  int64    `json:"createdAt"`
	UpdatedAt  int64    `json:"updatedAt"`
	ArchivedAt int64    `json:"archivedAt,omitempty"`
}

// FileExtractRequest is the input for async file extraction.
type FileExtractRequest struct {
	Path string `json:"path"`
}

// TaskLocator is the response when starting an async task.
type TaskLocator struct {
	ID    int               `json:"id"`
	Links map[string]string `json:"links,omitempty"`
}

// ActionResponse is the response for checking async task status.
type ActionResponse struct {
	Status      string            `json:"status"`
	RequestedAt *time.Time        `json:"requestedAt,omitempty"`
	StartedAt   *time.Time        `json:"startedAt,omitempty"`
	CompletedAt *time.Time        `json:"completedAt,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
	NumErrors   int               `json:"numErrors,omitempty"`
	Errors      []StandardError   `json:"errors,omitempty"`
}

// --- Site search models ---

// ContentSearchResult represents a single search result.
type ContentSearchResult struct {
	ID               int      `json:"id"`
	Score            float64  `json:"score"`
	Type             string   `json:"type"`
	Domain           string   `json:"domain"`
	URL              string   `json:"url"`
	Title            string   `json:"title,omitempty"`
	Description      string   `json:"description,omitempty"`
	Category         string   `json:"category,omitempty"`
	Subcategory      string   `json:"subcategory,omitempty"`
	AuthorFullName   string   `json:"authorFullName,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	FeaturedImageUrl string   `json:"featuredImageUrl,omitempty"`
	PublishedDate    *int64   `json:"publishedDate,omitempty"`
	CombinedID       string   `json:"combinedId,omitempty"`
	Language         string   `json:"language,omitempty"`
	RowID            *int     `json:"rowId,omitempty"`
	TableID          *int     `json:"tableId,omitempty"`
}

// SearchResults is the paginated response for site search.
type SearchResults struct {
	Total      int                    `json:"total"`
	Offset     int                    `json:"offset"`
	Limit      int                    `json:"limit"`
	Page       int                    `json:"page"`
	SearchTerm string                 `json:"searchTerm,omitempty"`
	Results    []*ContentSearchResult `json:"results"`
}

// IndexedData represents indexed data for a content item.
type IndexedData struct {
	ID     string                   `json:"id"`
	Type   string                   `json:"type"`
	Fields map[string]*IndexedField `json:"fields,omitempty"`
}

// IndexedField represents a single indexed field.
type IndexedField struct {
	Name   string `json:"name,omitempty"`
	Value  any    `json:"value,omitempty"`
	Values []any  `json:"values,omitempty"`
}

// Content type constants used by site search.
const (
	ContentTypeLandingPage      = "LANDING_PAGE"
	ContentTypeBlogPost         = "BLOG_POST"
	ContentTypeSitePage         = "SITE_PAGE"
	ContentTypeKnowledgeArticle = "KNOWLEDGE_ARTICLE"
	ContentTypeListingPage      = "LISTING_PAGE"
)

// --- Performance models ---

// PerformanceView represents performance data for a single time interval.
type PerformanceView struct {
	StartTimestamp        int64   `json:"startTimestamp"`
	EndTimestamp          int64   `json:"endTimestamp"`
	StartDatetime         string  `json:"startDatetime,omitempty"`
	EndDatetime           string  `json:"endDatetime,omitempty"`
	TotalRequests         int     `json:"totalRequests"`
	CacheHits             int     `json:"cacheHits"`
	CacheHitRate          float64 `json:"cacheHitRate"`
	TotalRequestTime      *int    `json:"totalRequestTime,omitempty"`
	AvgOriginResponseTime int     `json:"avgOriginResponseTime"`
	ResponseTimeMs        int     `json:"responseTimeMs"`
	Status100x            int     `json:"100X"`
	Status20x             int     `json:"20X"`
	Status30x             int     `json:"30X"`
	Status40x             int     `json:"40X"`
	Status50x             int     `json:"50X"`
	Status403             int     `json:"403"`
	Status404             int     `json:"404"`
	Status500             int     `json:"500"`
	Status504             int     `json:"504"`
	Percentile50th        int     `json:"50th"`
	Percentile95th        int     `json:"95th"`
	Percentile99th        int     `json:"99th"`
}

// PerformanceResponse is the response for performance data queries.
type PerformanceResponse struct {
	Data          []*PerformanceView `json:"data"`
	Domain        string             `json:"domain,omitempty"`
	Path          string             `json:"path,omitempty"`
	Period        string             `json:"period,omitempty"`
	Interval      string             `json:"interval"`
	StartInterval int64              `json:"startInterval"`
	EndInterval   int64              `json:"endInterval"`
}

// Performance period/interval constants.
const (
	IntervalOneMinute      = "ONE_MINUTE"
	IntervalFiveMinutes    = "FIVE_MINUTES"
	IntervalTenMinutes     = "TEN_MINUTES"
	IntervalFifteenMinutes = "FIFTEEN_MINUTES"
	IntervalThirtyMinutes  = "THIRTY_MINUTES"
	IntervalOneHour        = "ONE_HOUR"
	IntervalFourHours      = "FOUR_HOURS"
	IntervalTwelveHours    = "TWELVE_HOURS"
	IntervalOneDay         = "ONE_DAY"
	IntervalOneWeek        = "ONE_WEEK"
)

// --- Common list options ---

// ListOptions configures a list/getPage request.
type ListOptions struct {
	Limit    int
	After    string
	Sort     []string
	Archived bool
	Property string
}
