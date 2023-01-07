package assets

import _ "embed"

//go:embed bark.ico
var BackIcon []byte

//go:embed config_template.json
var ConfigTemplate []byte
