package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/markthub/apis/api/server/models"
	"github.com/smallnest/gen/dbmeta"
)

func GetAllPayments(c *gin.Context) {
	payments := []model.Payment{}
	DB.Find(&payments)
	writeJSON(w, &payments)

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

	payments := []*model.Payment{}

	if order != "" {
		err = DB.Model(&model.Payment{}).Order(order).Offset(offset).Limit(pagesize).Find(&payments).Error
	} else {
		err = DB.Model(&model.Payment{}).Offset(offset).Limit(pagesize).Find(&payments).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetPayment(c *gin.Context) {
	id := ps.ByName("id")
	payment := &model.Payment{}
	if DB.First(payment, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, payment)
}

func AddPayment(c *gin.Context) {
	payment := &model.Payment{}

	if err := readJSON(r, payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(payment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, payment)
}

func UpdatePayment(c *gin.Context) {
	id := ps.ByName("id")

	payment := &model.Payment{}
	if DB.First(payment, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Payment{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(payment, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(payment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, payment)
}

func DeletePayment(c *gin.Context) {
	id := ps.ByName("id")
	payment := &model.Payment{}

	if DB.First(payment, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(payment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
