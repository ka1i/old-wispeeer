package tools

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path"
	"strings"
	"time"

	"github.com/ka1i/wispeeer/pkg/config"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v3"
)

var (
	CONFIG_SPLIT = "------"
	MORE_SPLIT   = "<!--more-->"
)

type metadata struct {
	Title      string
	Posted     time.Time
	Categories []string
	Tags       []string
}

type plugin struct {
	Wordcount int
}

// Article ...
type Article struct {
	URL      string
	Metadata metadata
	Edited   time.Time
	Overview template.HTML
	Content  template.HTML
	Plugin   plugin
}

func ArticleScanner(fullName string, options config.Options) (Article, error) {
	var article Article

	var metadataStr string
	var ContentStr string

	// check article
	content, err := ioutil.ReadFile(fullName)
	if err != nil {
		return article, err
	}

	markdownStr := strings.SplitN(string(content), CONFIG_SPLIT, 2)
	contentLen := len(markdownStr)
	if contentLen > 0 {
		metadataStr = markdownStr[0]
	}
	if contentLen > 1 {
		ContentStr = markdownStr[1]
	}

	article.URL = options.Root + "/" + path.Join()
	// Parse article markdown content
	if err := yaml.Unmarshal([]byte(metadataStr), &article.Metadata); err != nil {
		return article, fmt.Errorf("%s:%v", fullName, err)
	}

	overviewStr := strings.SplitN(ContentStr, MORE_SPLIT, 2)
	if len(overviewStr) > 1 {
		article.Overview = MD2HTML(overviewStr[0])
	}

	article.Content = MD2HTML(ContentStr)

	return article, nil
}

func MD2HTML(origin string) template.HTML {
	unsafeOverview := blackfriday.Run([]byte(origin))
	htmlSourceOverview := bluemonday.UGCPolicy().SanitizeBytes(unsafeOverview)
	return template.HTML(htmlSourceOverview)
}
