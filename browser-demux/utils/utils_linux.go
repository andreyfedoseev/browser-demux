// +build linux

package utils

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"

	"github.com/adrg/xdg"
)

func ListDesktopFiles() []string {
	var paths []string
	for _, appsDir := range xdg.ApplicationDirs {
		files, err := ioutil.ReadDir(appsDir)
		if err != nil {
			continue
		}
		for _, file := range files {
			name := file.Name()
			if strings.HasSuffix(name, ".desktop") {
				paths = append(paths, path.Join(appsDir, name))
			}
		}
	}
	return paths
}

const desktopFile = `#!/usr/bin/env xdg-open
[Desktop Entry]
Version=1.0
Name=Browser Demux
Keywords=Internet;WWW;Browser;Web
Exec=browser-demux %u
Terminal=false
X-MultipleArgs=false
Type=Application
Categories=Network;WebBrowser;
MimeType=text/html;text/xml;application/xhtml+xml;x-scheme-handler/http;x-scheme-handler/https;x-scheme-handler/ftp;
StartupNotify=false
Icon=browser-demux
`

func CreateDesktopFile() (string, error) {
	path := path.Join(xdg.DataHome, "applications", "browser-demux.desktop")
	if err := ioutil.WriteFile(path, []byte(desktopFile), 0644); err != nil {
		return "", err
	}
	fmt.Printf("desktop file created: %v\n", path)
	return "browser-demux.desktop", nil
}
