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

// RefOrderItemStatusCode is a reference table
type RefOrderItemStatusCode struct {
	StatusCode  string      `gorm:"column:status_code;primary_key" json:"status_code"`
	Description null.String `gorm:"column:description" json:"description"`
}

// TableName sets the insert table name for this struct type
func (r *RefOrderItemStatusCode) TableName() string {
	return "ref_order_item_status_code"
}
