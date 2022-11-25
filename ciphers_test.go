package micropass_api_test

import (
	"errors"
	"testing"
	"time"

	"github.com/MedzikUser/go-micropass-api/types"
)

var ciphersId []string
var ciphers []types.Cipher

func TestCiphers(t *testing.T) {
	// insert cipher
	if err := insertCipher(); err != nil {
		t.Error(err)
	}

	// list ciphers
	if err := listCiphers(); err != nil {
		t.Error(err)
	}

	// sync ciphers
	if err := syncCiphers(); err != nil {
		t.Error(err)
	}

	// take cipher
	if err := takeCipher(); err != nil {
		t.Error(err)
	}

	// update cipher
	if err := updateCipher(); err != nil {
		t.Error(err)
	}

	// validate ciphers id
	for i, cipher := range ciphers {
		if cipher.Id != ciphersId[i] {
			t.Error("Invalid cipher id. (expected: " + ciphersId[i] + ", got: " + cipher.Id + ")")
		}
	}

	// delete ciphers
	if err := deleteCiphers(); err != nil {
		t.Error(err)
	}
}

var exprectedCipherData types.CipherData

func insertCipher() error {
	fakeUsername, fakePassword := fakeData()

	cipher := types.CipherData{
		Type: types.CipherTypeAccount,
		Name: "Example",
		Fields: map[string]string{
			"Custom": "something",
		},
		TypedFields: types.CipherTypedFields{
			URL:      "https://example.com",
			Username: fakeUsername,
			Password: fakePassword,
		},
	}

	_, err := client.InsertCipher(accessToken, encryptionKey, cipher)
	if err != nil {
		return err
	}

	exprectedCipherData = cipher

	return nil
}

func listCiphers() error {
	ciphers, err := client.ListCiphers(accessToken, nil)
	if err != nil {
		return err
	}

	if len(ciphers) == 0 {
		return errors.New("failed to list ciphers owned by the user.")
	}

	ciphersId = ciphers

	return nil
}

func syncCiphers() error {
	timeNow := time.Now().Add(1 * time.Second)
	ciphersSync, err := client.ListCiphers(accessToken, &timeNow)
	if err != nil {
		return err
	}

	if len(ciphersSync) != 0 {
		return errors.New("failed to sync ciphers (this should not return any cipher).")
	}

	return nil
}

func updateCipher() error {
	fakeUsername, fakePassword := fakeData()

	cipher := types.CipherData{
		Type: types.CipherTypeAccount,
		Name: "Example",
		TypedFields: types.CipherTypedFields{
			Username: fakeUsername,
			Password: fakePassword,
		},
	}

	err := client.UpdateCipher(accessToken, encryptionKey, ciphersId[0], cipher)
	if err != nil {
		return err
	}

	return nil
}

func takeCipher() error {
	for _, cipherId := range ciphersId {
		cipher, err := client.TakeCipher(accessToken, encryptionKey, cipherId)
		if err != nil {
			return err
		}

		ciphers = append(ciphers, *cipher)
	}

	return nil
}

func deleteCiphers() error {
	for _, id := range ciphersId {
		err := client.DeleteCipher(accessToken, id)
		if err != nil {
			return err
		}
	}

	return nil
}
