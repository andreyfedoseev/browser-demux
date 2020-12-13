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

var cfg *config.Config

func main() {
	args, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if opts.Register {
		// TODO
	} else if opts.Create {
		if err := config.CreateBlank(); err != nil {
			log.Fatal(err)
		}
	} else {
		cfg, err = config.Load()
		if err != nil {
			log.Fatal(err)
		}
		if len(args) > 0 {
			url := args[0]
			b, err := getBrowserForURL(cfg, url)
			if err != nil {
				log.Fatal(err)
			}
			if err := b.Launch(url); err != nil {
				log.Fatal(err)
			}
		} else {
			// TODO: launch config
		}
	}

}

func getBrowserForURL(cfg *config.Config, url string) (*Browser, error) {
	for _, p := range cfg.Patterns {
		if p.Matches(url) {
			return p.Browser, nil
		}
	}
	return cfg.GetDefaultBrowser()
}
