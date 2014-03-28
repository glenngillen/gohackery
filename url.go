package main

import (
	"fmt"
	"net/url"
)

func main() {
	databaseUrl := "postgres://dba:sekrets@host.name.co:5432/database?option=yeah"
	u, err := url.Parse(databaseUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("User:", u.User)
	fmt.Println("Username:", u.User.Username())
	password, _ := u.User.Password()
	fmt.Println("Password:", password)
	fmt.Println("Host:", u.Host)
	fmt.Println("Path:", u.Path)
	fmt.Println("Fragment:", u.Fragment)
	fmt.Println("Raw Query:", u.RawQuery)
	pq, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("Parsed Query:", pq)
}
