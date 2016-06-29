# go-zaycevnet
Golang bindings for the Zaycev.net API

[![Build Status](https://travis-ci.org/pixfid/go-zaycevnet.svg?branch=master)](https://travis-ci.org/pixfid/go-zaycevnet)
[![Go Report Card](https://goreportcard.com/badge/github.com/pixfid/go-zaycevnet)](https://goreportcard.com/report/github.com/pixfid/go-zaycevnet)

# Install: <br/>

go get github.com/pixfid/go-zaycevnet <br/>

# Usage:<br/>

```go
package main

import (
	"github.com/pixfid/go-zaycevnet/api"
	"net/url"
	"strconv"
)

func main() {
	client := api.NewZClient(nil, "", "static_key")
	client.Auth()
	params := url.Values{}
	params.Add("query", "ZZ TOP")
	params.Add("page", strconv.Itoa(1))
	params.Add("type", "all")
	params.Add("sort", "")
	params.Add("style", "")

	result, err := client.Search(params)
	if err != nil {

	}

	println(result.PagesCount)
	println(result.Artist.Name)
}
```

# Supported Methods<br/>
>client.Search(query string)<br/>
>client.AutoComplete(query string)<br/>
>client.Top(page int)<br/>
>client.MusicSetList(page int)<br/>
>client.MusicSetDetile(musicSetId int)<br/>
>client.Genre(genreName string, page int)<br/>
>client.Artist(artistID int)<br/>
>client.Track(trackID int)<br/>
>client.Options<br/>
>client.Download(trackID int)<br/>
>client.Play(trackID int)<br/>
