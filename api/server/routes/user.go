package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/markthub/apis/api/server/models"
	"github.com/smallnest/gen/dbmeta"
)

func GetAllUsers(c *gin.Context) {
	users := []model.User{}
	DB.Find(&users)
	writeJSON(w, &users)

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

	users := []*model.User{}

	if order != "" {
		err = DB.Model(&model.User{}).Order(order).Offset(offset).Limit(pagesize).Find(&users).Error
	} else {
		err = DB.Model(&model.User{}).Offset(offset).Limit(pagesize).Find(&users).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUser(c *gin.Context) {
	id := ps.ByName("id")
	user := &model.User{}
	if DB.First(user, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, user)
}

func AddUser(c *gin.Context) {
	user := &model.User{}

	if err := readJSON(r, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, user)
}

func UpdateUser(c *gin.Context) {
	id := ps.ByName("id")

	user := &model.User{}
	if DB.First(user, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.User{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(user, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, user)
}

func DeleteUser(c *gin.Context) {
	id := ps.ByName("id")
	user := &model.User{}

	if DB.First(user, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
