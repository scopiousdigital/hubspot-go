package webhooks

import "time"

// --- Settings models ---

// ThrottlingSettings controls concurrency for webhook delivery.
type ThrottlingSettings struct {
	MaxConcurrentRequests int `json:"maxConcurrentRequests"`
}

// SettingsChangeRequest is the input for configuring webhook settings.
type SettingsChangeRequest struct {
	Throttling ThrottlingSettings `json:"throttling"`
	TargetURL  string             `json:"targetUrl"`
}

// SettingsResponse is the response from the webhook settings API.
type SettingsResponse struct {
	CreatedAt  time.Time          `json:"createdAt"`
	Throttling ThrottlingSettings `json:"throttling"`
	TargetURL  string             `json:"targetUrl"`
	UpdatedAt  *time.Time         `json:"updatedAt,omitempty"`
}

// --- Subscription models ---

// SubscriptionCreateRequest is the input for creating a new webhook subscription.
type SubscriptionCreateRequest struct {
	ObjectTypeID string `json:"objectTypeId,omitempty"`
	PropertyName string `json:"propertyName,omitempty"`
	Active       *bool  `json:"active,omitempty"`
	EventType    string `json:"eventType"`
}

// SubscriptionPatchRequest is the input for updating a webhook subscription.
type SubscriptionPatchRequest struct {
	Active *bool `json:"active,omitempty"`
}

// SubscriptionResponse is a single webhook subscription returned by the API.
type SubscriptionResponse struct {
	CreatedAt    time.Time  `json:"createdAt"`
	ObjectTypeID string     `json:"objectTypeId,omitempty"`
	PropertyName string     `json:"propertyName,omitempty"`
	Active       bool       `json:"active"`
	EventType    string     `json:"eventType"`
	ID           string     `json:"id"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
}

// SubscriptionListResponse contains a list of webhook subscriptions.
type SubscriptionListResponse struct {
	Results []SubscriptionResponse `json:"results"`
}

// --- Batch models ---

// SubscriptionBatchUpdateRequest is a single item in a batch update.
type SubscriptionBatchUpdateRequest struct {
	Active bool `json:"active"`
	ID     int  `json:"id"`
}

// BatchInputSubscriptionBatchUpdateRequest is the input for a batch subscription update.
type BatchInputSubscriptionBatchUpdateRequest struct {
	Inputs []SubscriptionBatchUpdateRequest `json:"inputs"`
}

// BatchResponseSubscriptionResponse is the response from a batch subscription update.
type BatchResponseSubscriptionResponse struct {
	CompletedAt time.Time              `json:"completedAt"`
	RequestedAt *time.Time             `json:"requestedAt,omitempty"`
	StartedAt   time.Time              `json:"startedAt"`
	Links       map[string]string      `json:"links,omitempty"`
	Results     []SubscriptionResponse `json:"results"`
	Status      string                 `json:"status"`
}

// --- Event type constants ---
// These match the HubSpot webhook event types.

const (
	EventTypeContactPropertyChange     = "contact.propertyChange"
	EventTypeCompanyPropertyChange     = "company.propertyChange"
	EventTypeDealPropertyChange        = "deal.propertyChange"
	EventTypeTicketPropertyChange      = "ticket.propertyChange"
	EventTypeProductPropertyChange     = "product.propertyChange"
	EventTypeLineItemPropertyChange    = "line_item.propertyChange"
	EventTypeContactCreation           = "contact.creation"
	EventTypeContactDeletion           = "contact.deletion"
	EventTypeContactPrivacyDeletion    = "contact.privacyDeletion"
	EventTypeCompanyCreation           = "company.creation"
	EventTypeCompanyDeletion           = "company.deletion"
	EventTypeDealCreation              = "deal.creation"
	EventTypeDealDeletion              = "deal.deletion"
	EventTypeTicketCreation            = "ticket.creation"
	EventTypeTicketDeletion            = "ticket.deletion"
	EventTypeProductCreation           = "product.creation"
	EventTypeProductDeletion           = "product.deletion"
	EventTypeLineItemCreation          = "line_item.creation"
	EventTypeLineItemDeletion          = "line_item.deletion"
	EventTypeConversationCreation      = "conversation.creation"
	EventTypeConversationDeletion      = "conversation.deletion"
	EventTypeConversationNewMessage    = "conversation.newMessage"
	EventTypeConversationPrivacyDel    = "conversation.privacyDeletion"
	EventTypeConversationPropertyChange = "conversation.propertyChange"
	EventTypeContactMerge              = "contact.merge"
	EventTypeCompanyMerge              = "company.merge"
	EventTypeDealMerge                 = "deal.merge"
	EventTypeTicketMerge               = "ticket.merge"
	EventTypeProductMerge              = "product.merge"
	EventTypeLineItemMerge             = "line_item.merge"
	EventTypeContactRestore            = "contact.restore"
	EventTypeCompanyRestore            = "company.restore"
	EventTypeDealRestore               = "deal.restore"
	EventTypeTicketRestore             = "ticket.restore"
	EventTypeProductRestore            = "product.restore"
	EventTypeLineItemRestore           = "line_item.restore"
	EventTypeContactAssociationChange  = "contact.associationChange"
	EventTypeCompanyAssociationChange  = "company.associationChange"
	EventTypeDealAssociationChange     = "deal.associationChange"
	EventTypeTicketAssociationChange   = "ticket.associationChange"
	EventTypeLineItemAssociationChange = "line_item.associationChange"
	EventTypeObjectPropertyChange      = "object.propertyChange"
	EventTypeObjectCreation            = "object.creation"
	EventTypeObjectDeletion            = "object.deletion"
	EventTypeObjectMerge               = "object.merge"
	EventTypeObjectRestore             = "object.restore"
	EventTypeObjectAssociationChange   = "object.associationChange"
)
