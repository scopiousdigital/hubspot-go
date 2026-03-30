package crm

import (
	"encoding/json"
	"time"
)

// =============================================================================
// Owners models
// =============================================================================

// PublicOwner represents a HubSpot owner (user with CRM access).
type PublicOwner struct {
	ID                    string      `json:"id"`
	Email                 string      `json:"email,omitempty"`
	FirstName             string      `json:"firstName,omitempty"`
	LastName              string      `json:"lastName,omitempty"`
	Type                  string      `json:"type"` // PERSON, QUEUE
	UserID                *int64      `json:"userId,omitempty"`
	UserIDIncludingInactive *int64    `json:"userIdIncludingInactive,omitempty"`
	Teams                 []PublicTeam `json:"teams,omitempty"`
	CreatedAt             time.Time   `json:"createdAt"`
	UpdatedAt             time.Time   `json:"updatedAt"`
	Archived              bool        `json:"archived"`
}

// PublicTeam represents a HubSpot team an owner belongs to.
type PublicTeam struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Primary bool   `json:"primary"`
}

// OwnerListResult is returned by Owners.List.
type OwnerListResult struct {
	Results []*PublicOwner `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// OwnerListOptions configures an Owners.List request.
type OwnerListOptions struct {
	Email    string
	After    string
	Limit    int
	Archived bool
}

// OwnerGetByIDOptions configures an Owners.GetByID request.
type OwnerGetByIDOptions struct {
	IDProperty string // "id" or "userId"
	Archived   bool
}

// Owner type constants.
const (
	OwnerTypePerson = "PERSON"
	OwnerTypeQueue  = "QUEUE"
)

// =============================================================================
// Properties models
// =============================================================================

// Property represents a HubSpot CRM property definition.
type Property struct {
	Name                   string                      `json:"name"`
	Label                  string                      `json:"label"`
	Type                   string                      `json:"type"`
	FieldType              string                      `json:"fieldType"`
	Description            string                      `json:"description"`
	GroupName              string                      `json:"groupName"`
	Options                []PropertyOption            `json:"options"`
	DisplayOrder           int                         `json:"displayOrder,omitempty"`
	HasUniqueValue         bool                        `json:"hasUniqueValue,omitempty"`
	Hidden                 bool                        `json:"hidden,omitempty"`
	FormField              bool                        `json:"formField,omitempty"`
	Calculated             bool                        `json:"calculated,omitempty"`
	ExternalOptions        bool                        `json:"externalOptions,omitempty"`
	HubSpotDefined         bool                        `json:"hubspotDefined,omitempty"`
	ShowCurrencySymbol     bool                        `json:"showCurrencySymbol,omitempty"`
	CalculationFormula     string                      `json:"calculationFormula,omitempty"`
	ReferencedObjectType   string                      `json:"referencedObjectType,omitempty"`
	CreatedUserId          string                      `json:"createdUserId,omitempty"`
	UpdatedUserId          string                      `json:"updatedUserId,omitempty"`
	ModificationMetadata   *PropertyModificationMetadata `json:"modificationMetadata,omitempty"`
	CreatedAt              *time.Time                  `json:"createdAt,omitempty"`
	UpdatedAt              *time.Time                  `json:"updatedAt,omitempty"`
	ArchivedAt             *time.Time                  `json:"archivedAt,omitempty"`
	Archived               bool                        `json:"archived,omitempty"`
}

// PropertyOption represents a selectable option for enumeration properties.
type PropertyOption struct {
	Label        string `json:"label"`
	Value        string `json:"value"`
	Description  string `json:"description,omitempty"`
	DisplayOrder int    `json:"displayOrder,omitempty"`
	Hidden       bool   `json:"hidden"`
}

// PropertyOptionInput is used when creating/updating property options.
type PropertyOptionInput struct {
	Label        string `json:"label"`
	Value        string `json:"value"`
	Description  string `json:"description,omitempty"`
	DisplayOrder int    `json:"displayOrder,omitempty"`
	Hidden       bool   `json:"hidden"`
}

// PropertyModificationMetadata contains metadata about property mutability.
type PropertyModificationMetadata struct {
	Archivable         bool `json:"archivable"`
	ReadOnlyDefinition bool `json:"readOnlyDefinition"`
	ReadOnlyValue      bool `json:"readOnlyValue"`
	ReadOnlyOptions    bool `json:"readOnlyOptions,omitempty"`
}

// PropertyCreate is the input for creating a new property.
type PropertyCreate struct {
	Name                 string              `json:"name"`
	Label                string              `json:"label"`
	Type                 string              `json:"type"`      // string, number, date, datetime, enumeration, bool
	FieldType            string              `json:"fieldType"` // textarea, text, date, file, number, select, radio, checkbox, booleancheckbox, calculation_equation
	GroupName            string              `json:"groupName"`
	Description          string              `json:"description,omitempty"`
	Options              []PropertyOptionInput `json:"options,omitempty"`
	DisplayOrder         int                 `json:"displayOrder,omitempty"`
	HasUniqueValue       bool                `json:"hasUniqueValue,omitempty"`
	Hidden               bool                `json:"hidden,omitempty"`
	FormField            bool                `json:"formField,omitempty"`
	ExternalOptions      bool                `json:"externalOptions,omitempty"`
	ReferencedObjectType string              `json:"referencedObjectType,omitempty"`
	CalculationFormula   string              `json:"calculationFormula,omitempty"`
}

// PropertyUpdate is the input for updating an existing property.
type PropertyUpdate struct {
	Label              string              `json:"label,omitempty"`
	Type               string              `json:"type,omitempty"`
	FieldType          string              `json:"fieldType,omitempty"`
	GroupName          string              `json:"groupName,omitempty"`
	Description        string              `json:"description,omitempty"`
	Options            []PropertyOptionInput `json:"options,omitempty"`
	DisplayOrder       int                 `json:"displayOrder,omitempty"`
	Hidden             *bool               `json:"hidden,omitempty"`
	FormField          *bool               `json:"formField,omitempty"`
	CalculationFormula string              `json:"calculationFormula,omitempty"`
}

// PropertyGroup represents a group of CRM properties.
type PropertyGroup struct {
	Name         string `json:"name"`
	Label        string `json:"label"`
	DisplayOrder int    `json:"displayOrder"`
	Archived     bool   `json:"archived"`
}

// PropertyGroupCreate is the input for creating a property group.
type PropertyGroupCreate struct {
	Name         string `json:"name"`
	Label        string `json:"label"`
	DisplayOrder int    `json:"displayOrder,omitempty"`
}

// PropertyGroupUpdate is the input for updating a property group.
type PropertyGroupUpdate struct {
	Label        string `json:"label,omitempty"`
	DisplayOrder int    `json:"displayOrder,omitempty"`
}

// PropertyListResult is returned by Properties.GetAll.
type PropertyListResult struct {
	Results []*Property `json:"results"`
}

// PropertyGroupListResult is returned by Properties.Groups.GetAll.
type PropertyGroupListResult struct {
	Results []*PropertyGroup `json:"results"`
}

// BatchPropertyCreateInput is the input for batch creating properties.
type BatchPropertyCreateInput struct {
	Inputs []PropertyCreate `json:"inputs"`
}

// BatchPropertyReadInput is the input for batch reading properties by name.
type BatchPropertyReadInput struct {
	Inputs []PropertyNameInput `json:"inputs"`
}

// PropertyNameInput wraps a property name for batch operations.
type PropertyNameInput struct {
	Name string `json:"name"`
}

// BatchPropertyResult is the response from batch property operations.
type BatchPropertyResult struct {
	Status      string      `json:"status"`
	Results     []*Property `json:"results"`
	StartedAt   time.Time   `json:"startedAt"`
	CompletedAt time.Time   `json:"completedAt"`
	NumErrors   int         `json:"numErrors,omitempty"`
	Errors      []StandardError `json:"errors,omitempty"`
}

// =============================================================================
// Pipelines models
// =============================================================================

// Pipeline represents a HubSpot pipeline.
type Pipeline struct {
	ID           string          `json:"id"`
	Label        string          `json:"label"`
	DisplayOrder int             `json:"displayOrder"`
	Stages       []PipelineStage `json:"stages"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	Archived     bool            `json:"archived"`
	ArchivedAt   *time.Time      `json:"archivedAt,omitempty"`
}

