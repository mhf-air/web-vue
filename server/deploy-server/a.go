package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"../util"
)

const (
	PORT       = ":9000"
	LOCAL_ROOT = "/home/mhf/js/src/web-vue/www"
)

func main() {
	fmt.Println("deploy server started")

	log.Fatal(http.ListenAndServe(PORT, http.HandlerFunc(DeployServer)))
}

var routeMap = map[string]func(http.ResponseWriter, *http.Request){
	"/get-diff":     getDiff,
	"/deploy-html":  deployHtml,
	"/deploy-asset": deployAsset,
	"/clean":        clean,
}

func DeployServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)

	f, ok := routeMap[req.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	f(w, req)
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

	oldList := util.GetLocalFileList(LOCAL_ROOT)
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

func deployAsset(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	err = os.Chdir(LOCAL_ROOT)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	defer os.Chdir(pwd)

	m := util.Uncompress(body)
	for file, data := range m {
		err := ioutil.WriteFile(file, data, 0600)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
	}

	_, err = w.Write([]byte("ok"))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
}

func deployHtml(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	err = os.Chdir(LOCAL_ROOT)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	defer os.Chdir(pwd)

	// TODO
	m := util.Uncompress(body)
	for file, data := range m {
		err := ioutil.WriteFile(file, data, 0600)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
	}

	_, err = w.Write([]byte("ok"))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
}

// input: old files in www
func clean(w http.ResponseWriter, req *http.Request) {
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
	newMap := make(map[string]bool, len(list))
	for _, item := range list {
		newMap[item] = true
	}

	oldList := util.GetLocalFileList(LOCAL_ROOT)

	var result []string
	for _, item := range oldList {
		if _, ok := newMap[item]; !ok {
			result = append(result, item)
		}
	}

	// clean
	pwd, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	err = os.Chdir(LOCAL_ROOT)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	defer os.Chdir(pwd)

	for _, item := range result {
		os.Remove(item)
	}

	_, err = w.Write([]byte("ok"))
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
}
