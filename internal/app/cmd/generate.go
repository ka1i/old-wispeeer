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
	wispeeer        tools.Wispeeer
	articleList     []tools.Article
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
	loger.Task("page").Infof("Pages  : %d (Total)\n", pages)
	loger.Task("article").Infof("Articles  : %d (Total)\n", articles)

	dst := path.Join(c.Options.PublicDir, c.Options.PaginationDir)
	tmpl := path.Join(c.ThemeStr, c.Options.Theme, c.LayoutStr, "post.html")
	err = tools.ArticleListRender(articleList, &c.Options, tmpl, dst)
	if err != nil {
		return err
	}

	tmpl = path.Join(c.ThemeStr, c.Options.Theme, c.LayoutStr, "index.html")
	dst = path.Join(c.Options.PublicDir, "index.html")
	err = tools.IndexRender(tmpl, dst, &c.Options)
	if err != nil {
		return err
	}
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
				// process markdown
				err = c.processor(filefullName, title)
				if err != nil {
					return err
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

func (c *CMD) processor(src string, title string) error {
	article, err := tools.ArticleScanner(src)
	if err != nil {
		return err
	}
	assetsPath := path.Join(strings.TrimRight(src, ".md"))
	layoutPath := path.Join(c.ThemeStr, c.Options.Theme, c.LayoutStr)
	tmpl := path.Join(layoutPath, c.Options.ArticleDir+".html")

	wispeeer.Article = article
	wispeeer.Options = c.Options

	if title != c.IndexStr {
		// process articles
		articles++
		dst := path.Join(c.Options.PublicDir, c.Options.Permalink)
		err = tools.ArticleRender(wispeeer, tmpl, path.Join(dst, title+".html"))
		if err != nil {
			return err
		}
		if utils.IsDir(assetsPath) {
			err = tools.DirCopy(assetsPath, path.Join(dst, title))
			if err != nil {
				return err
			}
		}
		articleList = append(articleList, article)
	} else {
		// process other page
		pages++
		recomandTmpl := path.Join(layoutPath, article.Metadata.Title+".html")
		if utils.IsExist(recomandTmpl) {
			tmpl = recomandTmpl
		}

		dst := path.Join(c.Options.PublicDir, article.Metadata.Title)
		err = tools.PageRender(wispeeer, tmpl, path.Join(dst, "index.html"))
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
