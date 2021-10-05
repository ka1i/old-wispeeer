package tools

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

func EmbedUnpkg(fs *embed.FS, root, start, title string) error {
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
			err := EmbedUnpkg(fs, root, path.Join(start, file.Name()), title)
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
