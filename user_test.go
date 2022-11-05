package avapi_test

import (
	"testing"
)

func TestEncryptionKey(t *testing.T) {
	_, err := client.EncryptionKey(accessToken, password, email)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestWhoami(t *testing.T) {
	res, err := client.Whoami(accessToken)
	if err != nil {
		t.Error(err)
		return
	}

	if len(res.Id) == 0 {
		t.Error("The returned id is null")
	}

	if len(res.Email) == 0 {
		t.Error("The returned email is null")
	}

	if len(res.Username) == 0 {
		t.Error("The returned username is null")
	}
}
