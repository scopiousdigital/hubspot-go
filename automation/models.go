package automation

import "time"

// --- Callbacks models ---

// CallbackCompletionRequest is the body for completing a single callback.
type CallbackCompletionRequest struct {
	OutputFields map[string]string `json:"outputFields"`
}

// CallbackCompletionBatchRequest is a single callback completion entry within a batch.
type CallbackCompletionBatchRequest struct {
	OutputFields map[string]string `json:"outputFields"`
	CallbackID   string            `json:"callbackId"`
}

// BatchInputCallbackCompletionBatchRequest wraps multiple callback completions.
type BatchInputCallbackCompletionBatchRequest struct {
	Inputs []CallbackCompletionBatchRequest `json:"inputs"`
}

// --- Definitions models ---

// PublicActionDefinition represents a custom workflow action definition.
type PublicActionDefinition struct {
	Functions              []PublicActionFunctionIdentifier `json:"functions"`
	ActionURL              string                           `json:"actionUrl"`
	Published              bool                             `json:"published"`
	Labels                 map[string]PublicActionLabels    `json:"labels"`
	InputFields            []InputFieldDefinition           `json:"inputFields"`
	OutputFields           []OutputFieldDefinition          `json:"outputFields,omitempty"`
	RevisionID             string                           `json:"revisionId"`
	ArchivedAt             *int64                           `json:"archivedAt,omitempty"`
	InputFieldDependencies []map[string]any                 `json:"inputFieldDependencies,omitempty"`
	ExecutionRules         []PublicExecutionTranslationRule `json:"executionRules,omitempty"`
	ID                     string                           `json:"id"`
	ObjectTypes            []string                         `json:"objectTypes"`
	ObjectRequestOptions   *PublicObjectRequestOptions      `json:"objectRequestOptions,omitempty"`
}

// PublicActionDefinitionEgg is the input for creating a new action definition.
type PublicActionDefinitionEgg struct {
	InputFields            []InputFieldDefinition           `json:"inputFields"`
	OutputFields           []OutputFieldDefinition          `json:"outputFields,omitempty"`
	ArchivedAt             *int64                           `json:"archivedAt,omitempty"`
	Functions              []PublicActionFunction           `json:"functions"`
	ActionURL              string                           `json:"actionUrl"`
	InputFieldDependencies []map[string]any                 `json:"inputFieldDependencies,omitempty"`
	Published              bool                             `json:"published"`
	ExecutionRules         []PublicExecutionTranslationRule `json:"executionRules,omitempty"`
	ObjectTypes            []string                         `json:"objectTypes"`
	ObjectRequestOptions   *PublicObjectRequestOptions      `json:"objectRequestOptions,omitempty"`
	Labels                 map[string]PublicActionLabels    `json:"labels"`
}

// PublicActionDefinitionPatch is the input for updating an action definition.
type PublicActionDefinitionPatch struct {
	InputFields            []InputFieldDefinition           `json:"inputFields,omitempty"`
	OutputFields           []OutputFieldDefinition          `json:"outputFields,omitempty"`
	ActionURL              string                           `json:"actionUrl,omitempty"`
	InputFieldDependencies []map[string]any                 `json:"inputFieldDependencies,omitempty"`
	Published              *bool                            `json:"published,omitempty"`
	ExecutionRules         []PublicExecutionTranslationRule `json:"executionRules,omitempty"`
	ObjectTypes            []string                         `json:"objectTypes,omitempty"`
	ObjectRequestOptions   *PublicObjectRequestOptions      `json:"objectRequestOptions,omitempty"`
	Labels                 map[string]PublicActionLabels    `json:"labels,omitempty"`
}

// PublicActionLabels contains label information for an action.
type PublicActionLabels struct {
	InputFieldDescriptions map[string]string            `json:"inputFieldDescriptions,omitempty"`
	AppDisplayName         string                       `json:"appDisplayName,omitempty"`
	OutputFieldLabels      map[string]string            `json:"outputFieldLabels,omitempty"`
	InputFieldOptionLabels map[string]map[string]string `json:"inputFieldOptionLabels,omitempty"`
	ActionDescription      string                       `json:"actionDescription,omitempty"`
	ExecutionRules         map[string]string            `json:"executionRules,omitempty"`
	InputFieldLabels       map[string]string            `json:"inputFieldLabels,omitempty"`
	ActionName             string                       `json:"actionName"`
	ActionCardContent      string                       `json:"actionCardContent,omitempty"`
}