// PipelineStage represents a stage within a pipeline.
type PipelineStage struct {
	ID               string            `json:"id"`
	Label            string            `json:"label"`
	DisplayOrder     int               `json:"displayOrder"`
	Metadata         map[string]string `json:"metadata,omitempty"`
	WritePermissions string            `json:"writePermissions,omitempty"` // CRM_PERMISSIONS_ENFORCEMENT, READ_ONLY, INTERNAL_ONLY
	CreatedAt        time.Time         `json:"createdAt"`
	UpdatedAt        time.Time         `json:"updatedAt"`
	Archived         bool              `json:"archived"`
	ArchivedAt       *time.Time        `json:"archivedAt,omitempty"`
}

// PipelineInput is the input for creating a pipeline.
type PipelineInput struct {
	Label        string             `json:"label"`
	DisplayOrder int                `json:"displayOrder"`
	Stages       []PipelineStageInput `json:"stages"`
}

// PipelinePatchInput is the input for updating a pipeline.
type PipelinePatchInput struct {
	Label        string `json:"label,omitempty"`
	DisplayOrder int    `json:"displayOrder,omitempty"`
	Archived     *bool  `json:"archived,omitempty"`
}

// PipelineStageInput is the input for creating a pipeline stage.
type PipelineStageInput struct {
	Label        string            `json:"label"`
	DisplayOrder int               `json:"displayOrder"`
	Metadata     map[string]string `json:"metadata,omitempty"`
}

