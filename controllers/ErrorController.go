package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error404(c *gin.Context) {
	c.String(http.StatusNotFound, "Page not found")
}

func Error500(c *gin.Context) {
	c.String(http.StatusInternalServerError, "Internal server error")
}
