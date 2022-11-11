package micropass_api

import (
	"github.com/MedzikUser/go-micropass-api/types"
	"github.com/MedzikUser/libcrypto-go"
	"github.com/MedzikUser/libcrypto-go/aes"
	"github.com/MedzikUser/libcrypto-go/hash"
)

// * Endpoint /identity/token

// Login - login to MicroPass.
func (c *Client) Login(email string, password string) (*types.IdentityLoginResponse, error) {
	emailBytes := []byte(email)
	password = hash.Pbkdf2Hash256(password, emailBytes, PasswordIterations)
	password = hash.Pbkdf2Hash512(password, emailBytes, 1)

	body := types.IdentityTokenRequest{
		GrantType: types.IdentityGrantTypePassword,
		Email:     &email,
		Password:  &password,
	}

	var res types.IdentityLoginResponse
	_, err := c.Post("/identity/token", nil, body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// RefreshToken - returns a new access token.
func (c *Client) RefreshToken(refreshToken string) (*types.IdentityRefreshTokenResponse, error) {
	body := types.IdentityTokenRequest{
		GrantType:    types.IdentityGrantTypeAccessToken,
		RefreshToken: &refreshToken,
	}

	var res types.IdentityRefreshTokenResponse
	_, err := c.Post("/identity/token", nil, body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// * Endpoint /identity/register

// Register - creates a new account in MicroPass.
func (c *Client) Register(email string, password string, passwordHint *string) error {
	emailBytes := []byte(email)
	password = hash.Pbkdf2Hash256(password, emailBytes, PasswordIterations)

	// generate salt for encryption key
	encKeySalt, err := libcrypto.GenerateSalt(32)
	if err != nil {
		return err
	}

	// make one iteration of the password with a different salt
	encKey := hash.Pbkdf2Hash256(password, encKeySalt, 1)

	// encrypt the encryption key using the master password to pass it to MicroPass server
	encKeyAes, err := aes.EncryptAesCbc(password, encKey)
	if err != nil {
		return err
	}

	// do one more iteration because the previous key was used for encryption key
	password = hash.Pbkdf2Hash512(password, emailBytes, 1)

	body := types.IdentityRegisterRequest{
		Email:         email,
		Password:      password,
		EncryptionKey: encKeyAes,
		PasswordHint:  passwordHint,
	}

	_, err = c.Post("/identity/register", nil, body, nil)
	if err != nil {
		return err
	}

	return nil
}
