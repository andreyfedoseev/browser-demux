package config

import (
	"errors"
	"fmt"
	"sort"

	"github.com/adrg/xdg"
	. "github.com/andeyfedoseev/browser-demux/browser"
	. "github.com/andeyfedoseev/browser-demux/pattern"
	"gopkg.in/ini.v1"
)

const configFilename = "browser-demux.ini"
const browsersSectionName = "browsers"
const urlsSectionName = "urls"

type Config struct {
	Browsers []*Browser
	Patterns []*Pattern
}

var GetDefaultConfigPath = func() (string, error) {
	return xdg.ConfigFile(configFilename)
}

func Load(configPath string) (*Config, error) {
	iniCfg := ini.Empty()
	if err := iniCfg.Append(configPath); err != nil {
		return nil, err
	}
	browsersSection, err := iniCfg.GetSection(browsersSectionName)
	if err != nil {
		return nil, err
	}
	browsersMap := make(map[string]*Browser)
	c := &Config{}
	for _, key := range browsersSection.Keys() {
		b := &Browser{
			Name:       key.Name(),
			Executable: key.String(),
		}
		c.Browsers = append(c.Browsers, b)
		browsersMap[b.Name] = b
	}
	urlsSection, err := iniCfg.GetSection(urlsSectionName)
	if err != nil {
		return nil, err
	}
	for _, key := range urlsSection.Keys() {
		browserName := key.String()
		b, ok := browsersMap[browserName]
		if !ok {
			continue
		}
		p := &Pattern{
			Pattern: key.Name(),
			Browser: b,
		}
		c.Patterns = append(c.Patterns, p)
	}

	sort.Slice(c.Patterns[:], func(i, j int) bool {
		return len(c.Patterns[i].Pattern) > len(c.Patterns[j].Pattern)
	})

	return c, nil
}

func (c *Config) Save() error {
	cfg := ini.Empty()
	browsersSection, err := cfg.NewSection(browsersSectionName)
	if err != nil {
		return err
	}
	for _, b := range c.Browsers {
		if _, err := browsersSection.NewKey(b.Name, b.Executable); err != nil {
			return err
		}
	}
	urlsSection, err := cfg.NewSection(urlsSectionName)
	if err != nil {
		return err
	}
	for _, p := range c.Patterns {
		if _, err := urlsSection.NewKey(p.Pattern, p.Browser.Name); err != nil {
			return err
		}
	}
	path, err := xdg.ConfigFile(configFilename)
	if err != nil {
		return err
	}
	if err := cfg.SaveTo(path); err != nil {
		return err
	}
	fmt.Printf("Configuration file saved to %s\n", path)
	return nil
}

func CreateBlank() error {
	return (&Config{
		Browsers: GetInstalledBrowsers(),
	}).Save()
}

func (c *Config) GetDefaultBrowser() (*Browser, error) {
	if len(c.Browsers) == 0 {
		return nil, errors.New("could not determine the default browser executable")
	}
	return c.Browsers[0], nil
}
