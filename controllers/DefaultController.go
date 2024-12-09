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

	c.HTML(http.StatusOK, "layout.gohtml", gin.H{
		"Title": "Gouef Project",
		"Users": data,
	})
	//c.String(http.StatusOK, "OK")
	//renderer.RenderTemplate(c.Writer, "default", nil)
}

func (controller DefaultController) Ping(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
