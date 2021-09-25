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

	// clear old public
	tools.FileRemove(c.Options.PublicDir)

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
	err = c.render(c.Options.SourceDir)
	if err != nil {
		return err
	}
	//loger.Task("generate").Infof("Article  : %d (Total)\n", Total)

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
				oneLevelAsset := path.Join(c.Options.PublicDir, f.Name())
				err = tools.FileCopy(filefullName, oneLevelAsset)
				if err != nil {
					return err
				}
			}
			suffix := path.Ext(f.Name())
			title := strings.TrimSuffix(f.Name(), suffix)

			if pathLevel == 2 && suffix == ".md" {

				if pathLevelSlice[1] == c.Options.PostDir {
					// process post
					err = c.processor(filefullName, path.Join(c.Options.Permalink, title+".html"))
					if err != nil {
						return err
					}
					assetRoot := path.Join(startDIR, title)
					if utils.IsDir(assetRoot) {
						dst := path.Join(c.Options.PublicDir, c.Options.Permalink, title)
						err = tools.DirCopy(assetRoot, dst)
						if err != nil {
							return err
						}
					}
				} else {
					fmt.Printf("[PAGE] ")
					fmt.Println(pathLevel, "FILE", filefullName)

					// process page
					err := c.processor(filefullName, path.Join(pathLevelSlice[1], title+".html"))
					if err != nil {
						return err
					}
					assetRoot := path.Join(startDIR, c.Options.PageAsset)
					if utils.IsDir(assetRoot) {
						dst := path.Join(c.Options.PublicDir, pathLevelSlice[1], c.Options.PageAsset)
						err = tools.DirCopy(assetRoot, dst)
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

func (c *CMD) processor(src string, dst string) error {
	dstPath := path.Join(c.Options.PublicDir, dst)
	err := tools.FileCopy(src, dstPath)
	if err != nil {
		return err
	}
	article, err := tools.ArticleScanner(src)
	if err != nil {
		return err
	}
	c.detailsCheck(article)
	return nil
}

func (c *CMD) detailsCheck(article tools.Article) {
	fmt.Println("*************************************")
	fmt.Printf("Title: %s\nData: %v\nCategories: %s\nTags: %s\n", article.Metadata.Title,
		article.Metadata.Posted, article.Metadata.Categories, article.Metadata.Tags)
}
