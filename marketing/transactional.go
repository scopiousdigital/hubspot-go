package marketing

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const transactionalBasePath = "/marketing/v3/transactional"

// TransactionalService handles operations on HubSpot transactional emails
// and SMTP API tokens.
type TransactionalService struct {
	requester api.Requester
}

// --- Single Send API ---

// SendEmail sends a single transactional email.
func (s *TransactionalService) SendEmail(ctx context.Context, input *SingleSendRequest) (*EmailSendStatusView, error) {
	path := transactionalBasePath + "/single-email/send"
	var result EmailSendStatusView
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- Public SMTP Tokens API ---

// CreateToken creates a new SMTP API token.
func (s *TransactionalService) CreateToken(ctx context.Context, input *SmtpApiTokenCreateRequest) (*SmtpApiTokenView, error) {
	path := transactionalBasePath + "/smtp-tokens"
	var result SmtpApiTokenView
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTokenByID retrieves an SMTP API token by its ID.
func (s *TransactionalService) GetTokenByID(ctx context.Context, tokenID string) (*SmtpApiTokenView, error) {
	path := fmt.Sprintf("%s/smtp-tokens/%s", transactionalBasePath, tokenID)
	var result SmtpApiTokenView
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetTokensPage retrieves a paginated list of SMTP tokens.
func (s *TransactionalService) GetTokensPage(ctx context.Context, opts *SmtpTokenListOptions) (*SmtpTokenListResult, error) {
	path := transactionalBasePath + "/smtp-tokens"
	q := url.Values{}
	if opts != nil {
		if opts.CampaignName != "" {
			q.Set("campaignName", opts.CampaignName)
		}
		if opts.EmailCampaignID != "" {
			q.Set("emailCampaignId", opts.EmailCampaignID)
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
	}
	var result SmtpTokenListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchiveToken deletes an SMTP API token.
func (s *TransactionalService) ArchiveToken(ctx context.Context, tokenID string) error {
	path := fmt.Sprintf("%s/smtp-tokens/%s", transactionalBasePath, tokenID)
	return s.requester.Delete(ctx, path)
}

// ResetPassword resets the password for an SMTP API token.
func (s *TransactionalService) ResetPassword(ctx context.Context, tokenID string) (*SmtpApiTokenView, error) {
	path := fmt.Sprintf("%s/smtp-tokens/%s/password-reset", transactionalBasePath, tokenID)
	var result SmtpApiTokenView
	if err := s.requester.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
