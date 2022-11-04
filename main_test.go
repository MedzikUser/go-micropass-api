package avapi_test

import (
	"os"
	"testing"

	"github.com/MedzikUser/go-avapi"
	"github.com/ddosify/go-faker/faker"
)

var registered bool
var loggedIn bool

var email string
var password string

var accessToken string
var refreshToken string

var encryptionKey string

var client = avapi.NewClient()

func TestMain(m *testing.M) {
	if !registered {
		email, password = fakeData()

		err := client.Register(email, password, nil)
		if err != nil {
			panic(err)
		}

		registered = true
	}

	if !loggedIn {
		res, err := client.Login(email, password)
		if err != nil {
			panic(err)
		}

		accessToken = res.AccessToken
		refreshToken = res.RefreshToken

		loggedIn = true
	}

	if len(encryptionKey) == 0 {
		var err error
		encryptionKey, err = client.EncryptionKey(accessToken, password, email)
		if err != nil {
			panic(err)
		}
	}

	code := m.Run()
	os.Exit(code)
}

func fakeData() (string, string) {
	faker := faker.NewFaker()

	fakeEmail := "_demo_" + faker.RandomEmail()
	fakePassword := faker.RandomPassword()

	return fakeEmail, fakePassword
}
