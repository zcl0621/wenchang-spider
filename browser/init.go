package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

var RUNMODE = "release"

var LiveBrowser *rod.Browser

func InitLiveBrowser() {
	if RUNMODE != "debug" {
		if path, exists := launcher.LookPath(); exists {
			u := launcher.New().
				Delete("use-mock-keychain").
				Set("no-sandbox").
				Bin(path).
				MustLaunch()
			LiveBrowser = rod.New().ControlURL(u).MustConnect()
		} else {
			panic("boardLive: cannot find browser")
		}
	} else {
		u := launcher.New().
			Headless(false).
			Delete("use-mock-keychain").
			Bin("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome").
			MustLaunch()
		LiveBrowser = rod.New().ControlURL(u).MustConnect()
	}
}
