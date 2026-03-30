package marketing

import "time"

// --- Common pagination ---

// ForwardPaging holds a cursor to the next page of results.
type ForwardPaging struct {
	Next *NextPage `json:"next,omitempty"`
}

// NextPage contains the cursor value for the next page.
type NextPage struct {
	After string `json:"after"`
	Link  string `json:"link,omitempty"`
}

// --- Forms models ---

// FormDefinition represents a HubSpot form definition as returned by the API.
type FormDefinition struct {
	ID                  string              `json:"id"`
	FormType            string              `json:"formType"`
	Name                string              `json:"name"`
	CreatedAt           time.Time           `json:"createdAt"`
	UpdatedAt           time.Time           `json:"updatedAt"`
	Archived            bool                `json:"archived"`
	ArchivedAt          *time.Time          `json:"archivedAt,omitempty"`
	FieldGroups         []FieldGroup        `json:"fieldGroups"`
	Configuration       map[string]any      `json:"configuration,omitempty"`
	DisplayOptions      map[string]any      `json:"displayOptions,omitempty"`
	LegalConsentOptions map[string]any      `json:"legalConsentOptions,omitempty"`
}

// FieldGroup represents a group of fields in a form.
type FieldGroup struct {
	GroupType    string           `json:"groupType"`
	RichTextType string          `json:"richTextType"`
	RichText     string          `json:"richText,omitempty"`
	Fields       []FormField     `json:"fields"`
}

// FormField represents a single field in a form field group.
type FormField struct {
	Name             string           `json:"name"`
	Label            string           `json:"label"`
	FieldType        string           `json:"fieldType,omitempty"`
	ObjectTypeID     string           `json:"objectTypeId,omitempty"`
	Required         bool             `json:"required,omitempty"`
	Hidden           bool             `json:"hidden,omitempty"`
	Description      string           `json:"description,omitempty"`
	DefaultValue     string           `json:"defaultValue,omitempty"`
	Placeholder      string           `json:"placeholder,omitempty"`
	Options          []FieldOption    `json:"options,omitempty"`
	Validation       map[string]any   `json:"validation,omitempty"`
	DependentFields  []map[string]any `json:"dependentFields,omitempty"`
}

// FieldOption represents an option for enumerated form fields.
type FieldOption struct {
	Label       string `json:"label"`
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
	DisplayOrder int   `json:"displayOrder,omitempty"`
}

// FormCreateRequest is the input for creating a new form.
type FormCreateRequest struct {
	FormType            string         `json:"formType"`
	Name                string         `json:"name"`
	FieldGroups         []FieldGroup   `json:"fieldGroups,omitempty"`
	Configuration       map[string]any `json:"configuration,omitempty"`
	DisplayOptions      map[string]any `json:"displayOptions,omitempty"`
	LegalConsentOptions map[string]any `json:"legalConsentOptions,omitempty"`
}

// FormUpdateRequest is the input for updating (patching) an existing form.
type FormUpdateRequest struct {
	Name                string         `json:"name,omitempty"`
	FieldGroups         []FieldGroup   `json:"fieldGroups,omitempty"`
	Archived            *bool          `json:"archived,omitempty"`
	Configuration       map[string]any `json:"configuration,omitempty"`
	DisplayOptions      map[string]any `json:"displayOptions,omitempty"`
	LegalConsentOptions map[string]any `json:"legalConsentOptions,omitempty"`
}

// FormReplaceRequest is the input for fully replacing a form definition.
type FormReplaceRequest struct {
	FormType            string         `json:"formType"`
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	FieldGroups         []FieldGroup   `json:"fieldGroups"`
	Configuration       map[string]any `json:"configuration,omitempty"`
	DisplayOptions      map[string]any `json:"displayOptions,omitempty"`
	LegalConsentOptions map[string]any `json:"legalConsentOptions,omitempty"`
}

// FormListResult is a paginated list of form definitions.
type FormListResult struct {
	Results []FormDefinition `json:"results"`
	Paging  *ForwardPaging   `json:"paging,omitempty"`
}

