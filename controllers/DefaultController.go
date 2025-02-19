package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct {
}

func (controller DefaultController) Index(c *gin.Context) {
	data := struct {
		Users []string
	}{
		Users: []string{"Alice", "Bob", "Charlie"},
	}

	c.HTML(http.StatusOK, "default.gohtml", gin.H{
		"Title": "Gouef Project",
		"H1":    "Homepage",
		"Users": data,
	})
	//c.String(http.StatusOK, "OK")
	//renderer.RenderTemplate(c.Writer, "default", nil)
}

func (controller DefaultController) Ping(c *gin.Context) {
	data := struct {
		Users []string
	}{
		Users: []string{"Ping", "Ping"},
	}
	c.HTML(http.StatusOK, "ping.gohtml", gin.H{
		"Title": "Gouef Project",
		"H1":    "Ping",
		"Users": data,
	})
}
