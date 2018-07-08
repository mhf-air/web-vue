package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const domain = "http://localhost:9000"

	resp, err := http.Get(domain + "/get-diff")
	ck(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return
	}

	buf, err := ioutil.ReadAll(resp.Body)
	ck(err)
	fmt.Println(string(buf))
}

func ck(err error) {
	if err != nil {
		panic(err)
	}
}
