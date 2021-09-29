package cmd

import (
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/config"
	loger "github.com/ka1i/wispeeer/pkg/log"
)

type CMD struct {
	Options               config.Options
	IndexMarkdownTitleStr string
	IndexStr              string
	ThemeStr              string
	StaticStr             string
	LayoutStr             string
}

func Run() *CMD {
	loger.Task("wispeeer").Infof("Location : %v", utils.GetWorkspace())
	return &CMD{
		Options:               config.Configure.Options,
		IndexMarkdownTitleStr: "index.md",
		IndexStr:              "index",
		ThemeStr:              "themes",
		StaticStr:             "static",
		LayoutStr:             "layouts",
	}
}
