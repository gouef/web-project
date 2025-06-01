package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error404(c *gin.Context) {
	c.HTML(http.StatusNotFound, "error/404.gohtml", gin.H{
		"Title": "Gouef Project",
		"H1":    "Homepage",
	})
}

func Error500(c *gin.Context) {
	c.String(http.StatusInternalServerError, "Internal server error")
}
