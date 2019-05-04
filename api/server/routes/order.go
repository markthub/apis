package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/markthub/apis/api/server/models"
	"github.com/smallnest/gen/dbmeta"
)

func GetAllOrders(c *gin.Context) {
	orders := []model.Order{}
	DB.Find(&orders)
	writeJSON(w, &orders)

	page, err := readInt(r, "page", 1)
	if err != nil || page < 1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	offset := (page - 1) * pagesize

	order := r.FormValue("order")

	orders := []*model.Order{}

	if order != "" {
		err = DB.Model(&model.Order{}).Order(order).Offset(offset).Limit(pagesize).Find(&orders).Error
	} else {
		err = DB.Model(&model.Order{}).Offset(offset).Limit(pagesize).Find(&orders).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetOrder(c *gin.Context) {
	id := ps.ByName("id")
	order := &model.Order{}
	if DB.First(order, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, order)
}

func AddOrder(c *gin.Context) {
	order := &model.Order{}

	if err := readJSON(r, order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(order).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, order)
}

func UpdateOrder(c *gin.Context) {
	id := ps.ByName("id")

	order := &model.Order{}
	if DB.First(order, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Order{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(order, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(order).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, order)
}

func DeleteOrder(c *gin.Context) {
	id := ps.ByName("id")
	order := &model.Order{}

	if DB.First(order, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(order).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
