package tools

import (
	"fmt"
	"html/template"
	"os"
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

type Wispeeers struct {
	Article []Article
	Pages   []string
	Options config.Options
}

func PostRender(wispeeer Wispeeer, tmpl string, dst string) error {
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

func ListRender(wispeeers Wispeeers, tmpl string, dst string) error {
	sort.Slice(wispeeers.Article, func(i, j int) bool {
		return wispeeers.Article[i].Metadata.Posted.Before(wispeeers.Article[j].Metadata.Posted)
	})
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}

	err = t.Execute(f, wispeeers)
	if err != nil {
		return err
	}
	return nil
}
