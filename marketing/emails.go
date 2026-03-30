package marketing

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const emailsBasePath = "/marketing/v3/emails"
const emailStatisticsBasePath = "/marketing/v3/emails/statistics"

// EmailsService handles operations on HubSpot marketing emails.
type EmailsService struct {
	requester api.Requester
}

// Create creates a new marketing email.
func (s *EmailsService) Create(ctx context.Context, input *EmailCreateRequest) (*PublicEmail, error) {
	var result PublicEmail
	if err := s.requester.Post(ctx, emailsBasePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a marketing email by its ID.
func (s *EmailsService) GetByID(ctx context.Context, emailID string, opts *EmailGetByIDOptions) (*PublicEmail, error) {
	path := fmt.Sprintf("%s/%s", emailsBasePath, emailID)
	q := url.Values{}
	if opts != nil {
		if opts.IncludeStats != nil && *opts.IncludeStats {
			q.Set("includeStats", "true")
		}
		if opts.MarketingCampaignNames != nil && *opts.MarketingCampaignNames {
			q.Set("marketingCampaignNames", "true")
		}
		if opts.WorkflowNames != nil && *opts.WorkflowNames {
			q.Set("workflowNames", "true")
		}
		if len(opts.IncludedProperties) > 0 {
			q.Set("includedProperties", strings.Join(opts.IncludedProperties, ","))
		}
		if opts.Archived != nil && *opts.Archived {
			q.Set("archived", "true")
		}
	}
	var result PublicEmail
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPage retrieves a paginated list of marketing emails.
func (s *EmailsService) GetPage(ctx context.Context, opts *EmailListOptions) (*EmailListResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if len(opts.Sort) > 0 {
			q.Set("sort", strings.Join(opts.Sort, ","))
		}
		if opts.Type != "" {
			q.Set("type", opts.Type)
		}
		if opts.IsPublished != nil {
			q.Set("isPublished", strconv.FormatBool(*opts.IsPublished))
		}
		if opts.Archived != nil && *opts.Archived {
			q.Set("archived", "true")
		}
		if opts.Campaign != "" {
			q.Set("campaign", opts.Campaign)
		}
		if opts.IncludeStats != nil && *opts.IncludeStats {
			q.Set("includeStats", "true")
		}
		if opts.MarketingCampaignNames != nil && *opts.MarketingCampaignNames {
			q.Set("marketingCampaignNames", "true")
		}
		if opts.WorkflowNames != nil && *opts.WorkflowNames {
			q.Set("workflowNames", "true")
		}
		if len(opts.IncludedProperties) > 0 {
			q.Set("includedProperties", strings.Join(opts.IncludedProperties, ","))
		}
		if opts.CreatedAt != "" {
			q.Set("createdAt", opts.CreatedAt)
		}
		if opts.CreatedAfter != "" {
			q.Set("createdAfter", opts.CreatedAfter)
		}
		if opts.CreatedBefore != "" {
			q.Set("createdBefore", opts.CreatedBefore)
		}
		if opts.UpdatedAt != "" {
			q.Set("updatedAt", opts.UpdatedAt)
		}
		if opts.UpdatedAfter != "" {
			q.Set("updatedAfter", opts.UpdatedAfter)
		}
		if opts.UpdatedBefore != "" {
			q.Set("updatedBefore", opts.UpdatedBefore)
		}
	}
	var result EmailListResult
	if err := s.requester.Get(ctx, emailsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update patches an existing marketing email.
func (s *EmailsService) Update(ctx context.Context, emailID string, input *EmailUpdateRequest, archived *bool) (*PublicEmail, error) {
	path := fmt.Sprintf("%s/%s", emailsBasePath, emailID)
	if archived != nil && *archived {
		path += "?archived=true"
	}
	var result PublicEmail
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive soft-deletes a marketing email.
func (s *EmailsService) Archive(ctx context.Context, emailID string) error {
	path := fmt.Sprintf("%s/%s", emailsBasePath, emailID)
	return s.requester.Delete(ctx, path)
}

// Clone duplicates a marketing email.
func (s *EmailsService) Clone(ctx context.Context, input *ContentCloneRequest) (*PublicEmail, error) {
	path := emailsBasePath + "/clone"
	var result PublicEmail
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateAbTestVariation creates an A/B test variation of an email.
func (s *EmailsService) CreateAbTestVariation(ctx context.Context, input *AbTestCreateRequest) (*PublicEmail, error) {
	path := emailsBasePath + "/ab-test-create"
	var result PublicEmail
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAbTestVariation retrieves the A/B test variation of an email.
func (s *EmailsService) GetAbTestVariation(ctx context.Context, emailID string) (*PublicEmail, error) {
	path := fmt.Sprintf("%s/%s/ab-test-variation", emailsBasePath, emailID)
	var result PublicEmail
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetDraft retrieves the draft version of an email.
func (s *EmailsService) GetDraft(ctx context.Context, emailID string) (*PublicEmail, error) {
	path := fmt.Sprintf("%s/%s/draft", emailsBasePath, emailID)
	var result PublicEmail
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpsertDraft creates or updates the draft of an email.
func (s *EmailsService) UpsertDraft(ctx context.Context, emailID string, input *EmailUpdateRequest) (*PublicEmail, error) {
	path := fmt.Sprintf("%s/%s/draft", emailsBasePath, emailID)
	var result PublicEmail
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ResetDraft resets an email's draft to match the live version.
func (s *EmailsService) ResetDraft(ctx context.Context, emailID string) error {
	path := fmt.Sprintf("%s/%s/draft/reset", emailsBasePath, emailID)
	return s.requester.Post(ctx, path, nil, nil)
}

// PublishOrSend publishes or sends a marketing email.
func (s *EmailsService) PublishOrSend(ctx context.Context, emailID string) error {
	path := fmt.Sprintf("%s/%s/publish", emailsBasePath, emailID)
	return s.requester.Post(ctx, path, nil, nil)
}

// UnpublishOrCancel unpublishes or cancels a scheduled email send.
func (s *EmailsService) UnpublishOrCancel(ctx context.Context, emailID string) error {
	path := fmt.Sprintf("%s/%s/unpublish", emailsBasePath, emailID)
	return s.requester.Post(ctx, path, nil, nil)
}

// GetRevisions retrieves the revision history for an email.
func (s *EmailsService) GetRevisions(ctx context.Context, emailID string, opts *EmailRevisionsOptions) (*VersionListResult, error) {
	path := fmt.Sprintf("%s/%s/revisions", emailsBasePath, emailID)
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
	var result VersionListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetRevisionByID retrieves a specific revision of an email.
func (s *EmailsService) GetRevisionByID(ctx context.Context, emailID, revisionID string) (*VersionPublicEmail, error) {
	path := fmt.Sprintf("%s/%s/revisions/%s", emailsBasePath, emailID, revisionID)
	var result VersionPublicEmail
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RestoreDraftRevision restores a draft to a specific revision.
func (s *EmailsService) RestoreDraftRevision(ctx context.Context, emailID string, revisionID int) (*PublicEmail, error) {
	path := fmt.Sprintf("%s/%s/revisions/%d/restore-draft", emailsBasePath, emailID, revisionID)
	var result PublicEmail
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RestoreRevision restores a live email to a previous revision.
func (s *EmailsService) RestoreRevision(ctx context.Context, emailID, revisionID string) error {
	path := fmt.Sprintf("%s/%s/revisions/%s/restore", emailsBasePath, emailID, revisionID)
	return s.requester.Post(ctx, path, nil, nil)
}

// --- Statistics Service ---

// StatisticsService handles marketing email statistics operations.
type StatisticsService struct {
	requester api.Requester
}

// GetEmailsList retrieves aggregate statistics for marketing emails.
func (s *StatisticsService) GetEmailsList(ctx context.Context, opts *EmailStatisticsListOptions) (*AggregateEmailStatistics, error) {
	q := url.Values{}
	if opts != nil {
		if opts.StartTimestamp != "" {
			q.Set("startTimestamp", opts.StartTimestamp)
		}
		if opts.EndTimestamp != "" {
			q.Set("endTimestamp", opts.EndTimestamp)
		}
		if len(opts.EmailIDs) > 0 {
			ids := make([]string, len(opts.EmailIDs))
			for i, id := range opts.EmailIDs {
				ids[i] = strconv.Itoa(id)
			}
			q.Set("emailIds", strings.Join(ids, ","))
		}
		if opts.Property != "" {
			q.Set("property", opts.Property)
		}
	}
	var result AggregateEmailStatistics
	if err := s.requester.Get(ctx, emailStatisticsBasePath+"/list", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetHistogram retrieves statistics broken down by time intervals.
func (s *StatisticsService) GetHistogram(ctx context.Context, opts *EmailStatisticsHistogramOptions) (*EmailStatisticsHistogramResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Interval != "" {
			q.Set("interval", opts.Interval)
		}
		if opts.StartTimestamp != "" {
			q.Set("startTimestamp", opts.StartTimestamp)
		}
		if opts.EndTimestamp != "" {
			q.Set("endTimestamp", opts.EndTimestamp)
		}
		if len(opts.EmailIDs) > 0 {
			ids := make([]string, len(opts.EmailIDs))
			for i, id := range opts.EmailIDs {
				ids[i] = strconv.Itoa(id)
			}
			q.Set("emailIds", strings.Join(ids, ","))
		}
	}
	var result EmailStatisticsHistogramResult
	if err := s.requester.Get(ctx, emailStatisticsBasePath+"/histogram", q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
