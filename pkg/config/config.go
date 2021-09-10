package config

import (
	"fmt"

	"github.com/creasty/defaults"
)

type config struct {
	Filenme string
	Options Options
	Error   error
}

func (c *config) Init() {
	options, err := ConfigParser(c.Filenme)
	if err != nil {
		c.Error = fmt.Errorf("%v", err)
		return
	}
	if err := defaults.Set(&options); err != nil {
		c.Error = fmt.Errorf("initial configuration : %v", err)
		return
	}
	c.Options = options
}

var Configure = getConfig()

func getConfig() *config {
	return &config{
		Filenme: "config.yml",
		Options: Options{},
		Error:   nil,
	}
}
