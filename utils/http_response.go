package utils

import (
	"encoding/json"
	"io"
)

// HttpResponse parse the response returned by the API.
func HttpResponse(r io.Reader, v any) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)

	return err
}
