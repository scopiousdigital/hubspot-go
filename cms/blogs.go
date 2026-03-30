package cms

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const blogPostsBasePath = "/cms/v3/blogs/posts"
const blogAuthorsBasePath = "/cms/v3/blogs/authors"
const blogTagsBasePath = "/cms/v3/blogs/tags"

// --- BlogPostsService ---

// BlogPostsService handles blog post CRUD, batch, and multi-language operations.
type BlogPostsService struct {
	requester api.Requester
}

// Create creates a new blog post.
func (s *BlogPostsService) Create(ctx context.Context, post *BlogPost) (*BlogPost, error) {
	var result BlogPost
	if err := s.requester.Post(ctx, blogPostsBasePath, post, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a blog post by its ID.
func (s *BlogPostsService) GetByID(ctx context.Context, id string, archived bool) (*BlogPost, error) {
	path := fmt.Sprintf("%s/%s", blogPostsBasePath, id)
	q := url.Values{}
	if archived {
		q.Set("archived", "true")
	}
	var result BlogPost
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing blog post.
func (s *BlogPostsService) Update(ctx context.Context, id string, post *BlogPost) (*BlogPost, error) {
	path := fmt.Sprintf("%s/%s", blogPostsBasePath, id)
	var result BlogPost
	if err := s.requester.Patch(ctx, path, post, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive soft-deletes a blog post.
func (s *BlogPostsService) Archive(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s", blogPostsBasePath, id)
	return s.requester.Delete(ctx, path)
}

// List retrieves a page of blog posts.
func (s *BlogPostsService) List(ctx context.Context, opts *ListOptions) (*BlogPostListResult, error) {
	q := buildListQuery(opts)
	var result BlogPostListResult
	if err := s.requester.Get(ctx, blogPostsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Clone creates a clone of a blog post.
func (s *BlogPostsService) Clone(ctx context.Context, input *ContentCloneRequest) (*BlogPost, error) {
	path := blogPostsBasePath + "/clone"
	var result BlogPost
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDraftByID retrieves the draft version of a blog post.
func (s *BlogPostsService) GetDraftByID(ctx context.Context, id string) (*BlogPost, error) {
	path := fmt.Sprintf("%s/%s/draft", blogPostsBasePath, id)
	var result BlogPost
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateDraft updates the draft version of a blog post.
func (s *BlogPostsService) UpdateDraft(ctx context.Context, id string, post *BlogPost) (*BlogPost, error) {
	path := fmt.Sprintf("%s/%s/draft", blogPostsBasePath, id)
	var result BlogPost
	if err := s.requester.Patch(ctx, path, post, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PushLive publishes the draft of a blog post to live.
func (s *BlogPostsService) PushLive(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s/push-live", blogPostsBasePath, id)
	return s.requester.Post(ctx, path, nil, nil)
}

// ResetDraft discards the draft changes for a blog post.
func (s *BlogPostsService) ResetDraft(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s/reset-draft", blogPostsBasePath, id)
	return s.requester.Post(ctx, path, nil, nil)
}

// Schedule schedules a blog post for publishing.
func (s *BlogPostsService) Schedule(ctx context.Context, input *ContentScheduleRequest) error {
	path := blogPostsBasePath + "/schedule"
	return s.requester.Post(ctx, path, input, nil)
}

// GetPreviousVersion retrieves a specific previous version of a blog post.
func (s *BlogPostsService) GetPreviousVersion(ctx context.Context, id, revisionID string) (*VersionBlogPost, error) {
	path := fmt.Sprintf("%s/%s/revisions/%s", blogPostsBasePath, id, revisionID)
	var result VersionBlogPost
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPreviousVersions lists previous versions of a blog post.
func (s *BlogPostsService) GetPreviousVersions(ctx context.Context, id string, opts *VersionListOptions) (*VersionListResult, error) {
	path := fmt.Sprintf("%s/%s/revisions", blogPostsBasePath, id)
	q := buildVersionListQuery(opts)
	var result VersionListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RestorePreviousVersion restores a previous version of a blog post.
func (s *BlogPostsService) RestorePreviousVersion(ctx context.Context, id, revisionID string) (*BlogPost, error) {
	path := fmt.Sprintf("%s/%s/revisions/%s/restore", blogPostsBasePath, id, revisionID)
	var result BlogPost
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RestorePreviousVersionToDraft restores a previous version to the draft.
func (s *BlogPostsService) RestorePreviousVersionToDraft(ctx context.Context, id string, revisionID int64) (*BlogPost, error) {
	path := fmt.Sprintf("%s/%s/revisions/%d/restore-to-draft", blogPostsBasePath, id, revisionID)
	var result BlogPost
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Blog post batch operations ---

// BatchCreate creates multiple blog posts.
func (s *BlogPostsService) BatchCreate(ctx context.Context, input *BatchInputBlogPost) (*BatchResponseBlogPost, error) {
	path := blogPostsBasePath + "/batch/create"
	var result BatchResponseBlogPost
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchRead retrieves multiple blog posts by ID.
func (s *BlogPostsService) BatchRead(ctx context.Context, input *BatchInputString) (*BatchResponseBlogPost, error) {
	path := blogPostsBasePath + "/batch/read"
	var result BatchResponseBlogPost
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchUpdate updates multiple blog posts.
func (s *BlogPostsService) BatchUpdate(ctx context.Context, input *BatchInputJsonNode) (*BatchResponseBlogPost, error) {
	path := blogPostsBasePath + "/batch/update"
	var result BatchResponseBlogPost
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchArchive archives multiple blog posts.
func (s *BlogPostsService) BatchArchive(ctx context.Context, input *BatchInputString) error {
	path := blogPostsBasePath + "/batch/archive"
	return s.requester.Post(ctx, path, input, nil)
}

// --- Blog post multi-language operations ---

// AttachToLangGroup attaches a blog post to a multi-language group.
func (s *BlogPostsService) AttachToLangGroup(ctx context.Context, input *AttachToLangPrimaryRequest) error {
	path := blogPostsBasePath + "/multi-language/attach-to-lang-group"
	return s.requester.Post(ctx, path, input, nil)
}

// DetachFromLangGroup detaches a blog post from a multi-language group.
func (s *BlogPostsService) DetachFromLangGroup(ctx context.Context, input *DetachFromLangGroupRequest) error {
	path := blogPostsBasePath + "/multi-language/detach-from-lang-group"
	return s.requester.Post(ctx, path, input, nil)
}

// CreateLangVariation creates a new language variation of a blog post.
func (s *BlogPostsService) CreateLangVariation(ctx context.Context, input *LanguageCloneRequest) (*BlogPost, error) {
	path := blogPostsBasePath + "/multi-language/create-language-variation"
	var result BlogPost
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SetLangPrimary sets a new primary language for a blog post group.
func (s *BlogPostsService) SetLangPrimary(ctx context.Context, input *SetNewLanguagePrimaryRequest) error {
	path := blogPostsBasePath + "/multi-language/set-new-lang-primary"
	return s.requester.Post(ctx, path, input, nil)
}

// UpdateLangs updates language settings for blog posts.
func (s *BlogPostsService) UpdateLangs(ctx context.Context, input *UpdateLanguagesRequest) error {
	path := blogPostsBasePath + "/multi-language/update-languages"
	return s.requester.Post(ctx, path, input, nil)
}

// --- BlogAuthorsService ---

// BlogAuthorsService handles blog author CRUD and batch operations.
type BlogAuthorsService struct {
	requester api.Requester
}

// Create creates a new blog author.
func (s *BlogAuthorsService) Create(ctx context.Context, author *BlogAuthor) (*BlogAuthor, error) {
	var result BlogAuthor
	if err := s.requester.Post(ctx, blogAuthorsBasePath, author, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a blog author by ID.
func (s *BlogAuthorsService) GetByID(ctx context.Context, id string, archived bool) (*BlogAuthor, error) {
	path := fmt.Sprintf("%s/%s", blogAuthorsBasePath, id)
	q := url.Values{}
	if archived {
		q.Set("archived", "true")
	}
	var result BlogAuthor
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing blog author.
func (s *BlogAuthorsService) Update(ctx context.Context, id string, author *BlogAuthor) (*BlogAuthor, error) {
	path := fmt.Sprintf("%s/%s", blogAuthorsBasePath, id)
	var result BlogAuthor
	if err := s.requester.Patch(ctx, path, author, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive soft-deletes a blog author.
func (s *BlogAuthorsService) Archive(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s", blogAuthorsBasePath, id)
	return s.requester.Delete(ctx, path)
}

// List retrieves a page of blog authors.
func (s *BlogAuthorsService) List(ctx context.Context, opts *ListOptions) (*BlogAuthorListResult, error) {
	q := buildListQuery(opts)
	var result BlogAuthorListResult
	if err := s.requester.Get(ctx, blogAuthorsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchCreate creates multiple blog authors.
func (s *BlogAuthorsService) BatchCreate(ctx context.Context, input *BatchInputBlogAuthor) (*BatchResponseBlogAuthor, error) {
	path := blogAuthorsBasePath + "/batch/create"
	var result BatchResponseBlogAuthor
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchRead retrieves multiple blog authors by ID.
func (s *BlogAuthorsService) BatchRead(ctx context.Context, input *BatchInputString) (*BatchResponseBlogAuthor, error) {
	path := blogAuthorsBasePath + "/batch/read"
	var result BatchResponseBlogAuthor
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchUpdate updates multiple blog authors.
func (s *BlogAuthorsService) BatchUpdate(ctx context.Context, input *BatchInputJsonNode) (*BatchResponseBlogAuthor, error) {
	path := blogAuthorsBasePath + "/batch/update"
	var result BatchResponseBlogAuthor
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchArchive archives multiple blog authors.
func (s *BlogAuthorsService) BatchArchive(ctx context.Context, input *BatchInputString) error {
	path := blogAuthorsBasePath + "/batch/archive"
	return s.requester.Post(ctx, path, input, nil)
}

// AttachToLangGroup attaches a blog author to a multi-language group.
func (s *BlogAuthorsService) AttachToLangGroup(ctx context.Context, input *AttachToLangPrimaryRequest) error {
	path := blogAuthorsBasePath + "/multi-language/attach-to-lang-group"
	return s.requester.Post(ctx, path, input, nil)
}

// DetachFromLangGroup detaches a blog author from a multi-language group.
func (s *BlogAuthorsService) DetachFromLangGroup(ctx context.Context, input *DetachFromLangGroupRequest) error {
	path := blogAuthorsBasePath + "/multi-language/detach-from-lang-group"
	return s.requester.Post(ctx, path, input, nil)
}

// CreateLangVariation creates a new language variation of a blog author.
func (s *BlogAuthorsService) CreateLangVariation(ctx context.Context, input *LanguageCloneRequest) (*BlogAuthor, error) {
	path := blogAuthorsBasePath + "/multi-language/create-language-variation"
	var result BlogAuthor
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SetLangPrimary sets a new primary language for a blog author group.
func (s *BlogAuthorsService) SetLangPrimary(ctx context.Context, input *SetNewLanguagePrimaryRequest) error {
	path := blogAuthorsBasePath + "/multi-language/set-new-lang-primary"
	return s.requester.Post(ctx, path, input, nil)
}

// UpdateLangs updates language settings for blog authors.
func (s *BlogAuthorsService) UpdateLangs(ctx context.Context, input *UpdateLanguagesRequest) error {
	path := blogAuthorsBasePath + "/multi-language/update-languages"
	return s.requester.Post(ctx, path, input, nil)
}

// --- BlogTagsService ---

// BlogTagsService handles blog tag CRUD and batch operations.
type BlogTagsService struct {
	requester api.Requester
}

// Create creates a new blog tag.
func (s *BlogTagsService) Create(ctx context.Context, tag *BlogTag) (*BlogTag, error) {
	var result BlogTag
	if err := s.requester.Post(ctx, blogTagsBasePath, tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a blog tag by ID.
func (s *BlogTagsService) GetByID(ctx context.Context, id string, archived bool) (*BlogTag, error) {
	path := fmt.Sprintf("%s/%s", blogTagsBasePath, id)
	q := url.Values{}
	if archived {
		q.Set("archived", "true")
	}
	var result BlogTag
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing blog tag.
func (s *BlogTagsService) Update(ctx context.Context, id string, tag *BlogTag) (*BlogTag, error) {
	path := fmt.Sprintf("%s/%s", blogTagsBasePath, id)
	var result BlogTag
	if err := s.requester.Patch(ctx, path, tag, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive soft-deletes a blog tag.
func (s *BlogTagsService) Archive(ctx context.Context, id string) error {
	path := fmt.Sprintf("%s/%s", blogTagsBasePath, id)
	return s.requester.Delete(ctx, path)
}

// List retrieves a page of blog tags.
func (s *BlogTagsService) List(ctx context.Context, opts *ListOptions) (*BlogTagListResult, error) {
	q := buildListQuery(opts)
	var result BlogTagListResult
	if err := s.requester.Get(ctx, blogTagsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchCreate creates multiple blog tags.
func (s *BlogTagsService) BatchCreate(ctx context.Context, input *BatchInputBlogTag) (*BatchResponseBlogTag, error) {
	path := blogTagsBasePath + "/batch/create"
	var result BatchResponseBlogTag
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchRead retrieves multiple blog tags by ID.
func (s *BlogTagsService) BatchRead(ctx context.Context, input *BatchInputString) (*BatchResponseBlogTag, error) {
	path := blogTagsBasePath + "/batch/read"
	var result BatchResponseBlogTag
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchUpdate updates multiple blog tags.
func (s *BlogTagsService) BatchUpdate(ctx context.Context, input *BatchInputJsonNode) (*BatchResponseBlogTag, error) {
	path := blogTagsBasePath + "/batch/update"
	var result BatchResponseBlogTag
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchArchive archives multiple blog tags.
func (s *BlogTagsService) BatchArchive(ctx context.Context, input *BatchInputString) error {
	path := blogTagsBasePath + "/batch/archive"
	return s.requester.Post(ctx, path, input, nil)
}

// AttachToLangGroup attaches a blog tag to a multi-language group.
func (s *BlogTagsService) AttachToLangGroup(ctx context.Context, input *AttachToLangPrimaryRequest) error {
	path := blogTagsBasePath + "/multi-language/attach-to-lang-group"
	return s.requester.Post(ctx, path, input, nil)
}

// DetachFromLangGroup detaches a blog tag from a multi-language group.
func (s *BlogTagsService) DetachFromLangGroup(ctx context.Context, input *DetachFromLangGroupRequest) error {
	path := blogTagsBasePath + "/multi-language/detach-from-lang-group"
	return s.requester.Post(ctx, path, input, nil)
}

// CreateLangVariation creates a new language variation of a blog tag.
func (s *BlogTagsService) CreateLangVariation(ctx context.Context, input *LanguageCloneRequest) (*BlogTag, error) {
	path := blogTagsBasePath + "/multi-language/create-language-variation"
	var result BlogTag
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SetLangPrimary sets a new primary language for a blog tag group.
func (s *BlogTagsService) SetLangPrimary(ctx context.Context, input *SetNewLanguagePrimaryRequest) error {
	path := blogTagsBasePath + "/multi-language/set-new-lang-primary"
	return s.requester.Post(ctx, path, input, nil)
}

// UpdateLangs updates language settings for blog tags.
func (s *BlogTagsService) UpdateLangs(ctx context.Context, input *UpdateLanguagesRequest) error {
	path := blogTagsBasePath + "/multi-language/update-languages"
	return s.requester.Post(ctx, path, input, nil)
}

// --- Helpers ---

// VersionListOptions configures a version list request.
type VersionListOptions struct {
	After  string
	Before string
	Limit  int
}

func buildListQuery(opts *ListOptions) url.Values {
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
	if opts.Property != "" {
		q.Set("property", opts.Property)
	}
	return q
}

func buildVersionListQuery(opts *VersionListOptions) url.Values {
	q := url.Values{}
	if opts == nil {
		return q
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
	return q
}
