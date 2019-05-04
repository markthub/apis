package api

import (
	"github.com/gin-gonic/gin"
	"github.com/markthub/apis/api/pkg/database"
	"github.com/markthub/apis/api/server/middleware"
	"github.com/markthub/apis/api/server/routes"
)

// SetupRouter defines all the endpoints of the APIs
func SetupRouter(dbHost, dbPort, dbUser, dbPassword, dbName string) *gin.Engine {

	r := gin.Default()
	r.RedirectTrailingSlash = true

	// get db client
	db := database.GetDatabaseClient(dbUser, dbPassword, dbHost, dbPort, dbName)

	// Set Middleware
	r.Use(middleware.DBClient(db))

	// Grouping the Endpoints under the basic path /api
	a := r.Group("/api")

	a.GET("/version", routes.Version)
	a.Static("/docs", "docs/swagger")

	return r
}

// Serve will serve the APIs on a specific address
func Serve(addr, dbHost, dbPort, dbUser, dbPassword, dbName string) error {
	r := SetupRouter(dbHost, dbPort, dbUser, dbPassword, dbName)
	return r.Run(addr)
}
