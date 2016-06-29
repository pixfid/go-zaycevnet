package main

import (
	"github.com/pixfid/go-zaycevnet/api"
)

func main() {
	client := api.NewZClient(nil, "", "static_key")
	if err := client.Auth(); err != nil {
		return
	}
	r, err := client.AutoComplete("aa")
	if err != nil {
		return
	}
	println(r.Terms[0])

}
