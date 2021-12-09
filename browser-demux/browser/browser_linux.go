// +build linux

package browser

import (
	"path"
	"strings"

	"github.com/andeyfedoseev/browser-demux/utils"
	"gopkg.in/ini.v1"
)

func GetInstalledBrowsers() []*Browser {

	browsers := []*Browser{}
	for _, desktopFile := range utils.ListDesktopFiles() {
		browser := browserFromDesktop(desktopFile)
		if browser != nil && !strings.HasSuffix(browser.Executable, "browser-demux") {
			browsers = append(browsers, browser)
		}
	}
	return browsers
}

func browserFromDesktop(desktopFile string) *Browser {
	parsed, err := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, desktopFile)
	if err != nil {
		return nil
	}
	section, err := parsed.GetSection("Desktop Entry")
	if err != nil {
		return nil
	}
	var displayName string
	var exec string
	var mimeTypes []string
	if nameKey, err := section.GetKey("Name"); err != nil {
		return nil
	} else {
		displayName = nameKey.Value()
	}
	if execKey, err := section.GetKey("Exec"); err != nil {
		return nil
	} else {
		exec = execKey.Value()
	}
	exec = strings.ReplaceAll(exec, "%u", "")
	exec = strings.ReplaceAll(exec, "%U", "")
	exec = strings.TrimSpace(exec)
	if mimeKey, err := section.GetKey("MimeType"); err != nil {
		return nil
	} else {
		mimeTypes = mimeKey.Strings(";")
	}
	isBrowser := false
	for _, mimeType := range mimeTypes {
		if mimeType == "x-scheme-handler/http" {
			isBrowser = true
			break
		}
	}
	if !isBrowser {
		return nil
	}
	_, name := path.Split(exec)
	return &Browser{
		Name:        name,
		DisplayName: displayName,
		Executable:  exec,
	}
}
