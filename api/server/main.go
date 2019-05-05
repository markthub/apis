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
	a.GET("/customers/:customer_id", routes.GetCustomer)
	a.PUT("/customers/:customer_id", routes.UpdateCustomer)
	a.DELETE("/customers/:customer_id", routes.DeleteCustomer)

	// User endpoints
	a.POST("/users", routes.AddUser)
	a.GET("/users/:user_id", routes.GetUser)
	a.PUT("/users/:user_id", routes.UpdateUser)
	a.DELETE("/users/:user_id", routes.DeleteUser)

	// Store endpoints
	s := a.Group("/stores")
	s.GET("/", routes.GetAllStores)
	s.POST("/", routes.AddStore)
	s.GET("/:store_id", routes.GetStore)
	s.PUT("/:store_id", routes.UpdateStore)
	s.DELETE("/:store_id", routes.DeleteStore)

	// Product endpoints
	s.GET("/:store_id/products", routes.GetAllProducts)
	s.POST("/:store_id/products", routes.AddProduct)
	s.GET("/:store_id/products/:product_id", routes.GetProduct)
	s.PUT("/:store_id/products/:product_id", routes.UpdateProduct)
	s.DELETE("/:store_id/products/:product_id", routes.DeleteProduct)

	// Order endpoints
	o := a.Group("/orders")
	o.GET("/", routes.GetAllOrders)
	o.POST("/", routes.AddOrder)
	o.GET("/:order_id", routes.GetOrder)
	o.PUT("/:order_id", routes.UpdateOrder)
	o.DELETE("/:order_id", routes.DeleteOrder)

	// Order items endpoints
	it := o.Group("/:order_id/items")

	it.GET("/", routes.GetAllOrderItems)
	it.POST("/", routes.AddOrderItem)
	it.GET("/:item_id", routes.GetOrderItem)
	it.PUT("/:item_id", routes.UpdateOrderItem)
	it.DELETE("/:item_id", routes.DeleteOrderItem)

	// Shipment endpoints
	sh := o.Group("/:order_id/shipments")

	sh.GET("/", routes.GetAllShipments)
	sh.POST("/", routes.AddShipment)
	sh.GET("/:shipment_id", routes.GetShipment)
	sh.PUT("/:shipment_id", routes.UpdateShipment)
	sh.DELETE("/:shipment_id", routes.DeleteShipment)

	// Invoice endpoints
	i := o.Group("/:order_id/invoices")

	i.GET("/", routes.GetAllInvoices)
	i.POST("/", routes.AddInvoice)
	i.GET("/:invoice_id", routes.GetInvoice)
	i.PUT("/:invoice_id", routes.UpdateInvoice)
	i.DELETE("/:invoice_id", routes.DeleteInvoice)

	i.GET("/:invoice_id/payments", routes.GetAllPayments)
	i.POST("/:invoice_id/payments", routes.AddPayment)
	i.GET("/:invoice_id/payments/:payment_id", routes.GetPayment)
	i.PUT("/:invoice_id/payments/:payment_id", routes.UpdatePayment)
	i.DELETE("/:invoice_id/payments/:payment_id", routes.DeletePayment)

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
