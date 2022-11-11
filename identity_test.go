package micropass_api_test

import "testing"

func TestRefreshToken(t *testing.T) {
	_, err := client.RefreshToken(refreshToken)
	if err != nil {
		t.Error(err)
		return
	}
}

// Note that registration and login are done in the main test.
