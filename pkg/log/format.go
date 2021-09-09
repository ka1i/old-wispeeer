package logeer

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type CustomizeFormatter struct{}

func (s *CustomizeFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	level := entry.Level
	msg := fmt.Sprintf("[%s] %v %v\n", level, timestamp, entry.Message)
	return []byte(msg), nil
}
