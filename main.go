package main

import (
    "os"
    "fmt"
    "strings"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/url"
    "github.com/atotto/clipboard"
)

type Message struct {
    Url string
}

func help() {
    fmt.Println("The right gif, every time, in your command line!\nPowered by rightgif.com\n")
    fmt.Println("Usage: rgif [search]\n")
    fmt.Println("> rgif oh boy\n> rgif whatever\n> rgif no no no\n")
    os.Exit(1)
}

func search(query string) {
    var m Message
    
    resp, err := http.PostForm("https://rightgif.com/search/web",
        url.Values{"text": {query}})

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    err = json.Unmarshal(body, &m)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    clipboard.WriteAll(m.Url)
    fmt.Println("✓ gif copied to clipboard")
}

func main() {
    args := os.Args[1:]
    query := strings.Join(args, " ")

    if len(args) == 0 {
        help()
    } 

    search(query)
}
