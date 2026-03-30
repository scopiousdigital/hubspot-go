package crm_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/scopiousdigital/hubspot-go/crm"
)

// =============================================================================
// Owners tests
// =============================================================================

func TestOwners_GetByID(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/owners/12345", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.PublicOwner{
			ID:        "12345",
			Email:     "owner@test.com",
			FirstName: "Jane",
			LastName:  "Doe",
			Type:      crm.OwnerTypePerson,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	owner, err := svc.Owners.GetByID(context.Background(), "12345", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if owner.ID != "12345" {
		t.Errorf("ID = %q, want 12345", owner.ID)
	}
	if owner.Email != "owner@test.com" {
		t.Errorf("Email = %q", owner.Email)
	}
}

func TestOwners_List(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/owners", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		if email := r.URL.Query().Get("email"); email != "owner@test.com" {
			t.Errorf("email = %q", email)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.OwnerListResult{
			Results: []*crm.PublicOwner{
				{ID: "1", Email: "owner@test.com", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
		}))
	})

	result, err := svc.Owners.List(context.Background(), &crm.OwnerListOptions{Email: "owner@test.com"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

// =============================================================================
// Properties tests
// =============================================================================

func TestProperties_Create(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/properties/contacts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		body, _ := io.ReadAll(r.Body)
		var input crm.PropertyCreate
		json.Unmarshal(body, &input)
		if input.Name != "custom_prop" {
			t.Errorf("name = %q", input.Name)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.Property{
			Name:      "custom_prop",
			Label:     "Custom Property",
			Type:      "string",
			FieldType: "text",
			GroupName: "contactinformation",
		}))
	})

	prop, err := svc.Properties.Create(context.Background(), "contacts", &crm.PropertyCreate{
		Name:      "custom_prop",
		Label:     "Custom Property",
		Type:      "string",
		FieldType: "text",
		GroupName: "contactinformation",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if prop.Name != "custom_prop" {
		t.Errorf("name = %q", prop.Name)
	}
}

func TestProperties_GetAll(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/properties/contacts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.PropertyListResult{
			Results: []*crm.Property{
				{Name: "email", Label: "Email", Type: "string", FieldType: "text"},
				{Name: "firstname", Label: "First Name", Type: "string", FieldType: "text"},
			},
		}))
	})

	result, err := svc.Properties.GetAll(context.Background(), "contacts", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 2 {
		t.Errorf("results = %d, want 2", len(result.Results))
	}
}

func TestProperties_GetByName(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/properties/contacts/email", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.Property{
			Name: "email", Label: "Email", Type: "string", FieldType: "text",
		}))
	})

	prop, err := svc.Properties.GetByName(context.Background(), "contacts", "email", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if prop.Name != "email" {
		t.Errorf("name = %q", prop.Name)
	}
}

func TestProperties_Update(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/properties/contacts/custom_prop", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %s, want PATCH", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.Property{
			Name: "custom_prop", Label: "Updated Label", Type: "string", FieldType: "text",
		}))
	})

	prop, err := svc.Properties.Update(context.Background(), "contacts", "custom_prop", &crm.PropertyUpdate{
		Label: "Updated Label",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if prop.Label != "Updated Label" {
		t.Errorf("label = %q", prop.Label)
	}
}

func TestProperties_Archive(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/properties/contacts/custom_prop", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Properties.Archive(context.Background(), "contacts", "custom_prop")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestProperties_CreateGroup(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/properties/contacts/groups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.PropertyGroup{
			Name:  "my_group",
			Label: "My Group",
		}))
	})

	group, err := svc.Properties.CreateGroup(context.Background(), "contacts", &crm.PropertyGroupCreate{
		Name:  "my_group",
		Label: "My Group",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if group.Name != "my_group" {
		t.Errorf("name = %q", group.Name)
	}
}

func TestProperties_GetAllGroups(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/properties/contacts/groups", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.PropertyGroupListResult{
			Results: []*crm.PropertyGroup{
				{Name: "contactinformation", Label: "Contact Information"},
			},
		}))
	})

	result, err := svc.Properties.GetAllGroups(context.Background(), "contacts")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

