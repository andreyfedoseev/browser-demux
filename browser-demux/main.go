package main

import (
	. "github.com/andeyfedoseev/browser-demux/browser"
	"github.com/andeyfedoseev/browser-demux/config"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

var opts struct {
	Create   bool `long:"create" description:"Create an empty config file"`
	Register bool `long:"register" description:"Register as the default browser"`
}

func main() {
	args, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Register {
		if err := register(); err != nil {
			log.Fatal(err)
		}
	} else if opts.Create {
		if err := createBlankConfig(); err != nil {
			log.Fatal(err)
		}
	} else {
		cfg, err := loadConfig()
		if err != nil {
			log.Fatal(err)
		}
		if len(args) > 0 {
			url := args[0]
			if err := openURL(new(ExecLauncher), cfg, url); err != nil {
				log.Fatal(err)
			}
		} else {
			// TODO: launch config UI
		}
	}
}

func createBlankConfig() error {
	return config.CreateBlank()
}

func register() error {
	// TODO: implement
	return nil
}

func loadConfig() (*config.Config, error) {
	configPath, err := config.GetDefaultConfigPath()
	if err != nil {
		return nil, err
	}
	cfg, err := config.Load(configPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func openURL(launcher Launcher, cfg *config.Config, url string) error {
	b, err := getBrowserForURL(cfg, url)
	if err != nil {
		return err
	}
	if err := launcher.Launch(b, url); err != nil {
		return err
	}
	return nil
}

func getBrowserForURL(cfg *config.Config, url string) (*Browser, error) {
	for _, p := range cfg.Patterns {
		if p.Matches(url) {
			return p.Browser, nil
		}
	}
	return cfg.GetDefaultBrowser()
}
