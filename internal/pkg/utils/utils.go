package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	logeer "github.com/ka1i/wispeeer/pkg/log"
)

// Timer ...
func Timer(action string, start time.Time) {
	dis := time.Since(start)
	if dis > 1 {
		fmt.Printf("%s >>> Done in %v\n", action, dis)
	}
}

// IsExist ...
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsValid ...
func IsValid(str string) bool {
	reg := regexp.MustCompile(`[\\\\/:*?\"<>|]`)
	if reg == nil {
		logeer.Task("app").Error("Title Incorrect")
		return false
	}
	result := reg.FindAllString(str, -1)
	return len(result) <= 0
}

// GetWorkspace ...
func GetWorkspace() string {
	dir, err := os.Getwd()
	if err != nil {
		logeer.Task("app").Error(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// SafeFormat ...
func SafeFormat(origin string, spec string, join string, with string) string {
	spell := strings.Fields(origin)
	concat := strings.Join(spell, spec)
	newspell := []string{concat, join}
	r := strings.Join(newspell, with)
	return r
}

func CheckDir(dir string) {
	if !IsExist(dir) {
		os.Mkdir(dir, os.ModePerm)
	}
}

// IsDir ...
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile ...
func IsFile(path string) bool {
	return !IsDir(path)
}
