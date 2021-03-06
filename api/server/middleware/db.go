package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	// Necessary to set up the dialect of MySQL in gorm
	_ "github.com/go-sql-driver/mysql"
)

// DBClient returns the MySQL client for connecting to the database
func DBClient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