// =============================================================================
// Pipelines tests
// =============================================================================

func TestPipelines_Create(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/pipelines/deals", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.Pipeline{
			ID:    "p1",
			Label: "Sales Pipeline",
			Stages: []crm.PipelineStage{
				{ID: "s1", Label: "Qualified"},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	pipeline, err := svc.Pipelines.Create(context.Background(), "deals", &crm.PipelineInput{
		Label:        "Sales Pipeline",
		DisplayOrder: 0,
		Stages: []crm.PipelineStageInput{
			{Label: "Qualified", DisplayOrder: 0},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pipeline.ID != "p1" {
		t.Errorf("ID = %q", pipeline.ID)
	}
}

func TestPipelines_GetAll(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/pipelines/deals", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.PipelineListResult{
			Results: []*crm.Pipeline{
				{ID: "p1", Label: "Sales", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
		}))
	})

	result, err := svc.Pipelines.GetAll(context.Background(), "deals")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

func TestPipelines_GetByID(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/pipelines/deals/p1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.Pipeline{
			ID: "p1", Label: "Sales", CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}))
	})

	pipeline, err := svc.Pipelines.GetByID(context.Background(), "deals", "p1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pipeline.Label != "Sales" {
		t.Errorf("label = %q", pipeline.Label)
	}
}

func TestPipelines_Update(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/pipelines/deals/p1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %s, want PATCH", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.Pipeline{
			ID: "p1", Label: "Updated Pipeline", CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}))
	})

	pipeline, err := svc.Pipelines.Update(context.Background(), "deals", "p1", &crm.PipelinePatchInput{
		Label: "Updated Pipeline",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pipeline.Label != "Updated Pipeline" {
		t.Errorf("label = %q", pipeline.Label)
	}
}

func TestPipelines_Replace(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/pipelines/deals/p1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.Pipeline{
			ID: "p1", Label: "Replaced Pipeline", CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}))
	})

	pipeline, err := svc.Pipelines.Replace(context.Background(), "deals", "p1", &crm.PipelineInput{
		Label:        "Replaced Pipeline",
		DisplayOrder: 0,
		Stages:       []crm.PipelineStageInput{{Label: "New Stage", DisplayOrder: 0}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pipeline.Label != "Replaced Pipeline" {
		t.Errorf("label = %q", pipeline.Label)
	}
}

func TestPipelines_Archive(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/pipelines/deals/p1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Pipelines.Archive(context.Background(), "deals", "p1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPipelines_CreateStage(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/pipelines/deals/p1/stages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.PipelineStage{
			ID: "s1", Label: "New Stage", CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}))
	})

	stage, err := svc.Pipelines.CreateStage(context.Background(), "deals", "p1", &crm.PipelineStageInput{
		Label: "New Stage", DisplayOrder: 0,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if stage.ID != "s1" {
		t.Errorf("ID = %q", stage.ID)
	}
}

func TestPipelines_ArchiveStage(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/pipelines/deals/p1/stages/s1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Pipelines.ArchiveStage(context.Background(), "deals", "p1", "s1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// =============================================================================
// Schemas tests
// =============================================================================

func TestSchemas_Create(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/schemas", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.ObjectSchema{
			ID:   "schema1",
			Name: "cars",
			Labels: crm.ObjectTypeDefinitionLabels{
				Singular: "Car",
				Plural:   "Cars",
			},
			RequiredProperties: []string{"make"},
		}))
	})

	schema, err := svc.Schemas.Create(context.Background(), &crm.ObjectSchemaEgg{
		Name:               "cars",
		Labels:             crm.ObjectTypeDefinitionLabels{Singular: "Car", Plural: "Cars"},
		RequiredProperties: []string{"make"},
		Properties: []crm.ObjectTypePropertyCreate{
			{Name: "make", Label: "Make", Type: "string", FieldType: "text"},
		},
		AssociatedObjects: []string{"contacts"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if schema.Name != "cars" {
		t.Errorf("name = %q", schema.Name)
	}
}

func TestSchemas_GetAll(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/schemas", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ObjectSchemaListResult{
			Results: []*crm.ObjectSchema{
				{ID: "1", Name: "cars", Labels: crm.ObjectTypeDefinitionLabels{Singular: "Car", Plural: "Cars"}},
			},
		}))
	})

	result, err := svc.Schemas.GetAll(context.Background(), false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

func TestSchemas_GetByID(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/schemas/cars", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ObjectSchema{
			ID: "1", Name: "cars", Labels: crm.ObjectTypeDefinitionLabels{Singular: "Car"},
		}))
	})

	schema, err := svc.Schemas.GetByID(context.Background(), "cars")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if schema.ID != "1" {
		t.Errorf("ID = %q", schema.ID)
	}
}

