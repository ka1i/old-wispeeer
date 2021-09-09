package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
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
