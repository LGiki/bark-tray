package main

import (
	_ "embed"
	"fmt"
	"github.com/LGiki/bark-tray/assets"
	"github.com/LGiki/bark-tray/pkg/config"
	"github.com/LGiki/bark-tray/pkg/httpClient"
	"github.com/LGiki/bark-tray/pkg/logger"
	"github.com/LGiki/bark-tray/pkg/util"
	"github.com/emersion/go-autostart"
	"github.com/getlantern/systray"
	"github.com/ncruces/zenity"
	"golang.design/x/clipboard"
	"os"
	"strings"
)

const (
	configFilePath = "config.json"
)

var appConfig *config.Config

func pushMessageFromClipboard(device *config.Device) {
	clipboardTextBytes := clipboard.Read(clipboard.FmtText)
	if clipboardTextBytes == nil {
		logger.Warn(fmt.Sprintf("There is no text content in the clipboard, sending to device '%s' (%s) failed.", device.Name, device.Key))
		_ = zenity.Notify("There is no text content in the clipboard", zenity.InfoIcon)
		return
	}
	clipboardText := strings.TrimSpace(string(clipboardTextBytes))
	logger.Info(fmt.Sprintf("Start sending `%s` to '%s' (%s)", clipboardText, device.Name, device.Key))
	err := device.PushTextMessage(clipboardText)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to send `%s` to '%s' (%s): %s", clipboardText, device.Name, device.Key, err.Error()))
		_ = zenity.Notify(fmt.Sprintf("Failed to send to '%s': %s", device.Name, err.Error()), zenity.ErrorIcon)
		return
	}
	logger.Info(fmt.Sprintf("Successfully sent `%s` to '%s' (%s)", clipboardText, device.Name, device.Key))
}

func addPushMenuItems() {
	if len(appConfig.Devices) == 0 {
		noAnyDeviceMenuItem := systray.AddMenuItem("No device configured", "No device configured")
		noAnyDeviceMenuItem.Disable()
	} else {
		if appConfig.IsDefaultDeviceExist() {
			sendToDefaultDeviceMenuItem := systray.AddMenuItem("Send to default device", "Send to default device")
			go func() {
				for {
					select {
					case <-sendToDefaultDeviceMenuItem.ClickedCh:
						defaultDevice := appConfig.GetDefaultDevice()
						pushMessageFromClipboard(defaultDevice)
					}
				}
			}()
		}

		sendToAllDevicesMenuItem := systray.AddMenuItem("Send to all devices", "Send to all devices")
		go func() {
			for {
				select {
				case <-sendToAllDevicesMenuItem.ClickedCh:
					for i := 0; i < len(appConfig.Devices); i++ {
						device := appConfig.Devices[i]
						pushMessageFromClipboard(device)
					}
				}
			}
		}()

		sendToDeviceMenuItem := systray.AddMenuItem("Send to devices...", "Send to devices...")
		for i := 0; i < len(appConfig.Devices); i++ {
			device := appConfig.Devices[i]
			subMenuItem := sendToDeviceMenuItem.AddSubMenuItem(device.Name, device.Name)
			go func() {
				for range subMenuItem.ClickedCh {
					pushMessageFromClipboard(device)
				}
			}()
		}
	}
}

func addStartOnBootMenuItem() {
	barkTrayPath, err := os.Executable()
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to get bark tray path: %s", err.Error()))
	} else {
		systray.AddSeparator()
		app := &autostart.App{
			Name:        "Bark Tray",
			Exec:        []string{barkTrayPath},
			DisplayName: "A tray tool for sending clipboard text to iOS devices via Bark",
		}
		openOnStartupMenuItem := systray.AddMenuItemCheckbox("Start on boot", "Start Bark Tray on boot", app.IsEnabled())
		go func() {
			for range openOnStartupMenuItem.ClickedCh {
				if app.IsEnabled() {
					// Turn off
					if err := app.Disable(); err != nil {
						logMessage := fmt.Sprintf("Failed to remove Bark Tray from programs that start on boot: %s", err.Error())
						logger.Error(logMessage)
						_ = zenity.Notify(logMessage, zenity.ErrorIcon)
						continue
					}
					openOnStartupMenuItem.Uncheck()
				} else {
					// Turn on
					if err := app.Enable(); err != nil {
						logMessage := fmt.Sprintf("Failed to set Bark Tray to start on boot: %s", err.Error())
						logger.Error(logMessage)
						_ = zenity.Notify(logMessage, zenity.ErrorIcon)
						continue
					}
					openOnStartupMenuItem.Check()
				}
			}
		}()
	}
}

func onReady() {
	systray.SetIcon(assets.BackIcon)
	systray.SetTooltip("Bark Tray")

	addPushMenuItems()
	addStartOnBootMenuItem()

	systray.AddSeparator()
	githubMenuItem := systray.AddMenuItem("Github", "Github")
	systray.AddSeparator()
	exitMenuItem := systray.AddMenuItem("Exit", "Exit")
	go func() {
		for {
			select {
			case <-githubMenuItem.ClickedCh:
				_ = util.OpenUrl("https://github.com/LGiki/bark-tray")
			case <-exitMenuItem.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	if appConfig != nil && appConfig.EnableLog {
		_ = logger.Sync()
	}
}

func main() {
	var err error
	if !util.IsFileExists(configFilePath) {
		err = config.CreateConfigFileTemplate(configFilePath)
		if err != nil {
			_ = zenity.Error(
				"Failed to create config file template: "+err.Error()+"\nPlease check the write permission of the current directory.",
				zenity.Title("Bark Tray"),
				zenity.OKLabel("OK"),
			)
		} else {
			_ = zenity.Info("No config file found.\nA config file template has been created in the current directory.\nPlease fill in the configuration and restart the application.",
				zenity.Title("Bark Tray"),
				zenity.OKLabel("OK"),
			)
		}
		return
	}
	appConfig, err = config.LoadConfig(configFilePath)
	if err != nil {
		_ = zenity.Error(
			"Failed to load config file: "+err.Error()+"\nPlease check the config file and restart the application.",
			zenity.Title("Bark Tray"),
			zenity.OKLabel("OK"),
		)
		return
	}
	if appConfig.EnableLog {
		err = logger.InitLogger(appConfig.LogFilePath)
		if err != nil {
			_ = zenity.Error(
				"Failed to initialize logger: "+err.Error()+"\nPlease check the log file path.",
				zenity.Title("Bark Tray"),
				zenity.OKLabel("OK"),
			)
		}
	}

	err = clipboard.Init()
	if err != nil {
		logger.Error("Failed to initialize clipboard: " + err.Error())
		_ = zenity.Error(
			"Failed to initialize clipboard: "+err.Error()+".",
			zenity.Title("Bark Tray"),
			zenity.OKLabel("OK"),
		)
		return
	}

	httpClient.Setup(appConfig.UserAgent, appConfig.Timeout)
	systray.Run(onReady, onExit)
}
