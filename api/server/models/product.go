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

// Product is the generic object that represent a single item of the store
type Product struct {
	ID          int         `gorm:"column:id;primary_key" json:"id"`
	StoreID     int         `gorm:"column:store_id" json:"store_id"`
	Name        string      `gorm:"column:name" json:"name"`
	Price       string      `gorm:"column:price" json:"price"`
	Description null.String `gorm:"column:description" json:"description"`
	CreatedAt   null.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   null.Time   `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   null.Time   `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (p *Product) TableName() string {
	return "products"
}
