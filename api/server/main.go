package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/location"
	"github.com/markthub/apis/api/server/models"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	"github.com/markthub/apis/api/pkg/database"
	"github.com/markthub/apis/api/server/middleware"
	"github.com/markthub/apis/api/server/routes"
)

func runMigrations(db *gorm.DB, debug bool) {
	db.LogMode(debug)

	db.AutoMigrate(
		&models.Customer{}, &models.Store{},
		&models.User{}, &models.Invoice{},
		&models.Order{}, &models.OrderItem{},
		&models.Payment{}, &models.ShipmentItem{},
		&models.Product{}, &models.RefInvoiceStatusCode{},
		&models.RefOrderItemStatusCode{}, &models.Shipment{},
	)

	// BUG in AutoMigrate. Forced to run the foreign key manually
	// These lines will lead to an error when starting the APIs but I can safely ignore it
	db.Model(&models.Customer{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Store{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&models.Payment{}).AddForeignKey("invoice_number", "invoices(number)", "CASCADE", "CASCADE")
	db.Model(&models.Invoice{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "CASCADE")
	db.Model(&models.Invoice{}).AddForeignKey("status_code", "ref_invoice_status_codes(status_code)", "CASCADE", "CASCADE")
	db.Model(&models.Shipment{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "CASCADE")
	db.Model(&models.Shipment{}).AddForeignKey("invoice_number", "invoices(number)", "CASCADE", "CASCADE")
	db.Model(&models.ShipmentItem{}).AddForeignKey("shipment_id", "shipments(id)", "CASCADE", "CASCADE")
	db.Model(&models.ShipmentItem{}).AddForeignKey("order_item_id", "order_items(id)", "CASCADE", "CASCADE")
	db.Model(&models.Order{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "CASCADE")
	db.Model(&models.Product{}).AddForeignKey("store_id", "stores(id)", "CASCADE", "CASCADE")
	db.Model(&models.OrderItem{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	db.Model(&models.OrderItem{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "CASCADE")
	db.Model(&models.OrderItem{}).AddForeignKey("status_code", "ref_order_item_status_code(status_code)", "CASCADE", "CASCADE")
}

// SetupRouter defines all the endpoints of the APIs
func SetupRouter(dbHost, dbPort, dbUser, dbPassword, dbName, mollieAPIKey string, mollieTestMode bool) *gin.Engine {

	r := gin.Default()
	r.RedirectTrailingSlash = true

	// get db client
	db := database.GetDatabaseClient(dbUser, dbPassword, dbHost, dbPort, dbName)
	runMigrations(db, true)

	// Set Middleware
	r.Use(middleware.DBClient(db))
	r.Use(location.Default())

	// Grouping the Endpoints under the basic path /api
	a := r.Group("/api")

	a.GET("/version", routes.Version)
	a.Static("/docs", "docs/swagger")

	// Checkout endpoints
	// Checkout Endpoints
	checkout := a.Group("/checkout")
	checkout.Use(middleware.Mollie(mollieAPIKey, mollieTestMode))
	checkout.POST("/webhook", routes.PaymentWebhook)
	checkout.GET("/redirect", routes.PaymentRedirectURL)

	// User endpoints (both clients and stores)
	a.POST("/users", routes.AddUser)
	a.GET("/users/:user_id", routes.GetUser)
	a.PUT("/users/:user_id", routes.UpdateUser)
	a.DELETE("/users/:user_id", routes.DeleteUser)

	// Customer endpoints (clients that order grocery)
	a.POST("/customers", routes.AddCustomer)
	a.GET("/customers/:customer_id", routes.GetCustomer)
	a.PUT("/customers/:customer_id", routes.UpdateCustomer)
	a.DELETE("/customers/:customer_id", routes.DeleteCustomer)

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
	o.Use(middleware.Mollie(mollieAPIKey, mollieTestMode))
	o.GET("/", routes.GetAllOrders)
	o.POST("/", routes.NewOrder) // Guest-Login order
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
func Serve(addr, dbHost, dbPort, dbUser, dbPassword, dbName, mollieAPIKey string, mollieTestMode bool) error {
	r := SetupRouter(dbHost, dbPort, dbUser, dbPassword, dbName, mollieAPIKey, mollieTestMode)
	return r.Run(addr)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}
