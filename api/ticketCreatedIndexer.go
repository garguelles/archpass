package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/garguelles/archpass/internal/adapter/repository"
	"github.com/garguelles/archpass/internal/application/contract"
	"github.com/garguelles/archpass/internal/domain/dto"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

func startTicketIndexer() {
	godotenv.Load(".env")

	chainUrl, ok := os.LookupEnv("CHAIN_URL")
	if !ok {
		log.Fatal("Cannot get url for eth-client")
	}

	contractAddress, ok := os.LookupEnv("EVENT_FACTORY_ADDRESS")
	if !ok {
		log.Fatal("Cannot get contract address for eth-client")
	}

	lastPolledBlockEnv, ok := os.LookupEnv("LAST_POLLED_BLOCK")
	if !ok {
		log.Fatal("Cannot get last polled block in environment variables.")
	}

	ctx := context.Background()
	fmt.Println("Starting indexer")
	client, err := ethclient.Dial(chainUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	factoryContract, err := contract.NewEventFactory(common.HexToAddress(contractAddress), client)
	if err != nil {
		fmt.Println(err)
	}

	lastPolledBlock, err := strconv.ParseUint(lastPolledBlockEnv, 10, 64)
	if err != nil {
		log.Fatal("Failed to convert LAST_POLLED_BLOCK to a big.Int")
	}

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Indexing new blocks")
			filterOpts := &bind.FilterOpts{Start: uint64(lastPolledBlock), End: nil, Context: ctx}

			// Ticket created
			iterTicketCreate, err := factoryContract.FilterTicketCreated(filterOpts)
			if err != nil {
				fmt.Println("Error fetching logs for ticket created: ", err)
				continue
			}

			for iterTicketCreate.Next() {
				event := iterTicketCreate.Event
				saveTicketCreated(ctx, event)
			}

			if iterTicketCreate.Error() != nil {
				fmt.Println("Error during iteration: ", iterTicketCreate.Error())
			}

			if iterTicketCreate.Event != nil {
				lastPolledBlock = uint64(iterTicketCreate.Event.Raw.BlockNumber) + 1
			}
		}
	}
}

func saveTicketCreated(ctx context.Context, evt *contract.EventFactoryTicketCreated) {
	transactionRepo := repository.NewTransactionRepository(&ctx)
	input := dto.CreateTransactionInput{
		TransactionHash: evt.Raw.TxHash.Hex(),
		BlockNumber:     int64(evt.Raw.BlockNumber),
		WalletAddress:   evt.Sender.Hex(),
		EventType:       string("TicketCreated"),
	}

	id, err := transactionRepo.Create(input)
	if err != nil {
		fmt.Println("Unable to save transaction for wallet created: ", err.Error())
		return
	}

	fmt.Println("Successfully updated event: ", id)
}