// PipelineStagePatchInput is the input for updating a pipeline stage.
type PipelineStagePatchInput struct {
	Label        string            `json:"label,omitempty"`
	DisplayOrder int               `json:"displayOrder,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	Archived     *bool             `json:"archived,omitempty"`
}

// PublicAuditInfo represents an audit log entry for a pipeline or stage.
type PublicAuditInfo struct {
	Identifier string     `json:"identifier"`
	Action     string     `json:"action"`
	PortalID   int        `json:"portalId"`
	FromUserID *int64     `json:"fromUserId,omitempty"`
	Message    string     `json:"message,omitempty"`
	Timestamp  *time.Time `json:"timestamp,omitempty"`
	RawObject  any        `json:"rawObject,omitempty"`
}

// PipelineListResult is returned by Pipelines.GetAll.
type PipelineListResult struct {
	Results []*Pipeline `json:"results"`
}

// PipelineStageListResult is returned by PipelineStages.GetAll.
type PipelineStageListResult struct {
	Results []*PipelineStage `json:"results"`
}

// AuditInfoListResult is returned by PipelineAudits operations.
type AuditInfoListResult struct {
	Results []*PublicAuditInfo `json:"results"`
}

// =============================================================================
// Schemas models
// =============================================================================

// ObjectTypeDefinitionLabels defines singular/plural labels for a custom object.
type ObjectTypeDefinitionLabels struct {
	Singular string `json:"singular,omitempty"`
	Plural   string `json:"plural,omitempty"`
}

// ObjectSchema represents a complete custom object schema with associations and properties.
type ObjectSchema struct {
	ID                         string                     `json:"id"`
	Name                       string                     `json:"name"`
	Labels                     ObjectTypeDefinitionLabels `json:"labels"`
	Description                string                     `json:"description,omitempty"`
	ObjectTypeID               string                     `json:"objectTypeId,omitempty"`
	FullyQualifiedName         string                     `json:"fullyQualifiedName,omitempty"`
	PrimaryDisplayProperty     string                     `json:"primaryDisplayProperty,omitempty"`
	SecondaryDisplayProperties []string                   `json:"secondaryDisplayProperties,omitempty"`
	RequiredProperties         []string                   `json:"requiredProperties"`
	SearchableProperties       []string                   `json:"searchableProperties,omitempty"`
	Properties                 []Property                 `json:"properties"`
	Associations               []SchemaAssociationDefinition `json:"associations"`
	CreatedAt                  *time.Time                 `json:"createdAt,omitempty"`
	UpdatedAt                  *time.Time                 `json:"updatedAt,omitempty"`
	Archived                   bool                       `json:"archived,omitempty"`
	CreatedByUserID            *int64                     `json:"createdByUserId,omitempty"`
	UpdatedByUserID            *int64                     `json:"updatedByUserId,omitempty"`
}

// ObjectSchemaEgg is the input for creating a custom object schema.
type ObjectSchemaEgg struct {
	Name                       string                          `json:"name"`
	Labels                     ObjectTypeDefinitionLabels      `json:"labels"`
	Description                string                          `json:"description,omitempty"`
	PrimaryDisplayProperty     string                          `json:"primaryDisplayProperty,omitempty"`
	SecondaryDisplayProperties []string                        `json:"secondaryDisplayProperties,omitempty"`
	RequiredProperties         []string                        `json:"requiredProperties"`
	SearchableProperties       []string                        `json:"searchableProperties,omitempty"`
	Properties                 []ObjectTypePropertyCreate      `json:"properties"`
	AssociatedObjects          []string                        `json:"associatedObjects"`
}

// ObjectTypePropertyCreate is the input for creating a property within a schema.
type ObjectTypePropertyCreate struct {
	Name                    string              `json:"name"`
	Label                   string              `json:"label"`
	Type                    string              `json:"type"`
	FieldType               string              `json:"fieldType"`
	GroupName               string              `json:"groupName,omitempty"`
	Description             string              `json:"description,omitempty"`
	Options                 []PropertyOptionInput `json:"options,omitempty"`
	DisplayOrder            int                 `json:"displayOrder,omitempty"`
	HasUniqueValue          bool                `json:"hasUniqueValue,omitempty"`
	Hidden                  bool                `json:"hidden,omitempty"`
	FormField               bool                `json:"formField,omitempty"`
	ShowCurrencySymbol      bool                `json:"showCurrencySymbol,omitempty"`
	ReferencedObjectType    string              `json:"referencedObjectType,omitempty"`
	TextDisplayHint         string              `json:"textDisplayHint,omitempty"`
	NumberDisplayHint       string              `json:"numberDisplayHint,omitempty"`
	OptionSortStrategy      string              `json:"optionSortStrategy,omitempty"`
	SearchableInGlobalSearch bool               `json:"searchableInGlobalSearch,omitempty"`
}

// ObjectTypeDefinition represents a custom object type definition (returned by update).
type ObjectTypeDefinition struct {
	ID                         string                     `json:"id"`
	Name                       string                     `json:"name"`
	Labels                     ObjectTypeDefinitionLabels `json:"labels"`
	Description                string                     `json:"description,omitempty"`
	ObjectTypeID               string                     `json:"objectTypeId,omitempty"`
	FullyQualifiedName         string                     `json:"fullyQualifiedName,omitempty"`
	PrimaryDisplayProperty     string                     `json:"primaryDisplayProperty,omitempty"`
	SecondaryDisplayProperties []string                   `json:"secondaryDisplayProperties,omitempty"`
	RequiredProperties         []string                   `json:"requiredProperties"`
	SearchableProperties       []string                   `json:"searchableProperties,omitempty"`
	PortalID                   int                        `json:"portalId,omitempty"`
	CreatedAt                  *time.Time                 `json:"createdAt,omitempty"`
	UpdatedAt                  *time.Time                 `json:"updatedAt,omitempty"`
	Archived                   bool                       `json:"archived,omitempty"`
}

// ObjectTypeDefinitionPatch is the input for updating a custom object type.
type ObjectTypeDefinitionPatch struct {
	Labels                     *ObjectTypeDefinitionLabels `json:"labels,omitempty"`
	Description                string                      `json:"description,omitempty"`
	PrimaryDisplayProperty     string                      `json:"primaryDisplayProperty,omitempty"`
	SecondaryDisplayProperties []string                    `json:"secondaryDisplayProperties,omitempty"`
	RequiredProperties         []string                    `json:"requiredProperties,omitempty"`
	SearchableProperties       []string                    `json:"searchableProperties,omitempty"`
	ClearDescription           bool                        `json:"clearDescription,omitempty"`
	Restorable                 bool                        `json:"restorable,omitempty"`
}

// SchemaAssociationDefinition represents an association in a schema.
type SchemaAssociationDefinition struct {
	ID               string     `json:"id"`
	FromObjectTypeID string     `json:"fromObjectTypeId"`
	ToObjectTypeID   string     `json:"toObjectTypeId"`
	Name             string     `json:"name,omitempty"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty"`
}

// SchemaAssociationDefinitionEgg is the input for creating an association in a schema.
type SchemaAssociationDefinitionEgg struct {
	FromObjectTypeID string `json:"fromObjectTypeId"`
	ToObjectTypeID   string `json:"toObjectTypeId"`
	Name             string `json:"name,omitempty"`
}