func TestSchemas_Update(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/schemas/cars", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("method = %s, want PATCH", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ObjectTypeDefinition{
			ID: "1", Name: "cars",
			Labels:             crm.ObjectTypeDefinitionLabels{Singular: "Vehicle"},
			RequiredProperties: []string{"make"},
		}))
	})

	result, err := svc.Schemas.Update(context.Background(), "cars", &crm.ObjectTypeDefinitionPatch{
		Labels: &crm.ObjectTypeDefinitionLabels{Singular: "Vehicle"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Labels.Singular != "Vehicle" {
		t.Errorf("singular = %q", result.Labels.Singular)
	}
}

func TestSchemas_Archive(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/schemas/cars", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Schemas.Archive(context.Background(), "cars")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSchemas_CreateAssociation(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/schemas/cars/associations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.SchemaAssociationDefinition{
			ID:               "assoc1",
			FromObjectTypeID: "cars",
			ToObjectTypeID:   "contacts",
		}))
	})

	assoc, err := svc.Schemas.CreateAssociation(context.Background(), "cars", &crm.SchemaAssociationDefinitionEgg{
		FromObjectTypeID: "cars",
		ToObjectTypeID:   "contacts",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if assoc.ID != "assoc1" {
		t.Errorf("ID = %q", assoc.ID)
	}
}

// =============================================================================
// Associations tests
// =============================================================================

func TestAssociations_BatchCreate(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/associations/contacts/companies/batch/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.BatchPublicAssociationResult{
			Status:      "COMPLETE",
			StartedAt:   time.Now(),
			CompletedAt: time.Now(),
		}))
	})

	result, err := svc.Associations.BatchCreate(context.Background(), "contacts", "companies", &crm.BatchPublicAssociationInput{
		Inputs: []crm.PublicAssociation{
			{From: crm.ObjectID{ID: "1"}, To: crm.ObjectID{ID: "2"}, Type: "contact_to_company"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Status != "COMPLETE" {
		t.Errorf("status = %q", result.Status)
	}
}

func TestAssociations_BatchRead(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/associations/contacts/companies/batch/read", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.BatchPublicAssociationMultiResult{
			Status:      "COMPLETE",
			StartedAt:   time.Now(),
			CompletedAt: time.Now(),
		}))
	})

	result, err := svc.Associations.BatchRead(context.Background(), "contacts", "companies", &crm.BatchPublicObjectIDInput{
		Inputs: []crm.ObjectID{{ID: "1"}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Status != "COMPLETE" {
		t.Errorf("status = %q", result.Status)
	}
}

func TestAssociations_V4Create(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v4/associations/contacts/100/companies/200", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.LabelsBetweenObjectPair{
			FromObjectTypeID: "contacts",
			FromObjectID:     100,
			ToObjectTypeID:   "companies",
			ToObjectID:       200,
		}))
	})

	result, err := svc.Associations.V4Create(context.Background(), "contacts", "100", "companies", "200", []crm.AssociationV4Spec{
		{AssociationCategory: crm.AssociationCategoryHubSpotDefined, AssociationTypeID: 1},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ToObjectID != 200 {
		t.Errorf("toObjectId = %d", result.ToObjectID)
	}
}

func TestAssociations_V4GetPage(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v4/associations/contacts/100/companies", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.CollectionMultiAssociatedObjectWithLabel{
			Results: []*crm.MultiAssociatedObjectWithLabel{
				{ToObjectID: 200},
			},
		}))
	})

	result, err := svc.Associations.V4GetPage(context.Background(), "contacts", "100", "companies", "", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

func TestAssociations_V4GetDefinitions(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v4/associations/contacts/companies/labels", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.AssociationDefinitionSpecWithLabelResult{
			Results: []*crm.AssociationV4SpecWithLabel{
				{Category: "HUBSPOT_DEFINED", TypeID: 1, Label: "Primary"},
			},
		}))
	})

	result, err := svc.Associations.V4GetDefinitions(context.Background(), "contacts", "companies")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

