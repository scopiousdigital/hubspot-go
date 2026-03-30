package cms_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/scopiousdigital/hubspot-go/cms"
)

// mockRequester implements api.Requester by routing requests to a test server.
type mockRequester struct {
	server *httptest.Server
}

func (m *mockRequester) Get(ctx context.Context, path string, query url.Values, result any) error {
	u := m.server.URL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	return m.doRequest(req, result)
}

func (m *mockRequester) Post(ctx context.Context, path string, body, result any) error {
	var bodyReader io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		bodyReader = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, m.server.URL+path, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return m.doRequest(req, result)
}

func (m *mockRequester) Put(ctx context.Context, path string, body, result any) error {
	b, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, m.server.URL+path, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return m.doRequest(req, result)
}

func (m *mockRequester) Patch(ctx context.Context, path string, body, result any) error {
	b, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, m.server.URL+path, bytes.NewReader(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return m.doRequest(req, result)
}

func (m *mockRequester) Delete(ctx context.Context, path string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, m.server.URL+path, nil)
	if err != nil {
		return err
	}
	return m.doRequest(req, nil)
}

func (m *mockRequester) doRequest(req *http.Request, result any) error {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if result != nil {
		body, _ := io.ReadAll(resp.Body)
		if len(body) > 0 {
			return json.Unmarshal(body, result)
		}
	}
	return nil
}

func setupCMS(t *testing.T) (*cms.Service, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	requester := &mockRequester{server: server}
	svc := cms.NewService(requester)
	return svc, mux, server.Close
}

func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("failed to marshal JSON: %v", err)
	}
	return b
}

// --- Blog Posts Tests ---

func TestBlogPosts_Create(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input cms.BlogPost
		json.Unmarshal(body, &input)
		if input.Name != "My Blog Post" {
			t.Errorf("name = %q, want 'My Blog Post'", input.Name)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, cms.BlogPost{
			ID:   "101",
			Name: "My Blog Post",
			Slug: "my-blog-post",
		}))
	})

	post, err := svc.BlogPosts.Create(context.Background(), &cms.BlogPost{
		Name: "My Blog Post",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if post.ID != "101" {
		t.Errorf("ID = %q, want 101", post.ID)
	}
	if post.Name != "My Blog Post" {
		t.Errorf("Name = %q", post.Name)
	}
}

func TestBlogPosts_GetByID(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/posts/101", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.Write(mustJSON(t, cms.BlogPost{
			ID:   "101",
			Name: "My Blog Post",
		}))
	})

	post, err := svc.BlogPosts.GetByID(context.Background(), "101", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if post.ID != "101" {
		t.Errorf("ID = %q", post.ID)
	}
}

func TestBlogPosts_Update(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/posts/101", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %s, want PATCH", r.Method)
		}
		w.Write(mustJSON(t, cms.BlogPost{
			ID:   "101",
			Name: "Updated Post",
		}))
	})

	post, err := svc.BlogPosts.Update(context.Background(), "101", &cms.BlogPost{
		Name: "Updated Post",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if post.Name != "Updated Post" {
		t.Errorf("Name = %q", post.Name)
	}
}

func TestBlogPosts_Archive(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/posts/101", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.BlogPosts.Archive(context.Background(), "101")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBlogPosts_List(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if limit := r.URL.Query().Get("limit"); limit != "10" {
			t.Errorf("limit = %q, want 10", limit)
		}
		w.Write(mustJSON(t, cms.BlogPostListResult{
			Total: 2,
			Results: []*cms.BlogPost{
				{ID: "1", Name: "Post 1"},
				{ID: "2", Name: "Post 2"},
			},
			Paging: &cms.ForwardPaging{Next: &cms.NextPage{After: "2"}},
		}))
	})

	result, err := svc.BlogPosts.List(context.Background(), &cms.ListOptions{Limit: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d, want 2", len(result.Results))
	}
	if result.Paging.Next.After != "2" {
		t.Errorf("after = %q", result.Paging.Next.After)
	}
}

