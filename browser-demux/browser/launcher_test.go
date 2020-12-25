package browser

import (
	"reflect"
	"testing"
)

func TestBuildCommandLineArguments(t *testing.T) {

	type TestCase struct {
		rawExec string
		url     string
		exec    string
		args    []string
	}

	tests := []TestCase{
		{
			rawExec: "firefox",
			url:     "  ",
			exec:    "firefox",
			args:    []string{},
		},
		{
			rawExec: "firefox",
			url:     "https://google.com",
			exec:    "firefox",
			args:    []string{"https://google.com"},
		},
		{
			rawExec: "firefox --profile foo",
			url:     "https://google.com",
			exec:    "firefox",
			args:    []string{"--profile", "foo", "https://google.com"},
		},
		{
			rawExec: "start microsoft-edge:{url}",
			url:     "https://google.com",
			exec:    "start",
			args:    []string{"microsoft-edge:https://google.com"},
		},
	}

	for _, test := range tests {
		exec, args, err := buildCommandArgs(test.rawExec, test.url)
		if err != nil {
			t.Error(err)
		}
		if exec != test.exec {
			t.Errorf("Incorrect executable, expected %s, got %s", test.exec, exec)
		}
		if !reflect.DeepEqual(args, test.args) {
			t.Errorf("Incorrect args, expected %s, got %s", test.args, args)
		}
	}

	if _, _, err := buildCommandArgs("", "foo"); err == nil {
		t.Error("Error should be returned when executable is blank")
	}

}
