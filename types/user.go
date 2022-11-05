package types

type UserEncryptionKeyResponse struct {
	EncryptionKey string `json:"encryption_key"`
}

type UserWhoamiResponse struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