// ObjectSchemaListResult is returned by Schemas.GetAll.
type ObjectSchemaListResult struct {
	Results []*ObjectSchema `json:"results"`
}

// =============================================================================
// Associations models
// =============================================================================

// PublicAssociation represents a v3 association between two objects.
type PublicAssociation struct {
	From ObjectID `json:"from"`
	To   ObjectID `json:"to"`
	Type string   `json:"type"`
}

// BatchPublicAssociationInput is input for v3 batch association operations.
type BatchPublicAssociationInput struct {
	Inputs []PublicAssociation `json:"inputs"`
}

// BatchPublicObjectIDInput is input for v3 batch reading associations.
type BatchPublicObjectIDInput struct {
	Inputs []ObjectID `json:"inputs"`
}

// PublicAssociationResult is a single v3 association result.
type PublicAssociationResult struct {
	From    ObjectID `json:"from"`
	To      []AssociatedID `json:"to"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// BatchPublicAssociationResult is returned by v3 batch create/archive.
type BatchPublicAssociationResult struct {
	Status      string                  `json:"status"`
	Results     []*PublicAssociation    `json:"results"`
	StartedAt   time.Time              `json:"startedAt"`
	CompletedAt time.Time              `json:"completedAt"`
	NumErrors   int                    `json:"numErrors,omitempty"`
	Errors      []StandardError        `json:"errors,omitempty"`
}

// BatchPublicAssociationMultiResult is returned by v3 batch read.
type BatchPublicAssociationMultiResult struct {
	Status      string                     `json:"status"`
	Results     []*PublicAssociationResult `json:"results"`
	StartedAt   time.Time                 `json:"startedAt"`
	CompletedAt time.Time                 `json:"completedAt"`
	NumErrors   int                       `json:"numErrors,omitempty"`
	Errors      []StandardError           `json:"errors,omitempty"`
}

// --- Associations v4 models ---

// AssociationV4Spec defines the type/category of a v4 association.
type AssociationV4Spec struct {
	AssociationCategory string `json:"associationCategory"` // HUBSPOT_DEFINED, USER_DEFINED, INTEGRATOR_DEFINED
	AssociationTypeID   int32  `json:"associationTypeId"`
}

// AssociationV4SpecWithLabel is a v4 association spec with a label.
type AssociationV4SpecWithLabel struct {
	Category string `json:"category"`
	TypeID   int32  `json:"typeId"`
	Label    string `json:"label,omitempty"`
}

// LabelsBetweenObjectPair represents labeled associations between two objects.
type LabelsBetweenObjectPair struct {
	FromObjectTypeID string               `json:"fromObjectTypeId"`
	FromObjectID     int64                `json:"fromObjectId"`
	ToObjectTypeID   string               `json:"toObjectTypeId"`
	ToObjectID       int64                `json:"toObjectId"`
	Labels           []AssociationV4SpecWithLabel `json:"labels,omitempty"`
}

// MultiAssociatedObjectWithLabel represents a v4 association target with labels.
type MultiAssociatedObjectWithLabel struct {
	ToObjectID         int64                        `json:"toObjectId"`
	AssociationTypes   []AssociationV4SpecWithLabel `json:"associationTypes"`
}

// CollectionMultiAssociatedObjectWithLabel is the paged result for v4 getPage.
type CollectionMultiAssociatedObjectWithLabel struct {
	Results []*MultiAssociatedObjectWithLabel `json:"results"`
	Paging  *ForwardPaging                    `json:"paging,omitempty"`
}

// AssociationV4MultiPost is a single item in a v4 batch create.
type AssociationV4MultiPost struct {
	From ObjectID          `json:"from"`
	To   ObjectID          `json:"to"`
	Types []AssociationV4Spec `json:"types"`
}

// BatchAssociationV4MultiPostInput is the input for v4 batch create.
type BatchAssociationV4MultiPostInput struct {
	Inputs []AssociationV4MultiPost `json:"inputs"`
}

// AssociationV4MultiArchive is a single item in a v4 batch archive.
type AssociationV4MultiArchive struct {
	From ObjectID          `json:"from"`
	To   ObjectID          `json:"to"`
}

// BatchAssociationV4MultiArchiveInput is the input for v4 batch archive.
type BatchAssociationV4MultiArchiveInput struct {
	Inputs []AssociationV4MultiArchive `json:"inputs"`
}

// AssociationV4DefaultMultiPost is a single item in a v4 batch default create.
type AssociationV4DefaultMultiPost struct {
	From ObjectID `json:"from"`
	To   ObjectID `json:"to"`
}

// BatchAssociationV4DefaultMultiPostInput is the input for v4 batch default create.
type BatchAssociationV4DefaultMultiPostInput struct {
	Inputs []AssociationV4DefaultMultiPost `json:"inputs"`
}

// FetchAssociationsBatchRequest is a single item for batch fetching associations.
type FetchAssociationsBatchRequest struct {
	ID string `json:"id"`
}

// BatchFetchAssociationsInput is the input for v4 batch getPage.
type BatchFetchAssociationsInput struct {
	Inputs []FetchAssociationsBatchRequest `json:"inputs"`
}

// BatchLabelsBetweenObjectPairResult is returned by v4 batch create.
type BatchLabelsBetweenObjectPairResult struct {
	Status      string                     `json:"status"`
	Results     []*LabelsBetweenObjectPair `json:"results"`
	StartedAt   time.Time                  `json:"startedAt"`
	CompletedAt time.Time                  `json:"completedAt"`
	NumErrors   int                        `json:"numErrors,omitempty"`
	Errors      []StandardError            `json:"errors,omitempty"`
}

// PublicDefaultAssociation represents a default association.
type PublicDefaultAssociation struct {
	From              ObjectID `json:"from"`
	To                ObjectID `json:"to"`
	AssociationTypeID int32    `json:"associationTypeId,omitempty"`
}

// BatchPublicDefaultAssociationResult is returned by v4 createDefault.
type BatchPublicDefaultAssociationResult struct {
	Status      string                      `json:"status"`
	Results     []*PublicDefaultAssociation `json:"results"`
	StartedAt   time.Time                  `json:"startedAt"`
	CompletedAt time.Time                  `json:"completedAt"`
	NumErrors   int                        `json:"numErrors,omitempty"`
	Errors      []StandardError            `json:"errors,omitempty"`
}

// BatchAssociationMultiWithLabelResult is returned by v4 batch getPage.
type BatchAssociationMultiWithLabelResult struct {
	Status      string                                    `json:"status"`
	Results     []*CollectionMultiAssociatedObjectWithLabel `json:"results"`
	StartedAt   time.Time                                 `json:"startedAt"`
	CompletedAt time.Time                                 `json:"completedAt"`
	NumErrors   int                                       `json:"numErrors,omitempty"`
	Errors      []StandardError                           `json:"errors,omitempty"`
}

// --- Associations v4 schema models ---

// PublicAssociationDefinitionCreateRequest is input for creating a custom association definition.
type PublicAssociationDefinitionCreateRequest struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

// PublicAssociationDefinitionUpdateRequest is input for updating a custom association definition.
type PublicAssociationDefinitionUpdateRequest struct {
	Label                string `json:"label"`
	AssociationTypeID    int32  `json:"associationTypeId"`
}

// AssociationDefinitionSpecWithLabelResult is the result of listing association definitions.
type AssociationDefinitionSpecWithLabelResult struct {
	Results []*AssociationV4SpecWithLabel `json:"results"`
}

// =============================================================================
// Imports models
// =============================================================================

// PublicImportResponse represents a HubSpot import.
type PublicImportResponse struct {
	ID                 string              `json:"id"`
	State              string              `json:"state"` // STARTED, PROCESSING, DONE, FAILED, CANCELED, DEFERRED, REVERTED
	ImportName         string              `json:"importName,omitempty"`
	ImportSource       string              `json:"importSource,omitempty"` // API, CRM_UI, IMPORT, MOBILE_ANDROID, MOBILE_IOS, SALESFORCE
	OptOutImport       bool                `json:"optOutImport"`
	Metadata           ImportMetadata      `json:"metadata"`
	MappedObjectTypeIds []string           `json:"mappedObjectTypeIds"`
	ImportRequestJSON  json.RawMessage     `json:"importRequestJson,omitempty"`
	CreatedAt          time.Time           `json:"createdAt"`
	UpdatedAt          time.Time           `json:"updatedAt"`
}

// ImportMetadata contains counters and file info for an import.
type ImportMetadata struct {
	Counters    map[string]int          `json:"counters"`
	FileIDs     []string                `json:"fileIds"`
	ObjectLists []PublicObjectListRecord `json:"objectLists"`
}

// PublicObjectListRecord represents an object list within an import.
type PublicObjectListRecord struct {
	ListID       string `json:"listId"`
	ObjectTypeID string `json:"objectTypeId"`
}

// ImportListResult is returned by Imports.List.
type ImportListResult struct {
	Results []*PublicImportResponse `json:"results"`
	Paging  *ForwardPaging          `json:"paging,omitempty"`
}

// ActionResponse is a generic action response.
type ActionResponse struct {
	Status      string     `json:"status"`
	StartedAt   time.Time  `json:"startedAt"`
	CompletedAt time.Time  `json:"completedAt"`
	RequestedAt *time.Time `json:"requestedAt,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
}

