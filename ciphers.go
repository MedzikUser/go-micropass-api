package micropass_api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/MedzikUser/go-micropass-api/types"
	"github.com/MedzikUser/libcrypto-go/aes"
)

// InsertCipher inserts a new cipher.
func (c *Client) InsertCipher(accessToken string, encryptionKey string, cipher types.Cipher) (string, error) {
	var cipherId string

	cipherBytes, err := json.Marshal(cipher)
	if err != nil {
		return cipherId, err
	}
	clearText := string(cipherBytes)

	cipherText, err := aes.EncryptAesCbc(encryptionKey, clearText)
	if err != nil {
		return cipherId, err
	}

	body := types.CipherRequest{
		Data: cipherText,
	}

	var res types.CipherInsertResponse
	_, err = c.Post("/ciphers/insert", &accessToken, body, &res)
	if err != nil {
		return cipherId, err
	}

	cipherId = res.Id

	return cipherId, nil
}

// TakeCipher returns a cipher.
func (c *Client) TakeCipher(accessToken string, encryptionKey string, id string) (*types.Cipher, error) {
	var cipher types.Cipher

	var res types.CipherTakeResponse
	_, err := c.Get("/ciphers/get/"+id, &accessToken, &res)
	if err != nil {
		return nil, err
	}

	cipherDataString, err := aes.DecryptAesCbc(encryptionKey, res.Data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(cipherDataString), &cipher)
	if err != nil {
		return nil, err
	}

	cipher.Id = &res.Id

	return &cipher, nil
}

// UpdateCipher updates a cipher.
func (c *Client) UpdateCipher(accessToken string, encryptionKey string, id string, cipher types.Cipher) error {
	cipherBytes, err := json.Marshal(cipher)
	if err != nil {
		return err
	}
	clearText := string(cipherBytes)

	cipherText, err := aes.EncryptAesCbc(encryptionKey, clearText)
	if err != nil {
		return err
	}

	body := types.CipherUpdateRequest{
		Id:   id,
		Data: cipherText,
	}

	_, err = c.Patch("/ciphers/update", &accessToken, body, nil)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCipher deletes a cipher.
func (c *Client) DeleteCipher(accessToken string, id string) error {
	_, err := c.Delete("/ciphers/delete/"+id, &accessToken, nil)
	if err != nil {
		return err
	}

	return nil
}

// ListCiphers returns a list of ciphers.
func (c *Client) ListCiphers(accessToken string, lastSync *time.Time) ([]string, error) {
	var ciphers []string

	var url string
	if lastSync == nil {
		url = "/ciphers/list"
	} else {
		url = fmt.Sprintf("/ciphers/list?lastSync=%d", lastSync.Unix())
	}

	var res types.CipherListResponse
	_, err := c.Get(url, &accessToken, &res)
	if err != nil {
		return ciphers, err
	}

	ciphers = res.Updated

	return ciphers, nil
}
