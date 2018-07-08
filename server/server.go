package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	webRoot = "www"
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
	// "Content-Encoding":       "gzip",
	"X-Content-Type-Options": "nosniff",
	"X-XSS-Protection":       "1; mode=block",
}

var preRenderMap = map[string]bool{
	"/": true,
}

const htmlRootFile = "/index.html"

func WebServer(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	println(path)
	var f *os.File
	var err error

	if preRenderMap[path] {
		f, err = os.Open(webRoot + "/prerender" + path + htmlRootFile)
		if err != nil {
			panic(err)
		}
	} else {
		f, err = os.Open(webRoot + path)
		if err != nil {
			f, err = os.Open(webRoot + htmlRootFile)
			if err != nil {
				panic(err)
			}
		}
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil || d.IsDir() {
		f, err = os.Open(webRoot + htmlRootFile)
		if err != nil {
			panic(err)
		}
		d, err = f.Stat()
		if err != nil {
			panic(err)
		}
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
