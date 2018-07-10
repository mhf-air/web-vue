package util

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
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

func Compress(dir string, fileList []string) []byte {
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
	var (
		gzipBuf bytes.Buffer
		zw      = gzip.NewWriter(&gzipBuf)
	)

	zw.Name = "files"
	zw.Comment = "new files to be deployed"
	zw.ModTime = time.Now()

	_, err = zw.Write(tarBuf.Bytes())
	ck(err)
	err = zw.Close()
	ck(err)

	return gzipBuf.Bytes()
}

func Uncompress(data []byte) map[string][]byte {
	// gzip
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	ck(err)
	buf, err := ioutil.ReadAll(zr)
	ck(err)
	err = zr.Close()
	ck(err)

	// map
	m := map[string][]byte{}

	// tar
	tr := tar.NewReader(bytes.NewBuffer(buf))
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
