package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// OpenUrl opens the specified URL in the default browser of the user.
func OpenUrl(url string) error {
	var err error = nil
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

// IsFileExists checks if the specified file exists.
func IsFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
