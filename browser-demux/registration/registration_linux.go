// +build linux

package registration

import (
	"errors"
	"github.com/adrg/xdg"
	"github.com/andeyfedoseev/browser-demux/utils"
	"gopkg.in/ini.v1"
	"path"
	"strings"
)

const defaultApplications = "Default Applications"
const schemeHTTP = "x-scheme-handler/http"
const schemeHTTPS = "x-scheme-handler/https"

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
	iniCfg := ini.Empty(ini.LoadOptions{
		IgnoreInlineComment: true,
	})
	cfgPath := path.Join(xdg.ConfigHome, "mimeapps.list")
	if err := iniCfg.Append(cfgPath); err != nil {
		return err
	}
	section := iniCfg.Section(defaultApplications)
	if section == nil {
		if newSection, err := iniCfg.NewSection(defaultApplications); err != nil {
			return err
		} else {
			section = newSection
		}
	}
	for _, scheme := range [2]string{
		schemeHTTP, schemeHTTPS,
	} {
		key := section.Key(scheme)
		if key == nil {
			if _, err := section.NewKey(scheme, desktopFile); err != nil {
				return err
			}
		} else {
			key.SetValue(desktopFile)
		}
	}
	ini.PrettyFormat = false
	return iniCfg.SaveTo(cfgPath)
}
