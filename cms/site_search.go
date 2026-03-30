package cms

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const siteSearchBasePath = "/cms/v3/site-search"

// SiteSearchService handles CMS site search operations.
type SiteSearchService struct {
	requester api.Requester
}

// SiteSearchOptions configures a site search query.
type SiteSearchOptions struct {
	Query           string
	Limit           int
	Offset          int
	Language        string
	MatchPrefix     *bool
	Autocomplete    *bool
	PopularityBoost *float64
	BoostLimit      *float64
	BoostRecent     string
	TableID         *int
	HubDBQuery      string
	Domain          []string
	Type            []string
	PathPrefix      []string
	Property        []string
	Length          string // "SHORT" or "LONG"
	GroupID         []int
}

// Search performs a site search.
func (s *SiteSearchService) Search(ctx context.Context, opts *SiteSearchOptions) (*SearchResults, error) {
	path := siteSearchBasePath + "/search"
	q := url.Values{}
	if opts != nil {
		if opts.Query != "" {
			q.Set("q", opts.Query)
		}
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.Offset > 0 {
			q.Set("offset", strconv.Itoa(opts.Offset))
		}
		if opts.Language != "" {
			q.Set("language", opts.Language)
		}
		if opts.MatchPrefix != nil {
			q.Set("matchPrefix", strconv.FormatBool(*opts.MatchPrefix))
		}
		if opts.Autocomplete != nil {
			q.Set("autocomplete", strconv.FormatBool(*opts.Autocomplete))
		}
		if opts.PopularityBoost != nil {
			q.Set("popularityBoost", strconv.FormatFloat(*opts.PopularityBoost, 'f', -1, 64))
		}
		if opts.BoostLimit != nil {
			q.Set("boostLimit", strconv.FormatFloat(*opts.BoostLimit, 'f', -1, 64))
		}
		if opts.BoostRecent != "" {
			q.Set("boostRecent", opts.BoostRecent)
		}
		if opts.TableID != nil {
			q.Set("tableId", strconv.Itoa(*opts.TableID))
		}
		if opts.HubDBQuery != "" {
			q.Set("hubdbQuery", opts.HubDBQuery)
		}
		for _, d := range opts.Domain {
			q.Add("domain", d)
		}
		for _, t := range opts.Type {
			q.Add("type", t)
		}
		for _, p := range opts.PathPrefix {
			q.Add("pathPrefix", p)
		}
		for _, p := range opts.Property {
			q.Add("property", p)
		}
		if opts.Length != "" {
			q.Set("length", opts.Length)
		}
		for _, g := range opts.GroupID {
			q.Add("groupId", strconv.Itoa(g))
		}
	}
	var result SearchResults
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves indexed data for a specific content item.
func (s *SiteSearchService) GetByID(ctx context.Context, contentID string, contentType string) (*IndexedData, error) {
	path := fmt.Sprintf("%s/indexed/%s", siteSearchBasePath, contentID)
	q := url.Values{}
	if contentType != "" {
		q.Set("type", contentType)
	}
	var result IndexedData
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Search result length constants.
const (
	SearchLengthShort = "SHORT"
	SearchLengthLong  = "LONG"
)
