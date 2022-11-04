package types

type HttpError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
