package tools

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/config"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

type Wispeeer struct {
	Article        `yaml:",inline"`
	config.Options `yaml:",inline"`
}

func ArticleRender(wispeeer Wispeeer, tmpl string, dst string) error {
	unsafeOverview := blackfriday.Run([]byte(wispeeer.Overview))
	htmlSourceOverview := bluemonday.UGCPolicy().SanitizeBytes(unsafeOverview)
	wispeeer.Overview = template.HTML(htmlSourceOverview)

	unsafeContent := blackfriday.Run([]byte(wispeeer.Content))
	htmlSourceContent := bluemonday.UGCPolicy().SanitizeBytes(unsafeContent)
	wispeeer.Content = template.HTML(htmlSourceContent)

	filePath := filepath.Dir(dst)
	if !utils.IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("fail to create floder %v ", filePath)
		}
	}

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}

	err = t.Execute(f, wispeeer)
	if err != nil {
		return err
	}
	return nil
}

func ArticleListRender(articles []Article, options *config.Options, tmpl string, dst string) error {
	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Metadata.Posted.Before(articles[j].Metadata.Posted)
	})

	filePath := dst
	if !utils.IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("fail to create floder %v ", filePath)
		}
	}

	f, err := os.Create(path.Join(dst, "index.html"))
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	wispeeers := &struct {
		Article []Article
		Options config.Options
	}{Article: articles, Options: *options}

	err = t.Execute(f, wispeeers)
	if err != nil {
		return err
	}
	return nil
}

func PageRender(wispeeer Wispeeer, tmpl string, dst string) error {
	unsafeContent := blackfriday.Run([]byte(wispeeer.Content))
	htmlSourceContent := bluemonday.UGCPolicy().SanitizeBytes(unsafeContent)
	wispeeer.Content = template.HTML(htmlSourceContent)

	filePath := filepath.Dir(dst)
	if !utils.IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("fail to create floder %v ", filePath)
		}
	}

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}

	err = t.Execute(f, wispeeer)
	if err != nil {
		return err
	}
	return nil
}

func IndexRender(tmpl string, dst string, options *config.Options) error {
	var err error
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}

	err = t.Execute(f, options)
	if err != nil {
		return err
	}
	return nil
}
