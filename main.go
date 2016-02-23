package main

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/dustin/go-humanize"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type ApiResponse struct {
	Url string // gif url
}

func help() {
	fmt.Println("The right gif, every time, in your command line!\nPowered by rightgif.com\n")
	fmt.Println("Usage: rgif [query]\n")
	fmt.Println("$ rgif oh boy\n$ rgif whatever\n$ rgif no no no\n")
}

func search(query string) (string, error) {
	var data ApiResponse

	resp, err := http.PostForm("https://rightgif.com/search/web",
		url.Values{"text": {query}})
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err = json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	return data.Url, nil
}

func getContentLength(uri string) (uint64) {
	resp, err := http.Head(uri)
	defer resp.Body.Close()
	if err != nil {
		return 0
	}
	
	length, err := strconv.ParseUint(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return 0
	}
	return length
}

func main() {
	args := os.Args[1:]
	query := strings.Join(args, " ")

	if len(args) == 0 {
		help()
		return
	}

	uri, err := search(query)
	if err != nil {
		panic(err)
	}
	size := getContentLength(uri)

	clipboard.WriteAll(uri)
	fmt.Printf("✓ gif copied to clipboard (%s)\n",
		humanize.Bytes(size))
}