// FormListOptions configures a form list request.
type FormListOptions struct {
	After     string   `json:"after,omitempty"`
	Limit     int      `json:"limit,omitempty"`
	Archived  bool     `json:"archived,omitempty"`
	FormTypes []string `json:"formTypes,omitempty"`
}

// --- Marketing Emails models ---

// PublicEmail represents a marketing email as returned by the API.
type PublicEmail struct {
	ID                  string         `json:"id"`
	Name                string         `json:"name"`
	Subject             string         `json:"subject"`
	Type                string         `json:"type,omitempty"`
	State               string         `json:"state"`
	Subcategory         string         `json:"subcategory"`
	Language            string         `json:"language,omitempty"`
	SendOnPublish       bool           `json:"sendOnPublish"`
	IsPublished         *bool          `json:"isPublished,omitempty"`
	IsTransactional     *bool          `json:"isTransactional,omitempty"`
	Archived            *bool          `json:"archived,omitempty"`
	ActiveDomain        string         `json:"activeDomain,omitempty"`
	Campaign            string         `json:"campaign,omitempty"`
	CampaignName        string         `json:"campaignName,omitempty"`
	FeedbackSurveyID    string         `json:"feedbackSurveyId,omitempty"`
	BusinessUnitID      string         `json:"businessUnitId,omitempty"`
	ClonedFrom          string         `json:"clonedFrom,omitempty"`
	FolderID            *int           `json:"folderId,omitempty"`
	JitterSendTime      *bool          `json:"jitterSendTime,omitempty"`
	Content             map[string]any `json:"content,omitempty"`
	From                map[string]any `json:"from,omitempty"`
	To                  map[string]any `json:"to,omitempty"`
	Webversion          map[string]any `json:"webversion,omitempty"`
	RssData             map[string]any `json:"rssData,omitempty"`
	Testing             map[string]any `json:"testing,omitempty"`
	SubscriptionDetails map[string]any `json:"subscriptionDetails,omitempty"`
	Stats               map[string]any `json:"stats,omitempty"`
	WorkflowNames       []string       `json:"workflowNames,omitempty"`
	CreatedAt           *time.Time     `json:"createdAt,omitempty"`
	UpdatedAt           *time.Time     `json:"updatedAt,omitempty"`
	PublishDate         *time.Time     `json:"publishDate,omitempty"`
	PublishedAt         *time.Time     `json:"publishedAt,omitempty"`
	DeletedAt           *time.Time     `json:"deletedAt,omitempty"`
	CreatedByID         string         `json:"createdById,omitempty"`
	UpdatedByID         string         `json:"updatedById,omitempty"`
	PublishedByID       string         `json:"publishedById,omitempty"`
}

// EmailCreateRequest is the input for creating a marketing email.
type EmailCreateRequest struct {
	Name                string         `json:"name"`
	Subject             string         `json:"subject,omitempty"`
	Subcategory         string         `json:"subcategory,omitempty"`
	State               string         `json:"state,omitempty"`
	Language            string         `json:"language,omitempty"`
	SendOnPublish       *bool          `json:"sendOnPublish,omitempty"`
	Archived            *bool          `json:"archived,omitempty"`
	ActiveDomain        string         `json:"activeDomain,omitempty"`
	Campaign            string         `json:"campaign,omitempty"`
	FeedbackSurveyID    string         `json:"feedbackSurveyId,omitempty"`
	BusinessUnitID      *int           `json:"businessUnitId,omitempty"`
	JitterSendTime      *bool          `json:"jitterSendTime,omitempty"`
	Content             map[string]any `json:"content,omitempty"`
	From                map[string]any `json:"from,omitempty"`
	To                  map[string]any `json:"to,omitempty"`
	Webversion          map[string]any `json:"webversion,omitempty"`
	RssData             map[string]any `json:"rssData,omitempty"`
	Testing             map[string]any `json:"testing,omitempty"`
	SubscriptionDetails map[string]any `json:"subscriptionDetails,omitempty"`
	PublishDate         *time.Time     `json:"publishDate,omitempty"`
}

