package webdriver

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/linweiyuan/go-logger/logger"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
)

//goland:noinspection SpellCheckingInspection
const (
	checkArkoseTokenTimeout  = 5
	checkArkoseTokenInterval = 1
)

//goland:noinspection SpellCheckingInspection
var (
	driver              selenium.WebDriver
	chatgptArkoseServer string
)

//goland:noinspection GoUnhandledErrorResult,SpellCheckingInspection,HttpUrlsUsage
func init() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		if driver != nil {
			driver.Quit()
		}
		os.Exit(0)
	}()

	chatgptArkoseServer = os.Getenv("CHATGPT_ARKOSE_SERVER")
	caps := selenium.Capabilities{
		"browserName": "firefox",
	}
	firefoxOptions := firefox.Capabilities{
		Args: []string{
			"--headless",
		},
	}
	caps.AddFirefox(firefoxOptions)

	var err error
	driver, err = selenium.NewRemote(caps, fmt.Sprintf("%s:4444", chatgptArkoseServer))
	if err != nil {
		logger.Error("Failed to init webdriver.")
		return
	}

	logger.Info("Service arkose-token-api is ready.")
}

func IsReady() bool {
	if err := driver.Get(fmt.Sprintf("%s:8000/arkose.html", chatgptArkoseServer)); err != nil {
		return false
	}

	err := driver.WaitWithTimeoutAndInterval(func(driver selenium.WebDriver) (bool, error) {
		_, err := driver.FindElement(selenium.ByCSSSelector, ".arkose-35536E1E-65B4-4D96-9D97-6ADB7EFF8147-wrapper")
		if err != nil {
			return false, nil
		}

		return true, nil
	}, time.Second*checkArkoseTokenTimeout, time.Second*checkArkoseTokenInterval)

	if err != nil {
		return false
	}

	time.Sleep(time.Second)
	return true
}

//goland:noinspection GoUnhandledErrorResult
func GetArkoseToken() string {
	element, _ := driver.FindElement(selenium.ByID, "click")
	element.Click()
	return getToken()
}

func getToken() string {
	element, _ := driver.FindElement(selenium.ByID, "token")
	text, err := element.Text()
	if text == "" || err != nil {
		time.Sleep(time.Second)
		return getToken()
	}

	return text
}
