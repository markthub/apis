package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spaghettifunk/go-mollie/mollie"
)

// Mollie is a middleware to inject the Mollie client for creating payments
func Mollie(apiKey string, testMode bool) gin.HandlerFunc {

	client := mollie.NewClient(apiKey, testMode)

	return func(c *gin.Context) {
		c.Set("MollieClient", client)
		c.Next()
	}
}
