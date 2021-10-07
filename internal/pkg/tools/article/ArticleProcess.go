package article

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/config"
)

func (rt *rt) ArticleDetailRender(article Article, dst string) error {
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
	err = rt.Template.Execute(f, struct {
		Article Article
		Options config.Options
		Mode    uint64
	}{Article: article, Options: *rt.Options, Mode: 0})
	if err != nil {
		return err
	}
	return nil
}

func (rt *rt) PageDetailRender(page Article, dst string) error {
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
	err = rt.Template.Execute(f, struct {
		Article Article
		Options config.Options
		Mode    uint64
	}{Article: page, Options: *rt.Options, Mode: 1})
	if err != nil {
		return err
	}
	return nil
}

func (rt *rt) ArticlesPaginationRender() error {
	rt.articleSort()

	dstPrefix := path.Join(rt.Options.PublicDir, rt.Options.PaginationDir)
	if !utils.IsExist(dstPrefix) {
		err := os.MkdirAll(dstPrefix, os.ModePerm)
		if err != nil {
			return fmt.Errorf("fail to create floder %v ", dstPrefix)
		}
	}

	// process page split
	for i := 1; i <= utils.SplitPage(rt.totalArticle); i++ {
		tdst := path.Join(dstPrefix, strconv.Itoa(i)+".html")

		rules, err := rt.renderRuleCtx()
		if err != nil {
			return err
		}
		for _, v := range rules {
			rule := strings.Fields(v)
			if len(rule) == 3 {
				if rule[0] == "2" {
					match, err := filepath.Match(rule[1], tdst)
					if err != nil {
						return err
					}
					if match {
						tdst = rule[2]
					}
				}
			}
		}
		f, err := os.Create(tdst)
		if err != nil {
			return err
		}
		defer f.Close()
		err = rt.Template.Execute(f, struct {
			Article []Article
			Options config.Options
			Mode    uint64
		}{Article: rt.Articles, Options: *rt.Options, Mode: 2})
		if err != nil {
			return err
		}

	}
	return nil
}

func (rt *rt) articleSort() {
	sort.Slice(rt.Articles, func(i, j int) bool {
		return rt.Articles[i].Metadata.Posted.Before(rt.Articles[j].Metadata.Posted)
	})
}

func (rt *rt) renderRuleCtx() ([]string, error) {
	var err error
	var rule string
	ruleFile := path.Join("themes", rt.Options.Theme, "rule.txt")
	if utils.IsExist(ruleFile) {
		content, err := ioutil.ReadFile(ruleFile)
		if err != nil {
			return nil, err
		}
		rule = string(content)
	} else {
		rule = "2 {{ .PublicDir }}/{{ .PaginationDir }}/1.html {{ .PublicDir }}/index.html"
	}
	tmpl, err := template.New("rule").Parse(rule)
	if err != nil {
		return nil, err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, rt.Options)
	if err != nil {
		return nil, err
	}
	return strings.Split(tpl.String(), "\n"), nil
}
