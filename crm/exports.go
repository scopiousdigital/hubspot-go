package crm

import (
	"context"
	"fmt"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const exportsBasePath = "/crm/v3/exports/export/async"

// ExportsService handles operations on HubSpot CRM exports.
type ExportsService struct {
	requester api.Requester
}

// Start begins a new export task.
func (s *ExportsService) Start(ctx context.Context, input *ExportRequest) (*TaskLocator, error) {
	var result TaskLocator
	if err := s.requester.Post(ctx, exportsBasePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetStatus retrieves the status of an export task.
func (s *ExportsService) GetStatus(ctx context.Context, taskID string) (*ExportStatusResponse, error) {
	path := fmt.Sprintf("%s/tasks/%s/status", exportsBasePath, taskID)
	var result ExportStatusResponse
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
