package main

import (
	_ "embed"
	"github.com/getlantern/systray"
)

//go:embed assets/bark.ico
var backIcon []byte

func onReady() {
	systray.SetIcon(backIcon)
	systray.SetTooltip("Bark Tray")

	exitMenuItem := systray.AddMenuItem("Exit", "Exit")

	go func() {
		for {
			select {
			case <-exitMenuItem.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {

}

func main() {
	systray.Run(onReady, onExit)
}
