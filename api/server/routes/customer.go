package routes

import (
	"fmt"
	"net/http"

	"github.com/markthub/apis/api/server/utils"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	model "github.com/markthub/apis/api/server/models"
	"github.com/smallnest/gen/dbmeta"
)

// GetCustomer will get a single customer
func GetCustomer(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)
	id := c.Param("customer_id")

	customer := &model.Customer{}
	if err := db.First(customer, id).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, customer)
}

// AddCustomer will add a new customer
func AddCustomer(c *gin.Context) {
	customer := &model.Customer{}
	db := c.MustGet("DB").(*gorm.DB)

	if err := c.ShouldBindJSON(customer); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	if err := db.Save(customer).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusCreated, customer)
}

// UpdateCustomer will update the customer row
func UpdateCustomer(c *gin.Context) {
	id := c.Param("customer_id")
	db := c.MustGet("DB").(*gorm.DB)

	customer := &model.Customer{}
	if err := db.First(customer, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}

	updated := &model.Customer{}
	if err := c.ShouldBindJSON(updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := dbmeta.Copy(customer, updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := db.Save(customer).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, customer)
}

// DeleteCustomer will soft-delete a customer
func DeleteCustomer(c *gin.Context) {
	id := c.Param("customer_id")
	customer := &model.Customer{}
	db := c.MustGet("DB").(*gorm.DB)

	if err := db.First(customer, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	if err := db.Delete(customer).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, fmt.Sprintf("customer with id %s deleted", id))
}