// PublicImportError represents an error from an import.
type PublicImportError struct {
	ErrorType    string `json:"errorType"`
	ObjectType   string `json:"objectType,omitempty"`
	InvalidValue string `json:"invalidValue,omitempty"`
	ExtraContext string `json:"extraContext,omitempty"`
	ObjectTypeID string `json:"objectTypeId,omitempty"`
	KnownColumnNumber int `json:"knownColumnNumber,omitempty"`
	SourceData   ImportRowCore `json:"sourceData,omitempty"`
}

// ImportRowCore represents the source data of an import error row.
type ImportRowCore struct {
	LineNumber int      `json:"lineNumber"`
	Values     []string `json:"values"`
}

// ImportErrorListResult is returned by Imports.GetErrors.
type ImportErrorListResult struct {
	Results []*PublicImportError `json:"results"`
	Paging  *ForwardPaging       `json:"paging,omitempty"`
}

// ImportListOptions configures an Imports.List request.
type ImportListOptions struct {
	After  string
	Before string
	Limit  int
}

// ImportGetErrorsOptions configures an Imports.GetErrors request.
type ImportGetErrorsOptions struct {
	After              string
	Limit              int
	IncludeErrorMessage bool
	IncludeRowData     bool
}

// =============================================================================
// Exports models
// =============================================================================

// ExportRequest is the input for starting an export.
type ExportRequest struct {
	ExportType     string   `json:"exportType"`     // VIEW, LIST
	Format         string   `json:"format"`          // XLS, XLSX, CSV
	ExportName     string   `json:"exportName"`
	ObjectType     string   `json:"objectType"`
	ObjectProperties []string `json:"objectProperties"`
	Language       string   `json:"language"`
	AssociatedObjectType string `json:"associatedObjectType,omitempty"`
	ExportInternalValuesOptions []string `json:"exportInternalValuesOptions,omitempty"`
	OverrideAssociatedObjectsPerDefinitionPerRowLimit bool `json:"overrideAssociatedObjectsPerDefinitionPerRowLimit,omitempty"`
	ListID         string   `json:"listId,omitempty"`        // For LIST exports
	PublicCrmSearchRequest *ExportCrmSearchRequest `json:"publicCrmSearchRequest,omitempty"` // For VIEW exports
}

