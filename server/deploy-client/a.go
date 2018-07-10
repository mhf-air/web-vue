package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../util"
)

func main() {
	deployAsset()
	// deployHtml()
	// clean()
}

const (
	DOMAIN     = "http://localhost:9000"
	LOCAL_ROOT = "/home/mhf/js/src/web-vue/www"
)

func getDiff() []string {
	result := util.GetLocalFileList(LOCAL_ROOT)
	buf := post("/get-diff", result)
	list := []string{}
	err := json.Unmarshal(buf, &list)
	ck(err)
	return list
}

func deployAsset() {
	list := getDiff()
	if len(list) == 0 {
		fmt.Println("no new files to be deployed")
		return
	}

	result := util.Compress(LOCAL_ROOT, list)
	buf := post("/deploy-asset", result)
	if string(buf) != "ok" {
		log.Fatal(string(buf))
	}
}

func deployHtml() {
	list := util.GetLocalHtmlList(LOCAL_ROOT)
	list = append(list, "index.html")

	result := util.Compress(LOCAL_ROOT, list)
	buf := post("/deploy-html", result)
	if string(buf) != "ok" {
		log.Fatal(string(buf))
	}
}

func clean() {
	result := util.GetLocalFileList(LOCAL_ROOT)
	buf := post("/clean", result)
	if string(buf) != "ok" {
		log.Fatal(string(buf))
	}
}

// ================================================================================
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
