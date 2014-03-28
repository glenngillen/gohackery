package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil { panic(err) }
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "lorem")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("lorem ipsum dolor sit amet\nDonec non rhoncus nunc."))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))
}
