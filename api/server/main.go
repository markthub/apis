package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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

	// Customer endpoints
	a.POST("/customers", routes.AddCustomer)
	a.GET("/customers/:id", routes.GetCustomer)
	a.PUT("/customers/:id", routes.UpdateCustomer)
	a.DELETE("/customers/:id", routes.DeleteCustomer)

	// User endpoints
	a.GET("/users", routes.GetAllUsers)
	a.POST("/users", routes.AddUser)
	a.GET("/users/:id", routes.GetUser)
	a.PUT("/users/:id", routes.UpdateUser)
	a.DELETE("/users/:id", routes.DeleteUser)

	// Store endpoints
	s := a.Group("/stores")
	s.GET("/", routes.GetAllStores)
	s.POST("/", routes.AddStore)
	s.GET("/:id", routes.GetStore)
	s.PUT("/:id", routes.UpdateStore)
	s.DELETE("/:id", routes.DeleteStore)

	s.GET("/products", routes.GetAllProducts)
	s.POST("/products", routes.AddProduct)
	s.GET("/products/:id", routes.GetProduct)
	s.PUT("/products/:id", routes.UpdateProduct)
	s.DELETE("/products/:id", routes.DeleteProduct)

	// Order endpoints
	o := a.Group("/orders")
	o.GET("/", routes.GetAllOrders)
	o.POST("/", routes.AddOrder)
	o.GET("/:id", routes.GetOrder)
	o.PUT("/:id", routes.UpdateOrder)
	o.DELETE("/:id", routes.DeleteOrder)

	o.GET("/items", routes.GetAllOrderItems)
	o.POST("/items", routes.AddOrderItem)
	o.GET("/items/:id", routes.GetOrderItem)
	o.PUT("/items/:id", routes.UpdateOrderItem)
	o.DELETE("/items/:id", routes.DeleteOrderItem)

	// Shipment endpoints
	sh := o.Group("/shipments")

	sh.GET("/", routes.GetAllShipments)
	sh.POST("/", routes.AddShipment)
	sh.GET("/:id", routes.GetShipment)
	sh.PUT("/:id", routes.UpdateShipment)
	sh.DELETE("/:id", routes.DeleteShipment)

	sh.GET("/items", routes.GetAllShipmentItems)
	sh.POST("/items", routes.AddShipmentItem)
	sh.GET("/items/:id", routes.GetShipmentItem)
	sh.PUT("/items/:id", routes.UpdateShipmentItem)
	sh.DELETE("/items/:id", routes.DeleteShipmentItem)

	// Invoice endpoints
	i := o.Group("/invoices")

	i.GET("/", routes.GetAllInvoices)
	i.POST("/", routes.AddInvoice)
	i.GET("/:id", routes.GetInvoice)
	i.PUT("/:id", routes.UpdateInvoice)
	i.DELETE("/:id", routes.DeleteInvoice)

	i.GET("/payments", routes.GetAllPayments)
	i.POST("/payments", routes.AddPayment)
	i.GET("/payments/:id", routes.GetPayment)
	i.PUT("/payments/:id", routes.UpdatePayment)
	i.DELETE("/payments/:id", routes.DeletePayment)

	return r
}

// Serve will serve the APIs on a specific address
func Serve(addr, dbHost, dbPort, dbUser, dbPassword, dbName string) error {
	r := SetupRouter(dbHost, dbPort, dbUser, dbPassword, dbName)
	return r.Run(addr)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}
