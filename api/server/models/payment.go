package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

// Payment is the struct that contains the payment details
type Payment struct {
	ID            int       `gorm:"column:id;primary_key" json:"id"`
	InvoiceNumber string    `gorm:"column:invoice_number" json:"invoice_number"`
	Placed        time.Time `gorm:"column:placed" json:"placed"`
	Amount        float64   `gorm:"column:amount" json:"amount"`
}

// TableName sets the insert table name for this struct type
func (p *Payment) TableName() string {
	return "payments"
}
