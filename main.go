package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/chatgpt-arkose-token-api/api"
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

	router.Run(":8081")
}
