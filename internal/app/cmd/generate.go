package cmd

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/ka1i/wispeeer/internal/pkg/tools"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	loger "github.com/ka1i/wispeeer/pkg/log"
)

func (c *CMD) Generate() error {
	var err error
	loger.Task("generate").Infof("Location : %v", utils.GetWorkspace())

	staticAssets := path.Join(utils.GetWorkspace(), c.ThemeStr, c.Options.Theme, "static")
	if utils.IsExist(staticAssets) {
		loger.Task("generate").Info("copy static assets")
		err = tools.FileCopy(staticAssets, c.Options.PublicDir)
		if err != nil {
			return err
		}
	}

	loger.Task("generate").Infof("public in: %v", c.Options.PublicDir)

	// kids! run
	err = c.render(c.Options.SourceDir)
	if err != nil {
		return err
	}
	//loger.Task("generate").Infof("Article  : %d (Total)\n", Total)

	c.detailsCheck()

	return nil
}

func (c *CMD) render(startDIR string) error {
	files, err := ioutil.ReadDir(startDIR)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.Name()[0] == 46 {
			continue
		}
		filefullName := path.Join(startDIR, f.Name())
		pathLevelSlice := strings.Split(filepath.ToSlash((filepath.Dir(filefullName))), "/")
		pathLevel := len(pathLevelSlice)
		if utils.IsFile(filefullName) {
			if pathLevel == 1 {
				fmt.Printf("[COPY] ")
			}
			suffix := path.Ext(f.Name())
			title := strings.TrimSuffix(f.Name(), suffix)

			if pathLevel == 2 && suffix == ".md" {

				if pathLevelSlice[1] == c.Options.PostDir {
					fmt.Printf("[POST] ")
					fmt.Println(pathLevel, "FILE", title)

					adst := path.Join(c.Options.PublicDir, c.Options.Permalink, title+".index")
					c.processor(filefullName, adst)

					assetRoot := path.Join(startDIR, title)
					if utils.IsDir(assetRoot) {
						fmt.Printf("[COPY] ")
						fmt.Println(pathLevel, "DIR", assetRoot)
						dst := path.Join(c.Options.PublicDir, c.Options.Permalink, title)
						err = tools.FileCopy(assetRoot, dst)
						if err != nil {
							return err
						}
					}
				} else {
					fmt.Printf("[PAGE] ")
					fmt.Println(pathLevel, "FILE", filefullName)

					adst := path.Join(c.Options.PublicDir, pathLevelSlice[1], title+".html")
					c.processor(filefullName, adst)

					assetRoot := path.Join(startDIR, c.Options.PageAsset)
					if utils.IsDir(assetRoot) {
						fmt.Printf("[COPY] ")
						fmt.Println(pathLevel, "DIR", assetRoot)
						dst := path.Join(c.Options.PublicDir, pathLevelSlice[1], c.Options.PageAsset)
						err = tools.FileCopy(assetRoot, dst)
						if err != nil {
							return err
						}
					}
				}
			}
		} else {
			if pathLevel == 2 {
				continue
			}
			err := c.render(filefullName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (c *CMD) detailsCheck() {

}

func (c *CMD) processor(src string, dst string) error {
	err := tools.FileCopy(src, dst)
	if err != nil {
		return err
	}
	return nil
}
