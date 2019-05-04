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

// GetAllInvoices returns all the invoices
func GetAllInvoices(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	pageParam := c.DefaultQuery("page", "0")
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil || page < 1 {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	offset := (page - 1) * pagesize

	invoices := []*model.Invoice{}
	err = db.Model(&model.Invoice{}).Offset(offset).Limit(pagesize).Find(&invoices).Error

	if err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, invoices)
}

// GetInvoice returns a single invoice given the number
func GetInvoice(c *gin.Context) {
	number := c.Param("number")
	db := c.MustGet("DB").(*gorm.DB)

	invoice := &model.Invoice{}
	if err := db.First(invoice, number).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	utils.Response(c, http.StatusOK, invoice)
}

// AddInvoice creates a new invoice in the database
func AddInvoice(c *gin.Context) {
	invoice := &model.Invoice{}
	db := c.MustGet("DB").(*gorm.DB)

	if err := c.ShouldBindJSON(invoice); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	if err := db.Save(invoice).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, invoice)
}

// UpdateInvoice updates the invoice in the database
func UpdateInvoice(c *gin.Context) {
	number := c.Param("number")
	db := c.MustGet("DB").(*gorm.DB)

	invoice := &model.Invoice{}
	if err := db.First(invoice, number).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}

	updated := &model.Invoice{}
	if err := c.ShouldBindJSON(updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := dbmeta.Copy(invoice, updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := db.Save(invoice).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, invoice)
}

// DeleteInvoice will delete the invoice from the database
func DeleteInvoice(c *gin.Context) {
	number := c.Param("number")
	db := c.MustGet("DB").(*gorm.DB)

	invoice := &model.Invoice{}
	if err := db.First(invoice, number).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	if err := db.Delete(invoice).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, fmt.Sprintf("invoice number %s is deleted", number))
}
