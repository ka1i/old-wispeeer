package article

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v3"
)

var (
	HEADER_SPLIT = "------"
	MORE_SPLIT   = "<!--more-->"
)

type metadata struct {
	Title      string
	Posted     time.Time
	Categories []string
	Tags       []string
}

// Article ...
type Article struct {
	Metadata metadata
	Overview template.HTML
	Content  template.HTML
}

func (rt *rt) ArticleScanner(fullName string) (Article, error) {
	var article Article

	var metadataStr string
	var ContentStr string

	// check article
	content, err := ioutil.ReadFile(fullName)
	if err != nil {
		return article, err
	}

	markdownStr := strings.SplitN(string(content), HEADER_SPLIT, 2)
	contentLen := len(markdownStr)
	if contentLen > 0 {
		metadataStr = markdownStr[0]
	}
	if contentLen > 1 {
		ContentStr = markdownStr[1]
	}

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
