package main

import (
	"context"
	"fmt"
	"github.com/garguelles/archpass/internal/adapter/repository"
	"github.com/garguelles/archpass/internal/domain/dto"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/garguelles/archpass/internal/application/contract"
	"github.com/joho/godotenv"
)

func startIndexer() {
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

			// Ticket minted
			iterMint, err := factoryContract.FilterTicketMinted(filterOpts, nil, nil, nil)
			if err != nil {
				fmt.Println("Error fetching logs ticket minted: ", err)
				continue
			}

			for iterMint.Next() {
				event := iterMint.Event
				saveMint(ctx, event)
			}

			if iterMint.Error() != nil {
				fmt.Println("Error during iteration: ", iterMint.Error())
			}

			if iterMint.Event != nil {
				lastPolledBlock = uint64(iterMint.Event.Raw.BlockNumber) + 1
			}
		}
	}
}

func saveMint(ctx context.Context, event *contract.EventFactoryTicketMinted) {
	userRepo := repository.NewUserRepository(&ctx)
	ticketRepo := repository.NewTicketRepository(&ctx)
	attendeeRepo := repository.NewAttendeeRepository(&ctx)

	user, err := userRepo.FindByWalletAddress(event.Minter.Hex())
	if err != nil {
		fmt.Println("User not found")
		return
	}

	ticket, err := ticketRepo.GetByContractAddress(event.TicketAddress.Hex())
	if err != nil {
		fmt.Println("Ticket not found")
		return
	}
	input := dto.CreateAttendeeInput{
		BlockNumber:     int64(event.Raw.BlockNumber),
		TicketId:        ticket.ID,
		EventId:         ticket.EventID,
		UserId:          user.ID,
		TokenId:         int(event.TokenId.Int64()),
		TransactionHash: event.Raw.TxHash.Hex(),
	}

	id, err := attendeeRepo.Create(input)
	if err != nil {
		fmt.Println("Unable to save mint: ", err.Error())
		return
	}

	fmt.Println("Successfully created attendee record: ", id)
}
