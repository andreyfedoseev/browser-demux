// +build linux

package registration

import (
	"errors"
	"github.com/andeyfedoseev/browser-demux/utils"
	"os/exec"
	"path"
	"strings"
)

func Register() error {
	for _, filePath := range utils.ListDesktopFiles() {
		_, filename := path.Split(filePath)
		if strings.HasSuffix(filename, "browser-demux.desktop") {
			return setDefaultBrowser(filename)
		}
	}
	return errors.New("could not find the .desktop file")
}

func setDefaultBrowser(desktopFile string) error {
	if _, err := exec.LookPath("xdg-settings"); err != nil {
		return errors.New("xdg-settings command not found. make sure that you have xdg-utils package installed")
	}
	for _, args := range [][]string{
		{"set", "default-web-browser", desktopFile},
		{"set", "default-url-scheme-handler", "http", desktopFile},
		{"set", "default-url-scheme-handler", "https", desktopFile},
	} {
		cmd := exec.Command("xdg-settings", args...)
		if err := cmd.Start(); err != nil {
			return err
		}
		if _, err := cmd.Process.Wait(); err != nil {
			return err
		}
	}
	return nil
}
