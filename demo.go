package main

import (
	"github.com/pixfid/go-zaycevnet/api"
)

func main() {
	client := api.NewZClient(nil, "", "kmskoNkYHDnl3ol2") //60kQwLlpV3jv //d7DVaaELv
	if err := client.Auth(); err != nil {
		return
	}
	r, err := client.AutoComplete("aa")
	if err != nil {
		return
	}
	println(r.Terms[0])

}
