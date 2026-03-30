package cms

import (
	"context"
	"net/url"
	"strconv"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const performanceBasePath = "/cms/v3/performance"

// PerformanceService handles CMS performance data queries.
type PerformanceService struct {
	requester api.Requester
}

// PerformanceOptions configures a performance data query.
type PerformanceOptions struct {
	Domain   string
	Path     string
	Pad      *bool
	Sum      *bool
	Period   string
	Interval string
	Start    *int64
	End      *int64
}

func buildPerformanceQuery(opts *PerformanceOptions) url.Values {
	q := url.Values{}
	if opts == nil {
		return q
	}
	if opts.Domain != "" {
		q.Set("domain", opts.Domain)
	}
	if opts.Path != "" {
		q.Set("path", opts.Path)
	}
	if opts.Pad != nil {
		q.Set("pad", strconv.FormatBool(*opts.Pad))
	}
	if opts.Sum != nil {
		q.Set("sum", strconv.FormatBool(*opts.Sum))
	}
	if opts.Period != "" {
		q.Set("period", opts.Period)
	}
	if opts.Interval != "" {
		q.Set("interval", opts.Interval)
	}
	if opts.Start != nil {
		q.Set("start", strconv.FormatInt(*opts.Start, 10))
	}
	if opts.End != nil {
		q.Set("end", strconv.FormatInt(*opts.End, 10))
	}
	return q
}

// GetPage retrieves performance data for the site.
func (s *PerformanceService) GetPage(ctx context.Context, opts *PerformanceOptions) (*PerformanceResponse, error) {
	q := buildPerformanceQuery(opts)
	var result PerformanceResponse
	if err := s.requester.Get(ctx, performanceBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetUptime retrieves uptime performance data for the site.
func (s *PerformanceService) GetUptime(ctx context.Context, opts *PerformanceOptions) (*PerformanceResponse, error) {
	path := performanceBasePath + "/uptime"
	q := buildPerformanceQuery(opts)
	var result PerformanceResponse
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
