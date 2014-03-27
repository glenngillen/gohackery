package main

import (
	"io/ioutil"
)

func main() {
	contents, err := ioutil.ReadFile("inputfile.txt")
	if err != nil { panic(err) }
	println(string(contents))
}
