package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"os"
)

const (
	seleniumPath    = "/Users/pundix2022/Downloads/selenium-server-standalone-3.12.0.jar"
	geckoDriverPath = "/Users/pundix2022/Downloads/geckodriver"
	port            = 8080
	connectionBase  = "https://dappradar.com/rankings"
)

func main() {

	// start server instance
	options := []selenium.ServiceOption{
		//selenium.StartFrameBuffer(),
		selenium.GeckoDriver(geckoDriverPath),
		selenium.Output(os.Stderr),
	}
	selenium.SetDebug(true)

	service, err := selenium.NewSeleniumService(seleniumPath, port, options...)
	if err != nil {
		fmt.Printf("Could not start Selenium service, %v", err)
		return
	}

	firefoxBinary := "/Applications/Firefox.app/Contents/MacOS/firefox"
	firefoxOptions := map[string]interface{}{
		"args": []string{
			"--headless", // optional argument to run Firefox in headless mode
		},
		"binary": firefoxBinary,
	}

	defer service.Stop()

	// connect to the webDriver instance running locally
	caps := selenium.Capabilities{
		"browserName":        "firefox",
		"moz:firefoxOptions": firefoxOptions}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://192.168.20.79:%d/wd/hub", port))
	if err != nil {
		fmt.Printf("Could not connect to WebDriver instance, %v", err)
		return
	}

	// Navigate to a URL
	if err := wd.Get(connectionBase); err != nil {
		fmt.Printf("Could not retreive info from %v, %v", connectionBase, err)
		return
	}

	fmt.Printf("syerseresr")

}
