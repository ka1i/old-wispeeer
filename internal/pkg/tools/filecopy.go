package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/ka1i/wispeeer/internal/pkg/utils"
)

// FileCopy ...
func FileCopy(src string, dst string) error {
	srcStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("fail to read %v filestat", src)
	}

	if !utils.IsExist(filepath.Dir(dst)) {
		err := os.MkdirAll(filepath.Dir(dst), os.ModePerm)
		if err != nil {
			return fmt.Errorf("fail to create floder %v ", filepath.Dir(dst))
		}
	}

	if srcStat.IsDir() {
		files, err := ioutil.ReadDir(src)
		if err != nil {
			return fmt.Errorf("fail to read dir %v ", src)
		}
		for _, f := range files {
			// ignore dotfile
			if f.Name()[0] == 46 {
				continue
			}
			filenameWithSuffix := path.Base(f.Name())
			err = FileCopy(path.Join(src, filenameWithSuffix), path.Join(dst, filenameWithSuffix))
			if err != nil {
				return err
			}
		}
	} else {
		if !srcStat.Mode().IsRegular() {
			return fmt.Errorf("%s is not regular file", src)
		}

		source, err := os.Open(src)
		if err != nil {
			return fmt.Errorf("fail to read %v file", src)
		}
		defer source.Close()

		destination, err := os.Create(dst)
		if err != nil {
			return fmt.Errorf("fail to create file %v", dst)
		}
		defer destination.Close()
		_, err = io.Copy(destination, source)
		if err != nil {
			return fmt.Errorf("fail to copy file %v", dst)
		}
	}
	return nil
}
