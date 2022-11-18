package types

// HttpError is a struct that contains the error message from the API request.
type HttpError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
