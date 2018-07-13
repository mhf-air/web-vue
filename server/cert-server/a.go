package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	WEB_ROOT = "/home/mhf/web/www"
)

func main() {
	fmt.Println("cert server started")

	log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(WebServer)))
}

func WebServer(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	fmt.Println(path)

	var (
		f   *os.File
		err error
	)

	newPath := WEB_ROOT + path
	if !strings.HasPrefix(path, "/.") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	f, err = os.Open(newPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	http.ServeContent(w, req, d.Name(), d.ModTime(), f)
}
