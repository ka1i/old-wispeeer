package assets

import "embed"

//go:embed wisper/*
var fs embed.FS

var root = "wisper"

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
