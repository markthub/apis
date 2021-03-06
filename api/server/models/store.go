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

// Store is the object that represents the store itself
type Store struct {
	ID          int       `gorm:"column:id;primary_key" json:"id"`
	UserID      int       `gorm:"column:user_id" json:"user_id" binding:"required"`
	Name        string    `gorm:"column:name" json:"name" binding:"required"`
	Description string    `gorm:"column:description" json:"description" binding:"required"`
	Address     string    `gorm:"column:address" json:"address" binding:"required"`
	Zipcode     string    `gorm:"column:zipcode" json:"zipcode" binding:"required"`
	City        string    `gorm:"column:city" json:"city" binding:"required"`
	Latitude    float64   `gorm:"column:latitude" json:"latitude"`
	Longitude   float64   `gorm:"column:longitude" json:"longitude"`
	Picture     string    `gorm:"column:picture" json:"picture" binding:"required"`
	CreatedAt   null.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   null.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   null.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName sets the insert table name for this struct type
func (s *Store) TableName() string {
	return "stores"
}
