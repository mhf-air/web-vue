package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":9000"
)

func main() {
	fmt.Println("deploy server started")

	log.Fatal(http.ListenAndServe(port, http.HandlerFunc(DeployServer)))
}

func DeployServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)

	switch req.URL.Path {
	case "/get-diff":
		getDiff(w, req)
	case "/deploy":
		deploy(w, req)
	case "clean":
		clean(w, req)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// input: new www structure
// diff input with old www on server
func getDiff(w http.ResponseWriter, req *http.Request) {
}

// input: new files in www and index.html and prerender
func deploy(w http.ResponseWriter, req *http.Request) {
}

// input: old files in www
func clean(w http.ResponseWriter, req *http.Request) {
}
