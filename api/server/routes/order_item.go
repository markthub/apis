package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/markthub/apis/api/server/models"
	"github.com/smallnest/gen/dbmeta"
)

func GetAllOrderItems(c *gin.Context) {
	orderitems := []model.OrderItem{}
	DB.Find(&orderitems)
	writeJSON(w, &orderitems)

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

	orderitems := []*model.OrderItem{}

	if order != "" {
		err = DB.Model(&model.OrderItem{}).Order(order).Offset(offset).Limit(pagesize).Find(&orderitems).Error
	} else {
		err = DB.Model(&model.OrderItem{}).Offset(offset).Limit(pagesize).Find(&orderitems).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetOrderItem(c *gin.Context) {
	id := ps.ByName("id")
	orderitem := &model.OrderItem{}
	if DB.First(orderitem, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, orderitem)
}

func AddOrderItem(c *gin.Context) {
	orderitem := &model.OrderItem{}

	if err := readJSON(r, orderitem); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(orderitem).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, orderitem)
}

func UpdateOrderItem(c *gin.Context) {
	id := ps.ByName("id")

	orderitem := &model.OrderItem{}
	if DB.First(orderitem, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.OrderItem{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(orderitem, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(orderitem).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, orderitem)
}

func DeleteOrderItem(c *gin.Context) {
	id := ps.ByName("id")
	orderitem := &model.OrderItem{}

	if DB.First(orderitem, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(orderitem).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
