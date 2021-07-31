package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	f, err := os.Create("resp.html")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(f, resp.Body)
}
