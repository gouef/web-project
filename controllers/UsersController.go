package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gouef/renderer"
	"github.com/gouef/web-project/models"
	"net/http"
	"strconv"
)

type UserController struct {
}

func (u UserController) Homepage(c *gin.Context) {

	c.String(http.StatusOK, "OK")
}

func (u UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUsers()
	renderer.RenderTemplate(w, "users", users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user := models.GetUserByID(id)
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	renderer.RenderTemplate(w, "user", user)
}
