package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type App struct {
}

type PingParam struct {
}

func (a App) Ping(c *gin.Context, p *PingParam) {
	c.String(http.StatusOK, "OK", p)
}

func (a App) Homepage(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