func TestBlogPosts_BatchCreate(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/posts/batch/create", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var input cms.BatchInputBlogPost
		json.Unmarshal(body, &input)
		if len(input.Inputs) != 2 {
			t.Errorf("inputs = %d, want 2", len(input.Inputs))
		}
		w.Write(mustJSON(t, cms.BatchResponseBlogPost{
			Status:      "COMPLETE",
			Results:     []*cms.BlogPost{{ID: "1"}, {ID: "2"}},
			StartedAt:   time.Now(),
			CompletedAt: time.Now(),
		}))
	})

	result, err := svc.BlogPosts.BatchCreate(context.Background(), &cms.BatchInputBlogPost{
		Inputs: []cms.BlogPost{
			{Name: "Post 1"},
			{Name: "Post 2"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Status != "COMPLETE" {
		t.Errorf("status = %q", result.Status)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d", len(result.Results))
	}
}

func TestBlogPosts_Schedule(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/posts/schedule", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	publishDate := time.Now().Add(24 * time.Hour)
	err := svc.BlogPosts.Schedule(context.Background(), &cms.ContentScheduleRequest{
		ID:          "101",
		PublishDate: &publishDate,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestBlogPosts_MultiLanguage(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/posts/multi-language/attach-to-lang-group", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.BlogPosts.AttachToLangGroup(context.Background(), &cms.AttachToLangPrimaryRequest{
		ID:        "102",
		Language:  "fr",
		PrimaryID: "101",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// --- Blog Authors Tests ---

func TestBlogAuthors_Create(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/authors", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, cms.BlogAuthor{
			ID:       "201",
			FullName: "John Doe",
		}))
	})

	author, err := svc.BlogAuthors.Create(context.Background(), &cms.BlogAuthor{
		FullName: "John Doe",
		Email:    "john@example.com",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if author.ID != "201" {
		t.Errorf("ID = %q", author.ID)
	}
}

// --- Blog Tags Tests ---

func TestBlogTags_Create(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/tags", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, cms.BlogTag{
			ID:   "301",
			Name: "Go",
		}))
	})

	tag, err := svc.BlogTags.Create(context.Background(), &cms.BlogTag{
		Name: "Go",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tag.ID != "301" {
		t.Errorf("ID = %q", tag.ID)
	}
}

func TestBlogTags_List(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/blogs/tags", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.Write(mustJSON(t, cms.BlogTagListResult{
			Total:   1,
			Results: []*cms.BlogTag{{ID: "301", Name: "Go"}},
		}))
	})

	result, err := svc.BlogTags.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 1 {
		t.Errorf("total = %d", result.Total)
	}
}

// --- HubDB Tables Tests ---

func TestTables_Create(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/hubdb/tables", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, cms.HubDbTable{
			ID:    "401",
			Name:  "my_table",
			Label: "My Table",
		}))
	})

	table, err := svc.Tables.Create(context.Background(), &cms.HubDbTableRequest{
		Name:  "my_table",
		Label: "My Table",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if table.ID != "401" {
		t.Errorf("ID = %q", table.ID)
	}
}

func TestTables_GetDetails(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/hubdb/tables/my_table", func(w http.ResponseWriter, r *http.Request) {
		w.Write(mustJSON(t, cms.HubDbTable{
			ID:   "401",
			Name: "my_table",
		}))
	})

	table, err := svc.Tables.GetDetails(context.Background(), "my_table", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if table.Name != "my_table" {
		t.Errorf("Name = %q", table.Name)
	}
}

func TestTables_Publish(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/hubdb/tables/401/draft/publish", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.Write(mustJSON(t, cms.HubDbTable{
			ID:        "401",
			Published: true,
		}))
	})

	table, err := svc.Tables.Publish(context.Background(), "401")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !table.Published {
		t.Error("expected Published to be true")
	}
}

