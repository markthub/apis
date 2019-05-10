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

// ShipmentItem is a reference table
type ShipmentItem struct {
	ShipmentID  int `gorm:"column:shipment_id;primary_key" json:"shipment_id"`
	OrderItemID int `gorm:"column:order_item_id" json:"order_item_id"`
}

// TableName sets the insert table name for this struct type
func (s *ShipmentItem) TableName() string {
	return "shipment_items"
}
