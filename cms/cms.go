package cms

import (
	"github.com/scopiousdigital/hubspot-go/internal/api"
)

// Service provides access to all HubSpot CMS APIs.
type Service struct {
	requester api.Requester

	// Blogs
	BlogPosts   *BlogPostsService
	BlogAuthors *BlogAuthorsService
	BlogTags    *BlogTagsService

	// HubDB
	Tables    *TablesService
	Rows      *RowsService
	RowsBatch *RowsBatchService

	// Content management
	Domains      *DomainsService
	LandingPages *PageService
	SitePages    *PageService
	AuditLogs    *AuditLogsService
	UrlRedirects *UrlRedirectsService

	// Source code
	SourceCode *SourceCodeService

	// Search and performance
	SiteSearch  *SiteSearchService
	Performance *PerformanceService
}

// NewService creates a new CMS service. Called by the root hubspot package.
func NewService(r api.Requester) *Service {
	return &Service{
		requester: r,

		// Blogs
		BlogPosts:   &BlogPostsService{requester: r},
		BlogAuthors: &BlogAuthorsService{requester: r},
		BlogTags:    &BlogTagsService{requester: r},

		// HubDB
		Tables:    &TablesService{requester: r},
		Rows:      &RowsService{requester: r},
		RowsBatch: &RowsBatchService{requester: r},

		// Content management
		Domains:      &DomainsService{requester: r},
		LandingPages: &PageService{requester: r, basePath: landingPagesBasePath},
		SitePages:    &PageService{requester: r, basePath: sitePagesBasePath},
		AuditLogs:    &AuditLogsService{requester: r},
		UrlRedirects: &UrlRedirectsService{requester: r},

		// Source code
		SourceCode: &SourceCodeService{
			Content:    &SourceCodeContentService{requester: r},
			Extract:    &SourceCodeExtractService{requester: r},
			Metadata:   &SourceCodeMetadataService{requester: r},
			Validation: &SourceCodeValidationService{requester: r},
		},

		// Search and performance
		SiteSearch:  &SiteSearchService{requester: r},
		Performance: &PerformanceService{requester: r},
	}
}
