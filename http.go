package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

	postResp, err := http.PostForm("http://requestb.in/s5rh8ds5",
		url.Values{"name": {"Glenn Gillen"}, "age": {"32"}})
	if err != nil {
		panic(err)
	}
	postBody, err := ioutil.ReadAll(postResp.Body)
	if err != nil {
		panic(err)
	}
	println(string(postBody))
}
