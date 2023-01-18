package config

import (
	"encoding/json"
	"fmt"
	"github.com/LGiki/bark-tray/assets"
	"github.com/LGiki/bark-tray/pkg/logger"
	"github.com/LGiki/bark-tray/pkg/util"
	"io"
	"os"
)

type Config struct {
	Version     string    `json:"version"`
	EnableLog   bool      `json:"enableLog"`
	LogFilePath string    `json:"logFilePath"`
	UserAgent   string    `json:"userAgent"`
	Timeout     int       `json:"timeout"`
	Devices     []*Device `json:"devices"`
}

func LoadConfig(configFilePath string) (*Config, error) {
	var config Config
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	configFileBytes, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configFileBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// StripInvalidDevices removes all invalid devices in Config.Devices,
// invalid device means:
// 1. Device key is empty
// 2. The BarkBaseUrl of Device is not a valid url
// 3. Failed to strip query parameters using util.StripQueryParamFromUrl
func (c *Config) StripInvalidDevices() {
	newDevices := make([]*Device, 0, len(c.Devices))
	for i := 0; i < len(c.Devices); i++ {
		device := c.Devices[i]
		if device.Key == "" || !util.IsValidHttpUrl(device.BarkBaseUrl) {
			logger.Warn(fmt.Sprintf("Invalid device: %s", device.Name))
			continue
		}
		baseUrl, err := util.StripQueryParamFromUrl(device.BarkBaseUrl)
		if err != nil {
			logger.Warn(fmt.Sprintf("Invalid device: %s (%s)", device.Name, err.Error()))
			continue
		}
		device.BarkBaseUrl = baseUrl
		newDevices = append(newDevices, device)
	}
	c.Devices = newDevices
}

func CreateConfigFileTemplate(configFilePath string) error {
	return os.WriteFile(configFilePath, assets.ConfigTemplate, 0644)
}

func (c *Config) GetDefaultDevice() *Device {
	for _, device := range c.Devices {
		if device.IsDefault {
			return device
		}
	}
	return nil
}

func (c *Config) IsDefaultDeviceExist() bool {
	return c.GetDefaultDevice() != nil
}
