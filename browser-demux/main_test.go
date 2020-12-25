package main

import (
	"fmt"
	"github.com/andeyfedoseev/browser-demux/browser"
	"github.com/andeyfedoseev/browser-demux/config"
	"path"
	"testing"
)

type dummyLauncher struct {
	launchedBrowser *browser.Browser
}

func (l *dummyLauncher) Launch(b *browser.Browser, _ string) error {
	l.launchedBrowser = b
	return nil
}

type erroneousLauncher struct{}

func (l *erroneousLauncher) Launch(b *browser.Browser, _ string) error {
	return fmt.Errorf("failed to launch browser %s", b.Name)
}

func TestOpenURL(t *testing.T) {
	cfg, err := config.Load(path.Join("testdata", "config.ini"))
	if err != nil {
		t.Error(err)
	}

	tests := []struct {
		url     string
		browser string
	}{
		{"", "firefox"},
		{"http://google.com", "firefox"},
		{"https://google.com/", "firefox"},
		{"https://google.com/mail", "chrome"},
	}

	for _, test := range tests {
		launcher := new(dummyLauncher)
		if err := openURL(launcher, cfg, test.url); err != nil {
			t.Error(err)
		} else if launcher.launchedBrowser.Name != test.browser {
			t.Errorf("Incorrect browser opened for URL %s. Expected: %s, got %s", test.url, test.browser, launcher.launchedBrowser.Name)
		}
	}

	if err := openURL(new(dummyLauncher), new(config.Config), "foo"); err == nil {
		t.Errorf("Opening an URL with empty config should raise an error")
	}

	if err := openURL(new(erroneousLauncher), cfg, "foo"); err == nil {
		t.Errorf("failure in launcher should raise an error")
	}

}
