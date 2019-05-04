package middleware

import (
	"github.com/gin-gonic/gin"
)

// DBClient returns the MySQL client for connecting to the database
func DBClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", nil)
		c.Next()
	}
}
