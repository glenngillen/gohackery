package main

import (
	"io/ioutil"
)

func main() {
	contents := []byte("This is the output file\n")
	err := ioutil.WriteFile("outputfile.txt", contents, 0644)
	if err != nil { panic(err) }
}
