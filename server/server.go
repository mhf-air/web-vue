package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	WEB_ROOT       = "/home/mhf/js/src/web-vue/www"
	HTML_ROOT_FILE = "/index.html"
)

func main() {
	fmt.Println("web server started")

	log.Fatal(http.ListenAndServe(":9000", http.HandlerFunc(WebServer)))

	/* const (
			port     = ":8080"
			certFile = "fullchain.pem"
			keyFile  = "privkey.pem"
		)
	  log.Fatal(http.ListenAndServeTLS(port, certFile, keyFile, http.HandlerFunc(WebServer))) */
}

var htmlHeaderMap = map[string]string{
	"X-Content-Type-Options": "nosniff",
	"X-Frame-Options":        "SAMEORIGIN",
	"X-XSS-Protection":       "1; mode=block",
}

var preRenderMap = map[string]bool{
	"/": true,
}

func WebServer(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	fmt.Println(path)

	var (
		f        *os.File
		err      error
		needGzip = false
	)

	if strings.HasSuffix(path, ".js") ||
		strings.HasSuffix(path, ".css") ||
		strings.HasSuffix(path, ".html") {
		needGzip = true
	}

	newPath := WEB_ROOT + path
	/* if preRenderMap[path] {
		needGzip = true
		if path != "/" {
			newPath = WEB_ROOT + "/prerender" + path + HTML_ROOT_FILE
		} else {
			newPath = WEB_ROOT + "/prerender" + HTML_ROOT_FILE
		}
	} */

	f, err = os.Open(newPath)
	if err != nil {
		needGzip = true
		f, err = os.Open(WEB_ROOT + HTML_ROOT_FILE)
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil || (d.IsDir() && !strings.Contains(path, "/.")) {
		needGzip = true
		f, err = os.Open(WEB_ROOT + HTML_ROOT_FILE)
		if err != nil {
			panic(err)
		}
		d, err = f.Stat()
		if err != nil {
			panic(err)
		}
	}

	if needGzip {
		w.Header().Add("Content-Encoding", "gzip")
	}

	// set header must occur before write to body
	for k, v := range htmlHeaderMap {
		w.Header().Add(k, v)
	}
	if path == "/" || preRenderMap[path] {
		w.Header().Add("Cache-Control", "no-cache")
	} else {
		w.Header().Add("Cache-Control", "max-age=31536000") // one year
	}

	http.ServeContent(w, req, d.Name(), d.ModTime(), f)
}