func TestTables_Archive(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/hubdb/tables/401", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Tables.Archive(context.Background(), "401")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// --- HubDB Rows Tests ---

func TestRows_Create(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/hubdb/tables/401/rows", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, cms.HubDbTableRow{
			ID:     "501",
			Values: map[string]any{"name": "Row 1"},
		}))
	})

	row, err := svc.Rows.Create(context.Background(), "401", &cms.HubDbTableRowRequest{
		Values: map[string]any{"name": "Row 1"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if row.ID != "501" {
		t.Errorf("ID = %q", row.ID)
	}
}

func TestRows_List(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/hubdb/tables/401/rows", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.Write(mustJSON(t, cms.HubDbRowListResult{
			Total: 1,
			Results: []*cms.HubDbTableRow{
				{ID: "501", Values: map[string]any{"name": "Row 1"}},
			},
		}))
	})

	result, err := svc.Rows.List(context.Background(), "401", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 1 {
		t.Errorf("total = %d", result.Total)
	}
}

func TestRows_Purge(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/hubdb/tables/401/rows/501/draft", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Rows.Purge(context.Background(), "401", "501")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// --- Domains Tests ---

func TestDomains_GetByID(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/domains/601", func(w http.ResponseWriter, r *http.Request) {
		w.Write(mustJSON(t, cms.Domain{
			ID:     "601",
			Domain: "example.com",
		}))
	})

	domain, err := svc.Domains.GetByID(context.Background(), "601")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if domain.Domain != "example.com" {
		t.Errorf("Domain = %q", domain.Domain)
	}
}

func TestDomains_List(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/domains", func(w http.ResponseWriter, r *http.Request) {
		w.Write(mustJSON(t, cms.DomainListResult{
			Total: 1,
			Results: []*cms.Domain{
				{ID: "601", Domain: "example.com"},
			},
		}))
	})

	result, err := svc.Domains.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 1 {
		t.Errorf("total = %d", result.Total)
	}
}

// --- Pages Tests ---

func TestLandingPages_Create(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/pages/landing-pages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, cms.Page{
			ID:   "701",
			Name: "My Landing Page",
		}))
	})

	page, err := svc.LandingPages.Create(context.Background(), &cms.Page{
		Name: "My Landing Page",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if page.ID != "701" {
		t.Errorf("ID = %q", page.ID)
	}
}

func TestSitePages_List(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/pages/site-pages", func(w http.ResponseWriter, r *http.Request) {
		w.Write(mustJSON(t, cms.PageListResult{
			Total: 1,
			Results: []*cms.Page{
				{ID: "801", Name: "About"},
			},
		}))
	})

	result, err := svc.SitePages.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 1 {
		t.Errorf("total = %d", result.Total)
	}
}

func TestLandingPages_PushLive(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/pages/landing-pages/701/push-live", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.LandingPages.PushLive(context.Background(), "701")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLandingPages_ABTest(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/pages/landing-pages/ab-test/create-variation", func(w http.ResponseWriter, r *http.Request) {
		w.Write(mustJSON(t, cms.Page{
			ID:   "702",
			Name: "Variation B",
		}))
	})

	page, err := svc.LandingPages.CreateABTestVariation(context.Background(), &cms.ABTestCreateRequest{
		VariationID: "701",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if page.ID != "702" {
		t.Errorf("ID = %q", page.ID)
	}
}

// --- Audit Logs Tests ---

func TestAuditLogs_List(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/audit-logs", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("eventType") != "CREATED" {
			t.Errorf("eventType = %q, want CREATED", r.URL.Query().Get("eventType"))
		}
		w.Write(mustJSON(t, cms.AuditLogListResult{
			Results: []*cms.PublicAuditLog{
				{
					ObjectName: "My Post",
					Event:      "CREATED",
					ObjectType: "BLOG_POST",
					Timestamp:  time.Now(),
				},
			},
		}))
	})

	result, err := svc.AuditLogs.List(context.Background(), &cms.AuditLogListOptions{
		EventType: []string{"CREATED"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

// --- URL Redirects Tests ---

func TestUrlRedirects_Create(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/url-redirects", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, cms.UrlMapping{
			ID:            "901",
			RoutePrefix:   "/old-page",
			Destination:   "/new-page",
			RedirectStyle: 301,
		}))
	})

	redirect, err := svc.UrlRedirects.Create(context.Background(), &cms.UrlMappingCreateRequest{
		RoutePrefix:   "/old-page",
		Destination:   "/new-page",
		RedirectStyle: 301,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if redirect.ID != "901" {
		t.Errorf("ID = %q", redirect.ID)
	}
	if redirect.RedirectStyle != 301 {
		t.Errorf("RedirectStyle = %d", redirect.RedirectStyle)
	}
}

func TestUrlRedirects_GetByID(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/url-redirects/901", func(w http.ResponseWriter, r *http.Request) {
		w.Write(mustJSON(t, cms.UrlMapping{
			ID:          "901",
			RoutePrefix: "/old-page",
			Destination: "/new-page",
		}))
	})

	redirect, err := svc.UrlRedirects.GetByID(context.Background(), "901")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if redirect.RoutePrefix != "/old-page" {
		t.Errorf("RoutePrefix = %q", redirect.RoutePrefix)
	}
}

