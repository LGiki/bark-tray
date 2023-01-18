# Bark Tray

[![GitHub](https://img.shields.io/github/license/LGiki/bark-tray?style=flat-square)](https://github.com/LGiki/bark-tray/blob/master/LICENSE) ![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/LGiki/bark-tray/release.yml?style=flat-square) [![GitHub release (latest by date)](https://img.shields.io/github/v/release/LGiki/bark-tray?style=flat-square)](https://github.com/LGiki/bark-tray/releases)

A tray tool for sending **clipboard** text to iOS devices via [Bark](https://github.com/Finb/Bark).

# Screenshot

https://user-images.githubusercontent.com/20807713/213143098-73b28ce4-4b68-40b5-a42a-80f8411ae5e8.mp4

# Usage

- Download the latest version from [release](https://github.com/LGiki/bark-tray/releases) page and extract it.

- Edit the `config.json` file according to [Configuration file](#configuration-file). A valid configuration file is as follows.

  ```json
  {
    "version": "1.0.0",
    "enableLog": true,
    "logFilePath": "bark-tray.log",
    "userAgent": "Bark Tray/1.0",
    "timeout": 5,
    "devices": [
      {
        "name": "MY_PHONE",
        "barkBaseUrl": "https://api.day.app",
        "key": "REPLACE_WITH_YOUR_DEVICE_KEY",
        "isDefault": true
      }
    ]
  }
  ```
  
- Start the Bark Tray and enjoy it. :-)

# Configuration file

The configuration file of the program is `config.json`, if the file does not exist, the program will create a `config.json` file based on [config_template.json](assets/config_template.json).

The definitions of each item in the configuration file are as follows.

| Field       | Type     | Description                                              |
| ----------- | -------- | -------------------------------------------------------- |
| version     | string   | The Bark Tray version.                                   |
| enableLog   | boolean  | Enable logging or not.                                   |
| logFilePath | string   | Path to the log file.                                    |
| userAgent   | string   | The User Agent used to send requests to the Bark server. |
| timeout     | integer  | Request timeout in seconds.                              |
| devices     | []Device | See [Devices](#Devices).                                 |

## Devices

The `devices` field in the configuration file is an array of `Device` objects, and the `Device` objects are defined as follows.

| Field       | Type    | Description                                                  |
| ----------- | ------- | ------------------------------------------------------------ |
| name        | string  | Device name.                                                 |
| barkBaseUrl | string  | URL of Bark server, e.g. `https://api.day.app`.              |
| key         | string  | Key of the device.<br />Suppose the URL displayed on the Bark App homepage is: `https://api.day.app/abcdefghijklmnopqrstuv/example`, then `abcdefghijklmnopqrstuv` is the key of your device. |
| isDefault   | boolean | Whether the current device is the default device.<br />If there are multiple default devices, the first default device in the devices array will be the default device. |

# Build

This program uses [systray](https://github.com/getlantern/systray), which has some requirements for compiling on different platforms, you can [click here](https://github.com/getlantern/systray#platform-notes) to see the detailed requirements.

# License

MIT License

