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
	LOCAL_ROOT = "/home/mhf/web/www"

	HTML_READY = LOCAL_ROOT + "/html/ready"
	HTML_NEW   = LOCAL_ROOT + "/html/new"
	HTML_LAST  = LOCAL_ROOT + "/html/last"
)

func main() {
	fmt.Println("deploy server started")

	log.Fatal(http.ListenAndServe(util.DEPLOY_PORT, http.HandlerFunc(DeployServer)))
}

var routeMap = map[string]func(*Context) (*util.ApiResult, error){
	"/get-diff":     getDiff,
	"/deploy-html":  deployHtml,
	"/deploy-asset": deployAsset,
	"/clean":        clean,
}

type Context struct {
	w    http.ResponseWriter
	req  *http.Request
	body []byte
}

func (c *Context) api(r *util.ApiResult) {
	buf, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	_, err = c.w.Write(buf)
	if err != nil {
		panic(err)
	}
}

func apiResult(data interface{}) *util.ApiResult {
	return &util.ApiResult{
		Code: 0,
		Data: data,
	}
}
func apiError(errorMessage string) *util.ApiResult {
	return &util.ApiResult{
		Code:         1,
		ErrorMessage: errorMessage,
	}
}

func DeployServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.Path)

	// cd
	pwd, err := util.Cd(LOCAL_ROOT)
	if err != nil {
		panic(err)
	}
	defer os.Chdir(pwd)

	c := &Context{
		w:   w,
		req: req,
	}
	c.body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		c.api(apiError(err.Error()))
		return
	}

	f, ok := routeMap[req.URL.Path]
	if !ok {
		c.api(apiError("route not found"))
		return
	}

	// handler
	result, err := f(c)
	if err != nil {
		c.api(apiError(err.Error()))
	} else {
		c.api(result)
	}
}

// ================================================================================

// input: new www structure
// diff input with old www on server
func getDiff(c *Context) (*util.ApiResult, error) {
	list := []string{}
	err := json.Unmarshal(c.body, &list)
	if err != nil {
		return nil, err
	}

	oldList := util.GetLocalFileList(LOCAL_ROOT)
	oldMap := make(map[string]bool, len(oldList))
	for _, item := range oldList {
		oldMap[item] = true
	}

	result := []string{}
	for _, item := range list {
		if _, ok := oldMap[item]; !ok {
			result = append(result, item)
		}
	}

	return apiResult(result), nil
}

func deployAsset(c *Context) (*util.ApiResult, error) {
	www := util.WWW{}
	err := json.Unmarshal(c.body, &www)
	if err != nil {
		return nil, err
	}

	// ensure all the direcotries exist
	util.Mkdir(www.DirList)

	m := util.Uncompress(www.Data)
	for file, data := range m {
		err := ioutil.WriteFile(file, data, 0600)
		if err != nil {
			return nil, err
		}
	}

	return apiResult(nil), nil
}

func deployHtml(c *Context) (*util.ApiResult, error) {
	www := util.WWW{}
	err := json.Unmarshal(c.body, &www)
	if err != nil {
		return nil, err
	}

	// ensure all the html direcotries exist
	util.MkHtmlDir()

	// cd
	pwd, err := util.Cd(HTML_READY)
	if err != nil {
		panic(err)
	}
	defer os.Chdir(pwd)

	// ensure all the direcotries exist
	util.Mkdir(www.DirList)

	m := util.Uncompress(www.Data)
	for file, data := range m {
		err := ioutil.WriteFile(file, data, 0600)
		if err != nil {
			return nil, err
		}
	}

	return apiResult(nil), nil
}

// input: old files in www
func clean(c *Context) (*util.ApiResult, error) {
	list := []string{}
	err := json.Unmarshal(c.body, &list)
	if err != nil {
		return nil, err
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
	for _, item := range result {
		os.Remove(item)
	}

	return apiResult(nil), nil
}
