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

// GetAllOrderItems returns all the order_items objects
func GetAllOrderItems(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	pageParam := c.DefaultQuery("page", "0")
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil || page < 1 {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	offset := (page - 1) * pagesize

	orderitems := []*model.OrderItem{}
	if err = db.Model(&model.OrderItem{}).Offset(offset).Limit(pagesize).Find(&orderitems).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, orderitems)
}

// GetOrderItem return a single order_item
func GetOrderItem(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	orderitem := &model.OrderItem{}
	if err := db.First(orderitem, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	utils.Response(c, http.StatusOK, orderitem)
}

// AddOrderItem creates a new order_item
func AddOrderItem(c *gin.Context) {
	orderitem := &model.OrderItem{}
	db := c.MustGet("DB").(*gorm.DB)

	if err := c.ShouldBindJSON(orderitem); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	if err := db.Save(orderitem).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, orderitem)
}

// UpdateOrderItem updates an order_item
func UpdateOrderItem(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	orderitem := &model.OrderItem{}
	if err := db.First(orderitem, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}

	updated := &model.OrderItem{}
	if err := c.ShouldBindJSON(updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := dbmeta.Copy(orderitem, updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := db.Save(orderitem).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, orderitem)
}

// DeleteOrderItem deletes an order_item
func DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	orderitem := &model.OrderItem{}
	if err := db.First(orderitem, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	if err := db.Delete(orderitem).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, fmt.Sprintf("order_item with id %s is deleted", id))
}
