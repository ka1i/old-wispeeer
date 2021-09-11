package app

import (
	"fmt"
	"os"
	"time"

	"github.com/ka1i/wispeeer/internal/app/cmd"
	"github.com/ka1i/wispeeer/internal/pkg/usage"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	"github.com/ka1i/wispeeer/pkg/config"
	logeer "github.com/ka1i/wispeeer/pkg/log"
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
	defer utils.Timer("wispeeer ", time.Now())

	config.Configure.Init()
	run := cmd.Run()

	switch argv[0] {
	case "-i", "init":
		if argc > 2 {
			if utils.IsValid(argv[1]) {
				err = run.Initialzation(argv[1])
			} else {
				err = fmt.Errorf("invalid name")
			}
		} else {
			err = fmt.Errorf("wispeeer init <ka1i.github.io>")
		}
	case "-n", "new":
		if argc > 2 {
			if config.Configure.Error == nil {
				if argv[1] == "page" && argc > 3 {
					if utils.IsValid(argv[2]) {
						err = run.NewPage(argv[2])
					}
				} else {
					if utils.IsValid(argv[1]) {
						err = run.NewPost(argv[1])
					}
				}
			} else {
				err = config.Configure.Error
			}
		} else {
			err = fmt.Errorf("wispeeer new [post] <title>")
		}
	case "-g", "generate":
		err = run.Generate()
	case "-s", "server":
		fmt.Println("server")
	case "-d", "deploy":
		fmt.Println("deploy")
	case "-h", "--help", "help":
		usage.Usage()
	case "-v", "--version", "version":
		version.Version.Print()
	default:
		err = fmt.Errorf("wispeeer usage: wispeeer -h")
	}
	if err != nil {
		logeer.Task("app").Error(err)
	}
}

var App = GetApp()

func GetApp() *app {
	return &app{
		success: 0,
		failure: 137,
	}
}