func TestUrlRedirects_Archive(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/url-redirects/901", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.UrlRedirects.Archive(context.Background(), "901")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// --- Source Code Tests ---

func TestSourceCode_GetMetadata(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/source-code/publish/metadata/templates/home.html", func(w http.ResponseWriter, r *http.Request) {
		w.Write(mustJSON(t, cms.AssetFileMetadata{
			ID:        "f1",
			Name:      "home.html",
			Folder:    false,
			CreatedAt: 1704067200,
			UpdatedAt: 1704153600,
		}))
	})

	meta, err := svc.SourceCode.Metadata.Get(context.Background(), "publish", "templates/home.html", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if meta.Name != "home.html" {
		t.Errorf("Name = %q", meta.Name)
	}
}

func TestSourceCode_Archive(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/source-code/publish/content/templates/old.html", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.SourceCode.Content.Archive(context.Background(), "publish", "templates/old.html")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSourceCode_Extract(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/source-code/extract/async", func(w http.ResponseWriter, r *http.Request) {
		w.Write(mustJSON(t, cms.TaskLocator{ID: 12345}))
	})

	task, err := svc.SourceCode.Extract.Extract(context.Background(), &cms.FileExtractRequest{
		Path: "/templates/archive.zip",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.ID != 12345 {
		t.Errorf("ID = %d", task.ID)
	}
}

func TestSourceCode_Validate(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/source-code/validation/validate/templates/home.html", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.SourceCode.Validation.Validate(context.Background(), "templates/home.html")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// --- Site Search Tests ---

func TestSiteSearch_Search(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/site-search/search", func(w http.ResponseWriter, r *http.Request) {
		if q := r.URL.Query().Get("q"); q != "hubspot" {
			t.Errorf("q = %q, want hubspot", q)
		}
		w.Write(mustJSON(t, cms.SearchResults{
			Total:  1,
			Limit:  10,
			Offset: 0,
			Page:   0,
			Results: []*cms.ContentSearchResult{
				{
					ID:    1,
					Title: "Getting Started with HubSpot",
					URL:   "https://example.com/blog/getting-started",
					Type:  "BLOG_POST",
					Score: 0.95,
				},
			},
		}))
	})

	result, err := svc.SiteSearch.Search(context.Background(), &cms.SiteSearchOptions{
		Query: "hubspot",
		Limit: 10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 1 {
		t.Errorf("total = %d", result.Total)
	}
	if result.Results[0].Title != "Getting Started with HubSpot" {
		t.Errorf("title = %q", result.Results[0].Title)
	}
}

func TestSiteSearch_GetByID(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/site-search/indexed/abc123", func(w http.ResponseWriter, r *http.Request) {
		if ct := r.URL.Query().Get("type"); ct != "BLOG_POST" {
			t.Errorf("type = %q, want BLOG_POST", ct)
		}
		w.Write(mustJSON(t, cms.IndexedData{
			ID:   "abc123",
			Type: "BLOG_POST",
		}))
	})

	data, err := svc.SiteSearch.GetByID(context.Background(), "abc123", cms.ContentTypeBlogPost)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if data.Type != "BLOG_POST" {
		t.Errorf("type = %q", data.Type)
	}
}

