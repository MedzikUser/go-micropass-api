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
