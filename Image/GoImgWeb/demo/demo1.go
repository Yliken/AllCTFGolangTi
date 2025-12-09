package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("file:///C:/flag.txt")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
