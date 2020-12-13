package browser

import "os/exec"

type Browser struct {
	Name       string
	Executable string
}

func (browser *Browser) Launch(url string) error {
	command, args := buildCommandArgs(browser.Executable, url)
	cmd := exec.Command(command, args...)
	if err := cmd.Start(); err != nil {
		return err
	}
	if _, err := cmd.Process.Wait(); err != nil {
		return err
	}
	return nil
}

func buildCommandArgs(executable string, url string) (string, []string) {
	args := []string{url}
	// TODO: handle cases such as `google-chrome --profile=ABC`
	// TODO: handle cases such as `microsoft-edge:{url}`
	return executable, args
}

func GetInstalledBrowsers() []*Browser {
	// TODO: implement
	return []*Browser{{
		Name:       "firefox",
		Executable: "firefox",
	}}
}
