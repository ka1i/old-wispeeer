package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
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
		log.Println("Title Incorrect")
		return false
	}
	result := reg.FindAllString(str, -1)
	return len(result) <= 0
}

// GetWorkspace ...
func GetWorkspace() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("error")
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
