# /rgif

The right gif, every time, in your command-line.  

```
rgif so excited
✓ gif copied to clipboard (551 kB)
```

_Powered by [rightgif.com](https://rightgif.com)_

## How to use

```
Usage:
  rgif [flags] [query]

Flags:
  -u    print url only
  -o    open in default browser

Examples:
  rgif oh boy
  rgif -o whatever
  rgif -u "can't touch this!"
```

## Installation

[Install and set up Go](https://golang.org/doc/install), then run:

```
go get -u github.com/michaelhue/rgif
```

You'll find the binary in `$GOPATH/bin` and source files in `$GOPATH/src/github.com/michaelhue/rgif`. If everything is set up correctly you should be able to run `rgif` in your terminal.

---

**Please note:** This is written by a Go noob as a small learning excercise. Don't hate, share your wisdom and open a pull request!