// InputFieldDefinition describes an input field for an action.
type InputFieldDefinition struct {
	IsRequired          bool                `json:"isRequired"`
	AutomationFieldType string              `json:"automationFieldType,omitempty"`
	TypeDefinition      FieldTypeDefinition `json:"typeDefinition"`
	SupportedValueTypes []string            `json:"supportedValueTypes,omitempty"`
}

// OutputFieldDefinition describes an output field for an action.
type OutputFieldDefinition struct {
	TypeDefinition FieldTypeDefinition `json:"typeDefinition"`
}

// FieldTypeDefinition describes the type of a field.
type FieldTypeDefinition struct {
	HelpText                     string   `json:"helpText,omitempty"`
	ReferencedObjectType         string   `json:"referencedObjectType,omitempty"`
	Name                         string   `json:"name"`
	Options                      []Option `json:"options"`
	Description                  string   `json:"description,omitempty"`
	ExternalOptionsReferenceType string   `json:"externalOptionsReferenceType,omitempty"`
	Label                        string   `json:"label,omitempty"`
	Type                         string   `json:"type"`
	FieldType                    string   `json:"fieldType,omitempty"`
	OptionsURL                   string   `json:"optionsUrl,omitempty"`
	ExternalOptions              bool     `json:"externalOptions"`
}

// Option represents a selectable option for a field.
type Option struct {
	Hidden       bool    `json:"hidden"`
	DisplayOrder int     `json:"displayOrder"`
	DoubleData   float64 `json:"doubleData"`
	Description  string  `json:"description"`
	ReadOnly     bool    `json:"readOnly"`
	Label        string  `json:"label"`
	Value        string  `json:"value"`
}

// PublicExecutionTranslationRule defines an execution translation rule.
type PublicExecutionTranslationRule struct {
	LabelName  string         `json:"labelName"`
	Conditions map[string]any `json:"conditions"`
}

// PublicObjectRequestOptions describes which object properties to fetch.
type PublicObjectRequestOptions struct {
	Properties []string `json:"properties"`
}

// --- Functions models ---

// PublicActionFunction represents a serverless function attached to an action.
type PublicActionFunction struct {
	FunctionSource string `json:"functionSource"`
	FunctionType   string `json:"functionType"`
	ID             string `json:"id,omitempty"`
}

// PublicActionFunctionIdentifier identifies a function attached to an action.
type PublicActionFunctionIdentifier struct {
	FunctionType string `json:"functionType"`
	ID           string `json:"id,omitempty"`
}

// FunctionsResult is the response for listing functions (no paging).
type FunctionsResult struct {
	Results []*PublicActionFunctionIdentifier `json:"results"`
}

// --- Revisions models ---

// PublicActionRevision represents a revision of an action definition.
type PublicActionRevision struct {
	RevisionID string                 `json:"revisionId"`
	CreatedAt  time.Time              `json:"createdAt"`
	Definition PublicActionDefinition `json:"definition"`
	ID         string                 `json:"id"`
}

// --- List result types ---

// DefinitionsListResult is the paginated response for listing definitions.
type DefinitionsListResult struct {
	Results []*PublicActionDefinition `json:"results"`
	Paging  *ForwardPaging            `json:"paging,omitempty"`
}

// DefinitionsListOptions contains query parameters for listing definitions.
type DefinitionsListOptions struct {
	Limit    int    `json:"-"`
	After    string `json:"-"`
	Archived bool   `json:"-"`
}

// RevisionsListResult is the paginated response for listing revisions.
type RevisionsListResult struct {
	Results []*PublicActionRevision `json:"results"`
	Paging  *ForwardPaging          `json:"paging,omitempty"`
}

// RevisionsListOptions contains query parameters for listing revisions.
type RevisionsListOptions struct {
	Limit int    `json:"-"`
	After string `json:"-"`
}

// --- Shared paging ---

// ForwardPaging contains the next-page cursor.
type ForwardPaging struct {
	Next *NextPage `json:"next,omitempty"`
}

// NextPage contains the cursor for the next page of results.
type NextPage struct {
	Link  string `json:"link,omitempty"`
	After string `json:"after"`
}
