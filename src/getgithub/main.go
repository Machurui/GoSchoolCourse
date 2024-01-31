package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=2&per_page=3")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(resp.Body)
	writer := customWriter{}
	io.Copy(writer, resp.Body)
}

type GithubResponse []struct {
	FullName string `json:"full_name"`
}

type customWriter struct {
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GithubResponse
	json.Unmarshal(p, &resp)
	for _, r := range resp {
		fmt.Println(r.FullName)
	}

	return len(p), nil
}
