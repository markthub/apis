package mollie

import (
	"fmt"
	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

const baseAddress = "https://api.mollie.com/v2/"

var validate = validator.New()

// Client is a client for working with the Mollie API.
type Client struct {
	sling   *sling.Sling
	apiKey  string
	baseURL string

	TestMode bool
}

// ErrorMollie represents an error returned by the Mollie API.
type ErrorMollie struct {
	Status int         `json:"status"`
	Title  string      `json:"title"`
	Detail string      `json:"detail"`
	Field  string      `json:"field"`
	Links  interface{} `json:"_links"`
}

// Error string reformat
func (e ErrorMollie) Error() string {
	return fmt.Sprintf("mollie: %d %s %s", e.Status, e.Title, e.Detail)
}

// post makes a POST request to the specified Mollie endpoint
func (c *Client) post(url string, body interface{}, result interface{}) error {
	errorMollie := new(ErrorMollie)
	_, err := c.sling.Post(url).BodyJSON(body).Receive(result, errorMollie)
	if err != nil {
		return err
	}
	if errorMollie.Status != 0 {
		return errorMollie
	}
	return nil
}

// get makes a GET request to the specified Mollie endpoint
func (c *Client) get(url string, result interface{}) error {
	errorMollie := new(ErrorMollie)
	_, err := c.sling.Get(url).Receive(result, errorMollie)
	if err != nil {
		return err
	}
	if errorMollie.Status != 0 {
		return errorMollie
	}
	return nil
}

// delete makes a DELETE request to the specified Mollie endpoint
func (c *Client) delete(url string, result interface{}) error {
	errorMollie := new(ErrorMollie)
	_, err := c.sling.Delete(url).Receive(result, errorMollie)
	if err != nil {
		return err
	}
	if errorMollie.Status != 0 {
		return errorMollie
	}
	return nil
}
