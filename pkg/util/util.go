package util

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
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

// IsValidHttpUrl validates if the specified raw URL is a http URL.
func IsValidHttpUrl(rawUrl string) bool {
	if strings.HasPrefix(rawUrl, "http://") || strings.HasPrefix(rawUrl, "https://") {
		u, err := url.Parse(rawUrl)
		return err == nil && u.Host != ""
	}
	return false
}

// StripQueryParamFromUrl strips query parameters from specified URL.
func StripQueryParamFromUrl(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	u.RawQuery = ""
	u.ForceQuery = false
	return u.String(), nil
}

// GetExecutablePath returns the directory of the currently running file.
func GetExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	exPath := filepath.Dir(ex)
	return exPath, nil
}

// ToAbsolutePath determines if filePath is an absolute path.
// If it is an absolute path, it returns the original filePath directly.
// If it is a relative path, it will be concatenated with executablePath and returned.
func ToAbsolutePath(filePath, executablePath string) string {
	if filepath.IsAbs(filePath) {
		return filePath
	}
	return filepath.Join(executablePath, filePath)
}

// ExtractUrlFromText extracts the first URL from the specified text.
func ExtractUrlFromText(text string) string {
	re := regexp.MustCompile(`(http|https)://\S+`)
	extractedUrl := re.FindString(text)
	if extractedUrl != "" {
		return extractedUrl
	}
	return ""
}
