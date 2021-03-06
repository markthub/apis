package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	model "github.com/markthub/apis/api/server/models"
	"github.com/markthub/apis/api/server/utils"
	"github.com/smallnest/gen/dbmeta"
)

const (
	orderNumberLength = 10
)

// GetAllOrders return all the orders
func GetAllOrders(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	pageParam := c.DefaultQuery("page", "1")
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil || page < 1 {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	offset := (page - 1) * pagesize

	order := []*model.Order{}
	if err = db.Model(&model.Order{}).Offset(offset).Limit(pagesize).Find(&order).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, order)
}

// GetOrder returns a single order
func GetOrder(c *gin.Context) {
	id := c.Param("order_id")
	db := c.MustGet("DB").(*gorm.DB)

	order := &model.Order{}
	if err := db.First(order, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	utils.Response(c, http.StatusOK, order)
}

// NewOrder creates a new order
func NewOrder(c *gin.Context) {
	order := &model.Order{}
	db := c.MustGet("DB").(*gorm.DB)
	// m := c.MustGet("MollieClient").(mollie.Client)

	if err := c.ShouldBindJSON(order); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	// create order
	// orderNumber := utils.GenerateOrderNumber(orderNumberLength)

	// Create Mollie Payment here
	// TODO: REMOVE THIS!!!
	//url := location.Get(c)
	//baseUrl := url.Scheme + url.Host
	// baseURL := os.Getenv("NGROK_URL") + "/checkout"
	// paymentStatus, checkoutURL, err := AuthorizePayment(m, order, orderNumber, baseURL)

	if err := db.Save(order).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, order)
}

// UpdateOrder updates an order
func UpdateOrder(c *gin.Context) {
	id := c.Param("order_id")
	db := c.MustGet("DB").(*gorm.DB)

	order := &model.Order{}
	if err := db.First(order, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}

	updated := &model.Order{}
	if err := c.ShouldBindJSON(updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := dbmeta.Copy(order, updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := db.Save(order).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, order)
}

// DeleteOrder deletes an order
func DeleteOrder(c *gin.Context) {
	id := c.Param("order_id")
	db := c.MustGet("DB").(*gorm.DB)

	order := &model.Order{}
	if err := db.First(order, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	if err := db.Delete(order).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, fmt.Sprintf("order with id %s is deleted", id))
}
