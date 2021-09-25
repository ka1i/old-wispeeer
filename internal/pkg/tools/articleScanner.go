package tools

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"
	"time"

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

func ArticleScanner(fullName string) (Article, error) {
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
	// Parse article markdown content
	if err := yaml.Unmarshal([]byte(metadataStr), &article.Metadata); err != nil {
		return article, fmt.Errorf("%s:%v", fullName, err)
	}

	overviewStr := strings.SplitN(ContentStr, MORE_SPLIT, 2)
	if len(overviewStr) > 1 {
		article.Overview = template.HTML(overviewStr[0])
	}

	article.Content = template.HTML(ContentStr)

	return article, nil
}
