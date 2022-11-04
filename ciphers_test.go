package avapi_test

import (
	"errors"
	"testing"
	"time"

	"github.com/MedzikUser/go-avapi/types"
)

var ciphersId []string
var ciphers []types.Cipher

func TestCiphersList(t *testing.T) {
	// insert cipher
	if err := testInsert(); err != nil {
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

	// validate ciphers id
	for i, cipher := range ciphers {
		if *cipher.Id != ciphersId[i] {
			t.Error("Invalid cipher id. (expected: " + ciphersId[i] + ", got: " + *cipher.Id + ")")
		}
	}

	// delete ciphers
	if err := deleteCiphers(); err != nil {
		t.Error(err)
	}
}

func testInsert() error {
	fakeUsername, fakePassword := fakeData()

	cipher := types.Cipher{
		Type:     types.CipherTypeAccount,
		Name:     "test",
		Username: &fakeUsername,
		Password: &fakePassword,
	}

	_, err := client.InsertCipher(accessToken, cipher, encryptionKey)
	if err != nil {
		return err
	}

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
