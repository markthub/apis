package models

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type StatusCode string

const (
	Open     StatusCode = "OPEN"
	Pending  StatusCode = "PENDING"
	Paid     StatusCode = "PAID"
	Expired  StatusCode = "EXPIRED"
	Canceled StatusCode = "CANCELED"
	Failed   StatusCode = "FAILED"
)

// Order is the struct that creates the order from the customer
type Order struct {
	ID            int        `gorm:"column:id;primary_key" json:"id"`
	CustomerID    int        `gorm:"column:customer_id" json:"customer_id" binding:"required"`
	StatusCode    StatusCode `gorm:"column:status_code" json:"status_code" binding:"required" sql:"not null;type:ENUM('OPEN','PENDING','PAID','EXPIRED','CANCELED','FAILED')"`
	Products      []Product  `json:"products" validate:"required"`
	Amount        float64    `json:"amount" validate:"required"`
	Tip           float64    `json:"tip" validate:"required"`
	Remarks       string     `json:"remarks" validate:"required"`
	PaymentMethod string     `json:"paymentMethod" validate:"required"`
	Issuer        string     `json:"issuer"`
	Fee           float64    `json:"deliveryFee" validate:"required"`
	Time          time.Time  `json:"deliveryTime" validate:"required"`
	Address       string     `json:"deliveryAddress" validate:"required"`
	Postcode      string     `json:"deliveryPostcode" validate:"required"`
	City          string     `json:"deliveryCity" validate:"required"`
	PhoneNumber   string     `json:"phoneNumber" validate:"required"`
	OrderNumber   string     `gorm:"unique_index"`
	CreatedAt     null.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     null.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     null.Time  `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (o *Order) TableName() string {
	return "orders"
}

// Scan converts the object StatusCode into bytes
func (sc *StatusCode) Scan(value interface{}) error {
	asBytes, ok := value.([]byte)
	if !ok {
		return errors.New("scan source is not []byte")
	}
	*sc = StatusCode(string(asBytes))
	return nil
}

// Value returns the correct string representation of a StatusCode
func (sc StatusCode) Value() (driver.Value, error) {
	return string(sc), nil
}
