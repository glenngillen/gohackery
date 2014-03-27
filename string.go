package main

import (
	"strings"
)

func main() {
	s := "This"
	s += " is"
	s += " a"
	s += " sentence."
	println(s)

	sSlice := make([]string, 0)
	sSlice = append(sSlice, "And")
	sSlice = append(sSlice, "a")
	sSlice = append(sSlice, "second")
	sSlice = append(sSlice, "made")
	sSlice = append(sSlice, "of")
	sSlice = append(sSlice, "slices.")
	println(strings.Join(sSlice, " "))

	ssSlice := []string{"Finally", "shortcut"}
	ssSlice = append(ssSlice, "slice")
	ssSlice = append(ssSlice, "notation.")
	println(strings.Join(ssSlice, " "))
}
