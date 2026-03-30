package marketing

import (
	"github.com/scopiousdigital/hubspot-go/internal/api"
)

// Service provides access to all HubSpot Marketing APIs.
type Service struct {
	requester api.Requester

	Forms         *FormsService
	Emails        *EmailsService
	Statistics    *StatisticsService
	Events        *EventsService
	Transactional *TransactionalService
}

// NewService creates a new Marketing service. Called by the root hubspot package.
func NewService(r api.Requester) *Service {
	return &Service{
		requester:     r,
		Forms:         &FormsService{requester: r},
		Emails:        &EmailsService{requester: r},
		Statistics:    &StatisticsService{requester: r},
		Events:        &EventsService{requester: r},
		Transactional: &TransactionalService{requester: r},
	}
}
