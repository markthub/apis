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

// Order is the struct that creates the order from the customer
type Order struct {
	ID           int         `gorm:"column:id;primary_key" json:"id"`
	CustomerID   int         `gorm:"column:customer_id" json:"customer_id"`
	StatusCode   string      `gorm:"column:status_code" json:"status_code"`
	DatePlaced   time.Time   `gorm:"column:date_placed" json:"date_placed"`
	OrderDetails null.String `gorm:"column:order_details" json:"order_details"`
	CreatedAt    null.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    null.Time   `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    null.Time   `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (o *Order) TableName() string {
	return "orders"
}
