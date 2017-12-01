package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	user := "admin"
	pass := "admin"

	ip := "192.168.10.1"

	url := conctUrl(user, pass, ip)
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Password correct is ->", pass)
	}

	fmt.Println("Resposta->>", resp)
}

func conctUrl(username, password, url string) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://")
	buffer.WriteString(username)
	buffer.WriteString(":")
	buffer.WriteString(password)
	buffer.WriteString("@")
	buffer.WriteString(url)
	return buffer.String()
}
