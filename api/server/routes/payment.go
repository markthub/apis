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

// GetAllPayments returns all the payments
func GetAllPayments(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	pageParam := c.DefaultQuery("page", "0")
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil || page < 1 {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	offset := (page - 1) * pagesize

	payment := []*model.Payment{}
	if err = db.Model(&model.Payment{}).Offset(offset).Limit(pagesize).Find(&payment).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, payment)
}

// GetPayment returns a single payment
func GetPayment(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	payment := &model.Payment{}
	if err := db.First(payment, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	utils.Response(c, http.StatusOK, payment)
}

// AddPayment creates a new payment
func AddPayment(c *gin.Context) {
	payment := &model.Payment{}
	db := c.MustGet("DB").(*gorm.DB)

	if err := c.ShouldBindJSON(payment); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	if err := db.Save(payment).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, payment)
}

// UpdatePayment updates a payment
func UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	payment := &model.Payment{}
	if err := db.First(payment, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}

	updated := &model.Payment{}
	if err := c.ShouldBindJSON(updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := dbmeta.Copy(payment, updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := db.Save(payment).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, payment)
}

// DeletePayment deletes a payment
func DeletePayment(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	payment := &model.Payment{}
	if err := db.First(payment, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	if err := db.Delete(payment).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, fmt.Sprintf("payment with id %s is deleted", id))
}
