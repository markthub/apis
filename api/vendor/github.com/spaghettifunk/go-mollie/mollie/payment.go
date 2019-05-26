package mollie

import (
	"fmt"
	"net/url"
	"time"
)

const paymentsEndpoint = "payments"

// Address describes where the payee resides
type Address struct {
	StreetAndNumber string `json:"streetAndNumber,omitempty"`
	PostalCode      string `json:"postalCode,omitempty"`
	City            string `json:"city,omitempty"`
	Region          string `json:"region,omitempty"`
	Country         string `json:"country,omitempty"`
}

// Amount describe the currency and value of the payment
type Amount struct {
	Currency string `json:"currency"`
	Value    string `json:"value"`
}

// PaymentRequest describe the object necessary to create a Payment in Mollie
// https://docs.mollie.com/reference/v2/payments-api/create-payment
type PaymentRequest struct {
	Amount            Amount      `json:"amount" validate:"required"`
	Description       string      `json:"description" validate:"required"`
	RedirectURL       string      `json:"redirectUrl" validate:"required"`
	WebhookURL        string      `json:"webhookUrl" validate:"required"`
	Method            string      `json:"method" validate:"required"`
	Locale            string      `json:"locale,omitempty"`
	Metadata          interface{} `json:"metadata,omitempty"`
	SequenceType      string      `json:"sequenceType,omitempty"`
	CustomerID        string      `json:"customerId,omitempty"`
	MandateID         string      `json:"mandateId,omitempty"`
	BillingEmail      string      `json:"billingEmail,omitempty"`
	DueDate           string      `json:"dueDate,omitempty"`
	BillingAddress    Address     `json:"billingAddress,omitempty"`
	ShippingAddress   Address     `json:"shippingAddress,omitempty"`
	VoucherNumber     string      `json:"voucherNumber,omitempty"`
	VoucherPin        string      `json:"voucherPin,omitempty"`
	Issuer            string      `json:"issuer,omitempty"`
	CustomerReference string      `json:"customerReference,omitempty"`
}

// PaymentResponse describes the object obtained in the response from every Payment operation
// https://docs.mollie.com/reference/v2/payments-api/get-payment
type PaymentResponse struct {
	Resource         string      `json:"resource"`
	ID               string      `json:"id"`
	Mode             string      `json:"mode"`
	CreatedAt        time.Time   `json:"createdAt"`
	Amount           Amount      `json:"amount"`
	Description      string      `json:"description"`
	Method           interface{} `json:"method"`
	Metadata         interface{} `json:"metadata"`
	Status           string      `json:"status"`
	IsCancelable     bool        `json:"isCancelable"`
	ExpiresAt        time.Time   `json:"expiresAt"`
	Details          interface{} `json:"details"`
	ProfileID        string      `json:"profileId"`
	SettlementAmount Amount      `json:"settlementAmount"`
	SequenceType     string      `json:"sequenceType"`
	RedirectURL      string      `json:"redirectUrl"`
	WebhookURL       string      `json:"webhookUrl"`
	Links            interface{} `json:"_links"`
}

// CreatePayment creates an actual payment object in Mollie. After this call you can see it in the Mollie Dashboard
// https://docs.mollie.com/reference/v2/payments-api/create-payment
func (c *Client) CreatePayment(p *PaymentRequest) (*PaymentResponse, error) {

	if err := validate.Struct(p); err != nil {
		return nil, err
	}

	paymentURL := fmt.Sprintf("%s", paymentsEndpoint)

	var r PaymentResponse
	err := c.post(paymentURL, p, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// PaymentOptions are the options that can be passed as query parameters when getting a Payment(s)
// https://docs.mollie.com/reference/v2/payments-api/get-payment
type PaymentOptions struct {
	IncludeQrCode    string // value: "details.qrCode"
	EmbedRefunds     string // value: "refunds"
	EmbedChargebacks string // value: "chargebacks"
	From             string // value: "from"
	Limit            string // value: "limit"
}

// GetPayment gets a Payment object given the payment id from Mollie
// https://docs.mollie.com/reference/v2/payments-api/get-payment
func (c *Client) GetPayment(id string, options *PaymentOptions) (*PaymentResponse, error) {

	paymentURL := fmt.Sprintf("%s/%s", paymentsEndpoint, id)

	values := url.Values{}
	if options != nil {
		if options.IncludeQrCode != "" {
			values.Set("include", options.IncludeQrCode)
		}
		if options.EmbedRefunds != "" {
			values.Set("embed", options.EmbedRefunds)
		}
		if options.EmbedChargebacks != "" {
			values.Set("embed", options.EmbedChargebacks)
		}
	}
	if query := values.Encode(); query != "" {
		paymentURL += "?" + query
	}

	var r PaymentResponse
	err := c.get(paymentURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// CancelPayment cancel the payment with specific id. Not all the payments can be canceled.
// https://docs.mollie.com/reference/v2/payments-api/cancel-payment
func (c *Client) CancelPayment(id string) (*PaymentResponse, error) {
	paymentURL := fmt.Sprintf("%s/%s", paymentsEndpoint, id)

	var r PaymentResponse
	err := c.get(paymentURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// EmbeddedPayments is the object used for marshelling correctly the response
// when asking for more than one payment at the time
type EmbeddedPayments struct {
	Payments []PaymentResponse `json:"payments"`
}

// PaymentsResponse is an object returned when multiple Payment objects are requested
type PaymentsResponse struct {
	Count            int                    `json:"count"`
	EmbeddedPayments EmbeddedPayments       `json:"_embedded"`
	Links            map[string]interface{} `json:"_links"`
}

// ListPayments retrieve all the payments of your account
func (c *Client) ListPayments(options *PaymentOptions) (*PaymentsResponse, error) {
	paymentURL := fmt.Sprintf("%s", paymentsEndpoint)

	values := url.Values{}
	if options != nil {
		if options.IncludeQrCode != "" {
			values.Set("include", options.IncludeQrCode)
		}
		if options.EmbedRefunds != "" {
			values.Set("embed", options.EmbedRefunds)
		}
		if options.EmbedChargebacks != "" {
			values.Set("embed", options.EmbedChargebacks)
		}
	}
	if query := values.Encode(); query != "" {
		paymentURL += "?" + query
	}

	var r PaymentsResponse
	err := c.get(paymentURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
