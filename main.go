package main

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/dustin/go-humanize"
	"github.com/briandowns/spinner"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// Gif holds information about a gif response.
type Gif struct {
	Url  string // gif url
	Size uint64 // file size in bytes
}

// Print help text.
func printHelp() {
	fmt.Println("The right gif, every time, in your command line!\nPowered by rightgif.com\n")
	fmt.Println("Usage: rgif [query]\n")
	fmt.Println("$ rgif oh boy\n$ rgif whatever\n$ rgif no no no\n")
}

// Get content length from a HEAD request to given uri.
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

// Make a search request to the api and return a gif.
func search(query string) (Gif, error) {
	var gif Gif

	resp, err := http.PostForm("https://rightgif.com/search/web",
		url.Values{"text": {query}})
	defer resp.Body.Close()
	if err != nil {
		return gif, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return gif, err
	}
	if err = json.Unmarshal(body, &gif); err != nil {
		return gif, err
	}
	gif.Size = getContentLength(gif.Url)
	
	return gif, nil
}

// Main program.
func main() {
	args := os.Args[1:]
	query := strings.Join(args, " ")

	if len(args) == 0 {
		printHelp()
		return
	}

	spin := spinner.New(spinner.CharSets[4], 125*time.Millisecond)
	spin.Start()

	gif, err := search(query)
	if err != nil {
		panic(err)
	}

	spin.Stop()
	clipboard.WriteAll(gif.Url)
	fmt.Printf("✓ gif copied to clipboard (%s)\n",
		humanize.Bytes(gif.Size))
}
