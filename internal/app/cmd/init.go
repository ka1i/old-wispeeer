package cmd

import (
	"fmt"

	"github.com/ka1i/wispeeer/internal/pkg/tools"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	assets "github.com/ka1i/wispeeer/pkg/asset"
	loger "github.com/ka1i/wispeeer/pkg/log"
)

//Initialzation ...
func (c *CMD) Initialzation(title string) error {
	var err error
	loger.Task("init").Infof("Location: %s", utils.GetWorkspace())

	if utils.IsExist(title) {
		return fmt.Errorf("%s: File exists", title)
	}

	loger.Task("init").Infof("wispeeer init %s", title)

	loger.Task("init").Info("unpkg embed assets")

	var storage = assets.GetStorage()
	fs := storage.Fs
	root := storage.Root
	err = tools.AssetsUnpkg(&fs, root, root, title)
	if err != nil {
		return err
	}
	return nil
}
