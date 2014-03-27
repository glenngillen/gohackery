package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		key := parts[0]
		val := parts[1]
		fmt.Printf("'%s' has a value of '%s'\n", key, val)
	}
}
