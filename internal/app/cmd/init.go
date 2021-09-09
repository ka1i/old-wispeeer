package cmd

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/ka1i/wispeeer/internal/pkg/utils"
	assets "github.com/ka1i/wispeeer/pkg/asset"
	logeer "github.com/ka1i/wispeeer/pkg/log"
)

//Initialzation ...
func Initialzation(title string) error {
	var err error

	defer utils.Timer("wispeeer ", time.Now())
	if utils.IsExist(title) {
		return fmt.Errorf("%s: File exists", title)
	}

	logeer.Println("init").Infof("wispeeer init %s", title)

	logeer.Println("init").Info("unpkg embed assets")

	var storage = assets.GetStorage()
	fs := storage.Fs
	root := storage.Root
	err = assetsUnpkg(&fs, root, root, title)
	if err != nil {
		return err
	}
	return nil
}

func assetsUnpkg(fs *embed.FS, root, start, title string) error {
	assets, err := fs.ReadDir(start)
	if err != nil {
		return err
	}
	for _, file := range assets {
		src := path.Join(start, file.Name())
		dst := path.Join(title, src[len(root)+1:])
		// mkdir dst floder
		err = os.MkdirAll(filepath.Dir((dst)), os.ModePerm)
		if err != nil {
			return err
		}
		// process embed assets
		if file.IsDir() {
			err := assetsUnpkg(fs, root, path.Join(start, file.Name()), title)
			if err != nil {
				return err
			}
		} else if file.Name()[0] == 46 {
			continue
		} else {
			fmt.Printf("unpkg: %s\n", dst)
			in, err := fs.Open(src)
			if err != nil {
				return err
			}
			defer in.Close()
			out, err := os.Create(dst)
			if err != nil {
				return err
			}
			defer out.Close()
			// assets copy
			io.Copy(out, in)
		}
	}
	return nil
}
