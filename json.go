package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `{"id":"13314", "name": "Glenn Gillen", "handle":"@glenngillen"}`
	var obj map[string]interface{}

	if err := json.Unmarshal([]byte(str), &obj); err != nil {
		panic(err)
	}
	println("JSON string parsed into: ")
	fmt.Println(obj)

	fmt.Printf("id was '%s'\n", obj["id"])
	fmt.Printf("name was '%s'\n", obj["name"])
	fmt.Printf("handle was '%s'\n", obj["handle"])
}
