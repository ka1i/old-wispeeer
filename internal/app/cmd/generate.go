package cmd

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/ka1i/wispeeer/internal/pkg/tools"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	loger "github.com/ka1i/wispeeer/pkg/log"
)

var (
	total     uint64
	wispeeers tools.Wispeeers
)

func (c *CMD) Generate() error {
	var err error

	wispeeers.Options = c.Options
	wispeeers.Pages = append(wispeeers.Pages, "home")

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
	loger.Task("generate").Infof("Article  : %d (Total)\n", total)

	dst := path.Join(c.Options.PublicDir, "index.html")
	tmpl := path.Join(c.ThemeStr, c.Options.Theme, c.LayoutStr, "index.html")
	tools.ListRender(wispeeers, tmpl, dst)

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
					dstPath := path.Join(c.Options.PublicDir, c.Options.Permalink)
					err = c.processor(filefullName, path.Join(dstPath, title+".html"), "post")
					if err != nil {
						return err
					}
					assetRoot := path.Join(startDIR, title)
					if utils.IsDir(assetRoot) {
						dst := path.Join(dstPath, title)
						err = tools.DirCopy(assetRoot, dst)
						if err != nil {
							return err
						}
					}
				} else {
					// process page
					dstPath := path.Join(c.Options.PublicDir, pathLevelSlice[1])
					err := c.processor(filefullName, path.Join(dstPath, title+".html"), "page")
					if err != nil {
						return err
					}
					assetRoot := path.Join(startDIR, c.Options.PageAsset)
					if utils.IsDir(assetRoot) {
						dst := path.Join(dstPath, c.Options.PageAsset)
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

func (c *CMD) processor(src string, dst string, mode string) error {
	article, err := tools.ArticleScanner(src)
	if err != nil {
		return err
	}
	c.detailsCheck(article)

	wispeeer := tools.Wispeeer{
		Article: article,
		Options: c.Options,
	}

	if mode == "post" {
		total++
		tmpl := path.Join(c.ThemeStr, c.Options.Theme, c.LayoutStr, "post.html")
		err = tools.PostRender(wispeeer, tmpl, dst)
		if err != nil {
			return err
		}
		// save article info to mem
		wispeeers.Article = append(wispeeers.Article, article)
	} else if mode == "page" {
		pathSlice := strings.Split(filepath.ToSlash((src)), "/")
		tmpl := path.Join(c.ThemeStr, c.Options.Theme, c.LayoutStr, pathSlice[1]+".html")
		if !utils.IsExist(tmpl) {
			tmpl = path.Join(c.ThemeStr, c.Options.Theme, c.LayoutStr, "page.html")
		}
		err = tools.PageRender(wispeeer, tmpl, dst)
		if err != nil {
			return err
		}
		wispeeers.Pages = append(wispeeers.Pages, pathSlice[1])
	}

	return nil
}

func (c *CMD) detailsCheck(article tools.Article) {
	loger.Task("render").Infof("process %s\n", article.Metadata.Title)
	// fmt.Println("*************************************")
	// fmt.Printf("Title: %s\nPosted: %s\nCategories: %s\nTags: %s\n", article.Metadata.Title,
	// 	article.Metadata.Posted, article.Metadata.Categories, article.Metadata.Tags)
}
