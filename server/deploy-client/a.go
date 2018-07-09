package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	getDiff()
}

func getDiff() []string {
	result := getLocalFileList()
	buf := post("/get-diff", result)
	list := []string{}
	err := json.Unmarshal(buf, &list)
	ck(err)
	return list
}

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

// ================================================================================
const DOMAIN = "http://localhost:9000"

func post(path string, data interface{}) []byte {
	b, err := json.Marshal(data)
	ck(err)
	resp, err := http.Post(DOMAIN+path, "application/json", bytes.NewBuffer(b))
	ck(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return nil
	}

	buf, err := ioutil.ReadAll(resp.Body)
	ck(err)
	return buf
}

func ck(err error) {
	if err != nil {
		panic(err)
	}
}
