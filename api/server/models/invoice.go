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

// Invoice is the struct to represent the invoice sent to the customer
type Invoice struct {
	Number     string      `gorm:"column:number;primary_key" json:"number"`
	OrderID    int         `gorm:"column:order_id" json:"order_id" binding:"required"`
	StatusCode string      `gorm:"column:status_code" json:"status_code" binding:"required"`
	Placed     time.Time   `gorm:"column:placed" json:"placed" binding:"required"`
	Details    null.String `gorm:"column:details" json:"details"`
	CreatedAt  null.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  null.Time   `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  null.Time   `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (i *Invoice) TableName() string {
	return "invoices"
}
