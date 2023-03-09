package config

import (
	"fmt"
	"github.com/LGiki/bark-tray/pkg/bark"
	"github.com/LGiki/bark-tray/pkg/util"
)

type Device struct {
	Name        string `json:"name"`
	BarkBaseUrl string `json:"barkBaseUrl"`
	Key         string `json:"key"`
	IsDefault   bool   `json:"isDefault"`
}

func (d *Device) PushTextMessage(message string) error {
	messageUrl := util.ExtractUrlFromText(message)
	barkPushResponse, err := bark.PushTextMessage(d.BarkBaseUrl, d.Key, message, messageUrl)
	if err != nil {
		return err
	}
	if barkPushResponse.Code != 200 {
		return fmt.Errorf(barkPushResponse.Message)
	}
	return nil
}
