package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Handle string `json:"handle"`
	Age int `json:"age"`
	Dependents []Person `json:"dependents"`
	Measurements map[string]int `json:"measurements"`
}

func main() {
	byt := []byte(`{"name": "Glenn Gillen", "handle": "@glenngillen", "age": 32,
	                "dependents": [{"name": "Emma"}, {"name": "Eliott"}],
			"measurements": {"height": 177, "weight": 73}}`)
	var obj map[string]interface{}
	glenn := &Person{}
	json.Unmarshal(byt, &obj)
	json.Unmarshal(byt, &glenn)
	println("JSON string parsed into: ")
	fmt.Println(obj)
	println("Dependents came out as: ")
	fmt.Println(obj["dependents"])
	println("Measurements came out as: ")
	fmt.Println(obj["measurements"])

	println("Structured parsing came out as: ")
	fmt.Println(glenn)

	println("Dependents came out as: ")
	fmt.Println(glenn.Dependents)
	println("Measurements came out as: ")
	fmt.Println(glenn.Measurements)
}
