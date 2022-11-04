package utils

import (
	"encoding/json"
	"io"
)

func HttpResponse(r io.Reader, v any) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)

	return err
}
