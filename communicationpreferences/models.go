package communicationpreferences

import "time"

// --- Subscription Definitions ---

// SubscriptionDefinition represents a communication subscription type.
type SubscriptionDefinition struct {
	IsInternal          bool      `json:"isInternal"`
	CreatedAt           time.Time `json:"createdAt"`
	IsDefault           bool      `json:"isDefault"`
	CommunicationMethod string    `json:"communicationMethod,omitempty"`
	Purpose             string    `json:"purpose,omitempty"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	ID                  string    `json:"id"`
	IsActive            bool      `json:"isActive"`
	BusinessUnitID      *int64    `json:"businessUnitId,omitempty"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

// SubscriptionDefinitionsResponse contains a list of subscription definitions.
type SubscriptionDefinitionsResponse struct {
	SubscriptionDefinitions []*SubscriptionDefinition `json:"subscriptionDefinitions"`
}

// --- Subscription Status ---

// PublicSubscriptionStatus represents the subscription status for a single subscription type.
type PublicSubscriptionStatus struct {
	BrandID               *int64 `json:"brandId,omitempty"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	LegalBasis            string `json:"legalBasis,omitempty"`
	PreferenceGroupName   string `json:"preferenceGroupName,omitempty"`
	ID                    string `json:"id"`
	LegalBasisExplanation string `json:"legalBasisExplanation,omitempty"`
	Status                string `json:"status"`
	SourceOfStatus        string `json:"sourceOfStatus"`
}

// PublicSubscriptionStatusesResponse contains subscription statuses for a recipient.
type PublicSubscriptionStatusesResponse struct {
	Recipient            string                      `json:"recipient"`
	SubscriptionStatuses []*PublicSubscriptionStatus `json:"subscriptionStatuses"`
}

// PublicUpdateSubscriptionStatusRequest is the body for subscribing or unsubscribing.
type PublicUpdateSubscriptionStatusRequest struct {
	EmailAddress          string `json:"emailAddress"`
	LegalBasis            string `json:"legalBasis,omitempty"`
	SubscriptionID        string `json:"subscriptionId"`
	LegalBasisExplanation string `json:"legalBasisExplanation,omitempty"`
}

// Well-known values for LegalBasis.
const (
	LegalBasisLegitimateInterestPQL    = "LEGITIMATE_INTEREST_PQL"
	LegalBasisLegitimateInterestClient = "LEGITIMATE_INTEREST_CLIENT"
	LegalBasisPerformanceOfContract    = "PERFORMANCE_OF_CONTRACT"
	LegalBasisConsentWithNotice        = "CONSENT_WITH_NOTICE"
	LegalBasisNonGDPR                  = "NON_GDPR"
	LegalBasisProcessAndStore          = "PROCESS_AND_STORE"
	LegalBasisLegitimateInterestOther  = "LEGITIMATE_INTEREST_OTHER"
)

// Well-known values for Status.
const (
	StatusSubscribed    = "SUBSCRIBED"
	StatusNotSubscribed = "NOT_SUBSCRIBED"
)

// Well-known values for SourceOfStatus.
const (
	SourcePortalWideStatus   = "PORTAL_WIDE_STATUS"
	SourceBrandWideStatus    = "BRAND_WIDE_STATUS"
	SourceSubscriptionStatus = "SUBSCRIPTION_STATUS"
)
