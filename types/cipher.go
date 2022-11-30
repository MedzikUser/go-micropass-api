package types

import (
	"encoding/json"

	"github.com/MedzikUser/libcrypto-go/aes"
)

type Cipher struct {
	Id        string `json:"id"`
	Favorite  bool   `json:"favorite"`
	Directory string `json:"dir"`
	Data      CipherData
	CreatedAt string `json:"created,omitempty"`
	UpdatedAt string `json:"updated,omitempty"`
}

type EncryptedCipher struct {
	Id        string `json:"id"`
	Favorite  bool   `json:"favorite"`
	Directory string `json:"dir"`
	Data      string `json:"data"`
	CreatedAt string `json:"created,omitempty"`
	UpdatedAt string `json:"updated,omitempty"`
}

type CipherData struct {
	Type        int                `json:"type"`
	Name        string             `json:"name"`
	Note        *string            `json:"note,omitempty"`
	Fields      CipherFieldTypeMap `json:"fields,omitempty"`
	TypedFields CipherTypedFields  `json:"-"`
}

type CipherFieldType = struct {
	Type  int    `json:"typ"`
	Value string `json:"val"`
}
type CipherFieldTypeMap = map[string]CipherFieldType

type CipherTypedFields struct {
	Username string
	Password string
	OTPAuth  string
	URL      string
}

func (cipher *Cipher) Unmarshal(data string) error {
	err := json.Unmarshal([]byte(data), cipher)
	if err != nil {
		return err
	}

	cipher.Data.UnmarshalFields()

	return nil
}

func (cipher *Cipher) UnmarshalEncrypt(data string, encryptionKey string) error {
	encCipher := EncryptedCipher{}
	err := json.Unmarshal([]byte(data), &encCipher)
	if err != nil {
		return err
	}

	cipher.Id = encCipher.Id
	cipher.Favorite = encCipher.Favorite
	cipher.Directory = encCipher.Directory
	cipher.CreatedAt = encCipher.CreatedAt
	cipher.UpdatedAt = encCipher.UpdatedAt
	err = cipher.Data.UnmarshalEncrypt(encryptionKey, encCipher.Data)
	if err != nil {
		return err
	}

	cipher.Data.UnmarshalFields()

	return nil
}

func (cipher *Cipher) MarshalEncrypt(data string, encryptionKey string) error {
	encCipher := EncryptedCipher{
		Id:        cipher.Id,
		Favorite:  cipher.Favorite,
		Directory: cipher.Directory,
		CreatedAt: cipher.CreatedAt,
		UpdatedAt: cipher.UpdatedAt,
	}
	err := json.Unmarshal([]byte(data), &encCipher)
	if err != nil {
		return err
	}

	encCipher.Data, err = cipher.Data.MarshalEncrypt(encryptionKey)
	if err != nil {
		return err
	}

	return nil
}

// MarshalFields converts typed fields to custom fields.
func (cipher *CipherData) MarshalFields() {
	if cipher.Fields == nil {
		cipher.Fields = CipherFieldTypeMap{}
	}

	if cipher.TypedFields.Username != "" {
		cipher.Fields["user"] = CipherFieldType{
			Type:  -1,
			Value: cipher.TypedFields.Username,
		}
	}

	if cipher.TypedFields.Password != "" {
		cipher.Fields["pass"] = CipherFieldType{
			Type:  -1,
			Value: cipher.TypedFields.Password,
		}
	}

	if cipher.TypedFields.OTPAuth != "" {
		cipher.Fields["otpauth"] = CipherFieldType{
			Type:  -1,
			Value: cipher.TypedFields.OTPAuth,
		}
	}

	if cipher.TypedFields.URL != "" {
		cipher.Fields["url"] = CipherFieldType{
			Type:  -1,
			Value: cipher.TypedFields.URL,
		}
	}
}

// UnmarshalFields converts custom fields to typed fields.
func (cipher *CipherData) UnmarshalFields() {
	for k, v := range cipher.Fields {
		v := v.Value

		switch k {
		case "user":
			cipher.TypedFields.Username = v
		case "pass":
			cipher.TypedFields.Password = v
		case "otpauth":
			cipher.TypedFields.OTPAuth = v
		case "url":
			cipher.TypedFields.URL = v
		}
	}
}

// Marshal returns a JSON string of the cipher.
func (cipher *CipherData) Marshal() (string, error) {
	var jsonString string

	// Convert typed fields to custom fields

	cipher.MarshalFields()

	jsonBytes, err := json.Marshal(cipher)
	if err != nil {
		return jsonString, err
	}

	jsonString = string(jsonBytes)

	return jsonString, nil
}

// MarshalEncrypt returns an encrypted JSON string of the cipher.
func (cipher *CipherData) MarshalEncrypt(encryptionKey string) (string, error) {
	var cipherText string

	clearText, err := cipher.Marshal()
	if err != nil {
		return cipherText, err
	}

	cipherText, err = aes.EncryptAesCbc(encryptionKey, clearText)
	if err != nil {
		return cipherText, err
	}

	return cipherText, nil
}

// Unmarshal parses a JSON string of the cipher.
func (cipher *CipherData) Unmarshal(jsonString string) error {
	err := json.Unmarshal([]byte(jsonString), cipher)
	if err != nil {
		return err
	}

	cipher.UnmarshalFields()

	return nil
}

// Unmarshal decrypts and parses a JSON string of the cipher.
func (cipher *CipherData) UnmarshalEncrypt(encryptionKey string, cipherText string) error {
	jsonString, err := aes.DecryptAesCbc(encryptionKey, cipherText)
	if err != nil {
		return err
	}

	cipher.Unmarshal(jsonString)

	return nil
}

type CipherRequest struct {
	Data string `json:"data"`
}

type CipherUpdateRequest struct {
	Id        string `json:"id"`
	Favorite  bool   `json:"favorite"`
	Directory string `json:"dir"`
	Data      string `json:"data"`
}

type CipherInsertResponse struct {
	Id string `json:"id"`
}

type CipherTakeResponse struct {
	Id        string `json:"id"`
	Favorite  bool   `json:"favorite"`
	Directory string `json:"dir"`
	Data      string `json:"data"`
	CreatedAt int64  `json:"created"`
	UpdatedAt int64  `json:"updated"`
}

type CipherListResponse struct {
	Updated []string `json:"updated"`
	Deleted []string `json:"deleted"`
}

var (
	CipherTypeAccount = 0
)
