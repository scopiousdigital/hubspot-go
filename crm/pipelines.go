package crm

import (
	"context"
	"fmt"
	"net/url"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const pipelinesBasePath = "/crm/v3/pipelines"

// PipelinesService handles operations on HubSpot pipelines, stages, and audits.
type PipelinesService struct {
	requester api.Requester
}

// --- Pipelines API ---

// Create creates a new pipeline.
func (s *PipelinesService) Create(ctx context.Context, objectType string, input *PipelineInput) (*Pipeline, error) {
	path := fmt.Sprintf("%s/%s", pipelinesBasePath, objectType)
	var result Pipeline
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAll retrieves all pipelines for the given object type.
func (s *PipelinesService) GetAll(ctx context.Context, objectType string) (*PipelineListResult, error) {
	path := fmt.Sprintf("%s/%s", pipelinesBasePath, objectType)
	var result PipelineListResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a pipeline by its ID.
func (s *PipelinesService) GetByID(ctx context.Context, objectType, pipelineID string) (*Pipeline, error) {
	path := fmt.Sprintf("%s/%s/%s", pipelinesBasePath, objectType, pipelineID)
	var result Pipeline
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates an existing pipeline (PATCH).
func (s *PipelinesService) Update(ctx context.Context, objectType, pipelineID string, input *PipelinePatchInput) (*Pipeline, error) {
	path := fmt.Sprintf("%s/%s/%s", pipelinesBasePath, objectType, pipelineID)
	var result Pipeline
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Replace replaces a pipeline entirely (PUT).
func (s *PipelinesService) Replace(ctx context.Context, objectType, pipelineID string, input *PipelineInput) (*Pipeline, error) {
	path := fmt.Sprintf("%s/%s/%s", pipelinesBasePath, objectType, pipelineID)
	var result Pipeline
	if err := s.requester.Put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive deletes a pipeline.
func (s *PipelinesService) Archive(ctx context.Context, objectType, pipelineID string) error {
	path := fmt.Sprintf("%s/%s/%s", pipelinesBasePath, objectType, pipelineID)
	return s.requester.Delete(ctx, path)
}

// --- Pipeline Stages API ---

// CreateStage creates a new stage within a pipeline.
func (s *PipelinesService) CreateStage(ctx context.Context, objectType, pipelineID string, input *PipelineStageInput) (*PipelineStage, error) {
	path := fmt.Sprintf("%s/%s/%s/stages", pipelinesBasePath, objectType, pipelineID)
	var result PipelineStage
	if err := s.requester.Post(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAllStages retrieves all stages for a pipeline.
func (s *PipelinesService) GetAllStages(ctx context.Context, objectType, pipelineID string) (*PipelineStageListResult, error) {
	path := fmt.Sprintf("%s/%s/%s/stages", pipelinesBasePath, objectType, pipelineID)
	var result PipelineStageListResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetStageByID retrieves a stage by its ID.
func (s *PipelinesService) GetStageByID(ctx context.Context, objectType, pipelineID, stageID string) (*PipelineStage, error) {
	path := fmt.Sprintf("%s/%s/%s/stages/%s", pipelinesBasePath, objectType, pipelineID, stageID)
	var result PipelineStage
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateStage updates an existing stage (PATCH).
func (s *PipelinesService) UpdateStage(ctx context.Context, objectType, pipelineID, stageID string, input *PipelineStagePatchInput) (*PipelineStage, error) {
	path := fmt.Sprintf("%s/%s/%s/stages/%s", pipelinesBasePath, objectType, pipelineID, stageID)
	var result PipelineStage
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReplaceStage replaces a stage entirely (PUT).
func (s *PipelinesService) ReplaceStage(ctx context.Context, objectType, pipelineID, stageID string, input *PipelineStageInput) (*PipelineStage, error) {
	path := fmt.Sprintf("%s/%s/%s/stages/%s", pipelinesBasePath, objectType, pipelineID, stageID)
	var result PipelineStage
	if err := s.requester.Put(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchiveStage deletes a stage.
func (s *PipelinesService) ArchiveStage(ctx context.Context, objectType, pipelineID, stageID string) error {
	path := fmt.Sprintf("%s/%s/%s/stages/%s", pipelinesBasePath, objectType, pipelineID, stageID)
	return s.requester.Delete(ctx, path)
}

// --- Pipeline Audits API ---

// GetAuditByPipeline retrieves audit logs for a pipeline.
func (s *PipelinesService) GetAuditByPipeline(ctx context.Context, objectType, pipelineID string) (*AuditInfoListResult, error) {
	path := fmt.Sprintf("%s/%s/%s/audit", pipelinesBasePath, objectType, pipelineID)
	var result AuditInfoListResult
	if err := s.requester.Get(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAuditByPipelineStage retrieves audit logs for a pipeline stage.
func (s *PipelinesService) GetAuditByPipelineStage(ctx context.Context, objectType, pipelineID string) (*AuditInfoListResult, error) {
	path := fmt.Sprintf("%s/%s/%s/stages/audit", pipelinesBasePath, objectType, pipelineID)
	q := url.Values{}
	var result AuditInfoListResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
