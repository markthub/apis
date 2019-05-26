package mollie

import (
	"fmt"
	"net/url"
	"time"
)

const refundsEndpoint = "refunds"

// RefundRequest represents the object to create a Refund in Mollie API
type RefundRequest struct {
	Amount      map[string]string `json:"amount" validate:"required"`
	Description string            `json:"description" validate:"required"`
}

// RefundResponse represents a convenient object for every response of the refunds endpoint from Mollie APIs
type RefundResponse struct {
	Resource    string            `json:"resource"`
	ID          string            `json:"id"`
	Amount      map[string]string `json:"amount"`
	Status      string            `json:"status"`
	CreatedAt   time.Time         `json:"createdAt"`
	Description string            `json:"description"`
	PaymentID   string            `json:"paymentId"`
	Links       interface{}       `json:"_links"`
}

// CreateRefund creates a Refund for a given payment
func (c *Client) CreateRefund(r *RefundRequest, paymentID string) (*RefundResponse, error) {
	if err := validate.Struct(r); err != nil {
		return nil, err
	}

	refundURL := fmt.Sprintf("%s/%s/%s", paymentsEndpoint, paymentID, refundsEndpoint)

	var p RefundResponse
	err := c.post(refundURL, r, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// GetRefund returns a specific refund from a given payment and refund ids
func (c *Client) GetRefund(paymentID, refundID string) (*RefundResponse, error) {

	refundURL := fmt.Sprintf("%s/%s/%s/%s", paymentsEndpoint, paymentID, refundsEndpoint, refundID)

	// TODO: Missing the Embed -> don't understand how it works

	var r RefundResponse
	err := c.get(refundURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// CancelRefund cancels a specific refund for a given payment id
func (c *Client) CancelRefund(paymentID, refundID string) (*RefundResponse, error) {
	refundURL := fmt.Sprintf("%s/%s/%s/%s", paymentsEndpoint, paymentID, refundsEndpoint, refundID)

	var r RefundResponse
	err := c.delete(refundURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// RefundOptions is a convenient struct to add query parametes when getting a refund(s)
type RefundOptions struct {
	From  string
	Limit string
}

// EmbeddedRefunds is convenient struct for marshalling/unmarshalling when multiple Refunds
type EmbeddedRefunds struct {
	Refunds []RefundResponse `json:"refunds"`
}

// RefundsResponse is the object when asking for multiple refunds object
type RefundsResponse struct {
	Count           int                    `json:"count"`
	EmbeddedRefunds EmbeddedRefunds        `json:"_embedded"`
	Links           map[string]interface{} `json:"_links"`
}

// ListAllRefunds returns an object with all the reunds in your account
func (c *Client) ListAllRefunds(options *RefundOptions) (*RefundsResponse, error) {
	refundURL := fmt.Sprintf("%s", refundsEndpoint)

	values := url.Values{}
	if options != nil {
		if options.From != "" {
			values.Set("from", options.From)
		}
		if options.Limit != "" {
			values.Set("limit", options.Limit)
		}
	}
	if query := values.Encode(); query != "" {
		refundURL += "?" + query
	}

	var r RefundsResponse
	err := c.get(refundURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// ListRefundsOfPayment returns all the refunds for a given payment
func (c *Client) ListRefundsOfPayment(options *RefundOptions, paymentID string) (*RefundsResponse, error) {
	refundURL := fmt.Sprintf("%s/%s/%s", paymentsEndpoint, paymentID, refundsEndpoint)

	values := url.Values{}
	if options != nil {
		if options.From != "" {
			values.Set("from", options.From)
		}
		if options.Limit != "" {
			values.Set("limit", options.Limit)
		}
	}
	if query := values.Encode(); query != "" {
		refundURL += "?" + query
	}

	var r RefundsResponse
	err := c.get(refundURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
