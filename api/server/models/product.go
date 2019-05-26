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

// Product is the generic object that represent a single item of the store
type Product struct {
	ID          int         `gorm:"column:id;primary_key" json:"id"`
	StoreID     int         `gorm:"column:store_id;index" json:"store_id" binding:"required"`
	Name        string      `gorm:"column:name" json:"name" binding:"required"`
	Price       string      `gorm:"column:price" json:"price" binding:"required"`
	Description null.String `gorm:"column:description" json:"description"`
	Category    string      `gorm:"column:category" json:"category" binding:"required" validate:"required"`
	Weight      string      `gorm:"column:weight" json:"weight" binding:"required" validate:"required"`
	Frozen      bool        `gorm:"column:frozen" json:"frozen" binding:"required"`
	ImageURL    string      `gorm:"column:image_url" json:"imageUrl" binding:"required" validate:"required"`
	Quantity    uint        `gorm:"column:quantity" json:"quantity"`
	CreatedAt   null.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   null.Time   `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   null.Time   `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (p *Product) TableName() string {
	return "products"
}

// Products is a set of product
type Products struct {
	Data []Product `json:"products"`
}
