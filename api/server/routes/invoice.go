package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/markthub/apis/api/server/models"
	"github.com/smallnest/gen/dbmeta"
)

func GetAllInvoices(c *gin.Context) {
	invoices := []model.Invoice{}
	DB.Find(&invoices)
	writeJSON(w, &invoices)

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

	invoices := []*model.Invoice{}

	if order != "" {
		err = DB.Model(&model.Invoice{}).Order(order).Offset(offset).Limit(pagesize).Find(&invoices).Error
	} else {
		err = DB.Model(&model.Invoice{}).Offset(offset).Limit(pagesize).Find(&invoices).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetInvoice(c *gin.Context) {
	id := ps.ByName("id")
	invoice := &model.Invoice{}
	if DB.First(invoice, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, invoice)
}

func AddInvoice(c *gin.Context) {
	invoice := &model.Invoice{}

	if err := readJSON(r, invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(invoice).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, invoice)
}

func UpdateInvoice(c *gin.Context) {
	id := ps.ByName("id")

	invoice := &model.Invoice{}
	if DB.First(invoice, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Invoice{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(invoice, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(invoice).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, invoice)
}

func DeleteInvoice(c *gin.Context) {
	id := ps.ByName("id")
	invoice := &model.Invoice{}

	if DB.First(invoice, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(invoice).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