// EmailUpdateRequest is the input for updating a marketing email.
type EmailUpdateRequest struct {
	Name                string         `json:"name,omitempty"`
	Subject             string         `json:"subject,omitempty"`
	Subcategory         string         `json:"subcategory,omitempty"`
	State               string         `json:"state,omitempty"`
	Language            string         `json:"language,omitempty"`
	SendOnPublish       *bool          `json:"sendOnPublish,omitempty"`
	Archived            *bool          `json:"archived,omitempty"`
	ActiveDomain        string         `json:"activeDomain,omitempty"`
	Campaign            string         `json:"campaign,omitempty"`
	BusinessUnitID      *int           `json:"businessUnitId,omitempty"`
	JitterSendTime      *bool          `json:"jitterSendTime,omitempty"`
	Content             map[string]any `json:"content,omitempty"`
	From                map[string]any `json:"from,omitempty"`
	To                  map[string]any `json:"to,omitempty"`
	Webversion          map[string]any `json:"webversion,omitempty"`
	RssData             map[string]any `json:"rssData,omitempty"`
	Testing             map[string]any `json:"testing,omitempty"`
	SubscriptionDetails map[string]any `json:"subscriptionDetails,omitempty"`
	PublishDate         *time.Time     `json:"publishDate,omitempty"`
}

// ContentCloneRequest is used to clone a marketing email.
type ContentCloneRequest struct {
	ID        string `json:"id"`
	CloneName string `json:"cloneName,omitempty"`
}

// AbTestCreateRequest is used to create an A/B test variation.
type AbTestCreateRequest struct {
	ContentID     string `json:"contentId"`
	VariationName string `json:"variationName"`
}

