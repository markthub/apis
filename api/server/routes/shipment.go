package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/markthub/apis/api/server/models"
	"github.com/smallnest/gen/dbmeta"
)

func GetAllShipments(c *gin.Context) {
	shipments := []model.Shipment{}
	DB.Find(&shipments)
	writeJSON(w, &shipments)

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

	shipments := []*model.Shipment{}

	if order != "" {
		err = DB.Model(&model.Shipment{}).Order(order).Offset(offset).Limit(pagesize).Find(&shipments).Error
	} else {
		err = DB.Model(&model.Shipment{}).Offset(offset).Limit(pagesize).Find(&shipments).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetShipment(c *gin.Context) {
	id := ps.ByName("id")
	shipment := &model.Shipment{}
	if DB.First(shipment, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, shipment)
}

func AddShipment(c *gin.Context) {
	shipment := &model.Shipment{}

	if err := readJSON(r, shipment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(shipment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, shipment)
}

func UpdateShipment(c *gin.Context) {
	id := ps.ByName("id")

	shipment := &model.Shipment{}
	if DB.First(shipment, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Shipment{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(shipment, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(shipment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, shipment)
}

func DeleteShipment(c *gin.Context) {
	id := ps.ByName("id")
	shipment := &model.Shipment{}

	if DB.First(shipment, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(shipment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
