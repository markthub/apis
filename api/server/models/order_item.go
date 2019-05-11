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

// OrderItem represents the object that creates basically an order
type OrderItem struct {
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	ProductID  int       `gorm:"column:product_id" json:"product_id" binding:"required"`
	OrderID    int       `gorm:"column:order_id" json:"order_id" binding:"required"`
	StatusCode string    `gorm:"column:status_code" json:"status_code" binding:"required"`
	Quantity   int       `gorm:"column:quantity" json:"quantity" binding:"required"`
	Price      float64   `gorm:"column:price" json:"price" binding:"required"`
	CreatedAt  null.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  null.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  null.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (o *OrderItem) TableName() string {
	return "order_items"
}
