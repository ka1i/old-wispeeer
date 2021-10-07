package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	loger "github.com/ka1i/wispeeer/pkg/log"
)

// Server ...
func (c *CMD) Server() error {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", handler{}))

	server := &http.Server{
		Addr:    "localhost:4180",
		Handler: mux,
	}
	loger.Task("server").Printf("Running at http://%v\n", server.Addr)
	loger.Task("server").Println("Press Ctrl+c to quit")

	server.ListenAndServe()

	return nil
}

type handler struct{}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	etag := strconv.FormatInt(time.Now().Unix(), 10)
	w.Header().Set("Etag", etag)

	fs := http.FileServer(http.Dir("public"))
	fs.ServeHTTP(w, r)

	requestsURL, _ := url.QueryUnescape(r.RequestURI)
	fmt.Printf("%v ---> %v %v %s ", r.RemoteAddr, r.Proto, r.Method, requestsURL)
	fmt.Println(time.Since(start))
}
