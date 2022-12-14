package micropass_api

import (
	"github.com/MedzikUser/go-micropass-api/types"
	"github.com/MedzikUser/libcrypto-go/aes"
	"github.com/MedzikUser/libcrypto-go/hash"
)

// EncryptionKey returns the encryption key for the user.
func (c *Client) EncryptionKey(accessToken string, masterPassword string, email string) (string, error) {
	var encryptionKey string

	var res types.UserEncryptionKeyResponse
	_, err := c.Get("/user/encryption_key", &accessToken, &res)
	if err != nil {
		return encryptionKey, err
	}

	key := hash.Pbkdf2Hash256(masterPassword, []byte(email), PasswordIterations)

	encryptionKey, err = aes.DecryptAesCbc(key, res.EncryptionKey)
	if err != nil {
		return encryptionKey, err
	}

	return encryptionKey, nil
}

// Whoami returns the user's information.
func (c *Client) Whoami(accessToken string) (types.UserWhoamiResponse, error) {
	var res types.UserWhoamiResponse
	_, err := c.Get("/user/whoami", &accessToken, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
