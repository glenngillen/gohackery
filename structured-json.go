package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Handle string `json:"handle"`
	Age int `json:"age"`
}

func main() {
	byt := []byte(`{"name": "Glenn Gillen", "handle": "@glenngillen", "age": 32}`)
	glenn := &Person{}
	json.Unmarshal(byt, &glenn)
	fmt.Printf("Name is %s\n", glenn.Name)
	fmt.Printf("Handle is %s\n", glenn.Handle)
	fmt.Printf("Age is %d\n", glenn.Age)
}
