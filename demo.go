package main

import (
	"github.com/pixfid/go-zaycevnet/api"
)

func main() {
	client := api.NewZClient("", "kmskoNkYHDnl3ol2") //60kQwLlpV3jv //d7DVaaELv
	client.Auth()
}
