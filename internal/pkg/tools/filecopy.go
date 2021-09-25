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
		return fmt.Errorf("fail to read %v stat", src)
	}
	if srcStat.IsDir() {
		return fmt.Errorf("%v is dir", src)
	}
	filePath := filepath.Dir(dst)
	if !utils.IsExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("fail to create floder %v ", filePath)
		}
	}
	srcfile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("fail to read %v", src)
	}
	defer srcfile.Close()

	dstfile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("fail to create file %v:%v", dst, err)
	}
	defer dstfile.Close()
	_, err = io.Copy(dstfile, srcfile)
	if err != nil {
		return fmt.Errorf("fail copy %s -> %s", src, dst)
	}
	return nil
}

// DirCopy ...
func DirCopy(src string, dst string) error {
	srcStat, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("fail to read %v stat", src)
	}

	//fmt.Println(src, "--->", dst)

	if srcStat.IsDir() {
		if !utils.IsExist(dst) {
			err := os.MkdirAll(dst, os.ModePerm)
			if err != nil {
				return fmt.Errorf("fail to create floder %v ", dst)
			}
		}
		//fmt.Println(!utils.IsExist(filepath.Dir(dst)), filepath.Dir(dst))
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

		err := FileCopy(src, dst)
		if err != nil {
			return err
		}
	}
	return nil
}
