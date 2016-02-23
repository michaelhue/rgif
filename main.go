package main

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Response struct {
	Url string
}

func help() {
	fmt.Println("The right gif, every time, in your command line!\nPowered by rightgif.com\n")
	fmt.Println("Usage: rgif [query]\n")
	fmt.Println("> rgif oh boy\n> rgif whatever\n> rgif no no no\n")
}

func search(query string) string {
	var data Response

	resp, err := http.PostForm("https://rightgif.com/search/web",
		url.Values{"text": {query}})

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	return data.Url
}

func main() {
	args := os.Args[1:]
	query := strings.Join(args, " ")

	if len(args) == 0 {
		help()
		return
	}

	url := search(query)
	clipboard.WriteAll(url)
	fmt.Println("✓ gif copied to clipboard")
}
