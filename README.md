GopherCSS - A compiler from Go to CSS
-------------------------------------

[![Sourcegraph](https://sourcegraph.com/github.com/mateuspontes/gophercss/-/badge.svg)](https://sourcegraph.com/github.com/mateuspontes/gophercss?badge)

GopherCSS compiles Go code ([golang.org](https://golang.org/)) to pure CSS stylesheet code. Its main purpose is to give you the opportunity to write front-end code in Go which will still run in all browsers.

### What is supported?
You define your stylesheet inside Go code, compile and use. All CSS features are converted and minified for your confort.
See the ([can I use](https://caniuse.com/)).

### Installation and Usage
Get or update GopherCSS and dependencies with:

```
go get -u github.com/mateuspontes/gophercss
```

Now you can use `gophercss build [package]`, `gophercss build [files]` or `gophercss install [package]` which behave similar to the `go` tool. For `main` packages, these commands create a `.css` file and `.min.css` minified file in the current directory or in `$GOPATH/bin`. The generated CSS file can be used as usual in a website. Use `gophercss help [command]` to get a list of possible command line flags, e.g. for minification and automatically watching for changes.

`gophercss` uses your platform's default `GOOS` value when generating code. Supported `GOOS` values are: `linux`, `darwin`. If you're on a different platform (e.g., Windows or FreeBSD), you'll need to set the `GOOS` environment variable to a supported value. For example, `GOOS=linux gophercss build [package]`.

*Note: Gophercss will try to write compiled object files of the core packages to your $GOROOT/pkg directory. If that fails, it will fall back to $GOPATH/pkg.*

### Getting started

This is the initial version of the transpiler, you can write nested styles to your classes.

```go
package main

import "github.com/mateuspontes/gophercss"

func main() {
  css.Set("html", css.Style{
    "margin": 0,
    "padding": 0,
  })

  css.Set("body", css.Style{
    "backgroundColor": "#eeeeee",
    "fontFamily": "Arial",
    "fontSize": "16px",
  })

  css.Set(".wrapper", css.Style{
    "display": "flex",
    "justifyContent": "center",
    ".content": css.Style{
      "flex": 1,
      "maxWidth": "600px",

    }
  })
}
```

The output will be:


```css
html {
  margin: 0;
  padding: 0;
}
body {
  background-color: "#eee"
  font-family: Arial;
  font-size: 16px;
}
.wrapper {
  display: flex;
  justify-content: center;
}

.wrapper .content {
  flex: 1;
  max-width: "600px";
}
```

### Development status

Contribute with PR and comments to improve the code and documentation.
