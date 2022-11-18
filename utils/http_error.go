package utils

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/MedzikUser/go-micropass-api/types"
)

// HttpError parse the error message returned by the API.
func HttpError(r io.Reader) error {
	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	var res types.HttpError
	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	return errors.New(res.ErrorDescription)
}
