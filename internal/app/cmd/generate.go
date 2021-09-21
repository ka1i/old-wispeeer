package cmd

import (
	"path"

	"github.com/ka1i/wispeeer/internal/pkg/tools"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	logeer "github.com/ka1i/wispeeer/pkg/log"
)

func (c *CMD) Generate() error {
	var err error
	logeer.Task("generate").Infof("Location : %v", utils.GetWorkspace())

	staticAssets := path.Join(utils.GetWorkspace(), c.ThemeStr, c.Options.Theme, "static")
	if utils.IsExist(staticAssets) {
		logeer.Task("generate").Info("copy static assets")
		err = tools.FileCopy(staticAssets, c.Options.PublicDir)
		if err != nil {
			return err
		}
	}

	logeer.Task("generate").Infof("public in: %v", c.Options.PublicDir)
	return nil
}
