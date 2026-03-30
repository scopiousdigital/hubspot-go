package files_test

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

	"github.com/scopiousdigital/hubspot-go/files"
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
	b, _ := json.Marshal(body)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, m.server.URL+path, bytes.NewReader(b))
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

func setupFiles(t *testing.T) (*files.Service, *http.ServeMux, func()) {
	t.Helper()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	requester := &mockRequester{server: server}
	svc := files.NewService(requester)
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

// --- Files Tests ---

func TestFiles_GetByID(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/files/file-123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, files.File{
			ID:        "file-123",
			Name:      "test.pdf",
			Extension: "pdf",
			Access:    files.AccessPublicNotIndexable,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	file, err := svc.Files.GetByID(context.Background(), "file-123", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if file.ID != "file-123" {
		t.Errorf("ID = %q", file.ID)
	}
	if file.Name != "test.pdf" {
		t.Errorf("Name = %q", file.Name)
	}
}

func TestFiles_Search(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/files/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if name := r.URL.Query().Get("name"); name != "report" {
			t.Errorf("name = %q, want report", name)
		}
		if limit := r.URL.Query().Get("limit"); limit != "5" {
			t.Errorf("limit = %q, want 5", limit)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, files.CollectionResponseFile{
			Results: []files.File{
				{ID: "1", Name: "report.pdf", CreatedAt: time.Now(), UpdatedAt: time.Now()},
				{ID: "2", Name: "report.docx", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
		}))
	})

	result, err := svc.Files.Search(context.Background(), &files.FileSearchOptions{
		Name:  "report",
		Limit: 5,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d, want 2", len(result.Results))
	}
}

func TestFiles_Archive(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/files/file-123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Files.Archive(context.Background(), "file-123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFiles_UpdateProperties(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/files/file-123", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %s, want PATCH", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input files.FileUpdateInput
		_ = json.Unmarshal(body, &input)
		if input.Name != "renamed.pdf" {
			t.Errorf("Name = %q", input.Name)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, files.File{
			ID:        "file-123",
			Name:      "renamed.pdf",
			Access:    files.AccessPublicNotIndexable,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	file, err := svc.Files.UpdateProperties(context.Background(), "file-123", &files.FileUpdateInput{
		Name: "renamed.pdf",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if file.Name != "renamed.pdf" {
		t.Errorf("Name = %q", file.Name)
	}
}

func TestFiles_CheckImportStatus(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/files/import-from-url/async/tasks/task-abc/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, files.FileActionResponse{
			TaskID:      "task-abc",
			Status:      files.StatusComplete,
			CompletedAt: time.Now(),
			StartedAt:   time.Now(),
		}))
	})

	resp, err := svc.Files.CheckImportStatus(context.Background(), "task-abc")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Status != files.StatusComplete {
		t.Errorf("Status = %q", resp.Status)
	}
	if resp.TaskID != "task-abc" {
		t.Errorf("TaskID = %q", resp.TaskID)
	}
}

func TestFiles_ImportFromURL(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/files/import-from-url/async", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input files.ImportFromURLInput
		_ = json.Unmarshal(body, &input)
		if input.URL != "https://example.com/image.png" {
			t.Errorf("URL = %q", input.URL)
		}
		w.WriteHeader(http.StatusAccepted)
		_, _ = w.Write(mustJSON(t, files.ImportFromURLTaskLocator{
			ID:    "task-xyz",
			Links: map[string]string{"status": "/files/v3/files/import-from-url/async/tasks/task-xyz/status"},
		}))
	})

	locator, err := svc.Files.ImportFromURL(context.Background(), &files.ImportFromURLInput{
		URL:    "https://example.com/image.png",
		Access: files.AccessPublicNotIndexable,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if locator.ID != "task-xyz" {
		t.Errorf("ID = %q", locator.ID)
	}
}

func TestFiles_GetSignedURL(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/files/file-123/signed-url", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if size := r.URL.Query().Get("size"); size != "thumb" {
			t.Errorf("size = %q, want thumb", size)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, files.SignedURL{
			URL:       "https://cdn.hubspot.com/signed/file-123",
			Name:      "test.png",
			Extension: "png",
			Type:      "IMG",
			Size:      1024,
			ExpiresAt: time.Now().Add(time.Hour),
		}))
	})

	signed, err := svc.Files.GetSignedURL(context.Background(), "file-123", "thumb", 0, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if signed.URL != "https://cdn.hubspot.com/signed/file-123" {
		t.Errorf("URL = %q", signed.URL)
	}
}

// --- Folders Tests ---

func TestFolders_Create(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/folders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input files.FolderInput
		_ = json.Unmarshal(body, &input)
		if input.Name != "my-folder" {
			t.Errorf("Name = %q", input.Name)
		}
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write(mustJSON(t, files.Folder{
			ID:        "folder-1",
			Name:      "my-folder",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	folder, err := svc.Folders.Create(context.Background(), &files.FolderInput{
		Name: "my-folder",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if folder.ID != "folder-1" {
		t.Errorf("ID = %q", folder.ID)
	}
	if folder.Name != "my-folder" {
		t.Errorf("Name = %q", folder.Name)
	}
}

func TestFolders_Search(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/folders/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, files.CollectionResponseFolder{
			Results: []files.Folder{
				{ID: "folder-1", Name: "images", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
		}))
	})

	result, err := svc.Folders.Search(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

func TestFolders_GetByID(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/folders/folder-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, files.Folder{
			ID:        "folder-1",
			Name:      "images",
			Path:      "/images",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	folder, err := svc.Folders.GetByID(context.Background(), "folder-1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if folder.ID != "folder-1" {
		t.Errorf("ID = %q", folder.ID)
	}
	if folder.Path != "/images" {
		t.Errorf("Path = %q", folder.Path)
	}
}

func TestFolders_Update(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/folders/folder-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %s, want PATCH", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input files.FolderUpdateInput
		_ = json.Unmarshal(body, &input)
		if input.Name != "renamed-folder" {
			t.Errorf("Name = %q", input.Name)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(mustJSON(t, files.Folder{
			ID:        "folder-1",
			Name:      "renamed-folder",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	folder, err := svc.Folders.Update(context.Background(), "folder-1", &files.FolderUpdateInput{
		Name: "renamed-folder",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if folder.Name != "renamed-folder" {
		t.Errorf("Name = %q", folder.Name)
	}
}

func TestFolders_Archive(t *testing.T) {
	svc, mux, teardown := setupFiles(t)
	defer teardown()

	mux.HandleFunc("/files/v3/folders/folder-1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Folders.Archive(context.Background(), "folder-1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestFile_JSONRoundTrip(t *testing.T) {
	input := `{
		"id": "file-1",
		"name": "test.png",
		"extension": "png",
		"access": "PUBLIC_NOT_INDEXABLE",
		"createdAt": "2024-01-15T10:30:00.000Z",
		"updatedAt": "2024-01-16T11:00:00.000Z",
		"archived": false,
		"size": 2048,
		"width": 800,
		"height": 600
	}`
	var file files.File
	if err := json.Unmarshal([]byte(input), &file); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}
	if file.ID != "file-1" {
		t.Errorf("ID = %q", file.ID)
	}
	if file.Access != files.AccessPublicNotIndexable {
		t.Errorf("Access = %q", file.Access)
	}
	if file.Size == nil || *file.Size != 2048 {
		t.Errorf("Size = %v", file.Size)
	}
	if file.Width == nil || *file.Width != 800 {
		t.Errorf("Width = %v", file.Width)
	}
}