// =============================================================================
// Imports tests
// =============================================================================

func TestImports_List(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/imports", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ImportListResult{
			Results: []*crm.PublicImportResponse{
				{ID: "imp1", State: "DONE", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
		}))
	})

	result, err := svc.Imports.List(context.Background(), nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

func TestImports_GetByID(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/imports/imp1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.PublicImportResponse{
			ID: "imp1", State: "DONE", CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}))
	})

	imp, err := svc.Imports.GetByID(context.Background(), "imp1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if imp.State != "DONE" {
		t.Errorf("state = %q", imp.State)
	}
}

func TestImports_Cancel(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/imports/imp1/cancel", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ActionResponse{
			Status:      "COMPLETE",
			StartedAt:   time.Now(),
			CompletedAt: time.Now(),
		}))
	})

	result, err := svc.Imports.Cancel(context.Background(), "imp1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Status != "COMPLETE" {
		t.Errorf("status = %q", result.Status)
	}
}

func TestImports_GetErrors(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/imports/imp1/errors", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ImportErrorListResult{
			Results: []*crm.PublicImportError{
				{ErrorType: "INVALID_EMAIL"},
			},
		}))
	})

	result, err := svc.Imports.GetErrors(context.Background(), "imp1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

// =============================================================================
// Exports tests
// =============================================================================

func TestExports_Start(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/exports/export/async", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusAccepted)
		w.Write(mustJSON(t, crm.TaskLocator{ID: "task1"}))
	})

	result, err := svc.Exports.Start(context.Background(), &crm.ExportRequest{
		ExportType:       "VIEW",
		Format:           "CSV",
		ExportName:       "test-export",
		ObjectType:       "contacts",
		ObjectProperties: []string{"email", "firstname"},
		Language:         "EN",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "task1" {
		t.Errorf("ID = %q", result.ID)
	}
}

func TestExports_GetStatus(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/exports/export/async/tasks/task1/status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ExportStatusResponse{
			Status:      "COMPLETE",
			Result:      "https://download.url/file.csv",
			StartedAt:   time.Now(),
			CompletedAt: time.Now(),
		}))
	})

	result, err := svc.Exports.GetStatus(context.Background(), "task1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Status != "COMPLETE" {
		t.Errorf("status = %q", result.Status)
	}
}

// =============================================================================
// Lists tests
// =============================================================================

func TestLists_Create(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/lists", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ListCreateResponse{
			ListID:         "lst1",
			Name:           "Test List",
			ObjectTypeID:   "0-1",
			ProcessingType: "MANUAL",
		}))
	})

	result, err := svc.Lists.Create(context.Background(), &crm.ListCreateRequest{
		Name:           "Test List",
		ObjectTypeID:   "0-1",
		ProcessingType: "MANUAL",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ListID != "lst1" {
		t.Errorf("listId = %q", result.ListID)
	}
}

func TestLists_GetByID(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/lists/lst1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ListFetchResponse{
			List: &crm.HubSpotList{ListID: "lst1", Name: "Test List"},
		}))
	})

	result, err := svc.Lists.GetByID(context.Background(), "lst1", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.List.Name != "Test List" {
		t.Errorf("name = %q", result.List.Name)
	}
}

