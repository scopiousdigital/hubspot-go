package api

import (
	"context"
	"net/url"
)

// Requester is the interface that sub-packages use to make HTTP requests.
// Implemented by the root package's httpClient.
type Requester interface {
	Get(ctx context.Context, path string, query url.Values, result any) error
	Post(ctx context.Context, path string, body, result any) error
	Put(ctx context.Context, path string, body, result any) error
	Patch(ctx context.Context, path string, body, result any) error
	Delete(ctx context.Context, path string) error
}
