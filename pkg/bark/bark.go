package bark

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
	// Title is the notification title which font size would be larger than the body. Required.
	Title string `json:"title"`
	// Body is the notification content. Required.
	Body string `json:"body"`
	// Category is a reserved field, no use yet. Required.
	Category string `json:"category"`
	// DeviceKey is the key for each device. Required.
	DeviceKey string `json:"device_key"`
	// Level is the interruption level of the push message. Optional.
	Level PushLevel `json:"level"`
	// Badge is the number displayed next to the app icon. Optional.
	// See <https://developer.apple.com/documentation/usernotifications/unnotificationcontent/1649864-badge>.
	Badge int `json:"badge"`
	// AutomaticallyCopy must be 1. Optional.
	AutomaticallyCopy string `json:"automaticallyCopy"`
	// Copy is the value to be copied. Optional.
	Copy string `json:"copy"`
	// Sound is the sound of the push notification. Optional.
	// See <https://github.com/Finb/Bark/tree/master/Sounds>.
	Sound string `json:"sound"`
	// Icon is an url to the icon, available only on iOS 15 or later. Optional.
	Icon string `json:"icon"`
	// Group is the group of the notification. Optional.
	Group string `json:"group"`
	// IsArchive must be 1. Optional.
	IsArchive string `json:"isArchive "`
	// Url is the url that will jump when click the notification. Optional.
	Url string `json:"url"`
}

// PushResponse is the response struct for Bark API.
type PushResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
