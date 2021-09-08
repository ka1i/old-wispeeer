package main

import (
	"os"

	"github.com/ka1i/wispeeer/internal/app"
)

func main() {
	wispeeer := app.App
	os.Exit(wispeeer.Wispeeer())
}