// ExportCrmSearchRequest defines a search filter for VIEW exports.
type ExportCrmSearchRequest struct {
	Filters []ExportFilter `json:"filters,omitempty"`
	Sorts   []string       `json:"sorts,omitempty"`
	Query   string         `json:"query,omitempty"`
}

// ExportFilter is a single filter for export search.
type ExportFilter struct {
	PropertyName string `json:"propertyName"`
	Operator     string `json:"operator"`
	Value        string `json:"value,omitempty"`
}

// TaskLocator is returned when starting an export.
type TaskLocator struct {
	ID    string            `json:"id"`
	Links map[string]string `json:"links,omitempty"`
}

// ExportStatusResponse is returned when checking export status.
type ExportStatusResponse struct {
	Status      string     `json:"status"` // PENDING, PROCESSING, CANCELED, COMPLETE
	Result      string     `json:"result,omitempty"` // Download URI
	StartedAt   time.Time  `json:"startedAt"`
	CompletedAt time.Time  `json:"completedAt"`
	RequestedAt *time.Time `json:"requestedAt,omitempty"`
	NumErrors   int        `json:"numErrors,omitempty"`
	Errors      []StandardError `json:"errors,omitempty"`
	Links       map[string]string `json:"links,omitempty"`
}

// =============================================================================
// Lists models
// =============================================================================

// ListCreateRequest is the input for creating a list.
type ListCreateRequest struct {
	Name            string `json:"name"`
	ObjectTypeID    string `json:"objectTypeId"`
	ProcessingType  string `json:"processingType"` // MANUAL, DYNAMIC, SNAPSHOT
	FilterBranch    any    `json:"filterBranch,omitempty"` // Complex filter definition
	ListFolderID    *int64 `json:"listFolderId,omitempty"`
}

// ListCreateResponse is returned by Lists.Create.
type ListCreateResponse struct {
	ListID          string `json:"listId"`
	ListVersion     int    `json:"listVersion"`
	Name            string `json:"name"`
	ObjectTypeID    string `json:"objectTypeId"`
	ProcessingType  string `json:"processingType"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
}

// ListFetchResponse is returned by Lists.GetByID.
type ListFetchResponse struct {
	List *HubSpotList `json:"list"`
}

// HubSpotList represents a HubSpot list.
type HubSpotList struct {
	ListID          string `json:"listId"`
	ListVersion     int    `json:"listVersion"`
	Name            string `json:"name"`
	ObjectTypeID    string `json:"objectTypeId"`
	ProcessingType  string `json:"processingType"`
	Size            int    `json:"size,omitempty"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
	FilterBranch    any    `json:"filterBranch,omitempty"`
	ListFolderID    *int64 `json:"listFolderId,omitempty"`
}

// ListsByIDResponse is returned by Lists.GetAll.
type ListsByIDResponse struct {
	Lists map[string]*HubSpotList `json:"lists"`
}

// ListSearchRequest is the input for searching lists.
type ListSearchRequest struct {
	Query           string `json:"query,omitempty"`
	ListIds         []string `json:"listIds,omitempty"`
	Offset          int    `json:"offset,omitempty"`
	Count           int    `json:"count,omitempty"`
	ProcessingTypes []string `json:"processingTypes,omitempty"`
	AdditionalProperties []string `json:"additionalProperties,omitempty"`
}

// ListSearchResponse is returned by Lists.Search.
type ListSearchResponse struct {
	Lists   []*HubSpotList `json:"lists"`
	HasMore bool           `json:"hasMore"`
	Offset  int            `json:"offset"`
	Total   int            `json:"total"`
}

// ListFilterUpdateRequest is the input for updating list filters.
type ListFilterUpdateRequest struct {
	FilterBranch any `json:"filterBranch"`
}

// ListUpdateResponse is returned by list update operations.
type ListUpdateResponse struct {
	List *HubSpotList `json:"list"`
}

// MembershipChangeRequest is the input for add-and-remove operations.
type MembershipChangeRequest struct {
	RecordIdsToAdd    []string `json:"recordIdsToAdd,omitempty"`
	RecordIdsToRemove []string `json:"recordIdsToRemove,omitempty"`
}

// MembershipsUpdateResponse is returned by membership add/remove operations.
type MembershipsUpdateResponse struct {
	RecordIdsAdded    []string `json:"recordIdsAdded,omitempty"`
	RecordIdsRemoved  []string `json:"recordIdsRemoved,omitempty"`
}

// JoinTimeAndRecordID represents a membership entry with join time.
type JoinTimeAndRecordID struct {
	RecordID  string `json:"recordId"`
	AddedAt   string `json:"addedAt"`
}

// MembershipPageResult is returned by Memberships.GetPage.
type MembershipPageResult struct {
	Results []*JoinTimeAndRecordID `json:"results"`
	Paging  *ForwardPaging          `json:"paging,omitempty"`
}

// RecordListMembership represents a list membership for a record.
type RecordListMembership struct {
	ListID    string `json:"listId"`
	ListName  string `json:"listName,omitempty"`
}

// RecordListMembershipResult is returned by Memberships.GetLists.
type RecordListMembershipResult struct {
	Results []*RecordListMembership `json:"results"`
}

// ListFolderCreateRequest is the input for creating a list folder.
type ListFolderCreateRequest struct {
	Name           string `json:"name"`
	ParentFolderID string `json:"parentFolderId,omitempty"`
}

// ListFolderCreateResponse is returned by Folders.Create.
type ListFolderCreateResponse struct {
	Folder *ListFolder `json:"folder"`
}

