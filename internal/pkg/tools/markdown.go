package tools

import (
	"bufio"
	"os"
	"time"
)

func CreateMarkdown(fileName string, title string, tags string) error {
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
	fileWrite.WriteString("tags: " + tags + "\n")
	fileWrite.WriteString("categories: " + tags + "\n")
	fileWrite.WriteString("------\n")
	fileWrite.WriteString("\n\n")
	fileWrite.WriteString("# Absract\n")
	fileWrite.WriteString("<!-- more -->\n\n")
	fileWrite.WriteString("# Reference\n\n")

	//Flush buffer
	fileWrite.Flush()
	return nil
}
