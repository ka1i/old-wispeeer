package tools

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/ka1i/wispeeer/internal/pkg/utils"
)

func FileRemove(publicDIR string) {
	if !utils.IsExist(publicDIR) {
		os.Mkdir(publicDIR, os.ModePerm)
	} else {
		dir, _ := ioutil.ReadDir(publicDIR)
		for _, d := range dir {
			if d.Name() != ".git" {
				os.RemoveAll(path.Join([]string{publicDIR, d.Name()}...))
			}
		}
	}
}
