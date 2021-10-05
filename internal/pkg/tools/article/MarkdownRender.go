package article

import (
	"html/template"
	"path"
	"strings"

	"github.com/ka1i/wispeeer/internal/pkg/tools"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/config"
)

type rt struct {
	Articles     []Article
	Options      *config.Options
	Template     *template.Template
	totalArticle int
	GP           GlobalVariable
}

func InitRT() *rt {
	// config.yml options
	options := &config.Configure.Options
	// template parser
	prefix := path.Join("themes", options.Theme, "layouts")
	t, err := utils.HTMLParse(template.New("layouts"), prefix, "*.html")
	if err != nil {
		panic(err)
	}
	template := t.Lookup("index.html")

	return &rt{
		Articles: make([]Article, 1),
		Options:  options,
		Template: template,
	}
}

func MarkdownRender(articles []string) (int, int, error) {
	var total int = 0
	var articlesN int = 0

	rt := InitRT()

	options := rt.Options

	for k, fullPath := range articles {
		// parser article file info
		suffix := path.Ext(fullPath)
		fileName := path.Base(fullPath)
		fileNameWithoutSuffix := strings.TrimSuffix(fileName, suffix)

		// parser
		assetsPath := path.Join(strings.TrimRight(fullPath, ".md"))
		article, err := rt.ArticleScanner(fullPath)
		if err != nil {
			return 0, 0, err
		}
		// post & page
		if fileName != "index.md" {
			// render articles
			articlesN++
			rt.Articles = append(rt.Articles, article)
			dstPrefix := path.Join(options.PublicDir, options.Permalink)
			err = rt.ArticleDetailRender(article, path.Join(dstPrefix, fileNameWithoutSuffix+".html"))
			if err != nil {
				return 0, 0, err
			}
			if utils.IsDir(assetsPath) {
				err := tools.DirCopy(assetsPath, path.Join(dstPrefix, fileNameWithoutSuffix))
				if err != nil {
					return 0, 0, err
				}
			}
		} else {
			// render other page
			dstPrefix := path.Join(options.PublicDir, article.Metadata.Title)
			err = rt.PageDetailRender(article, path.Join(dstPrefix, "index.html"))
			if err != nil {
				return 0, 0, err
			}

			if utils.IsDir(assetsPath) {
				err := tools.DirCopy(assetsPath, path.Join(dstPrefix, article.Metadata.Title))
				if err != nil {
					return 0, 0, err
				}
			}
		}
		total = k
	}
	// article pagination
	rt.totalArticle = articlesN
	err := rt.ArticlesPaginationRender()
	if err != nil {
		return 0, 0, err
	}

	return total + 1, articlesN, nil
}