// EmailListResult is a paginated list of marketing emails.
type EmailListResult struct {
	Total   int            `json:"total"`
	Results []PublicEmail   `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// EmailListOptions configures a marketing email list request.
type EmailListOptions struct {
	After                  string   `json:"after,omitempty"`
	Limit                  int      `json:"limit,omitempty"`
	Sort                   []string `json:"sort,omitempty"`
	Type                   string   `json:"type,omitempty"`
	IsPublished            *bool    `json:"isPublished,omitempty"`
	Archived               *bool    `json:"archived,omitempty"`
	Campaign               string   `json:"campaign,omitempty"`
	IncludeStats           *bool    `json:"includeStats,omitempty"`
	MarketingCampaignNames *bool    `json:"marketingCampaignNames,omitempty"`
	WorkflowNames          *bool    `json:"workflowNames,omitempty"`
	IncludedProperties     []string `json:"includedProperties,omitempty"`
	CreatedAt              string   `json:"createdAt,omitempty"`
	CreatedAfter           string   `json:"createdAfter,omitempty"`
	CreatedBefore          string   `json:"createdBefore,omitempty"`
	UpdatedAt              string   `json:"updatedAt,omitempty"`
	UpdatedAfter           string   `json:"updatedAfter,omitempty"`
	UpdatedBefore          string   `json:"updatedBefore,omitempty"`
}

// EmailGetByIDOptions configures how a single email is fetched.
type EmailGetByIDOptions struct {
	IncludeStats           *bool    `json:"includeStats,omitempty"`
	MarketingCampaignNames *bool    `json:"marketingCampaignNames,omitempty"`
	WorkflowNames          *bool    `json:"workflowNames,omitempty"`
	IncludedProperties     []string `json:"includedProperties,omitempty"`
	Archived               *bool    `json:"archived,omitempty"`
}

// VersionPublicEmail represents a versioned email revision.
type VersionPublicEmail struct {
	ID        string      `json:"id"`
	User      VersionUser `json:"user"`
	Object    PublicEmail `json:"object"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

// VersionUser identifies the user who created a revision.
type VersionUser struct {
	ID    string `json:"id"`
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

// VersionListResult is a paginated list of email revisions.
type VersionListResult struct {
	Total   int                  `json:"total"`
	Results []VersionPublicEmail `json:"results"`
}

// EmailRevisionsOptions configures a revisions list request.
type EmailRevisionsOptions struct {
	After  string `json:"after,omitempty"`
	Before string `json:"before,omitempty"`
	Limit  int    `json:"limit,omitempty"`
}

// --- Email statistics models ---

// EmailStatisticsData contains counters, ratios, etc. for email performance.
type EmailStatisticsData struct {
	Counters        map[string]int                `json:"counters,omitempty"`
	Ratios          map[string]float64            `json:"ratios,omitempty"`
	DeviceBreakdown map[string]map[string]int     `json:"deviceBreakdown,omitempty"`
	QualifierStats  map[string]map[string]int     `json:"qualifierStats,omitempty"`
}

// AggregateEmailStatistics contains aggregate statistics for one or more emails.
type AggregateEmailStatistics struct {
	Emails               []int                          `json:"emails,omitempty"`
	CampaignAggregations map[string]EmailStatisticsData `json:"campaignAggregations,omitempty"`
	Aggregate            *EmailStatisticsData           `json:"aggregate,omitempty"`
}

// EmailStatisticInterval represents a statistics interval.
type EmailStatisticInterval struct {
	Interval     map[string]any       `json:"interval,omitempty"`
	Aggregations *EmailStatisticsData `json:"aggregations,omitempty"`
}

// EmailStatisticsHistogramResult is a list of statistics intervals.
type EmailStatisticsHistogramResult struct {
	Total   int                      `json:"total"`
	Results []EmailStatisticInterval `json:"results"`
}

// EmailStatisticsListOptions configures a statistics list request.
type EmailStatisticsListOptions struct {
	StartTimestamp string `json:"startTimestamp,omitempty"`
	EndTimestamp   string `json:"endTimestamp,omitempty"`
	EmailIDs       []int  `json:"emailIds,omitempty"`
	Property       string `json:"property,omitempty"`
}

// EmailStatisticsHistogramOptions configures a statistics histogram request.
type EmailStatisticsHistogramOptions struct {
	Interval       string `json:"interval,omitempty"` // YEAR, QUARTER, MONTH, WEEK, DAY, HOUR, etc.
	StartTimestamp string `json:"startTimestamp,omitempty"`
	EndTimestamp   string `json:"endTimestamp,omitempty"`
	EmailIDs       []int  `json:"emailIds,omitempty"`
}

// --- Marketing Events models ---

// PropertyValue represents a custom property with a name, value, and optional metadata.
type PropertyValue struct {
	Name             string `json:"name"`
	Value            string `json:"value"`
	Timestamp        string `json:"timestamp,omitempty"`
	SourceID         string `json:"sourceId,omitempty"`
	SourceLabel      string `json:"sourceLabel,omitempty"`
	Source           string `json:"source,omitempty"`
	SelectedByUser   bool   `json:"selectedByUser,omitempty"`
	SelectedByUserTS int64  `json:"selectedByUserTimestamp,omitempty"`
}

// MarketingEventCreateRequest is the input for creating a marketing event.
type MarketingEventCreateRequest struct {
	ExternalEventID   string          `json:"externalEventId"`
	ExternalAccountID string          `json:"externalAccountId"`
	EventName         string          `json:"eventName"`
	EventOrganizer    string          `json:"eventOrganizer"`
	EventDescription  string          `json:"eventDescription,omitempty"`
	EventURL          string          `json:"eventUrl,omitempty"`
	EventType         string          `json:"eventType,omitempty"`
	EventCancelled    *bool           `json:"eventCancelled,omitempty"`
	EventCompleted    *bool           `json:"eventCompleted,omitempty"`
	StartDateTime     *time.Time      `json:"startDateTime,omitempty"`
	EndDateTime       *time.Time      `json:"endDateTime,omitempty"`
	CustomProperties  []PropertyValue `json:"customProperties,omitempty"`
}

// MarketingEventUpdateRequest is the input for updating a marketing event.
type MarketingEventUpdateRequest struct {
	EventName        string          `json:"eventName,omitempty"`
	EventOrganizer   string          `json:"eventOrganizer,omitempty"`
	EventDescription string          `json:"eventDescription,omitempty"`
	EventURL         string          `json:"eventUrl,omitempty"`
	EventType        string          `json:"eventType,omitempty"`
	EventCancelled   *bool           `json:"eventCancelled,omitempty"`
	EventCompleted   *bool           `json:"eventCompleted,omitempty"`
	StartDateTime    *time.Time      `json:"startDateTime,omitempty"`
	EndDateTime      *time.Time      `json:"endDateTime,omitempty"`
	CustomProperties []PropertyValue `json:"customProperties,omitempty"`
}

// MarketingEventDefaultResponse is the response when creating/updating a marketing event.
type MarketingEventDefaultResponse struct {
	ObjectID         string          `json:"objectId,omitempty"`
	EventName        string          `json:"eventName"`
	EventOrganizer   string          `json:"eventOrganizer"`
	EventDescription string          `json:"eventDescription,omitempty"`
	EventURL         string          `json:"eventUrl,omitempty"`
	EventType        string          `json:"eventType,omitempty"`
	EventCancelled   *bool           `json:"eventCancelled,omitempty"`
	EventCompleted   *bool           `json:"eventCompleted,omitempty"`
	StartDateTime    *time.Time      `json:"startDateTime,omitempty"`
	EndDateTime      *time.Time      `json:"endDateTime,omitempty"`
	CustomProperties []PropertyValue `json:"customProperties,omitempty"`
}

// MarketingEventPublicDefaultResponse is the full public response for an event.
type MarketingEventPublicDefaultResponse struct {
	ID               string          `json:"id"`
	ObjectID         string          `json:"objectId,omitempty"`
	EventName        string          `json:"eventName"`
	EventOrganizer   string          `json:"eventOrganizer"`
	EventDescription string          `json:"eventDescription,omitempty"`
	EventURL         string          `json:"eventUrl,omitempty"`
	EventType        string          `json:"eventType,omitempty"`
	EventCancelled   *bool           `json:"eventCancelled,omitempty"`
	EventCompleted   *bool           `json:"eventCompleted,omitempty"`
	StartDateTime    *time.Time      `json:"startDateTime,omitempty"`
	EndDateTime      *time.Time      `json:"endDateTime,omitempty"`
	CustomProperties []PropertyValue `json:"customProperties,omitempty"`
	CreatedAt        time.Time       `json:"createdAt"`
	UpdatedAt        time.Time       `json:"updatedAt"`
}

// MarketingEventReadResponse is the response for reading an event by external IDs.
type MarketingEventReadResponse struct {
	ID                string          `json:"id"`
	ObjectID          string          `json:"objectId,omitempty"`
	ExternalEventID   string          `json:"externalEventId"`
	EventName         string          `json:"eventName"`
	EventOrganizer    string          `json:"eventOrganizer"`
	EventDescription  string          `json:"eventDescription,omitempty"`
	EventURL          string          `json:"eventUrl,omitempty"`
	EventType         string          `json:"eventType,omitempty"`
	EventCancelled    *bool           `json:"eventCancelled,omitempty"`
	EventCompleted    *bool           `json:"eventCompleted,omitempty"`
	StartDateTime     *time.Time      `json:"startDateTime,omitempty"`
	EndDateTime       *time.Time      `json:"endDateTime,omitempty"`
	CustomProperties  []PropertyValue `json:"customProperties,omitempty"`
	Registrants       int             `json:"registrants"`
	Attendees         int             `json:"attendees"`
	Cancellations     int             `json:"cancellations"`
	NoShows           int             `json:"noShows"`
	CreatedAt         time.Time       `json:"createdAt"`
	UpdatedAt         time.Time       `json:"updatedAt"`
}

// MarketingEventReadResponseV2 is the V2 response for reading an event by object ID.
type MarketingEventReadResponseV2 struct {
	ObjectID         string           `json:"objectId"`
	EventName        string           `json:"eventName"`
	ExternalEventID  string           `json:"externalEventId,omitempty"`
	EventOrganizer   string           `json:"eventOrganizer,omitempty"`
	EventDescription string           `json:"eventDescription,omitempty"`
	EventURL         string           `json:"eventUrl,omitempty"`
	EventType        string           `json:"eventType,omitempty"`
	EventStatus      string           `json:"eventStatus,omitempty"`
	EventCancelled   *bool            `json:"eventCancelled,omitempty"`
	EventCompleted   *bool            `json:"eventCompleted,omitempty"`
	StartDateTime    *time.Time       `json:"startDateTime,omitempty"`
	EndDateTime      *time.Time       `json:"endDateTime,omitempty"`
	CustomProperties []map[string]any `json:"customProperties,omitempty"`
	AppInfo          map[string]any   `json:"appInfo,omitempty"`
	Registrants      *int             `json:"registrants,omitempty"`
	Attendees        *int             `json:"attendees,omitempty"`
	Cancellations    *int             `json:"cancellations,omitempty"`
	NoShows          *int             `json:"noShows,omitempty"`
	CreatedAt        time.Time        `json:"createdAt"`
	UpdatedAt        time.Time        `json:"updatedAt"`
}

// MarketingEventV2UpdateRequest is the V2 input for updating a marketing event.
type MarketingEventV2UpdateRequest struct {
	EventName        string          `json:"eventName,omitempty"`
	EventOrganizer   string          `json:"eventOrganizer,omitempty"`
	EventDescription string          `json:"eventDescription,omitempty"`
	EventURL         string          `json:"eventUrl,omitempty"`
	EventType        string          `json:"eventType,omitempty"`
	EventCancelled   *bool           `json:"eventCancelled,omitempty"`
	StartDateTime    *time.Time      `json:"startDateTime,omitempty"`
	EndDateTime      *time.Time      `json:"endDateTime,omitempty"`
	CustomProperties []PropertyValue `json:"customProperties,omitempty"`
}

// MarketingEventV2DefaultResponse is the V2 response for event mutations.
type MarketingEventV2DefaultResponse struct {
	ObjectID         string           `json:"objectId"`
	EventName        string           `json:"eventName"`
	EventOrganizer   string           `json:"eventOrganizer,omitempty"`
	EventDescription string           `json:"eventDescription,omitempty"`
	EventURL         string           `json:"eventUrl,omitempty"`
	EventType        string           `json:"eventType,omitempty"`
	EventCancelled   *bool            `json:"eventCancelled,omitempty"`
	EventCompleted   *bool            `json:"eventCompleted,omitempty"`
	StartDateTime    *time.Time       `json:"startDateTime,omitempty"`
	EndDateTime      *time.Time       `json:"endDateTime,omitempty"`
	CustomProperties []map[string]any `json:"customProperties,omitempty"`
	AppInfo          map[string]any   `json:"appInfo,omitempty"`
	CreatedAt        time.Time        `json:"createdAt"`
	UpdatedAt        time.Time        `json:"updatedAt"`
}

// MarketingEventCompleteRequest sets start and end times to mark an event as complete.
type MarketingEventCompleteRequest struct {
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`
}

// MarketingEventListResult is a paginated list of marketing events.
type MarketingEventListResult struct {
	Results []MarketingEventReadResponseV2 `json:"results"`
	Paging  *ForwardPaging                 `json:"paging,omitempty"`
}

// MarketingEventExternalUniqueIdentifier identifies an event by external IDs and app.
type MarketingEventExternalUniqueIdentifier struct {
	ExternalAccountID string `json:"externalAccountId"`
	ExternalEventID   string `json:"externalEventId"`
	AppID             int    `json:"appId"`
}

// MarketingEventEmailSubscriber represents a subscriber identified by email.
type MarketingEventEmailSubscriber struct {
	Email               string            `json:"email"`
	InteractionDateTime int64             `json:"interactionDateTime"`
	ContactProperties   map[string]string `json:"contactProperties,omitempty"`
	Properties          map[string]string `json:"properties,omitempty"`
}

// MarketingEventSubscriber represents a subscriber identified by contact ID (vid).
type MarketingEventSubscriber struct {
	VID                 *int              `json:"vid,omitempty"`
	InteractionDateTime int64             `json:"interactionDateTime"`
	Properties          map[string]string `json:"properties,omitempty"`
}

// BatchInputEmailSubscribers is a batch of email-identified subscribers.
type BatchInputEmailSubscribers struct {
	Inputs []MarketingEventEmailSubscriber `json:"inputs"`
}

// BatchInputSubscribers is a batch of contact-ID-identified subscribers.
type BatchInputSubscribers struct {
	Inputs []MarketingEventSubscriber `json:"inputs"`
}

// SubscriberEmailResponse is the response for recording attendance by email.
type SubscriberEmailResponse struct {
	Email  string `json:"email"`
	Status string `json:"status"`
}

// SubscriberVidResponse is the response for recording attendance by contact ID.
type SubscriberVidResponse struct {
	VID    int    `json:"vid"`
	Status string `json:"status"`
}

// BatchResponseSubscriberEmail is a batch response for email subscribers.
type BatchResponseSubscriberEmail struct {
	Status      string                    `json:"status"`
	Results     []SubscriberEmailResponse `json:"results"`
	StartedAt   time.Time                 `json:"startedAt"`
	CompletedAt time.Time                 `json:"completedAt"`
}

// BatchResponseSubscriberVid is a batch response for VID subscribers.
type BatchResponseSubscriberVid struct {
	Status      string                  `json:"status"`
	Results     []SubscriberVidResponse `json:"results"`
	StartedAt   time.Time               `json:"startedAt"`
	CompletedAt time.Time               `json:"completedAt"`
}

// BatchInputMarketingEventCreateRequests is a batch of event create requests.
type BatchInputMarketingEventCreateRequests struct {
	Inputs []MarketingEventCreateRequest `json:"inputs"`
}

// BatchInputMarketingEventExternalUniqueIdentifiers is a batch of external IDs.
type BatchInputMarketingEventExternalUniqueIdentifiers struct {
	Inputs []MarketingEventExternalUniqueIdentifier `json:"inputs"`
}

// BatchResponseMarketingEventPublicDefault is a batch response for event operations.
type BatchResponseMarketingEventPublicDefault struct {
	Status      string                                `json:"status"`
	Results     []MarketingEventPublicDefaultResponse `json:"results"`
	StartedAt   time.Time                             `json:"startedAt"`
	CompletedAt time.Time                             `json:"completedAt"`
}

// AttendanceCounters reports attendance statistics for a marketing event.
type AttendanceCounters struct {
	Attended   int `json:"attended"`
	Registered int `json:"registered"`
	Cancelled  int `json:"cancelled"`
	NoShows    int `json:"noShows"`
}

// ParticipationBreakdown represents a single participation record.
type ParticipationBreakdown struct {
	ID           string         `json:"id"`
	Associations map[string]any `json:"associations,omitempty"`
	Properties   map[string]any `json:"properties,omitempty"`
	CreatedAt    time.Time      `json:"createdAt"`
}

// ParticipationBreakdownResult is a paginated list of participation breakdowns.
type ParticipationBreakdownResult struct {
	Total   int                      `json:"total"`
	Results []ParticipationBreakdown `json:"results"`
	Paging  *ForwardPaging           `json:"paging,omitempty"`
}

// ParticipationBreakdownOptions configures a participation breakdown query.
type ParticipationBreakdownOptions struct {
	ContactIdentifier string `json:"contactIdentifier,omitempty"`
	State             string `json:"state,omitempty"`
	Limit             int    `json:"limit,omitempty"`
	After             string `json:"after,omitempty"`
}

// EventDetailSettings holds settings for marketing event details page.
type EventDetailSettings struct {
	AppID           int    `json:"appId"`
	EventDetailsURL string `json:"eventDetailsUrl"`
}

// EventDetailSettingsURL is the input for updating event detail settings.
type EventDetailSettingsURL struct {
	EventDetailsURL string `json:"eventDetailsUrl"`
}

// SearchResponseWrapper wraps a marketing event search result.
type SearchResponseWrapper struct {
	ID                string `json:"id,omitempty"`
	EventName         string `json:"eventName,omitempty"`
	ExternalEventID   string `json:"externalEventId,omitempty"`
	ExternalAccountID string `json:"externalAccountId,omitempty"`
}

// SearchResult is a collection of search response wrappers.
type EventSearchResult struct {
	Results []SearchResponseWrapper `json:"results"`
}

// MarketingEventIdentifiersResponse contains identifiers for a marketing event.
type MarketingEventIdentifiersResponse struct {
	ExternalEventID   string `json:"externalEventId"`
	ExternalAccountID string `json:"externalAccountId"`
	AppID             int    `json:"appId"`
}

// MarketingEventIdentifiersResult is a collection of event identifiers.
type MarketingEventIdentifiersResult struct {
	Total   int                                 `json:"total"`
	Results []MarketingEventIdentifiersResponse `json:"results"`
}

// PublicListEntry represents a list associated with a marketing event.
type PublicListEntry struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// PublicListResult is a collection of lists for an event.
type PublicListResult struct {
	Total   int               `json:"total"`
	Results []PublicListEntry `json:"results"`
}

// --- Transactional models ---

// SingleSendEmail represents the email message details for a transactional send.
type SingleSendEmail struct {
	To      string   `json:"to"`
	From    string   `json:"from,omitempty"`
	SendID  string   `json:"sendId,omitempty"`
	CC      []string `json:"cc,omitempty"`
	BCC     []string `json:"bcc,omitempty"`
	ReplyTo []string `json:"replyTo,omitempty"`
}

// SingleSendRequest is the input for sending a single transactional email.
type SingleSendRequest struct {
	EmailID            int               `json:"emailId"`
	Message            SingleSendEmail   `json:"message"`
	ContactProperties  map[string]string `json:"contactProperties,omitempty"`
	CustomProperties   map[string]any    `json:"customProperties,omitempty"`
}

// EmailSendStatusView is the response for a transactional email send.
type EmailSendStatusView struct {
	StatusID    string     `json:"statusId"`
	Status      string     `json:"status"` // PENDING, PROCESSING, CANCELED, COMPLETE
	SendResult  string     `json:"sendResult,omitempty"`
	RequestedAt *time.Time `json:"requestedAt,omitempty"`
	StartedAt   *time.Time `json:"startedAt,omitempty"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	EventID     *EventID   `json:"eventId,omitempty"`
}

// EventID identifies a send event.
type EventID struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created"`
}

// SmtpApiTokenCreateRequest is the input for creating an SMTP token.
type SmtpApiTokenCreateRequest struct {
	CampaignName  string `json:"campaignName"`
	CreateContact bool   `json:"createContact"`
}

// SmtpApiTokenView represents an SMTP API token.
type SmtpApiTokenView struct {
	ID              string    `json:"id"`
	CampaignName    string    `json:"campaignName"`
	EmailCampaignID string    `json:"emailCampaignId"`
	CreatedBy       string    `json:"createdBy"`
	CreateContact   bool      `json:"createContact"`
	Password        string    `json:"password,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
}

// SmtpTokenListResult is a paginated list of SMTP tokens.
type SmtpTokenListResult struct {
	Results []SmtpApiTokenView `json:"results"`
	Paging  *ForwardPaging     `json:"paging,omitempty"`
}

// SmtpTokenListOptions configures an SMTP token list request.
type SmtpTokenListOptions struct {
	CampaignName    string `json:"campaignName,omitempty"`
	EmailCampaignID string `json:"emailCampaignId,omitempty"`
	After           string `json:"after,omitempty"`
	Limit           int    `json:"limit,omitempty"`
}
