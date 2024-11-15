package main

import (
	"github.com/pragma-collective/0xStarter-api/internal/adapter"
)

func main() {
	r := adapter.Router()

	r.Logger.Fatal(r.Start(":8000"))
}
