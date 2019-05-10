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

// Shipment is the object that represents the tracking of the order
type Shipment struct {
	ID             int         `gorm:"column:id;primary_key" json:"id"`
	OrderID        int         `gorm:"column:order_id" json:"order_id"`
	InvoiceNumber  string      `gorm:"column:invoice_number" json:"invoice_number"`
	TrackingNumber string      `gorm:"column:tracking_number" json:"tracking_number"`
	Placed         time.Time   `gorm:"column:placed" json:"placed"`
	Details        null.String `gorm:"column:details" json:"details"`
	CreatedAt      null.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      null.Time   `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      null.Time   `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (s *Shipment) TableName() string {
	return "shipments"
}
