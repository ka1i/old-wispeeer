package assets

import "embed"

//go:embed github.io/*
var fs embed.FS

var root = "github.io"

type Storage struct {
	Fs   embed.FS
	Root string
}

func GetStorage() Storage {
	var storage Storage
	storage.Fs = fs
	storage.Root = root
	return storage
}
