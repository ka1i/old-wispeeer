package cmd

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/ka1i/wispeeer/internal/pkg/tools"
	"github.com/ka1i/wispeeer/internal/pkg/tools/article"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	loger "github.com/ka1i/wispeeer/pkg/log"
)

func (c *CMD) Generate() error {
	var err error

	// clear old public
	tools.FileRemove(c.Options.PublicDir)

	// copt static asset
	staticAssets := path.Join(c.ThemeStr, c.Options.Theme, c.StaticStr)
	if utils.IsExist(staticAssets) {
		loger.Task("generate").Info("copy static assets")
		err = tools.DirCopy(staticAssets, c.Options.PublicDir)
		if err != nil {
			return err
		}
	}

	loger.Task("generate").Infof("public in: %v", c.Options.PublicDir)

	// kids! run
	err = c.processor(c.Options.SourceDir)
	if err != nil {
		return err
	}
	// render markdown
	total, articles, err := article.MarkdownRender(c.Articles)
	if err != nil {
		return err
	}
	loger.Task("articles").Infof("Total:%2d (Articles:%d)\n", total, articles)

	return nil
}

func (c *CMD) processor(startDIR string) error {
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
				oneLevelAsset := path.Join(c.Options.PublicDir, f.Name())
				err = tools.FileCopy(filefullName, oneLevelAsset)
				if err != nil {
					return err
				}
			}

			suffix := path.Ext(f.Name())
			if pathLevel == 2 && suffix == ".md" {
				c.Articles = append(c.Articles, filefullName)
			}
		} else {
			if pathLevel == 2 {
				continue
			}
			err := c.processor(filefullName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
