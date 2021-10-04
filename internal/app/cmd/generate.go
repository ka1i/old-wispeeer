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
	articles, pages uint64
)

func (c *CMD) Generate() error {
	var err error

	// Init Articles
	c.Wispeeer.Blog.Article = make([]tools.Article, 0)
	c.Wispeeer.Blog.Options = c.Options
	c.Wispeeer.Init()

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
	loger.Task("articles").Infof("Total:%2d (Articles:%d)\n", articles+pages, articles)

	// article pagination
	c.Wispeeer.Blog.Articles = articles
	c.Wispeeer.Blog.Pages = pages
	dst := path.Join(c.Options.PublicDir, c.Options.PaginationDir)
	err = c.Wispeeer.ArticlesPaginationRender(dst)
	if err != nil {
		return err
	}
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
			title := strings.TrimSuffix(f.Name(), suffix)
			if pathLevel == 2 && suffix == ".md" {
				// render markdown
				err = c.render(filefullName, title)
				if err != nil {
					return err
				}
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

func (c *CMD) render(src string, title string) error {
	article, err := tools.ArticleScanner(src, c.Options)
	if err != nil {
		return err
	}
	assetsPath := path.Join(strings.TrimRight(src, ".md"))
	// post & page
	if title != c.IndexStr {
		// render articles
		articles++
		c.Wispeeer.Blog.Article = append(c.Wispeeer.Blog.Article, article)

		dst := path.Join(c.Options.PublicDir, c.Options.Permalink)
		err = c.Wispeeer.ArticleDetailRender(article, path.Join(dst, title+".html"))
		if err != nil {
			return err
		}

		if utils.IsDir(assetsPath) {
			err = tools.DirCopy(assetsPath, path.Join(dst, title))
			if err != nil {
				return err
			}
		}
	} else {
		// render other page
		pages++
		dst := path.Join(c.Options.PublicDir, article.Metadata.Title)
		err = c.Wispeeer.PageDetailRender(article, path.Join(dst, "index.html"))
		if err != nil {
			return err
		}

		if utils.IsDir(assetsPath) {
			err = tools.DirCopy(assetsPath, path.Join(dst, article.Metadata.Title))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
