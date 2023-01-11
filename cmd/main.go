package main

import (
	"github.com/spriigan/RPCognito/infrastructure"
	"github.com/spriigan/RPCognito/infrastructure/router"
)

func main() {
	app := infrastructure.Application()

	if err := app.Serve(router.Router()); err != nil {
		panic(err)
	}
}
