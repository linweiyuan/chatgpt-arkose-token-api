package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/chatgpt-arkose-token-api/webdriver"
)

var (
	lock sync.Mutex
)

//goland:noinspection GoUnhandledErrorResult
func GetArkoseToken(c *gin.Context) {
	lock.Lock()
	defer lock.Unlock()

	if !webdriver.IsReady() {
		time.Sleep(time.Second * 3)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"errorMessage": "Failed to get arkose token.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": webdriver.GetArkoseToken(),
	})
}
