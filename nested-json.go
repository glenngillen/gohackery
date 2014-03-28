package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	byt := []byte(`{"name": "Glenn Gillen", "handle": "@glenngillen", "age": 32,
	                "dependents": [{"name": "Emma"}, {"name": "Eliott"}],
			"measurements": {"height": 177, "weight": 73}}`)
	var obj map[string]interface{}
	json.Unmarshal(byt, &obj)
	println("JSON string parsed into: ")
	fmt.Println(obj)
	println("Dependents came out as: ")
	fmt.Println(obj["dependents"])
	println("Measurements came out as: ")
	fmt.Println(obj["measurements"])
}
