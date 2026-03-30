package oauth

// AccessTokenInfo contains metadata about an OAuth access token.
type AccessTokenInfo struct {
	HubID     int      `json:"hubId"`
	UserID    int      `json:"userId"`
	Scopes    []string `json:"scopes"`
	TokenType string   `json:"tokenType"`
	User      string   `json:"user,omitempty"`
	HubDomain string   `json:"hubDomain,omitempty"`
	AppID     int      `json:"appId"`
	ExpiresIn int      `json:"expiresIn"`
	Token     string   `json:"token"`
}

// RefreshTokenInfo contains metadata about an OAuth refresh token.
type RefreshTokenInfo struct {
	HubID     int      `json:"hubId"`
	UserID    int      `json:"userId"`
	Scopes    []string `json:"scopes"`
	TokenType string   `json:"tokenType"`
	User      string   `json:"user,omitempty"`
	HubDomain string   `json:"hubDomain,omitempty"`
	ClientID  string   `json:"clientId"`
	Token     string   `json:"token"`
}

// TokenResponse is the response when exchanging an authorization code or
// refreshing an OAuth token.
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token,omitempty"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}

// TokenCreateRequest is the input for exchanging an authorization code or
// refreshing an OAuth token. Fields are sent as form-encoded values, but
// we model them here as a struct for the JSON-based Requester. The HubSpot
// OAuth token endpoint accepts JSON as well.
type TokenCreateRequest struct {
	GrantType    string `json:"grant_type"`
	Code         string `json:"code,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

