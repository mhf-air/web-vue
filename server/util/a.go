package util

import (
	"archive/tar"
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func ck(err error) {
	if err != nil {
		panic(err)
	}
}

func GetLocalFileList(root string) []string {
	prefixLen := len(root)
	if !strings.HasSuffix(root, "/") {
		prefixLen++
	}

	result := getLocalFile(prefixLen, root, func(name string) bool {
		if name == "index.html" || name == "prerender" || strings.HasPrefix(name, ".") {
			return true
		}
		return false
	})
	return result
}

func GetLocalHtmlList(root string) []string {
	prefixLen := len(root)
	if !strings.HasSuffix(root, "/") {
		prefixLen++
	}

	result := getLocalFile(prefixLen, root+"/prerender", nil)
	result = append(result, "index.html")
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

func Compress(dir string, fileList []string) string {
	pwd, err := os.Getwd()
	ck(err)
	err = os.Chdir(dir)
	ck(err)
	defer os.Chdir(pwd)

	// tar
	var (
		tarBuf bytes.Buffer
		tw     = tar.NewWriter(&tarBuf)
		f      []byte
		header = &tar.Header{
			Mode: 0600,
		}
	)

	for _, file := range fileList {
		f, err = ioutil.ReadFile(file)
		ck(err)
		header.Name = file
		header.Size = int64(len(f))
		err = tw.WriteHeader(header)
		ck(err)
		_, err = tw.Write(f)
		ck(err)
	}

	err = tw.Close()
	ck(err)

	// gzip
	/* var (
		gzipBuf bytes.Buffer
		zw      = gzip.NewWriter(&gzipBuf)
	)

	zw.Name = "files"
	zw.Comment = "new files to be deployed"
	zw.ModTime = time.Now()

	_, err = zw.Write(tarBuf.Bytes())
	ck(err)
	err = zw.Close()
	ck(err) */

	// base64
	result := base64.StdEncoding.EncodeToString(tarBuf.Bytes())

	return result
}

func Uncompress(data string) map[string][]byte {
	// base64
	decoded, err := base64.StdEncoding.DecodeString(data)
	ck(err)

	// gzip
	/* zr, err := gzip.NewReader(bytes.NewBuffer(data))
	ck(err)
	buf, err := ioutil.ReadAll(zr)
	ck(err)
	err = zr.Close()
	ck(err) */

	// map
	m := map[string][]byte{}

	// tar
	tr := tar.NewReader(bytes.NewBuffer(decoded))
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		ck(err)
		buf, err := ioutil.ReadAll(tr)
		ck(err)
		m[header.Name] = buf
	}

	return m
}

func DirList(pathList []string) []string {
	m := map[string]bool{}
	for _, p := range pathList {
		list := subDirList(p)
		for _, item := range list {
			m[item] = true
		}
	}

	result := make([]string, len(m))
	i := 0
	for k := range m {
		result[i] = k
		i++
	}

	sort.Strings(result)
	return result
}

func subDirList(s string) []string {
	indexList := []int{}
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		if s[i] == '/' {
			indexList = append(indexList, i)
		}
	}

	indexLen := len(indexList)
	list := make([]string, indexLen)
	for i := 0; i < indexLen; i++ {
		list[i] = s[:indexList[i]]
	}
	return list
}

func Mkdir(list []string) {
	for _, dir := range list {
		err := os.MkdirAll(dir, 0700)
		ck(err)
	}
}

func MkHtmlDir() {
	list := []string{
		"html", "html/ready", "html/run", "html/last",
	}
	for _, dir := range list {
		err := os.MkdirAll(dir, 0700)
		ck(err)
	}
}

type WWW struct {
	DirList []string `json:"dir_list"`
	Data    string   `json:"data"`
}

func Cd(dir string) (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	err = os.Chdir(dir)
	if err != nil {
		return "", err
	}
	return pwd, nil
}

type ApiResult struct {
	Code         int         `json:"code"`
	Data         interface{} `json:"data"`
	ErrorMessage string      `json:"error_message"`
}
