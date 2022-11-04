package utils

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/MedzikUser/go-avapi/types"
)

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
