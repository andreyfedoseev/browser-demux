// +build linux

package utils

import (
	"github.com/adrg/xdg"
	"io/ioutil"
	"path"
	"strings"
)

func ListDesktopFiles() []string {
	var paths []string
	for _, dir := range append([]string{xdg.DataHome}, xdg.DataDirs...) {
		dir = path.Join(dir, "applications")
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, file := range files {
			name := file.Name()
			if strings.HasSuffix(name, ".desktop") {
				paths = append(paths, path.Join(dir, name))
			}
		}
	}
	return paths
}
