package types

type Cipher struct {
	Id          *string `json:"id"`
	Type        int     `json:"type"`
	Name        string  `json:"name"`
	Username    *string `json:"username"`
	Password    *string `json:"password"`
	Favorite    bool    `json:"favorite"`
	DirectoryId string  `json:"dir"`
}

type CipherRequest struct {
	Data string `json:"data"`
}

type CipherInsertResponse struct {
	Id string `json:"id"`
}

type CipherTakeResponse struct {
	Id        string `json:"id"`
	Data      string `json:"data"`
	CreatedAt int64  `json:"created"`
	UpdatedAt int64  `json:"updated"`
}

type CipherListResponse struct {
	Ciphers []string `json:"ciphers"`
}

var (
	CipherTypeAccount = 0
)
