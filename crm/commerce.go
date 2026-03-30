package crm

import (
	"github.com/scopiousdigital/hubspot-go/internal/api"
)

// CommerceService provides access to HubSpot commerce APIs.
// Commerce objects (invoices, etc.) use the same ObjectService pattern
// as standard CRM objects, just with a commerce-specific object type.
type CommerceService struct {
	requester api.Requester

	// Invoices follows the standard ObjectService pattern at /crm/v3/objects/invoices.
	Invoices *ObjectService
}

// newCommerceService creates a new CommerceService.
func newCommerceService(r api.Requester) *CommerceService {
	return &CommerceService{
		requester: r,
		Invoices:  newObjectService(r, "invoices"),
	}
}
