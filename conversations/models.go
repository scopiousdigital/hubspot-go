package conversations

// IdentificationTokenGenerationRequest is the body for generating a visitor identification token.
type IdentificationTokenGenerationRequest struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email"`
}

// IdentificationTokenResponse contains the generated visitor identification token.
type IdentificationTokenResponse struct {
	Token string `json:"token"`
}
