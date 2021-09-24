package cmd

import (
	"github.com/ka1i/wispeeer/pkg/config"
)

type CMD struct {
	Options   config.Options
	IndexStr  string
	ThemeStr  string
	StaticStr string
}

func Run() *CMD {
	return &CMD{
		Options:   config.Configure.Options,
		IndexStr:  "index.md",
		ThemeStr:  "themes",
		StaticStr: "static",
	}
}
