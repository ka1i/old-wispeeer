package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/config"
	logeer "github.com/ka1i/wispeeer/pkg/log"
)

// NewPost ...
func NewPost(title string) error {
	logeer.Println("new").Infof("Location: %s", utils.GetWorkspace())
	options := config.Configure.Options

	// 检查发布文件夹状态
	checkAndFixDIR(path.Join(utils.GetWorkspace(), options.SourceDir))
	checkAndFixDIR(path.Join(utils.GetWorkspace(), options.SourceDir, options.PostDir))

	title = utils.SafeFormat(title, " ", "", "")
	var safeName = utils.SafeFormat(title, "_", "md", ".")
	var filePath = path.Join(utils.GetWorkspace(), options.SourceDir, options.PostDir, safeName)
	if utils.IsExist(filePath) {
		return fmt.Errorf("article %v is exist", safeName)
	}
	// 创建文章文件
	err := createMarkdown(filePath, title)
	if err != nil {
		return fmt.Errorf("create article %s is failed", safeName)
	}
	fmt.Printf("title  : %s\n", title)
	fmt.Printf("posted : %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("Created: %s\n", safeName)
	return nil
}

func NewPage(title string) error {
	fmt.Println("page")
	return nil
}

func checkAndFixDIR(dir string) {
	if !utils.IsExist(dir) {
		os.Mkdir(dir, os.ModePerm)
	}
}

func createMarkdown(fileName string, title string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	fileWrite := bufio.NewWriter(file)
	//Markdown header
	fileWrite.WriteString("------\n")
	fileWrite.WriteString("title: " + title + "\n")
	fileWrite.WriteString("posted: " + time.Now().Format("2006-01-02 15:04:05") + "\n")
	fileWrite.WriteString("tags: void\n")
	fileWrite.WriteString("categories: void\n")
	fileWrite.WriteString("------\n")
	fileWrite.WriteString("\n\n")
	fileWrite.WriteString("# Absract\n")
	fileWrite.WriteString("<!-- more -->\n\n")
	fileWrite.WriteString("# Reference\n\n")

	//Flush buffer
	fileWrite.Flush()
	return nil
}
