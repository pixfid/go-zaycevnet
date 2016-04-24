# go-zaycevnet
Golang bindings for the Zaycev.net API

# Usage:

```go
package main

import (
	"github.com/pixfid/go-zaycevnet/api"
	"net/url"
	"strconv"
)

func main() {
	client := api.NewZClient("", "kmskoNkYHDnl3ol2") //60kQwLlpV3jv //d7DVaaELv
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
# Supported Methods
>client.Search(query string)<br/>
>client.AutoComplete(query string)<br/>
>client.TOP(page int)<br/>
>client.MusicSetList(page int)<br/>
>client.MusicSetDetile(musicSetId int)<br/>
>client.Genre(genreName string, page int)<br/>
>client.Artist(artistID int)<br/>
>client.Track(trackID int)<br/>
>client.Options<br/>
>client.Download(trackID int)<br/>
>client.Play(trackID int)<br/>
