package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/chatgpt-arkose-token-api/api"
	"github.com/linweiyuan/chatgpt-arkose-token-api/browser"
)

func init() {
	gin.ForceConsoleColor()
	gin.SetMode(gin.ReleaseMode)
}

//goland:noinspection GoUnhandledErrorResult
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arkose.html", nil)
	})

	router.GET("/token", api.GetArkoseToken)
	router.GET("/bx", api.GetBX)

	router.Run(fmt.Sprintf(":%d", browser.Port))
}
