package crm

import (
	"encoding/json"
	"time"
)

// Properties is a map of HubSpot object property names to values.
// Handles JSON null values gracefully by omitting them from the map.
type Properties map[string]string

func (p *Properties) UnmarshalJSON(data []byte) error {
	var raw map[string]*string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*p = make(Properties, len(raw))
	for k, v := range raw {
		if v != nil {
			(*p)[k] = *v
		}
	}
	return nil
}

// --- Core object types ---

// SimplePublicObject is the core CRM object returned by HubSpot.
type SimplePublicObject struct {
	ID                    string                          `json:"id"`
	Properties            Properties                      `json:"properties"`
	PropertiesWithHistory map[string][]ValueWithTimestamp  `json:"propertiesWithHistory,omitempty"`
	CreatedAt             time.Time                       `json:"createdAt"`
	UpdatedAt             time.Time                       `json:"updatedAt"`
	Archived              bool                            `json:"archived"`
	ArchivedAt            *time.Time                      `json:"archivedAt,omitempty"`
	ObjectWriteTraceID    string                          `json:"objectWriteTraceId,omitempty"`
}

// SimplePublicObjectWithAssociations includes association data.
type SimplePublicObjectWithAssociations struct {
	ID                    string                                      `json:"id"`
	Properties            Properties                                  `json:"properties"`
	PropertiesWithHistory map[string][]ValueWithTimestamp              `json:"propertiesWithHistory,omitempty"`
	Associations          map[string]CollectionResponseAssociatedID    `json:"associations,omitempty"`
	CreatedAt             time.Time                                   `json:"createdAt"`
	UpdatedAt             time.Time                                   `json:"updatedAt"`
	Archived              bool                                        `json:"archived"`
	ArchivedAt            *time.Time                                  `json:"archivedAt,omitempty"`
}

// SimplePublicObjectInput is used for updating an existing object.
type SimplePublicObjectInput struct {
	Properties         Properties `json:"properties"`
	ObjectWriteTraceID string     `json:"objectWriteTraceId,omitempty"`
}

// SimplePublicObjectInputForCreate is used for creating a new object.
type SimplePublicObjectInputForCreate struct {
	Properties         Properties                   `json:"properties"`
	Associations       []PublicAssociationsForObject `json:"associations,omitempty"`
	ObjectWriteTraceID string                        `json:"objectWriteTraceId,omitempty"`
}

// SimplePublicUpsertObject is returned by batch upsert operations.
type SimplePublicUpsertObject struct {
	ID                    string                          `json:"id"`
	Properties            Properties                      `json:"properties"`
	PropertiesWithHistory map[string][]ValueWithTimestamp  `json:"propertiesWithHistory,omitempty"`
	CreatedAt             time.Time                       `json:"createdAt"`
	UpdatedAt             time.Time                       `json:"updatedAt"`
	Archived              bool                            `json:"archived"`
	ArchivedAt            *time.Time                      `json:"archivedAt,omitempty"`
	New                   bool                            `json:"new"`
}

// ValueWithTimestamp represents a property value at a point in time.
type ValueWithTimestamp struct {
	Value           string `json:"value"`
	Timestamp       string `json:"timestamp"`
	SourceType      string `json:"sourceType"`
	SourceID        string `json:"sourceId"`
	SourceLabel     string `json:"sourceLabel,omitempty"`
	UpdatedByUserID int64  `json:"updatedByUserId,omitempty"`
}

// --- Association types ---

// AssociatedID represents an associated object reference.
type AssociatedID struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// AssociationSpec defines the type of an association.
type AssociationSpec struct {
	AssociationCategory string `json:"associationCategory"` // HUBSPOT_DEFINED, USER_DEFINED, INTEGRATOR_DEFINED
	AssociationTypeID   int32  `json:"associationTypeId"`
}

// PublicAssociationsForObject is used when creating objects with associations.
type PublicAssociationsForObject struct {
	Types []AssociationSpec `json:"types"`
	To    ObjectID          `json:"to"`
}

