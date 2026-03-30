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

// values encodes the search options as URL query parameters.
func (o *SiteSearchOptions) values() url.Values {
	q := url.Values{}
	if o == nil {
		return q
	}
	if o.Query != "" {
		q.Set("q", o.Query)
	}
	if o.Limit > 0 {
		q.Set("limit", strconv.Itoa(o.Limit))
	}
	if o.Offset > 0 {
		q.Set("offset", strconv.Itoa(o.Offset))
	}
	if o.Language != "" {
		q.Set("language", o.Language)
	}
	if o.MatchPrefix != nil {
		q.Set("matchPrefix", strconv.FormatBool(*o.MatchPrefix))
	}
	if o.Autocomplete != nil {
		q.Set("autocomplete", strconv.FormatBool(*o.Autocomplete))
	}
	if o.PopularityBoost != nil {
		q.Set("popularityBoost", strconv.FormatFloat(*o.PopularityBoost, 'f', -1, 64))
	}
	if o.BoostLimit != nil {
		q.Set("boostLimit", strconv.FormatFloat(*o.BoostLimit, 'f', -1, 64))
	}
	if o.BoostRecent != "" {
		q.Set("boostRecent", o.BoostRecent)
	}
	if o.TableID != nil {
		q.Set("tableId", strconv.Itoa(*o.TableID))
	}
	if o.HubDBQuery != "" {
		q.Set("hubdbQuery", o.HubDBQuery)
	}
	for _, d := range o.Domain {
		q.Add("domain", d)
	}
	for _, t := range o.Type {
		q.Add("type", t)
	}
	for _, p := range o.PathPrefix {
		q.Add("pathPrefix", p)
	}
	for _, p := range o.Property {
		q.Add("property", p)
	}
	if o.Length != "" {
		q.Set("length", o.Length)
	}
	for _, g := range o.GroupID {
		q.Add("groupId", strconv.Itoa(g))
	}
	return q
}

// Search performs a site search.
func (s *SiteSearchService) Search(ctx context.Context, opts *SiteSearchOptions) (*SearchResults, error) {
	var result SearchResults
	if err := s.requester.Get(ctx, siteSearchBasePath+"/search", opts.values(), &result); err != nil {
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
