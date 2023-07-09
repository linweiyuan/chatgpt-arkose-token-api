package webdriver

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/linweiyuan/go-logger/logger"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

//goland:noinspection SpellCheckingInspection
var (
	webDriver            selenium.WebDriver
	ChatgptArkoseServer  string
	WebSocketDebuggerUrl string
)

//goland:noinspection GoUnhandledErrorResult,SpellCheckingInspection,HttpUrlsUsage
func init() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		<-signalChannel
		if webDriver != nil {
			webDriver.Quit()
		}
		os.Exit(0)
	}()

	chromeArgs := []string{
		"--no-sandbox",
		"--disable-gpu",
		"--disable-dev-shm-usage",
		"--disable-blink-features=AutomationControlled",
		"--incognito",
		"--headless=new",
		"--remote-debugging-port=9222",
	}

	ChatgptArkoseServer = os.Getenv("CHATGPT_ARKOSE_SERVER")
	webDriver, _ = selenium.NewRemote(selenium.Capabilities{
		"chromeOptions": chrome.Capabilities{
			Args:            chromeArgs,
			ExcludeSwitches: []string{"enable-automation"},
		},
	}, fmt.Sprintf("%s:9515", ChatgptArkoseServer))

	// Host header is specified and is not an IP address or localhost.
	ips, _ := net.LookupHost(strings.ReplaceAll(ChatgptArkoseServer, "http://", ""))
	ip := ips[0]
	resp, _ := http.Get(fmt.Sprintf("http://%s:60709/json/version", ip))
	responseMap := make(map[string]string)
	json.NewDecoder(resp.Body).Decode(&responseMap)
	fmt.Println(responseMap)
	WebSocketDebuggerUrl = responseMap["webSocketDebuggerUrl"]
	resp.Body.Close()

	logger.Info("Service arkose-token-api is ready.")
}
