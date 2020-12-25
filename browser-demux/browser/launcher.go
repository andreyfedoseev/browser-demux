package browser

import (
	"errors"
	"os/exec"
	"strings"
)

type Launcher interface {
	Launch(browser *Browser, url string) error
}

type ExecLauncher struct{}

func (launcher *ExecLauncher) Launch(browser *Browser, url string) error {
	command, args, err := buildCommandArgs(browser.Executable, url)
	if err != nil {
		return err
	}
	cmd := exec.Command(command, args...)
	if err := cmd.Start(); err != nil {
		return err
	}
	if _, err := cmd.Process.Wait(); err != nil {
		return err
	}
	return nil
}

func buildCommandArgs(executable string, url string) (string, []string, error) {
	url = strings.TrimSpace(url)
	words := strings.Fields(executable)
	if len(words) == 0 {
		return "", []string{}, errors.New("executable must not be an empty string")
	}
	urlConsumed := false
	for i, word := range words {
		if strings.Contains(word, "{url}") {
			words[i] = strings.Replace(word, "{url}", url, 1)
			urlConsumed = true
			break
		}
	}
	executable = words[0]
	args := words[1:]
	if len(url) > 0 && !urlConsumed {
		args = append(args, url)
	}
	return executable, args, nil
}
