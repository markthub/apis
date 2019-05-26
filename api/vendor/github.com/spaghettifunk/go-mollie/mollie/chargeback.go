package mollie

import (
	"fmt"
	"time"
)

const chargebackEndpoint = "chargebacks"

// ChargebackResponse defines the objecy for every response from the Mollie APIs regarding chargebacks
// https://docs.mollie.com/reference/v2/chargebacks-api/get-chargeback
type ChargebackResponse struct {
	Resource         string                 `json:"resource"`
	ID               string                 `json:"id"`
	Amount           Amount                 `json:"amount"`
	SettlementAmount Amount                 `json:"settlementAmount"`
	CreatedAt        time.Time              `json:"createdAt"`
	ReversedAt       interface{}            `json:"reversedAt"`
	PaymentID        string                 `json:"paymentId"`
	Links            map[string]interface{} `json:"_links"`
}

// GetChargeBack retrieve the charge back object given both a payment id and the chargeback id
func (c *Client) GetChargeBack(paymentID, chargebackID string) (*ChargebackResponse, error) {

	chargebackURL := fmt.Sprintf("%s/%s/%s/%s", paymentsEndpoint, paymentID, chargebackEndpoint, chargebackID)

	var r ChargebackResponse
	err := c.get(chargebackURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// EmbeddedChargeback is the object used for marshelling correctly the response
// when asking for more than one chargeback at the time
type EmbeddedChargeback struct {
	Chargeback []ChargebackResponse `json:"chargebacks"`
}

// ChargebacksResponse is an object returned when multiple Chargeback objects are requested
type ChargebacksResponse struct {
	Count              int                    `json:"count"`
	EmbeddedChargeback EmbeddedChargeback     `json:"_embedded"`
	Links              map[string]interface{} `json:"_links"`
}

// ListAllChargeBacks returns all the chargebacks in your account
func (c *Client) ListAllChargeBacks() (*ChargebacksResponse, error) {
	chargebackURL := fmt.Sprintf("%s", refundsEndpoint)

	var r ChargebacksResponse
	err := c.get(chargebackURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// ListChargeBacksOfPayment returns all the chargebacks of a specific payment
func (c *Client) ListChargeBacksOfPayment(paymentID string) (*ChargebacksResponse, error) {
	chargebackURL := fmt.Sprintf("%s/%s/%s", paymentsEndpoint, paymentID, refundsEndpoint)

	var r ChargebacksResponse
	err := c.get(chargebackURL, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