// ObjectID is a simple wrapper for an object ID used in associations.
type ObjectID struct {
	ID string `json:"id"`
}

// CollectionResponseAssociatedID wraps a list of associated IDs.
type CollectionResponseAssociatedID struct {
	Results []AssociatedID `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// --- Pagination types ---

// ForwardPaging contains the cursor for the next page of results.
type ForwardPaging struct {
	Next *NextPage `json:"next,omitempty"`
}

// NextPage holds the cursor and optional link for pagination.
type NextPage struct {
	After string `json:"after"`
	Link  string `json:"link,omitempty"`
}

// --- Collection response types ---

// ListResult is returned by List (GetPage) operations.
type ListResult struct {
	Results []*SimplePublicObjectWithAssociations `json:"results"`
	Paging  *ForwardPaging                        `json:"paging,omitempty"`
}

// SearchResult is returned by Search operations.
type SearchResult struct {
	Total   int                    `json:"total"`
	Results []*SimplePublicObject  `json:"results"`
	Paging  *ForwardPaging         `json:"paging,omitempty"`
}

// --- Batch input types ---

// BatchCreateInput is the input for batch create operations.
type BatchCreateInput struct {
	Inputs []SimplePublicObjectInputForCreate `json:"inputs"`
}

// BatchUpdateInput is the input for batch update operations.
type BatchUpdateInput struct {
	Inputs []BatchObjectInput `json:"inputs"`
}

// BatchObjectInput is a single item in a batch update.
type BatchObjectInput struct {
	ID                 string     `json:"id"`
	Properties         Properties `json:"properties"`
	IDProperty         string     `json:"idProperty,omitempty"`
	ObjectWriteTraceID string     `json:"objectWriteTraceId,omitempty"`
}

// BatchUpsertInput is the input for batch upsert operations.
type BatchUpsertInput struct {
	Inputs []BatchObjectUpsertInput `json:"inputs"`
}

// BatchObjectUpsertInput is a single item in a batch upsert.
type BatchObjectUpsertInput struct {
	ID                 string     `json:"id"`
	Properties         Properties `json:"properties"`
	IDProperty         string     `json:"idProperty,omitempty"`
	ObjectWriteTraceID string     `json:"objectWriteTraceId,omitempty"`
}

// BatchArchiveInput is the input for batch archive operations.
type BatchArchiveInput struct {
	Inputs []ObjectID `json:"inputs"`
}

// BatchReadInput is the input for batch read operations.
type BatchReadInput struct {
	Properties            []string   `json:"properties"`
	PropertiesWithHistory []string   `json:"propertiesWithHistory,omitempty"`
	IDProperty            string     `json:"idProperty,omitempty"`
	Inputs                []ObjectID `json:"inputs"`
}

// --- Batch response types ---

// BatchResult is returned by batch create/read/update operations.
type BatchResult struct {
	Status      string                 `json:"status"` // PENDING, PROCESSING, CANCELED, COMPLETE
	Results     []*SimplePublicObject  `json:"results"`
	RequestedAt *time.Time             `json:"requestedAt,omitempty"`
	StartedAt   time.Time              `json:"startedAt"`
	CompletedAt time.Time              `json:"completedAt"`
	Links       map[string]string      `json:"links,omitempty"`
	NumErrors   int                    `json:"numErrors,omitempty"`
	Errors      []StandardError        `json:"errors,omitempty"`
}

// BatchUpsertResult is returned by batch upsert operations.
type BatchUpsertResult struct {
	Status      string                      `json:"status"`
	Results     []*SimplePublicUpsertObject `json:"results"`
	RequestedAt *time.Time                  `json:"requestedAt,omitempty"`
	StartedAt   time.Time                   `json:"startedAt"`
	CompletedAt time.Time                   `json:"completedAt"`
	Links       map[string]string           `json:"links,omitempty"`
	NumErrors   int                         `json:"numErrors,omitempty"`
	Errors      []StandardError             `json:"errors,omitempty"`
}

// StandardError represents an error within a batch response.
type StandardError struct {
	Status      string              `json:"status"`
	ID          string              `json:"id,omitempty"`
	Category    string              `json:"category"`
	SubCategory string              `json:"subCategory,omitempty"`
	Message     string              `json:"message"`
	Errors      []ErrorDetail       `json:"errors,omitempty"`
	Context     map[string][]string `json:"context,omitempty"`
	Links       map[string]string   `json:"links,omitempty"`
}

// ErrorDetail represents a single validation or processing error within a batch.
type ErrorDetail struct {
	Message     string              `json:"message"`
	In          string              `json:"in"`
	Code        string              `json:"code"`
	SubCategory string              `json:"subCategory"`
	Context     map[string][]string `json:"context,omitempty"`
}

// --- Merge and GDPR types ---

// PublicMergeInput is used to merge two objects.
type PublicMergeInput struct {
	PrimaryObjectID string `json:"primaryObjectId"`
	ObjectIDToMerge string `json:"objectIdToMerge"`
}

// PublicGdprDeleteInput is used for GDPR-compliant deletion.
type PublicGdprDeleteInput struct {
	ObjectID   string `json:"objectId"`
	IDProperty string `json:"idProperty,omitempty"`
}

// --- Search types ---

// PublicObjectSearchRequest defines a search query.
type PublicObjectSearchRequest struct {
	FilterGroups []FilterGroup `json:"filterGroups,omitempty"`
	Sorts        []string      `json:"sorts,omitempty"`
	Query        string        `json:"query,omitempty"`
	Properties   []string      `json:"properties,omitempty"`
	Limit        int           `json:"limit,omitempty"`
	After        int           `json:"after,omitempty"`
}

// FilterGroup is a group of filters combined with AND.
type FilterGroup struct {
	Filters []Filter `json:"filters"`
}

// Filter is a single search filter.
type Filter struct {
	PropertyName string   `json:"propertyName"`
	Operator     string   `json:"operator"` // Use FilterOperator* constants
	Value        string   `json:"value,omitempty"`
	HighValue    string   `json:"highValue,omitempty"`
	Values       []string `json:"values,omitempty"`
}

// Filter operator constants.
const (
	FilterOperatorEQ               = "EQ"
	FilterOperatorNEQ              = "NEQ"
	FilterOperatorLT               = "LT"
	FilterOperatorLTE              = "LTE"
	FilterOperatorGT               = "GT"
	FilterOperatorGTE              = "GTE"
	FilterOperatorBetween          = "BETWEEN"
	FilterOperatorIn               = "IN"
	FilterOperatorNotIn            = "NOT_IN"
	FilterOperatorHasProperty      = "HAS_PROPERTY"
	FilterOperatorNotHasProperty   = "NOT_HAS_PROPERTY"
	FilterOperatorContainsToken    = "CONTAINS_TOKEN"
	FilterOperatorNotContainsToken = "NOT_CONTAINS_TOKEN"
)

// Association category constants.
const (
	AssociationCategoryHubSpotDefined    = "HUBSPOT_DEFINED"
	AssociationCategoryUserDefined       = "USER_DEFINED"
	AssociationCategoryIntegratorDefined = "INTEGRATOR_DEFINED"
)

// --- Option types ---

// GetByIDOptions configures a GetByID request.
type GetByIDOptions struct {
	Properties            []string
	PropertiesWithHistory []string
	Associations          []string
	Archived              bool
	IDProperty            string
}

// UpdateOptions configures an Update request.
type UpdateOptions struct {
	IDProperty string
}

// ListOptions configures a List (GetPage) request.
type ListOptions struct {
	Limit                 int
	After                 string
	Properties            []string
	PropertiesWithHistory []string
	Associations          []string
	Archived              bool
}

// GetAllOptions configures a GetAll request.
type GetAllOptions struct {
	Limit                 int // per-page limit (default 100)
	Properties            []string
	PropertiesWithHistory []string
	Associations          []string
	Archived              bool
}
