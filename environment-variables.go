package main

import (
	"os"
)

func main() {
	for _, e := range os.Environ() {
		println(e)
	}
}