// --- Performance Tests ---

func TestPerformance_GetPage(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/performance", func(w http.ResponseWriter, r *http.Request) {
		if d := r.URL.Query().Get("domain"); d != "example.com" {
			t.Errorf("domain = %q, want example.com", d)
		}
		w.Write(mustJSON(t, cms.PerformanceResponse{
			Domain:        "example.com",
			Interval:      "ONE_HOUR",
			StartInterval: 1704067200,
			EndInterval:   1704153600,
			Data: []*cms.PerformanceView{
				{
					StartTimestamp: 1704067200,
					EndTimestamp:   1704070800,
					TotalRequests:  1500,
					CacheHitRate:   0.85,
				},
			},
		}))
	})

	result, err := svc.Performance.GetPage(context.Background(), &cms.PerformanceOptions{
		Domain: "example.com",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Domain != "example.com" {
		t.Errorf("Domain = %q", result.Domain)
	}
	if len(result.Data) != 1 {
		t.Errorf("data = %d, want 1", len(result.Data))
	}
	if result.Data[0].TotalRequests != 1500 {
		t.Errorf("TotalRequests = %d", result.Data[0].TotalRequests)
	}
}

func TestPerformance_GetUptime(t *testing.T) {
	svc, mux, teardown := setupCMS(t)
	defer teardown()

	mux.HandleFunc("/cms/v3/performance/uptime", func(w http.ResponseWriter, r *http.Request) {
		w.Write(mustJSON(t, cms.PerformanceResponse{
			Interval:      "ONE_DAY",
			StartInterval: 1704067200,
			EndInterval:   1704153600,
			Data:          []*cms.PerformanceView{},
		}))
	})

	result, err := svc.Performance.GetUptime(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Interval != "ONE_DAY" {
		t.Errorf("Interval = %q", result.Interval)
	}
}

// --- Service Container Test ---

func TestNewService_AllFieldsInitialized(t *testing.T) {
	svc, _, teardown := setupCMS(t)
	defer teardown()

	if svc.BlogPosts == nil {
		t.Error("BlogPosts is nil")
	}
	if svc.BlogAuthors == nil {
		t.Error("BlogAuthors is nil")
	}
	if svc.BlogTags == nil {
		t.Error("BlogTags is nil")
	}
	if svc.Tables == nil {
		t.Error("Tables is nil")
	}
	if svc.Rows == nil {
		t.Error("Rows is nil")
	}
	if svc.RowsBatch == nil {
		t.Error("RowsBatch is nil")
	}
	if svc.Domains == nil {
		t.Error("Domains is nil")
	}
	if svc.LandingPages == nil {
		t.Error("LandingPages is nil")
	}
	if svc.SitePages == nil {
		t.Error("SitePages is nil")
	}
	if svc.AuditLogs == nil {
		t.Error("AuditLogs is nil")
	}
	if svc.UrlRedirects == nil {
		t.Error("UrlRedirects is nil")
	}
	if svc.SourceCode == nil {
		t.Error("SourceCode is nil")
	}
	if svc.SourceCode.Content == nil {
		t.Error("SourceCode.Content is nil")
	}
	if svc.SourceCode.Extract == nil {
		t.Error("SourceCode.Extract is nil")
	}
	if svc.SourceCode.Metadata == nil {
		t.Error("SourceCode.Metadata is nil")
	}
	if svc.SourceCode.Validation == nil {
		t.Error("SourceCode.Validation is nil")
	}
	if svc.SiteSearch == nil {
		t.Error("SiteSearch is nil")
	}
	if svc.Performance == nil {
		t.Error("Performance is nil")
	}
}
