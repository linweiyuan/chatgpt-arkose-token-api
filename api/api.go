package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/chatgpt-arkose-token-api/browser"
	"github.com/linweiyuan/funcaptcha"
	"github.com/linweiyuan/go-logger/logger"
)

//goland:noinspection GoUnhandledErrorResult
func GetArkoseToken(c *gin.Context) {
	data, _ := json.Marshal(browser.BX)
	token, err := funcaptcha.GetOpenAITokenWithBx(string(data))
	if err != nil {
		logger.Error(err.Error())

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to get arkose token.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

//goland:noinspection GoUnhandledErrorResult
func GetBX(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"bx": browser.BX,
	})
}
