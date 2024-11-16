package main

import (
	"github.com/garguelles/archpass/internal/adapter"
)

func main() {
	r := adapter.Router()

	go startIndexer()
	//go startWalletIndexer()
	//go startTicketIndexer()

	r.Logger.Fatal(r.Start(":8000"))
}
