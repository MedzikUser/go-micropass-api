package types

type IdentityLoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type IdentityRefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type IdentityTokenRequest struct {
	GrantType string `form:"grant_type" json:"grant_type" binding:"required"` // refresh_token or password

	// Needed for grant_type="refresh_token"
	RefreshToken *string `form:"refresh_token" json:"refresh_token"`

	// Needed for grant_type="password"
	Email    *string `form:"email"    json:"email"`
	Password *string `form:"password" json:"password"`
}

type IdentityRegisterRequest struct {
	Email         string  `form:"email"          json:"email"          binding:"required"`
	Password      string  `form:"password"       json:"password"       binding:"required"`
	EncryptionKey string  `form:"encryption_key" json:"encryption_key" binding:"required"`
	PasswordHint  *string `form:"password_hint"  json:"password_hint"`
}

var (
	IdentityGrantTypePassword    = "password"
	IdentityGrantTypeAccessToken = "refresh_token"
)