func TestLists_Search(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/lists/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ListSearchResponse{
			Lists: []*crm.HubSpotList{
				{ListID: "lst1", Name: "Test List"},
			},
			Total: 1,
		}))
	})

	result, err := svc.Lists.Search(context.Background(), &crm.ListSearchRequest{
		Query: "Test",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 1 {
		t.Errorf("total = %d", result.Total)
	}
}

func TestLists_Remove(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/lists/lst1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Lists.Remove(context.Background(), "lst1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLists_AddMembers(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/lists/lst1/memberships/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.MembershipsUpdateResponse{
			RecordIdsAdded: []string{"100", "200"},
		}))
	})

	result, err := svc.Lists.AddMembers(context.Background(), "lst1", []string{"100", "200"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.RecordIdsAdded) != 2 {
		t.Errorf("added = %d, want 2", len(result.RecordIdsAdded))
	}
}

func TestLists_RemoveAllMembers(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/lists/lst1/memberships", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Lists.RemoveAllMembers(context.Background(), "lst1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLists_CreateFolder(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/lists/folders", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.ListFolderCreateResponse{
			Folder: &crm.ListFolder{ID: "f1", Name: "My Folder"},
		}))
	})

	result, err := svc.Lists.CreateFolder(context.Background(), &crm.ListFolderCreateRequest{
		Name: "My Folder",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Folder.ID != "f1" {
		t.Errorf("ID = %q", result.Folder.ID)
	}
}

// =============================================================================
// Timeline tests
// =============================================================================

func TestTimeline_CreateEvent(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/timeline/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.TimelineEventResponse{
			ID:              "evt1",
			EventTemplateID: "tmpl1",
			ObjectType:      "contacts",
			Tokens:          map[string]string{"key": "value"},
		}))
	})

	result, err := svc.Timeline.CreateEvent(context.Background(), &crm.TimelineEvent{
		EventTemplateID: "tmpl1",
		Tokens:          map[string]string{"key": "value"},
		ObjectID:        "123",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "evt1" {
		t.Errorf("ID = %q", result.ID)
	}
}

func TestTimeline_GetEventByID(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/timeline/events/tmpl1/evt1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.TimelineEventResponse{
			ID: "evt1", EventTemplateID: "tmpl1", ObjectType: "contacts",
			Tokens: map[string]string{},
		}))
	})

	result, err := svc.Timeline.GetEventByID(context.Background(), "tmpl1", "evt1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "evt1" {
		t.Errorf("ID = %q", result.ID)
	}
}

func TestTimeline_CreateTemplate(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/timeline/12345/event-templates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.TimelineEventTemplate{
			ID:         "tmpl1",
			Name:       "Test Template",
			ObjectType: "contacts",
			Tokens: []crm.TimelineEventTemplateToken{
				{Name: "key", Label: "Key", Type: "string"},
			},
		}))
	})

	result, err := svc.Timeline.CreateTemplate(context.Background(), 12345, &crm.TimelineEventTemplateCreateRequest{
		Name:       "Test Template",
		ObjectType: "contacts",
		Tokens: []crm.TimelineEventTemplateToken{
			{Name: "key", Label: "Key", Type: "string"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "tmpl1" {
		t.Errorf("ID = %q", result.ID)
	}
}

func TestTimeline_GetAllTemplates(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/timeline/12345/event-templates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.TimelineEventTemplateListResult{
			Results: []*crm.TimelineEventTemplate{
				{ID: "tmpl1", Name: "Template 1", ObjectType: "contacts"},
			},
		}))
	})

	result, err := svc.Timeline.GetAllTemplates(context.Background(), 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

func TestTimeline_ArchiveTemplate(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/timeline/12345/event-templates/tmpl1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Timeline.ArchiveTemplate(context.Background(), "tmpl1", 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestTimeline_CreateToken(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/timeline/12345/event-templates/tmpl1/tokens", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.TimelineEventTemplateToken{
			Name:  "new_token",
			Label: "New Token",
			Type:  "string",
		}))
	})

	result, err := svc.Timeline.CreateToken(context.Background(), "tmpl1", 12345, &crm.TimelineEventTemplateToken{
		Name:  "new_token",
		Label: "New Token",
		Type:  "string",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "new_token" {
		t.Errorf("name = %q", result.Name)
	}
}

// =============================================================================
// Extensions tests
// =============================================================================

func TestExtensions_CreateCallingSettings(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/extensions/calling/12345/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.CallingSettingsResponse{
			Name:      "My App",
			URL:       "https://myapp.com/calling",
			IsReady:   true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	result, err := svc.Extensions.CreateCallingSettings(context.Background(), 12345, &crm.CallingSettingsRequest{
		Name: "My App",
		URL:  "https://myapp.com/calling",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Name != "My App" {
		t.Errorf("name = %q", result.Name)
	}
}

func TestExtensions_GetCallingSettings(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/extensions/calling/12345/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.CallingSettingsResponse{
			Name: "My App", URL: "https://myapp.com/calling",
			CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}))
	})

	result, err := svc.Extensions.GetCallingSettings(context.Background(), 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.URL != "https://myapp.com/calling" {
		t.Errorf("url = %q", result.URL)
	}
}

func TestExtensions_ArchiveCallingSettings(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/extensions/calling/12345/settings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("method = %s, want DELETE", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	err := svc.Extensions.ArchiveCallingSettings(context.Background(), 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExtensions_CreateCard(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/extensions/cards/12345", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.PublicCardResponse{
			ID:        "card1",
			Title:     "Test Card",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}))
	})

	result, err := svc.Extensions.CreateCard(context.Background(), 12345, &crm.CardCreateRequest{
		Title: "Test Card",
		Fetch: crm.CardFetchBody{TargetURL: "https://example.com"},
		Display: crm.CardDisplayBody{Properties: []crm.CardDisplayProperty{
			{Name: "email", Label: "Email", DataType: "STRING"},
		}},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.ID != "card1" {
		t.Errorf("ID = %q", result.ID)
	}
}

func TestExtensions_GetAllCards(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/extensions/cards/12345", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.PublicCardListResponse{
			Results: []*crm.PublicCardResponse{
				{ID: "card1", Title: "Test Card", CreatedAt: time.Now(), UpdatedAt: time.Now()},
			},
		}))
	})

	result, err := svc.Extensions.GetAllCards(context.Background(), 12345)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Results) != 1 {
		t.Errorf("results = %d, want 1", len(result.Results))
	}
}

