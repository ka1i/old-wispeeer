package cmd

import (
	"github.com/ka1i/wispeeer/internal/pkg/tools"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	logeer "github.com/ka1i/wispeeer/pkg/log"
)

func (c *CMD) Generate() error {
	var err error
	logeer.Task("generate").Infof("Location : %v", utils.GetWorkspace())

	logeer.Task("generate").Info("copy static assets")

	err = tools.FileCopy("static", c.Options.PublicDir)
	if err != nil {
		return err
	}

	logeer.Task("generate").Infof("public in: %v", c.Options.PublicDir)
	return nil
}
