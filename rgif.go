// Package main implements a command-line interface
// for quickly fetching a gif matching a search query.
package main

import (
	_ "crypto/sha512"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/briandowns/spinner"
	"github.com/dustin/go-humanize"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// URL to Realgif API endpoint.
var EndpointUrl = "https://rightgif.com/search/web"

// Command-line flags.
var (
	urlFlag  bool // print url only
	openFlag bool // open url
)

// Gif holds information about a gif response.
type Gif struct {
	Url  string // gif url
	Size uint64 // file size in bytes
}

// Print help text.
func printHelp() {
	fmt.Println("The right gif, every time, in your command-line")
	fmt.Println("\nUsage:\n  rgif [flags] [query]")
	fmt.Println("\nFlags:")
	fmt.Println("  -u    print url only")
	fmt.Println("  -o    open in default browser")
	fmt.Println("\nExamples:")
	fmt.Println("  rgif oh boy\n  rgif -o whatever\n  rgif -u \"can't touch this!\"\n")
}

// Get content length from a HEAD request to given uri.
func getContentLength(uri string) uint64 {
	resp, err := http.Head(uri)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	length := resp.Header.Get("Content-Length")
	bytes, err := strconv.ParseUint(length, 10, 64)
	if err != nil {
		return 0
	}
	return bytes
}

// Make a search request to the api and return a gif.
func search(query string) (Gif, error) {
	var gif Gif

	resp, err := http.PostForm(EndpointUrl, url.Values{"text": {query}})
	if err != nil {
		return gif, err
	}
	defer resp.Body.Close()

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

// Init flags.
func init() {
	flag.BoolVar(&urlFlag, "u", false, "print url only")
	flag.BoolVar(&openFlag, "o", false, "open url in browser")
	flag.Parse()
}

// Main program.
func main() {
	args := os.Args[1:]
	query := strings.Join(args, " ")

	if len(args) == 0 {
		printHelp()
		return
	}

	spin := spinner.New(spinner.CharSets[28], 120*time.Millisecond)
	if !urlFlag {
		spin.Start()
	}

	gif, err := search(query)
	if err != nil {
		panic(err)
	}

	spin.Stop()
	clipboard.WriteAll(gif.Url)

	if urlFlag {
		fmt.Printf(gif.Url)
		return
	}
	fmt.Printf("✓ gif copied to clipboard (%s)\n",
		humanize.Bytes(gif.Size))

	if openFlag {
		open.Run(gif.Url)
	}
}
