package controllers_module

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

	c.HTML(http.StatusOK, "default/content.gohtml", gin.H{
		"Title": "Gouef Project",
		"H1":    "Homepage",
		"Users": data,
	})
	//c.String(http.StatusOK, "OK")
	//renderer.RenderTemplate(c.Writer, "default", nil)
	//renderer.Render()
}

func (controller DefaultController) Ping(c *gin.Context) {
	data := struct {
		Users []string
	}{
		Users: []string{"Ping", "Ping"},
	}
	c.HTML(http.StatusOK, "ping/content.gohtml", gin.H{
		"Title": "Gouef Project",
		"H1":    "Ping",
		"Users": data,
	})
}
