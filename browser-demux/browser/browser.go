package browser

type Browser struct {
	Name       string
	Executable string
}

func GetInstalledBrowsers() []*Browser {
	// TODO: implement
	return []*Browser{{
		Name:       "firefox",
		Executable: "firefox",
	}}
}
