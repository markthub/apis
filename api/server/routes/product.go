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

// GetAllProducts returns all the products
func GetAllProducts(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	pageParam := c.DefaultQuery("page", "0")
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil || page < 1 {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	offset := (page - 1) * pagesize

	product := []*model.Product{}
	if err = db.Model(&model.Product{}).Offset(offset).Limit(pagesize).Find(&product).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, product)
}

// GetProduct returns a single product
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	product := &model.Product{}
	if err := db.First(product, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	utils.Response(c, http.StatusOK, product)
}

// AddProduct creates a new product
func AddProduct(c *gin.Context) {
	product := &model.Product{}
	db := c.MustGet("DB").(*gorm.DB)

	if err := c.ShouldBindJSON(product); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	if err := db.Save(product).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, product)
}

// UpdateProduct updates a product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	product := &model.Product{}
	if err := db.First(product, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}

	updated := &model.Product{}
	if err := c.ShouldBindJSON(updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := dbmeta.Copy(product, updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := db.Save(product).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, product)
}

// DeleteProduct deletes a product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	product := &model.Product{}
	if err := db.First(product, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	if err := db.Delete(product).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, fmt.Sprintf("product with id %s is deleted", id))
}
