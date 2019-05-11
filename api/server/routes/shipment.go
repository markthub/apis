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

// GetAllShipments returns all the shipments
func GetAllShipments(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	pageParam := c.DefaultQuery("page", "1")
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil || page < 1 {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	offset := (page - 1) * pagesize

	shipment := []*model.Shipment{}
	if err = db.Model(&model.Shipment{}).Offset(offset).Limit(pagesize).Find(&shipment).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, shipment)
}

// GetShipment returns a shipment
func GetShipment(c *gin.Context) {
	id := c.Param("shipment_id")
	db := c.MustGet("DB").(*gorm.DB)

	shipment := &model.Shipment{}
	if err := db.First(shipment, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	utils.Response(c, http.StatusOK, shipment)
}

// AddShipment creates a new shipment
func AddShipment(c *gin.Context) {
	shipment := &model.Shipment{}
	db := c.MustGet("DB").(*gorm.DB)

	if err := c.ShouldBindJSON(shipment); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	if err := db.Save(shipment).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, shipment)
}

// UpdateShipment updates a shipment
func UpdateShipment(c *gin.Context) {
	id := c.Param("shipment_id")
	db := c.MustGet("DB").(*gorm.DB)

	shipment := &model.Shipment{}
	if err := db.First(shipment, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}

	updated := &model.Shipment{}
	if err := c.ShouldBindJSON(updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := dbmeta.Copy(shipment, updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := db.Save(shipment).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, shipment)
}

// DeleteShipment deletes a shipment
func DeleteShipment(c *gin.Context) {
	id := c.Param("shipment_id")
	db := c.MustGet("DB").(*gorm.DB)

	shipment := &model.Shipment{}
	if err := db.First(shipment, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	if err := db.Delete(shipment).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, fmt.Sprintf("shipment with id %s is deleted", id))
}
