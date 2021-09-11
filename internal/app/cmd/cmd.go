package cmd

import (
	"github.com/ka1i/wispeeer/pkg/config"
)

type CMD struct {
	Options config.Options
}

func Run() *CMD {
	return &CMD{
		Options: config.Configure.Options,
	}
}
