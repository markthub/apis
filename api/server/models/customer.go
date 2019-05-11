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

// Customer represents the customer that purchase the grocery
type Customer struct {
	ID        int       `gorm:"column:id;primary_key" json:"id"`
	UserID    int       `gorm:"column:user_id" json:"user_id" binding:"required"`
	Name      string    `gorm:"column:name" json:"name" binding:"required"`
	Lastname  string    `gorm:"column:lastname" json:"lastname" binding:"required"`
	Email     string    `gorm:"column:email" json:"email" binding:"required"`
	Phone     string    `gorm:"column:phone" json:"phone" binding:"required"`
	Address   string    `gorm:"column:address" json:"address" binding:"required"`
	Zipcode   string    `gorm:"column:zipcode" json:"zipcode" binding:"required"`
	City      string    `gorm:"column:city" json:"city" binding:"required"`
	Country   string    `gorm:"column:country" json:"country" binding:"required"`
	CreatedAt null.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt null.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt null.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (c *Customer) TableName() string {
	return "customers"
}
