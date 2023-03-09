package bark

import (
	"bytes"
	"encoding/json"
	"github.com/LGiki/bark-tray/pkg/httpClient"
	"io"
	"net/http"
	"net/url"
)

// PushLevel is the interruption level of the push message.
// See <https://developer.apple.com/documentation/usernotifications/unnotificationinterruptionlevel>.
type PushLevel string

const (
	// PushLevelActive means the system presents the notification immediately,
	// lights up the screen, and can play a sound.
	PushLevelActive PushLevel = "active"
	// PushLevelTimeSensitive means the system presents the notification immediately,
	// lights up the screen, and can play a sound, but won’t break through system notification controls.
	PushLevelTimeSensitive PushLevel = "timeSensitive"
	// PushLevelPassive means the system adds the notification to the notification list
	// without lighting up the screen or playing a sound.
	PushLevelPassive PushLevel = "passive"
)

// PushSound is the sound of the Bark push notification.
// See <https://github.com/Finb/Bark/tree/master/Sounds>.
type PushSound string

const (
	PushSoundAlarm              PushSound = "alarm.caf"
	PushSoundAnticipate         PushSound = "anticipate.caf"
	PushSoundBell               PushSound = "bell.caf"
	PushSoundBirdSong           PushSound = "birdsong.caf"
	PushSoundBloom              PushSound = "bloom.caf"
	PushSoundCalypso            PushSound = "calypso.caf"
	PushSoundChime              PushSound = "chime.caf"
	PushSoundChoo               PushSound = "choo.caf"
	PushSoundDescent            PushSound = "descent.caf"
	PushSoundElectronic         PushSound = "electronic.caf"
	PushSoundFanfare            PushSound = "fanfare.caf"
	PushSoundGlass              PushSound = "glass.caf"
	PushSoundGotoSleep          PushSound = "gotosleep.caf"
	PushSoundHealthNotification PushSound = "healthnotification.caf"
	PushSoundHorn               PushSound = "horn.caf"
	PushSoundLadder             PushSound = "ladder.caf"
	PushSoundMailSent           PushSound = "mailsent.caf"
	PushSoundMinuet             PushSound = "minuet.caf"
	PushSoundMultiwayInvitation PushSound = "multiwayinvitation.caf"
	PushSoundNewMail            PushSound = "newmail.caf"
	PushSoundNewsflash          PushSound = "newsflash.caf"
	PushSoundNoir               PushSound = "noir.caf"
	PushSoundPaymentSuccess     PushSound = "paymentsuccess.caf"
	PushSoundShake              PushSound = "shake.caf"
	PushSoundSherwoodForest     PushSound = "sherwoodforest.caf"
	PushSoundSilence            PushSound = "silence.caf"
	PushSoundSpell              PushSound = "spell.caf"
	PushSoundSuspense           PushSound = "suspense.caf"
	PushSoundTelegraph          PushSound = "telegraph.caf"
	PushSoundTiptoes            PushSound = "tiptoes.caf"
	PushSoundTypewriters        PushSound = "typewriters.caf"
	PushSoundUpdate             PushSound = "update.caf"
)

// PushRequest is the request struct for Bark API.
// See <https://github.com/Finb/bark-server/blob/master/docs/API_V2.md#push>.
type PushRequest struct {
	// Title is the notification title which font size would be larger than the body.
	Title string `json:"title,omitempty"`
	// Body is the notification content. Required.
	Body string `json:"body"`
	// Category is a reserved field, no use yet.
	Category string `json:"category,omitempty"`
	// DeviceKey is the key for each device. Required.
	DeviceKey string `json:"device_key"`
	// Level is the interruption level of the push message. Optional.
	Level PushLevel `json:"level,omitempty"`
	// Badge is the number displayed next to the app icon. Optional.
	// See <https://developer.apple.com/documentation/usernotifications/unnotificationcontent/1649864-badge>.
	Badge int `json:"badge,omitempty"`
	// AutomaticallyCopy must be 1. Optional.
	AutomaticallyCopy string `json:"automaticallyCopy,omitempty"`
	// Copy is the value to be copied. Optional.
	Copy string `json:"copy,omitempty"`
	// Sound is the sound of the push notification. Optional.
	// See <https://github.com/Finb/Bark/tree/master/Sounds>.
	Sound string `json:"sound,omitempty"`
	// Icon is an url to the icon, available only on iOS 15 or later. Optional.
	Icon string `json:"icon,omitempty"`
	// Group is the group of the notification. Optional.
	Group string `json:"group,omitempty"`
	// IsArchive must be 1. Optional.
	IsArchive string `json:"isArchive,omitempty"`
	// Url is the url that will jump when click the notification. Optional.
	Url string `json:"url,omitempty"`
}

// PushResponse is the response struct for Bark API.
type PushResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func GetBarkPushUrl(barkBaseUrl string) (string, error) {
	barkPushUrl, err := url.JoinPath(barkBaseUrl, "/push")
	if err != nil {
		return "", err
	}
	return barkPushUrl, nil
}

func Push(barkBaseUrl string, pushRequest *PushRequest) (*PushResponse, error) {
	barkPushUrl, err := GetBarkPushUrl(barkBaseUrl)
	if err != nil {
		return nil, err
	}
	pushRequestBytes, err := json.Marshal(pushRequest)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", barkPushUrl, bytes.NewReader(pushRequestBytes))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := httpClient.MustGetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var pushResp PushResponse
	err = json.Unmarshal(body, &pushResp)
	if err != nil {
		return nil, err
	}
	return &pushResp, nil
}

func PushTextMessage(barkBaseUrl string, deviceKey string, message string, messageUrl string) (*PushResponse, error) {
	pushRequest := &PushRequest{
		Body:      message,
		DeviceKey: deviceKey,
	}
	if messageUrl != "" {
		pushRequest.Url = messageUrl
	}
	return Push(barkBaseUrl, pushRequest)
}
