package cms

import (
	"context"
	"net/url"
	"strconv"
	"strings"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const auditLogsBasePath = "/cms/v3/audit-logs"

// AuditLogsService handles CMS audit log queries.
type AuditLogsService struct {
	requester api.Requester
}

// AuditLogListOptions configures an audit log query.
type AuditLogListOptions struct {
	UserID     []string
	EventType  []string
	ObjectType []string
	ObjectID   []string
	After      string
	Before     string
	Limit      int
	Sort       []string
}

// List retrieves a page of audit log entries.
func (s *AuditLogsService) List(ctx context.Context, opts *AuditLogListOptions) (*AuditLogListResult, error) {
	q := url.Values{}
	if opts != nil {
		for _, v := range opts.UserID {
			q.Add("userId", v)
		}
		for _, v := range opts.EventType {
			q.Add("eventType", v)
		}
		for _, v := range opts.ObjectType {
			q.Add("objectType", v)
		}
		for _, v := range opts.ObjectID {
			q.Add("objectId", v)
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
	}
	var result AuditLogListResult
	if err := s.requester.Get(ctx, auditLogsBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
