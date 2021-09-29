package cmd

import (
	"net/http"

	loger "github.com/ka1i/wispeeer/pkg/log"
)

// Server ...
func (c *CMD) Server() error {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))

	server := &http.Server{
		Addr:    "0.0.0.0:1080",
		Handler: mux,
	}
	loger.Task("server").Printf("Running at %v\n", server.Addr)
	server.ListenAndServe()

	return nil
}
