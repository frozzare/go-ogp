# go-ogp [![Build Status](https://travis-ci.org/frozzare/go-ogp.svg?branch=master)](https://travis-ci.org/frozzare/go-ogp)

 Fetch [Open Graph protocol](http://ogp.me/) meta tags.

 View the [docs](http://godoc.org/github.com/frozzare/go-ogp).

## Installation

```
$ go get github.com/frozzare/go-ogp
```

## Example

  Note that you don't have to write `og:` in the keys

```go
list := ogp.Fetch("http://ogp.me")
title := list["title"]
image_width := list["image:width"]

// Or you can print out all Open Graph meta tags
ogp.Print("http://ogp.me")
```

# License

 MIT © [Fredrik Forsmo](https://github.com/frozzare)