// ListFolder represents a list folder.
type ListFolder struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ParentFolderID string `json:"parentFolderId,omitempty"`
	ChildCount     int    `json:"childCount,omitempty"`
}

// ListFolderFetchResponse is returned by Folders.Get.
type ListFolderFetchResponse struct {
	Folder   *ListFolder   `json:"folder,omitempty"`
	Folders  []*ListFolder `json:"folders,omitempty"`
}

// ListMoveRequest is the input for moving a list to a folder.
type ListMoveRequest struct {
	ListID   string `json:"listId"`
	FolderID string `json:"folderId"`
}

// PublicMigrationMapping represents an ILS to v3 list ID mapping.
type PublicMigrationMapping struct {
	LegacyListID string `json:"legacyListId"`
	ListID       string `json:"listId"`
}

// PublicBatchMigrationMapping is returned by batch mapping translation.
type PublicBatchMigrationMapping struct {
	Results []*PublicMigrationMapping `json:"results"`
}

// MembershipListOptions configures a Memberships.GetPage request.
type MembershipListOptions struct {
	After  string
	Before string
	Limit  int
}

// =============================================================================
// Timeline models
// =============================================================================

// TimelineEvent represents a timeline event to create.
type TimelineEvent struct {
	EventTemplateID string            `json:"eventTemplateId"`
	Tokens          map[string]string `json:"tokens"`
	ID              string            `json:"id,omitempty"`
	Email           string            `json:"email,omitempty"`
	UTK             string            `json:"utk,omitempty"`
	Domain          string            `json:"domain,omitempty"`
	ObjectID        string            `json:"objectId,omitempty"`
	Timestamp       *time.Time        `json:"timestamp,omitempty"`
	ExtraData       any               `json:"extraData,omitempty"`
	TimelineIFrame  *TimelineEventIFrame `json:"timelineIFrame,omitempty"`
}

// TimelineEventIFrame defines an iframe to display with a timeline event.
type TimelineEventIFrame struct {
	HeaderLabel string `json:"headerLabel"`
	URL         string `json:"url"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

// TimelineEventResponse represents a created/retrieved timeline event.
type TimelineEventResponse struct {
	ID              string            `json:"id"`
	EventTemplateID string            `json:"eventTemplateId"`
	ObjectType      string            `json:"objectType"`
	Tokens          map[string]string `json:"tokens"`
	Email           string            `json:"email,omitempty"`
	UTK             string            `json:"utk,omitempty"`
	Domain          string            `json:"domain,omitempty"`
	ObjectID        string            `json:"objectId,omitempty"`
	Timestamp       *time.Time        `json:"timestamp,omitempty"`
	CreatedAt       *time.Time        `json:"createdAt,omitempty"`
	ExtraData       any               `json:"extraData,omitempty"`
	TimelineIFrame  *TimelineEventIFrame `json:"timelineIFrame,omitempty"`
}

// EventDetail is detailed event information.
type EventDetail struct {
	json.RawMessage
}

// BatchTimelineEventInput is the input for batch creating timeline events.
type BatchTimelineEventInput struct {
	Inputs []TimelineEvent `json:"inputs"`
}

// BatchTimelineEventResult is returned by batch timeline event creation.
type BatchTimelineEventResult struct {
	Status      string                    `json:"status"`
	Results     []*TimelineEventResponse `json:"results"`
	StartedAt   time.Time                `json:"startedAt"`
	CompletedAt time.Time                `json:"completedAt"`
	NumErrors   int                      `json:"numErrors,omitempty"`
	Errors      []StandardError          `json:"errors,omitempty"`
}

// TimelineEventTemplate represents a timeline event template.
type TimelineEventTemplate struct {
	ID             string                       `json:"id"`
	Name           string                       `json:"name"`
	ObjectType     string                       `json:"objectType"`
	HeaderTemplate string                       `json:"headerTemplate,omitempty"`
	DetailTemplate string                       `json:"detailTemplate,omitempty"`
	Tokens         []TimelineEventTemplateToken `json:"tokens"`
	CreatedAt      *time.Time                   `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time                   `json:"updatedAt,omitempty"`
}

// TimelineEventTemplateCreateRequest is the input for creating a template.
type TimelineEventTemplateCreateRequest struct {
	Name           string                       `json:"name"`
	ObjectType     string                       `json:"objectType"`
	HeaderTemplate string                       `json:"headerTemplate,omitempty"`
	DetailTemplate string                       `json:"detailTemplate,omitempty"`
	Tokens         []TimelineEventTemplateToken `json:"tokens"`
}

// TimelineEventTemplateUpdateRequest is the input for updating a template.
type TimelineEventTemplateUpdateRequest struct {
	ID             string                       `json:"id"`
	Name           string                       `json:"name"`
	HeaderTemplate string                       `json:"headerTemplate,omitempty"`
	DetailTemplate string                       `json:"detailTemplate,omitempty"`
	Tokens         []TimelineEventTemplateToken `json:"tokens"`
}

// TimelineEventTemplateToken represents a token in a template.
type TimelineEventTemplateToken struct {
	Name               string                             `json:"name"`
	Label              string                             `json:"label"`
	Type               string                             `json:"type"` // date, enumeration, number, string
	ObjectPropertyName string                             `json:"objectPropertyName,omitempty"`
	Options            []TimelineEventTemplateTokenOption `json:"options,omitempty"`
	CreatedAt          *time.Time                         `json:"createdAt,omitempty"`
	UpdatedAt          *time.Time                         `json:"updatedAt,omitempty"`
}

// TimelineEventTemplateTokenUpdateRequest is the input for updating a token.
type TimelineEventTemplateTokenUpdateRequest struct {
	Label              string                             `json:"label"`
	ObjectPropertyName string                             `json:"objectPropertyName,omitempty"`
	Options            []TimelineEventTemplateTokenOption `json:"options,omitempty"`
}

