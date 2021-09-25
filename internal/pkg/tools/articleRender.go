package tools

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

func PostRender(article Article, tmpl string, dst string) error {
	unsafeOverview := blackfriday.Run([]byte(article.Overview))
	htmlSourceOverview := bluemonday.UGCPolicy().SanitizeBytes(unsafeOverview)
	article.Overview = template.HTML(htmlSourceOverview)

	unsafeContent := blackfriday.Run([]byte(article.Content))
	htmlSourceContent := bluemonday.UGCPolicy().SanitizeBytes(unsafeContent)
	article.Content = template.HTML(htmlSourceContent)

	filePath := filepath.Dir(dst)
	if !utils.IsExist(filePath) {
		fmt.Println(filePath)
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

	err = t.Execute(f, article)
	if err != nil {
		return err
	}
	return nil
}
