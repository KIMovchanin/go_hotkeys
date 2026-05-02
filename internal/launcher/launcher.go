package launcher

import (
	"os/exec"
	"strings"
)

func Launch(target string) error {
	if isURL(target) {
		return openURL(target)
	} else {
		return openEXE(target)
	}
}

func isURL(target string) bool {
	return strings.HasPrefix(target, "https://") || strings.HasPrefix(target, "http://")
}

func openURL(url string) error {
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	return cmd.Start()
}

func openEXE(path string) error {
	cmd := exec.Command(path)
	return cmd.Start()
}
