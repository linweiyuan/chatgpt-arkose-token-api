package main

import (
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

	router.GET("/", api.GenerateArkoseToken)

	router.Run(":8081")
}
