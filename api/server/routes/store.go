package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/markthub/apis/api/server/models"
	"github.com/smallnest/gen/dbmeta"
)

func GetAllStores(c *gin.Context) {
	stores := []model.Store{}
	DB.Find(&stores)
	writeJSON(w, &stores)

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

	stores := []*model.Store{}

	if order != "" {
		err = DB.Model(&model.Store{}).Order(order).Offset(offset).Limit(pagesize).Find(&stores).Error
	} else {
		err = DB.Model(&model.Store{}).Offset(offset).Limit(pagesize).Find(&stores).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetStore(c *gin.Context) {
	id := ps.ByName("id")
	store := &model.Store{}
	if DB.First(store, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, store)
}

func AddStore(c *gin.Context) {
	store := &model.Store{}

	if err := readJSON(r, store); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(store).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, store)
}

func UpdateStore(c *gin.Context) {
	id := ps.ByName("id")

	store := &model.Store{}
	if DB.First(store, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Store{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(store, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(store).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, store)
}

func DeleteStore(c *gin.Context) {
	id := ps.ByName("id")
	store := &model.Store{}

	if DB.First(store, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(store).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
