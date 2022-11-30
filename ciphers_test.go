package micropass_api_test

import (
	"errors"
	"testing"
	"time"

	"github.com/MedzikUser/go-micropass-api/types"
	"github.com/MedzikUser/libcrypto-go/hash"
)

var ciphersId []string
var ciphers []types.Cipher

func TestCipherDeserialization(t *testing.T) {
	cipherJson := `
{
    "id": "aa770bed-e199-41f1-90b1-c4578104e22b",
    "favorite": false,
    "dir": "622e5baf-f4b4-427b-b1dd-d54cded668e3",
    "data": {
        "type": 1,
        "name": "Example",
        "favorite": true,
        "fields": {
            "user": {
                "typ": -1,
                "val": "medzik@example.com"
            },
            "pass": {
                "typ": -1,
                "val": "SecretPassword"
            },
            "totp": {
                "typ": -1,
                "val": "otpauth://totp/medzik@example.com?secret=JBSWY3DPEHPK3PXP&issuer=example.com"
            },
            "url": {
                "typ": -1,
                "val": "https://example.com"
            },
            "note": {
                "typ": -1,
                "val": "my note about this cipher"
            },
            "Custom Field": {
                "typ": 0,
                "val": "This is a text in your custom field"
            },
            "Custom Secret Field": {
                "typ": 1,
                "val": "This is a secret text in your secret custom field"
            }
        }
    },
    "attachments": []
}`
	cipherJsonEnc := "{\"id\":\"aa770bed-e199-41f1-90b1-c4578104e22b\",\"favorite\":false,\"dir\":\"622e5baf-f4b4-427b-b1dd-d54cded668e3\",\"data\":\"cc52ce36d8e1d47aada7dbc4735852437f43718f05ba23eac2e111ec82a32bfe6b5e300d8f74a758dec89fcd77a30cefff4ab82477a8e800274c93800b53751925f8bf02522682fe313574d79eb38bfab691c207c825cc192c2e2869ccd5a1b1ef457e7a394b30ebfd6c486dd87cb8a203aa630365de4368a71e2cfd1a61f11a67f82b842dd15266d49a7314c46c578807b890b3415f2271ee4751e778603463c4f9fa0b442f16774684c93c8981286997b52df3c5fa8f1b5685e86cb7254a926d94d85503d40a6b0fbcd226f9c666dcab4c683f46a47fb73ef7e1374f42e27d37831b243b80c7e4c70536b93f1af9aaa38550c99aef03216af8990c567786c2c4f0c7d23b87eadea9a7f5e7040176e0645be16e94ea8e6b7df14f1275d8bea92c6664adfcd549954d84bc7fadbf496fbbcd7e9ad6d05556dd530e0483ec176f963722064df3c189ed766b960d8b4bfce2170536c393dbbf2b706a4bca712ec27157b218a485cd1ff449bb5b85acda659f0f0c057ea4bda5006c2e9afca51c5758e3b341731616e40a6357a4fa7c499f464ca8aa48b5ddf58e9f7ec84d0030374290f4f486a40c3d84339fcc90cc32a3a7e85b109ad9f02a4620a6e24bd2c93c9a31fda2077160b2e888ef92c28a34a4d9b9f3b71e539184869acf98cada0590074c9456536f5b0299dc6ea946ffeafe54102caff6b7f27208552401436b1ab3d235673722f6fbb3319caf80dd429625\"}"
	encKey := hash.Pbkdf2Hash256("hello world", []byte("salt"), 1000)

	cipherOne := types.Cipher{}

	err := cipherOne.Unmarshal(cipherJson)
	if err != nil {
		t.Error(err)
	}

	cipherTwo := types.Cipher{}

	err = cipherTwo.UnmarshalEncrypt(cipherJsonEnc, encKey)
	if err != nil {
		t.Error(err)
	}
}

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
		Fields: types.CipherFieldTypeMap{
			"Custom": types.CipherFieldType{
				Type:  0,
				Value: "This is a text in your custom field",
			},
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
