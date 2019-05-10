package models

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

// RefInvoiceStatusCode is a reference table
type RefInvoiceStatusCode struct {
	StatusCode string      `gorm:"column:status_code;primary_key" json:"status_code"`
	Details    null.String `gorm:"column:details" json:"details"`
}

// TableName sets the insert table name for this struct type
func (r *RefInvoiceStatusCode) TableName() string {
	return "ref_invoice_status_codes"
}
