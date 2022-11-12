package micropass_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MedzikUser/go-micropass-api/utils"
)

type Client struct {
	Http *http.Client
}

func NewClient() *Client {
	return &Client{
		Http: new(http.Client),
	}
}

func (c *Client) Get(path string, accessToken *string, v any) (*http.Response, error) {
	url := ApiUrl + path
	// make a request instance
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// send the request and parse body
	res, err := c.sendRequest(req, accessToken, v)

	return res, err
}

func (c *Client) Delete(path string, accessToken *string, v any) (*http.Response, error) {
	url := ApiUrl + path
	// make a request instance
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	// send the request and parse body
	res, err := c.sendRequest(req, accessToken, v)

	return res, err
}

func (c *Client) Post(path string, accessToken *string, body any, v any) (*http.Response, error) {
	dataBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// convert data bytes slice to io.Reader
	bodyReader := bytes.NewReader(dataBytes)

	url := ApiUrl + path
	// make a request instance
	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	// send the request and parse body
	res, err := c.sendRequest(req, accessToken, v)

	return res, err
}

func (c *Client) Patch(path string, accessToken *string, body any, v any) (*http.Response, error) {
	dataBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// convert data bytes slice to io.Reader
	bodyReader := bytes.NewReader(dataBytes)

	url := ApiUrl + path
	// make a request instance
	req, err := http.NewRequest("PATCH", url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	// send the request and parse body
	res, err := c.sendRequest(req, accessToken, v)

	return res, err
}

func (c *Client) sendRequest(req *http.Request, accessToken *string, v any) (*http.Response, error) {
	// add bearer authorization to the request if access token isn't nil
	if accessToken != nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *accessToken))
	}

	// send the request
	res, err := c.Http.Do(req)
	if err != nil {
		return nil, err
	}

	// check for API errors
	if res.StatusCode != 200 {
		return nil, utils.HttpError(res.Body)
	}

	// parse response body
	if v != nil {
		err = utils.HttpResponse(res.Body, v)
		if err != nil {
			return nil, err
		}
	}

	// return the http response
	return res, nil
}
