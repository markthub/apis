package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/markthub/apis/api/server/models"
	"github.com/smallnest/gen/dbmeta"
)

func GetAllProducts(c *gin.Context) {
	products := []model.Product{}
	DB.Find(&products)
	writeJSON(w, &products)

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

	products := []*model.Product{}

	if order != "" {
		err = DB.Model(&model.Product{}).Order(order).Offset(offset).Limit(pagesize).Find(&products).Error
	} else {
		err = DB.Model(&model.Product{}).Offset(offset).Limit(pagesize).Find(&products).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetProduct(c *gin.Context) {
	id := ps.ByName("id")
	product := &model.Product{}
	if DB.First(product, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, product)
}

func AddProduct(c *gin.Context) {
	product := &model.Product{}

	if err := readJSON(r, product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, product)
}

func UpdateProduct(c *gin.Context) {
	id := ps.ByName("id")

	product := &model.Product{}
	if DB.First(product, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Product{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(product, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, product)
}

func DeleteProduct(c *gin.Context) {
	id := ps.ByName("id")
	product := &model.Product{}

	if DB.First(product, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
