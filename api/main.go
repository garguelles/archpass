package main

import (
	"github.com/garguelles/archpass/internal/adapter"
)

func main() {
	r := adapter.Router()

	r.Logger.Fatal(r.Start(":8000"))
}