// TimelineEventTemplateTokenOption is a selectable option for an enumeration token.
type TimelineEventTemplateTokenOption struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// TimelineEventTemplateListResult is returned by Templates.GetAll.
type TimelineEventTemplateListResult struct {
	Results []*TimelineEventTemplate `json:"results"`
}

// =============================================================================
// Extensions models
// =============================================================================

// --- Calling extension models ---

// CallingSettingsRequest is the input for creating calling settings.
type CallingSettingsRequest struct {
	Name                     string `json:"name"`
	URL                      string `json:"url"`
	Width                    int    `json:"width,omitempty"`
	Height                   int    `json:"height,omitempty"`
	IsReady                  bool   `json:"isReady,omitempty"`
	SupportsCustomObjects    bool   `json:"supportsCustomObjects,omitempty"`
}

// CallingSettingsPatchRequest is the input for updating calling settings.
type CallingSettingsPatchRequest struct {
	Name                     string `json:"name,omitempty"`
	URL                      string `json:"url,omitempty"`
	Width                    int    `json:"width,omitempty"`
	Height                   int    `json:"height,omitempty"`
	IsReady                  *bool  `json:"isReady,omitempty"`
	SupportsCustomObjects    *bool  `json:"supportsCustomObjects,omitempty"`
}

// CallingSettingsResponse is returned by calling settings operations.
type CallingSettingsResponse struct {
	Name                     string     `json:"name"`
	URL                      string     `json:"url"`
	Width                    int        `json:"width,omitempty"`
	Height                   int        `json:"height,omitempty"`
	IsReady                  bool       `json:"isReady"`
	SupportsCustomObjects    bool       `json:"supportsCustomObjects"`
	CreatedAt                time.Time  `json:"createdAt"`
	UpdatedAt                time.Time  `json:"updatedAt"`
}

// RecordingSettingsRequest is the input for registering recording URL format.
type RecordingSettingsRequest struct {
	URLToRetrieveAuthedRecording string `json:"urlToRetrieveAuthedRecording"`
}

// RecordingSettingsPatchRequest is the input for updating recording settings.
type RecordingSettingsPatchRequest struct {
	URLToRetrieveAuthedRecording string `json:"urlToRetrieveAuthedRecording,omitempty"`
}

// RecordingSettingsResponse is returned by recording settings operations.
type RecordingSettingsResponse struct {
	URLToRetrieveAuthedRecording string    `json:"urlToRetrieveAuthedRecording"`
	CreatedAt                    time.Time `json:"createdAt"`
	UpdatedAt                    time.Time `json:"updatedAt"`
}

// MarkRecordingAsReadyRequest is the input for marking a recording as ready.
type MarkRecordingAsReadyRequest struct {
	EngagementID int64 `json:"engagementId"`
}

// --- Cards extension models ---

// CardCreateRequest is the input for creating an extension card.
type CardCreateRequest struct {
	Title    string         `json:"title"`
	Fetch    CardFetchBody  `json:"fetch"`
	Display  CardDisplayBody `json:"display"`
	Actions  map[string]any `json:"actions,omitempty"`
}

// CardPatchRequest is the input for updating a card.
type CardPatchRequest struct {
	Title    string          `json:"title,omitempty"`
	Fetch    *CardFetchBody  `json:"fetch,omitempty"`
	Display  *CardDisplayBody `json:"display,omitempty"`
	Actions  map[string]any  `json:"actions,omitempty"`
}

// CardFetchBody describes how to fetch card data.
type CardFetchBody struct {
	TargetURL        string   `json:"targetUrl"`
	ObjectTypes      []CardObjectType `json:"objectTypes"`
}

// CardObjectType maps an object type to a property list for a card.
type CardObjectType struct {
	Name             string   `json:"name"`
	PropertiesToSend []string `json:"propertiesToSend"`
}

// CardDisplayBody describes how to display the card.
type CardDisplayBody struct {
	Properties []CardDisplayProperty `json:"properties"`
}

// CardDisplayProperty defines a single property in a card display.
type CardDisplayProperty struct {
	Name     string `json:"name"`
	Label    string `json:"label"`
	DataType string `json:"dataType"` // STRING, NUMBER, DATE, DATETIME, EMAIL, PHONE, CURRENCY, STATUS, LINK
}

// PublicCardResponse is returned by card create/update/getById.
type PublicCardResponse struct {
	ID       string         `json:"id"`
	Title    string         `json:"title"`
	Fetch    CardFetchBody  `json:"fetch"`
	Display  CardDisplayBody `json:"display"`
	Actions  map[string]any `json:"actions,omitempty"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

// PublicCardListResponse is returned by cards.GetAll.
type PublicCardListResponse struct {
	Results []*PublicCardResponse `json:"results"`
}

// --- Video conferencing extension models ---

// VideoConferencingExternalSettings is the settings for video conferencing.
type VideoConferencingExternalSettings struct {
	CreateMeetingURL string `json:"createMeetingUrl"`
	UpdateMeetingURL string `json:"updateMeetingUrl,omitempty"`
	DeleteMeetingURL string `json:"deleteMeetingUrl,omitempty"`
	UserVerifyURL    string `json:"userVerifyUrl,omitempty"`
	FetchAccountsURL string `json:"fetchAccountsUrl,omitempty"`
}

// =============================================================================
// Commerce models
// =============================================================================

// CommerceObjectService wraps an ObjectService for commerce objects
// (invoices, etc.) using the commerce-specific base path.
// Commerce objects use the same API shape as standard CRM objects
// but at /crm/v3/objects/invoices.
// We reuse the ObjectService since the API is identical.