func TestExtensions_ReplaceVideoConferencingSettings(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/extensions/videoconferencing/settings/12345", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method = %s, want PUT", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.VideoConferencingExternalSettings{
			CreateMeetingURL: "https://example.com/create",
		}))
	})

	result, err := svc.Extensions.ReplaceVideoConferencingSettings(context.Background(), 12345, &crm.VideoConferencingExternalSettings{
		CreateMeetingURL: "https://example.com/create",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.CreateMeetingURL != "https://example.com/create" {
		t.Errorf("url = %q", result.CreateMeetingURL)
	}
}

// =============================================================================
// Commerce tests
// =============================================================================

func TestCommerce_Invoices_Create(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/invoices", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSON(t, crm.SimplePublicObject{
			ID:         "inv1",
			Properties: crm.Properties{"hs_invoice_number": "INV-001"},
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}))
	})

	invoice, err := svc.Commerce.Invoices.Create(context.Background(), &crm.SimplePublicObjectInputForCreate{
		Properties: crm.Properties{"hs_invoice_number": "INV-001"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if invoice.ID != "inv1" {
		t.Errorf("ID = %q", invoice.ID)
	}
}

func TestCommerce_Invoices_GetByID(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/invoices/inv1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.SimplePublicObjectWithAssociations{
			ID: "inv1", Properties: crm.Properties{}, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		}))
	})

	invoice, err := svc.Commerce.Invoices.GetByID(context.Background(), "inv1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if invoice.ID != "inv1" {
		t.Errorf("ID = %q", invoice.ID)
	}
}

func TestCommerce_Invoices_Search(t *testing.T) {
	svc, mux, teardown := setupCRM(t)
	defer teardown()

	mux.HandleFunc("/crm/v3/objects/invoices/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s, want POST", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSON(t, crm.SearchResult{
			Total:   1,
			Results: []*crm.SimplePublicObject{{ID: "inv1", Properties: crm.Properties{}, CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		}))
	})

	result, err := svc.Commerce.Invoices.Search(context.Background(), &crm.PublicObjectSearchRequest{
		Query: "INV-001",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 1 {
		t.Errorf("total = %d", result.Total)
	}
}
