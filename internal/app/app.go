package app

import (
	"fmt"
	"log"
	"os"

	"github.com/ka1i/wispeeer/internal/app/cmd"
	"github.com/ka1i/wispeeer/internal/pkg/usage"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/version"
)

type app struct {
	success int
	failure int
}

func (a *app) Wispeeer() int {
	if len(os.Args) > 1 {
		var argc = len(os.Args)
		var argv = os.Args[1:]
		barry(argc, argv)
	} else {
		usage.Usage()
		return a.failure
	}
	return a.success
}

func barry(argc int, argv []string) {
	var err error
	switch argv[0] {
	case "-i", "init":
		if argc > 2 {
			if utils.IsValid(argv[1]) {
				err = cmd.Initialzation(argv[1])
			} else {
				err = fmt.Errorf("invalid name")
			}
		} else {
			err = fmt.Errorf("wispeeer init <ka1i.github.io>")
		}
	case "-n", "new":
		log.Println("new")
	case "-g", "generate":
		log.Println("generate")
	case "-s", "server":
		log.Println("server")
	case "-d", "deploy":
		log.Println("deploy")
	case "-h", "--help", "help":
		usage.Usage()
	case "-v", "--version", "version":
		version.Version.Print()
	default:
		err = fmt.Errorf("wispeeer usage: wispeeer -h")
	}
	if err != nil {
		log.Println(err)
	}
}

var App = GetApp()

func GetApp() *app {
	return &app{
		success: 0,
		failure: 137,
	}
}
