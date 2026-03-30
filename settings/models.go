package settings

// --- Business Units models ---

// PublicBusinessUnit represents a HubSpot business unit.
type PublicBusinessUnit struct {
	LogoMetadata *PublicBusinessUnitLogoMetadata `json:"logoMetadata,omitempty"`
	Name         string                          `json:"name"`
	ID           string                          `json:"id"`
}

// PublicBusinessUnitLogoMetadata contains logo information for a business unit.
type PublicBusinessUnitLogoMetadata struct {
	LogoAltText string `json:"logoAltText,omitempty"`
	ResizedURL  string `json:"resizedUrl,omitempty"`
	LogoURL     string `json:"logoUrl,omitempty"`
}

// BusinessUnitsResult is the response for listing business units (no paging).
type BusinessUnitsResult struct {
	Results []*PublicBusinessUnit `json:"results"`
}

// BusinessUnitListOptions contains query parameters for listing business units.
type BusinessUnitListOptions struct {
	Properties []string `json:"-"`
	Name       []string `json:"-"`
}

// --- Users models ---

// PublicUser represents a HubSpot user.
type PublicUser struct {
	FirstName        string   `json:"firstName,omitempty"`
	LastName         string   `json:"lastName,omitempty"`
	PrimaryTeamID    string   `json:"primaryTeamId,omitempty"`
	RoleIDs          []string `json:"roleIds,omitempty"`
	SendWelcomeEmail *bool    `json:"sendWelcomeEmail,omitempty"`
	RoleID           string   `json:"roleId,omitempty"`
	SecondaryTeamIDs []string `json:"secondaryTeamIds,omitempty"`
	ID               string   `json:"id"`
	SuperAdmin       *bool    `json:"superAdmin,omitempty"`
	Email            string   `json:"email"`
}

// UserProvisionRequest is the body for creating a new user.
type UserProvisionRequest struct {
	FirstName        string   `json:"firstName,omitempty"`
	LastName         string   `json:"lastName,omitempty"`
	PrimaryTeamID    string   `json:"primaryTeamId,omitempty"`
	SendWelcomeEmail *bool    `json:"sendWelcomeEmail,omitempty"`
	RoleID           string   `json:"roleId,omitempty"`
	SecondaryTeamIDs []string `json:"secondaryTeamIds,omitempty"`
	Email            string   `json:"email"`
}

// PublicUserUpdate is the body for updating a user.
type PublicUserUpdate struct {
	FirstName        string   `json:"firstName,omitempty"`
	LastName         string   `json:"lastName,omitempty"`
	PrimaryTeamID    string   `json:"primaryTeamId,omitempty"`
	RoleID           string   `json:"roleId,omitempty"`
	SecondaryTeamIDs []string `json:"secondaryTeamIds,omitempty"`
}

// UsersListResult is the paginated response for listing users.
type UsersListResult struct {
	Results []*PublicUser  `json:"results"`
	Paging  *ForwardPaging `json:"paging,omitempty"`
}

// UsersListOptions contains query parameters for listing users.
type UsersListOptions struct {
	Limit int    `json:"-"`
	After string `json:"-"`
}

// --- Roles models ---

// PublicPermissionSet represents a role/permission set.
type PublicPermissionSet struct {
	RequiresBillingWrite bool   `json:"requiresBillingWrite"`
	Name                 string `json:"name"`
	ID                   string `json:"id"`
}

// RolesResult is the response for listing roles (no paging).
type RolesResult struct {
	Results []*PublicPermissionSet `json:"results"`
}

// --- Teams models ---

// PublicTeam represents a HubSpot team.
type PublicTeam struct {
	UserIDs          []string `json:"userIds"`
	Name             string   `json:"name"`
	ID               string   `json:"id"`
	SecondaryUserIDs []string `json:"secondaryUserIds"`
}

// TeamsResult is the response for listing teams (no paging).
type TeamsResult struct {
	Results []*PublicTeam `json:"results"`
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
