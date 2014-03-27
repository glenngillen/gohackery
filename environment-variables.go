package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	envMap := make(map[string]string)
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		key := parts[0]
		val := parts[1]
		envMap[key] = val
		fmt.Printf("'%s' has a value of '%s'\n", key, val)
	}

	envJson, err := json.Marshal(envMap)
	if err != nil { panic(err) }
	fmt.Println(string(envJson))
}
