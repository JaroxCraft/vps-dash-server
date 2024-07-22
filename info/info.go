package info

var info = Info{
	Name:    "vps-dash-server",
	Version: "v0.0.1", // TODO: get Version from env, file or something
}

type Info struct {
	Name    string
	Version string
}

func GetInfo() Info {
	return info
}
