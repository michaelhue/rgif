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

func search(query string) string {
	var data ApiResponse

	resp, err := http.PostForm("https://rightgif.com/search/web",
		url.Values{"text": {query}})
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, &data); err != nil {
		panic(err)
	}

	return data.Url
}

func getContentLength(url string) (uint64, error) {
	resp, err := http.Head(url)
	defer resp.Body.Close()
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(resp.Header.Get("Content-Length"), 10, 64)
}

func main() {
	args := os.Args[1:]
	query := strings.Join(args, " ")

	if len(args) == 0 {
		help()
		return
	}

	url := Search(query)
	length, err := getContentLength(url)
	if err != nil {
		panic(err)
	}

	clipboard.WriteAll(url)
	fmt.Printf("✓ gif copied to clipboard (%s)\n",
		humanize.Bytes(length))
}
