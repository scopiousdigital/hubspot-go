package crm

import (
	"github.com/scopiousdigital/hubspot-go/internal/api"
)

// Service provides access to all HubSpot CRM APIs.
type Service struct {
	requester api.Requester

	// Standard CRM object services — all share the same CRUD/batch/search pattern.
	Contacts  *ObjectService
	Companies *ObjectService
	Deals     *ObjectService
	Tickets   *ObjectService
	Products  *ObjectService
	LineItems *ObjectService
	Quotes    *ObjectService

	// Activity objects
	Calls          *ObjectService
	Emails         *ObjectService
	Meetings       *ObjectService
	Notes          *ObjectService
	Tasks          *ObjectService
	Communications *ObjectService
	PostalMail     *ObjectService

	// Other CRM objects
	FeedbackSubmissions *ObjectService
	Goals               *ObjectService
	Leads               *ObjectService
	DealSplits          *ObjectService
	Taxes               *ObjectService

	// Non-standard CRM services
	Owners       *OwnersService
	Properties   *PropertiesService
	Pipelines    *PipelinesService
	Schemas      *SchemasService
	Associations *AssociationsService
	Imports      *ImportsService
	Exports      *ExportsService
	Lists        *ListsService
	Timeline     *TimelineService
	Extensions   *ExtensionsService
	Commerce     *CommerceService
}

// NewService creates a new CRM service. Called by the root hubspot package.
func NewService(r api.Requester) *Service {
	s := &Service{requester: r}

	// Standard CRM object services
	s.Contacts = newObjectService(r, "contacts")
	s.Companies = newObjectService(r, "companies")
	s.Deals = newObjectService(r, "deals")
	s.Tickets = newObjectService(r, "tickets")
	s.Products = newObjectService(r, "products")
	s.LineItems = newObjectService(r, "line_items")
	s.Quotes = newObjectService(r, "quotes")

	// Activity objects
	s.Calls = newObjectService(r, "calls")
	s.Emails = newObjectService(r, "emails")
	s.Meetings = newObjectService(r, "meetings")
	s.Notes = newObjectService(r, "notes")
	s.Tasks = newObjectService(r, "tasks")
	s.Communications = newObjectService(r, "communications")
	s.PostalMail = newObjectService(r, "postal_mail")

	// Other objects
	s.FeedbackSubmissions = newObjectService(r, "feedback_submissions")
	s.Goals = newObjectService(r, "goals")
	s.Leads = newObjectService(r, "leads")
	s.DealSplits = newObjectService(r, "deal_splits")
	s.Taxes = newObjectService(r, "taxes")

	// Non-standard services
	s.Owners = &OwnersService{requester: r}
	s.Properties = &PropertiesService{requester: r}
	s.Pipelines = &PipelinesService{requester: r}
	s.Schemas = &SchemasService{requester: r}
	s.Associations = &AssociationsService{requester: r}
	s.Imports = &ImportsService{requester: r}
	s.Exports = &ExportsService{requester: r}
	s.Lists = &ListsService{requester: r}
	s.Timeline = &TimelineService{requester: r}
	s.Extensions = &ExtensionsService{requester: r}
	s.Commerce = newCommerceService(r)

	return s
}
