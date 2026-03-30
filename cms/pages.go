package cms

import (
	"context"
	"fmt"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const landingPagesBasePath = "/cms/v3/pages/landing-pages"
const sitePagesBasePath = "/cms/v3/pages/site-pages"

// --- PageService (shared between landing pages and site pages) ---

// PageService handles CMS page operations. Used for both landing pages and site pages
// since they share the same API shape at different base paths.
type PageService struct {
	requester api.Requester
	basePath  string
}

// Create creates a new page.
func (s *PageService) Create(ctx context.Context, page *Page) (*Page, error) {
	var result Page
	if err := s.requester.Post(ctx, s.basePath, page, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a page by its ID.
func (s *PageService) GetByID(ctx context.Context, id string, archived bool) (*Page, error) {
	path := fmt.Sprintf("%s/%s", s.basePath, id)
	q := buildListQuery(&ListOptions{Archived: archived})
	var result Page
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing page.
func (s *PageService) Update(ctx context.Context, id string, page *Page) (*Page, error) {
	path := fmt.Sprintf("%s/%s", s.basePath, id)
	var result Page
	if err := s.requester.Patch(ctx, path, page, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive soft-deletes a page.
func (s *PageService) Archive(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s", s.basePath, id)
	return s.requester.Delete(ctx, path)
}

// List retrieves a page of pages.
func (s *PageService) List(ctx context.Context, opts *ListOptions) (*PageListResult, error) {
	q := buildListQuery(opts)
	var result PageListResult
	if err := s.requester.Get(ctx, s.basePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Clone creates a clone of a page.
func (s *PageService) Clone(ctx context.Context, input *ContentCloneRequest) (*Page, error) {
	path := s.basePath + "/clone"
	var result Page
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDraftByID retrieves the draft version of a page.
func (s *PageService) GetDraftByID(ctx context.Context, id string) (*Page, error) {
	path := fmt.Sprintf("%s/%s/draft", s.basePath, id)
	var result Page
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDraft updates the draft version of a page.
func (s *PageService) UpdateDraft(ctx context.Context, id string, page *Page) (*Page, error) {
	path := fmt.Sprintf("%s/%s/draft", s.basePath, id)
	var result Page
	if err := s.requester.Patch(ctx, path, page, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PushLive publishes the draft of a page to live.
func (s *PageService) PushLive(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s/push-live", s.basePath, id)
	return s.requester.Post(ctx, path, nil, nil)
}

// ResetDraft discards the draft changes for a page.
func (s *PageService) ResetDraft(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s/reset-draft", s.basePath, id)
	return s.requester.Post(ctx, path, nil, nil)
}

// Schedule schedules a page for publishing.
func (s *PageService) Schedule(ctx context.Context, input *ContentScheduleRequest) error {
	path := s.basePath + "/schedule"
	return s.requester.Post(ctx, path, input, nil)
}

// GetPreviousVersion retrieves a specific previous version.
func (s *PageService) GetPreviousVersion(ctx context.Context, id, revisionID string) (*VersionPage, error) {
	path := fmt.Sprintf("%s/%s/revisions/%s", s.basePath, id, revisionID)
	var result VersionPage
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPreviousVersions lists previous versions.
func (s *PageService) GetPreviousVersions(ctx context.Context, id string, opts *VersionListOptions) (*VersionPageListResult, error) {
	path := fmt.Sprintf("%s/%s/revisions", s.basePath, id)
	q := buildVersionListQuery(opts)
	var result VersionPageListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RestorePreviousVersion restores a previous version.
func (s *PageService) RestorePreviousVersion(ctx context.Context, id, revisionID string) (*Page, error) {
	path := fmt.Sprintf("%s/%s/revisions/%s/restore", s.basePath, id, revisionID)
	var result Page
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RestorePreviousVersionToDraft restores a previous version to the draft.
func (s *PageService) RestorePreviousVersionToDraft(ctx context.Context, id string, revisionID int64) (*Page, error) {
	path := fmt.Sprintf("%s/%s/revisions/%d/restore-to-draft", s.basePath, id, revisionID)
	var result Page
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Batch operations ---

// BatchCreate creates multiple pages.
func (s *PageService) BatchCreate(ctx context.Context, input *BatchInputPage) (*BatchResponsePage, error) {
	path := s.basePath + "/batch/create"
	var result BatchResponsePage
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchRead retrieves multiple pages by ID.
func (s *PageService) BatchRead(ctx context.Context, input *BatchInputString) (*BatchResponsePage, error) {
	path := s.basePath + "/batch/read"
	var result BatchResponsePage
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchUpdate updates multiple pages.
func (s *PageService) BatchUpdate(ctx context.Context, input *BatchInputJsonNode) (*BatchResponsePage, error) {
	path := s.basePath + "/batch/update"
	var result BatchResponsePage
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchArchive archives multiple pages.
func (s *PageService) BatchArchive(ctx context.Context, input *BatchInputString) error {
	path := s.basePath + "/batch/archive"
	return s.requester.Post(ctx, path, input, nil)
}

// --- A/B test operations ---

// CreateABTestVariation creates an A/B test variation for a page.
func (s *PageService) CreateABTestVariation(ctx context.Context, input *ABTestCreateRequest) (*Page, error) {
	path := s.basePath + "/ab-test/create-variation"
	var result Page
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// EndActiveABTest ends an active A/B test.
func (s *PageService) EndActiveABTest(ctx context.Context, input *ABTestEndRequest) error {
	path := s.basePath + "/ab-test/end"
	return s.requester.Post(ctx, path, input, nil)
}

// RerunPreviousABTest reruns a previous A/B test.
func (s *PageService) RerunPreviousABTest(ctx context.Context, input *ABTestRerunRequest) error {
	path := s.basePath + "/ab-test/rerun"
	return s.requester.Post(ctx, path, input, nil)
}

// --- Multi-language operations ---

// AttachToLangGroup attaches a page to a multi-language group.
func (s *PageService) AttachToLangGroup(ctx context.Context, input *AttachToLangPrimaryRequest) error {
	path := s.basePath + "/multi-language/attach-to-lang-group"
	return s.requester.Post(ctx, path, input, nil)
}

// DetachFromLangGroup detaches a page from a multi-language group.
func (s *PageService) DetachFromLangGroup(ctx context.Context, input *DetachFromLangGroupRequest) error {
	path := s.basePath + "/multi-language/detach-from-lang-group"
	return s.requester.Post(ctx, path, input, nil)
}

// CreateLangVariation creates a new language variation of a page.
func (s *PageService) CreateLangVariation(ctx context.Context, input *LanguageCloneRequest) (*Page, error) {
	path := s.basePath + "/multi-language/create-language-variation"
	var result Page
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SetLangPrimary sets a new primary language for a page group.
func (s *PageService) SetLangPrimary(ctx context.Context, input *SetNewLanguagePrimaryRequest) error {
	path := s.basePath + "/multi-language/set-new-lang-primary"
	return s.requester.Post(ctx, path, input, nil)
}

// UpdateLangs updates language settings for pages.
func (s *PageService) UpdateLangs(ctx context.Context, input *UpdateLanguagesRequest) error {
	path := s.basePath + "/multi-language/update-languages"
	return s.requester.Post(ctx, path, input, nil)
}

// --- Folder operations (landing pages only in Node client, but same pattern) ---

// CreateFolder creates a content folder.
func (s *PageService) CreateFolder(ctx context.Context, folder *ContentFolder) (*ContentFolder, error) {
	path := s.basePath + "/folders"
	var result ContentFolder
	if err := s.requester.Post(ctx, path, folder, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFolderByID retrieves a content folder by ID.
func (s *PageService) GetFolderByID(ctx context.Context, id string) (*ContentFolder, error) {
	path := fmt.Sprintf("%s/folders/%s", s.basePath, id)
	var result ContentFolder
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateFolder updates a content folder.
func (s *PageService) UpdateFolder(ctx context.Context, id string, folder *ContentFolder) (*ContentFolder, error) {
	path := fmt.Sprintf("%s/folders/%s", s.basePath, id)
	var result ContentFolder
	if err := s.requester.Patch(ctx, path, folder, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchiveFolder deletes a content folder.
func (s *PageService) ArchiveFolder(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/folders/%s", s.basePath, id)
	return s.requester.Delete(ctx, path)
}

// ListFolders retrieves content folders.
func (s *PageService) ListFolders(ctx context.Context, opts *ListOptions) (*ContentFolderListResult, error) {
	path := s.basePath + "/folders"
	q := buildListQuery(opts)
	var result ContentFolderListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
