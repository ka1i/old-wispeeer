package tools

import (
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/config"
)

type Wispeeer struct {
	Blog struct {
		Articles uint64
		Pages    uint64
		Article  []Article
		Options  config.Options
		Mode     uint64
	}
	Template *template.Template
}

func (w *Wispeeer) Init() {
	w.parserTemplate()
}

func (w *Wispeeer) articleSort() {
	sort.Slice(w.Blog.Article, func(i, j int) bool {
		return w.Blog.Article[i].Metadata.Posted.Before(w.Blog.Article[j].Metadata.Posted)
	})
}

func (w *Wispeeer) parserTemplate() {
	main := path.Join("themes", w.Blog.Options.Theme, "layouts")

	t, err := HTMLParse(template.New("layouts"), main, "*.html")

	if err != nil {
		panic(err)
	}

	w.Template = t.Lookup("index.html")
}

func (w *Wispeeer) ArticleDetailRender(article Article, dst string) error {
	w.Blog.Mode = 0
	fmt.Printf("Process post: %s ---> %s \n", article.Metadata.Title, dst)
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
	err = w.Template.Execute(f, struct {
		Article Article
		Options config.Options
		Mode    uint64
	}{Article: article, Options: w.Blog.Options, Mode: w.Blog.Mode})
	if err != nil {
		return err
	}
	return nil
}

func (w *Wispeeer) PageDetailRender(page Article, dst string) error {
	w.Blog.Mode = 1
	fmt.Printf("Process page: %s ---> %s \n", page.Metadata.Title, dst)
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
	err = w.Template.Execute(f, struct {
		Article Article
		Options config.Options
		Mode    uint64
	}{Article: page, Options: w.Blog.Options, Mode: w.Blog.Mode})
	if err != nil {
		return err
	}
	return nil
}

func (w *Wispeeer) ArticlesPaginationRender(dst string) error {
	w.articleSort()
	w.Blog.Mode = 2
	filePath := dst
	if !utils.IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("fail to create floder %v ", filePath)
		}
	}
	for i := uint64(1); i <= splitpage(w.Blog.Articles); i++ {
		tdst := path.Join(dst, strconv.FormatUint(i, 10)+".html")

		match, err := filepath.Match("public/articles/1.html", tdst)
		if err != nil {
			return err
		}
		if match {
			tdst = path.Join("public", "index.html")
		}
		f, err := os.Create(tdst)
		if err != nil {
			return err
		}
		defer f.Close()
		err = w.Template.Execute(f, w.Blog)
		if err != nil {
			return err
		}

	}

	return nil
}

func splitpage(t uint64) uint64 {
	n := uint64(t / 9)
	if t%9 == 0 {
		return n
	}
	return n + 1
}
