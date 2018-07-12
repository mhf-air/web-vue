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

/*
-> means rename

deployHtml
  local   =>    ready

runNew
  ready   ->    run

runNewWithBackup
  run     ->    last
  ready   ->    run

runLast
  run     ->    ready
  last    ->    run
*/

func main() {
	deployAsset()
	deployHtml()

	// clean()
}

const (
	DOMAIN     = "http://localhost:9000"
	LOCAL_ROOT = "/home/mhf/js/src/web-vue/www"
)

func getDiff() []string {
	result := util.GetLocalFileList(LOCAL_ROOT)
	api := post("/get-diff", result).([]interface{})
	list := make([]string, len(api))
	for i, item := range api {
		list[i] = item.(string)
	}
	return list
}

func deployAsset() {
	list := getDiff()
	if len(list) == 0 {
		fmt.Println("no new files to be deployed")
		return
	}

	dirList := util.DirList(list)
	data := util.Compress(LOCAL_ROOT, list)

	post("/deploy-asset", &util.WWW{
		DirList: dirList,
		Data:    data,
	})
}

func deployHtml() {
	list := util.GetLocalHtmlList(LOCAL_ROOT)

	dirList := util.DirList(list)
	data := util.Compress(LOCAL_ROOT, list)

	post("/deploy-html", &util.WWW{
		DirList: dirList,
		Data:    data,
	})
}

func clean() {
	result := util.GetLocalFileList(LOCAL_ROOT)
	post("/clean", result)
}

// ================================================================================
func post(path string, data interface{}) interface{} {
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

	api := &util.ApiResult{}
	err = json.Unmarshal(buf, api)
	ck(err)

	if api.Code != 0 {
		log.Fatal(api.ErrorMessage)
	}

	return api.Data
}

func ck(err error) {
	if err != nil {
		panic(err)
	}
}
