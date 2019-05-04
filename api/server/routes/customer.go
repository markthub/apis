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
	id := c.Param("id")

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

	if err := readJSON(r, customer); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	if err := DB.Save(customer).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusCreated, customer)
}

// UpdateCustomer will update the customer row
func UpdateCustomer(c *gin.Context) {
	id := ps.ByName("id")

	customer := &model.Customer{}
	if DB.First(customer, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Customer{}
	if err := readJSON(r, updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := dbmeta.Copy(customer, updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := DB.Save(customer).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, customer)
}

// DeleteCustomer will soft-delete a customer
func DeleteCustomer(c *gin.Context) {
	id := ps.ByName("id")
	customer := &model.Customer{}

	if DB.First(customer, id).Error != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	if err := DB.Delete(customer).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, fmt.Sprintf("customer with id %s deleted", id))
}
