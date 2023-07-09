package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"github.com/linweiyuan/chatgpt-arkose-token-api/webdriver"
)

//goland:noinspection GoUnhandledErrorResult
func GenerateArkoseToken(c *gin.Context) {
	remoteAllocator, _ := chromedp.NewRemoteAllocator(c, webdriver.WebSocketDebuggerUrl)
	ctx, cancel := chromedp.NewContext(remoteAllocator)
	defer cancel()

	chromedp.Run(ctx, chromedp.ActionFunc(func(ctx context.Context) error {
		chromedp.Navigate(fmt.Sprintf("%s:8000/arkose.html", webdriver.ChatgptArkoseServer)).Do(ctx)
		chromedp.WaitVisible(".arkose-35536E1E-65B4-4D96-9D97-6ADB7EFF8147-wrapper", chromedp.ByQuery)

		chromedp.Sleep(time.Second).Do(ctx) // if no sleep, token will be empty

		chromedp.Click("#click", chromedp.ByID).Do(ctx)

		c.JSON(http.StatusOK, gin.H{
			"token": getToken(ctx),
		})

		return nil
	}))
}

//goland:noinspection GoUnhandledErrorResult
func getToken(ctx context.Context) string {
	chromedp.Sleep(time.Second).Do(ctx)

	token := ""
	chromedp.InnerHTML("#token", &token).Do(ctx)
	if token == "" {
		return getToken(ctx)
	}

	return token
}
