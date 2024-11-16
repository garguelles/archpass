package main

import (
	"context"
	"fmt"
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
				fmt.Println("Error fetching logs: ", err)
				continue
			}

			for iterMint.Next() {
				event := iterMint.Event
				// saveMint(ctx, event)
				// updateProjectFundingGoal(ctx, event, client)
			}

			if iterMint.Error() != nil {
				fmt.Println("Error during iteration: ", iter.Error())
			}

			if iterMint.Event != nil {
				lastPolledBlock = uint64(iterMint.Event.Raw.BlockNumber) + 1
			}

			// Event created
			iterEventCreate, err := factoryContract.FilterEventCreated(filterOpts)
			if err != nil {
				fmt.Println("Error fetching logs: ", err)
				continue
			}

			for iterEventCreate.Next() {
				event := iterEventCreate.Event
				// saveMint(ctx, event)
				// updateProjectFundingGoal(ctx, event, client)
			}

			if iterEventCreate.Error() != nil {
				fmt.Println("Error during iteration: ", iter.Error())
			}

			if iterEventCreate.Event != nil {
				lastPolledBlock = uint64(iterMint.Event.Raw.BlockNumber) + 1
			}

			// Ticket created
			iterTicketCreate, err := factoryContract.Filter

		}
	}
}

// func updateProjectFundingGoal(ctx context.Context, event *contract.ContractTierMinted, client *ethclient.Client) {
// 	projectRepo := repository.NewProjectRepository(&ctx)
// 	projectAddress := event.ProjectAddress

// 	balance, err := client.BalanceAt(ctx, projectAddress, nil)
// 	if err != nil {
// 		fmt.Println("Unable to get balance for contract")
// 		return
// 	}

// 	ethValue := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
// 	ethStringValue := ethValue.Text('f', -1)

// 	fmt.Printf("Balance in ETH: %s\n", ethStringValue)

// 	projectId, err := projectRepo.UpdateCurrentFunding(projectAddress.Hex(), ethStringValue)
// 	if err != nil {
// 		fmt.Printf("Unable to update funding for project ID %d: %v\n", projectId, err.Error())
// 		return
// 	}
// }

// func saveMint(ctx context.Context, event *contract.ContractTierMinted) {
// 	mintRepo := repository.NewMintRepository(&ctx)
// 	tierRepo := repository.NewTierRepository(&ctx)
// 	userRepo := repository.NewUserRepository(&ctx)

// 	user, err := userRepo.FindByWalletAddress(event.Minter.Hex())
// 	if err != nil {
// 		fmt.Println("User not found")
// 		return
// 	}

// 	fmt.Println(event.TierAddress.Hex())
// 	tier, err := tierRepo.FindByContractAddress(event.TierAddress.Hex())
// 	if err != nil {
// 		fmt.Println("Tier not found")
// 		return
// 	}
// 	input := dto.CreateMintInput{
// 		BlockNumber:     int64(event.Raw.BlockNumber),
// 		TierId:          tier.ID,
// 		ProjectId:       tier.ProjectID,
// 		TransactionHash: event.Raw.TxHash.Hex(),
// 		TokenId:         int(event.TokenId.Int64()),
// 		Message:         &event.Message,
// 	}

// 	id, err := mintRepo.Create(input, user.ID)
// 	if err != nil {
// 		fmt.Println("Unable to save mint: ", err.Error())
// 	}

// 	fmt.Println("Successfully created mint record: ", id)
// }
