// +build linux

package registration

import (
	"errors"
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/andeyfedoseev/browser-demux/utils"
)

func Register() error {
	for _, filePath := range utils.ListDesktopFiles() {
		_, filename := path.Split(filePath)
		filenameLower := strings.ToLower(filename)
		if strings.HasSuffix(filenameLower, "browser-demux.desktop") || strings.HasSuffix(filenameLower, "browser_demux.desktop") {
			return setDefaultBrowser(filename)
		}
	}
	if create, err := utils.ConfirmAction("Could not find a desktop file for browser-demux. Do you want to create one? Type 'y' to confirm: "); err != nil {
		return err
	} else if create {
		if desktopFile, err := utils.CreateDesktopFile(); err != nil {
			return err
		} else {
			return setDefaultBrowser(desktopFile)
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
	fmt.Printf("%s is set as the default browser\n", desktopFile)
	return nil
}
