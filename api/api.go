package api

import (
	"fmt"

	"github.com/go-resty/resty"
)

type Client struct {
	restyClient *resty.Client
}

const (
	BaseUrl string = "https://api.typeform.com/"
)

func DefaultClient(apiKey string) *Client {
	c := resty.DefaultClient
	c.SetHostURL(BaseUrl)
	c.SetHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	return &Client{restyClient: c}
}

// Form resource as retrieved by a client editing the form.
func (c *Client) RetrieveForm(formId string) (*Form, error)  {
	r, err := c.restyClient.R().
		SetPathParams(map[string]string{"form_id": formId}).
		SetResult(Form{}).
		Get("/forms/{form_id}")
	if err != nil {
		return nil, err
	}

	switch r.StatusCode() {
	case 200:
		return r.Result().(*Form), nil
	default:
		return nil, errorResponseFormatter(r)
	}
}

// RestyClient provides access to the underlying Resty client
func (c *Client) RestyClient() (*resty.Client)  {
	return c.restyClient
}

func errorResponseFormatter(resp *resty.Response) error {
	return fmt.Errorf("api returned an error response (%v): %s", resp.StatusCode(), string(resp.Body()))
}
