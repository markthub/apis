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

// GetAllStores returns all the stores in the database
func GetAllStores(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)

	pageParam := c.DefaultQuery("page", "0")
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil || page < 1 {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	offset := (page - 1) * pagesize

	store := []*model.Store{}
	if err = db.Model(&model.Store{}).Offset(offset).Limit(pagesize).Find(&store).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, store)
}

// GetStore returns a single store
func GetStore(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	store := &model.Store{}
	if err := db.First(store, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	utils.Response(c, http.StatusOK, store)
}

// AddStore add a new store
func AddStore(c *gin.Context) {
	store := &model.Store{}
	db := c.MustGet("DB").(*gorm.DB)

	if err := c.ShouldBindJSON(store); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	if err := db.Save(store).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, store)
}

// UpdateStore updates a store
func UpdateStore(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	store := &model.Store{}
	if err := db.First(store, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}

	updated := &model.Store{}
	if err := c.ShouldBindJSON(updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := dbmeta.Copy(store, updated); err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}

	if err := db.Save(store).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, store)
}

// DeleteStore deletes a store
func DeleteStore(c *gin.Context) {
	id := c.Param("id")
	db := c.MustGet("DB").(*gorm.DB)

	store := &model.Store{}
	if err := db.First(store, id).Error; err != nil {
		utils.ResponseError(c, http.StatusNotFound, err)
		return
	}
	if err := db.Delete(store).Error; err != nil {
		utils.ResponseError(c, http.StatusInternalServerError, err)
		return
	}
	utils.Response(c, http.StatusOK, fmt.Sprintf("store with id %s is deleted", id))
}
