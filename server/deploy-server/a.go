package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	list := []string{}
	err = json.Unmarshal(body, &list)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	oldList := getLocalFileList()
	oldMap := make(map[string]bool, len(oldList))
	for _, item := range oldList {
		oldMap[item] = true
	}

	var result []string
	for _, item := range list {
		if _, ok := oldMap[item]; !ok {
			result = append(result, item)
		}
	}

	b, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
}

// input: new files in www and index.html and prerender
func deploy(w http.ResponseWriter, req *http.Request) {
}

// input: old files in www
func clean(w http.ResponseWriter, req *http.Request) {
}

// ================================================================================

func getLocalFileList() []string {
	const localRoot = "/home/mhf/js/src/web-vue/www"
	result := getLocalFile(len(localRoot), localRoot, func(name string) bool {
		if name == "index.html" || name == "prerender" {
			return true
		}
		return false
	})
	return result
}

func getLocalFile(prefixLen int, dir string, shouldSkip func(string) bool) []string {
	list, err := ioutil.ReadDir(dir)
	ck(err)

	var result []string

	for _, item := range list {
		name := item.Name()
		if shouldSkip != nil && shouldSkip(name) {
			continue
		}

		fullName := dir + "/" + name
		if item.IsDir() {
			result = append(result, getLocalFile(prefixLen, fullName, nil)...)
		} else {
			result = append(result, fullName[prefixLen:])
		}
	}

	return result
}

func ck(err error) {
	if err != nil {
		panic(err)
	}
}
