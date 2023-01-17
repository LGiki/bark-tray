package config

import (
	"bark-tray-new/assets"
	"encoding/json"
	"io"
	"os"
)

type Device struct {
	Name        string `json:"name"`
	BarkBaseUrl string `json:"barkBaseUrl"`
	Key         string `json:"key"`
	IsDefault   bool   `json:"isDefault"`
}

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
