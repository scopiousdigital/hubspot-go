package settings

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/scopiousdigital/hubspot-go/internal/api"
)

const (
	businessUnitsBasePath = "/settings/v3/business-units/users"
	usersBasePath         = "/settings/v3/users"
	rolesBasePath         = "/settings/v3/users/roles"
	teamsBasePath         = "/settings/v3/users/teams"
)

// Service provides access to HubSpot Settings APIs.
type Service struct {
	requester api.Requester

	// BusinessUnits provides methods for managing business units.
	BusinessUnits *BusinessUnitsService
	// Users provides methods for managing users.
	Users *UsersService
	// Roles provides methods for listing roles.
	Roles *RolesService
	// Teams provides methods for listing teams.
	Teams *TeamsService
}

// NewService creates a new Settings service.
func NewService(r api.Requester) *Service {
	return &Service{
		requester:     r,
		BusinessUnits: &BusinessUnitsService{requester: r},
		Users:         &UsersService{requester: r},
		Roles:         &RolesService{requester: r},
		Teams:         &TeamsService{requester: r},
	}
}

// --- BusinessUnitsService ---

// BusinessUnitsService provides methods for the Business Units API.
type BusinessUnitsService struct {
	requester api.Requester
}

// GetByUserID retrieves business units for a given user ID.
func (s *BusinessUnitsService) GetByUserID(ctx context.Context, userID string, opts *BusinessUnitListOptions) (*BusinessUnitsResult, error) {
	path := fmt.Sprintf("%s/%s", businessUnitsBasePath, userID)
	q := url.Values{}
	if opts != nil {
		if len(opts.Properties) > 0 {
			q.Set("properties", strings.Join(opts.Properties, ","))
		}
		if len(opts.Name) > 0 {
			for _, n := range opts.Name {
				q.Add("name", n)
			}
		}
	}
	var result BusinessUnitsResult
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- UsersService ---

// UsersService provides methods for managing users.
type UsersService struct {
	requester api.Requester
}

// Create provisions a new user.
func (s *UsersService) Create(ctx context.Context, input *UserProvisionRequest) (*PublicUser, error) {
	var result PublicUser
	if err := s.requester.Post(ctx, usersBasePath, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetByID retrieves a user by user ID or email.
// idProperty can be "USER_ID" or "EMAIL"; leave empty for default (USER_ID).
func (s *UsersService) GetByID(ctx context.Context, userID string, idProperty string) (*PublicUser, error) {
	path := fmt.Sprintf("%s/%s", usersBasePath, userID)
	q := url.Values{}
	if idProperty != "" {
		q.Set("idProperty", idProperty)
	}
	var result PublicUser
	if err := s.requester.Get(ctx, path, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List retrieves a page of users.
func (s *UsersService) List(ctx context.Context, opts *UsersListOptions) (*UsersListResult, error) {
	q := url.Values{}
	if opts != nil {
		if opts.Limit > 0 {
			q.Set("limit", strconv.Itoa(opts.Limit))
		}
		if opts.After != "" {
			q.Set("after", opts.After)
		}
	}
	var result UsersListResult
	if err := s.requester.Get(ctx, usersBasePath, q, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Replace updates a user (full replacement).
// idProperty can be "USER_ID" or "EMAIL"; leave empty for default (USER_ID).
func (s *UsersService) Replace(ctx context.Context, userID string, input *PublicUserUpdate, idProperty string) (*PublicUser, error) {
	path := fmt.Sprintf("%s/%s", usersBasePath, userID)
	if idProperty != "" {
		path += "?idProperty=" + url.QueryEscape(idProperty)
	}
	var result PublicUser
	if err := s.requester.Patch(ctx, path, input, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive removes a user.
// idProperty can be "USER_ID" or "EMAIL"; leave empty for default (USER_ID).
func (s *UsersService) Archive(ctx context.Context, userID string, idProperty string) error {
	path := fmt.Sprintf("%s/%s", usersBasePath, userID)
	if idProperty != "" {
		path += "?idProperty=" + url.QueryEscape(idProperty)
	}
	return s.requester.Delete(ctx, path)
}

// --- RolesService ---

// RolesService provides methods for listing roles/permission sets.
type RolesService struct {
	requester api.Requester
}

// GetAll retrieves all roles (permission sets).
func (s *RolesService) GetAll(ctx context.Context) (*RolesResult, error) {
	var result RolesResult
	if err := s.requester.Get(ctx, rolesBasePath, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// --- TeamsService ---

// TeamsService provides methods for listing teams.
type TeamsService struct {
	requester api.Requester
}

// GetAll retrieves all teams.
func (s *TeamsService) GetAll(ctx context.Context) (*TeamsResult, error) {
	var result TeamsResult
	if err := s.requester.Get(ctx, teamsBasePath, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
