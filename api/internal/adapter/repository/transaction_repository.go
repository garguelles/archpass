package repository

import (
	"context"
	"github.com/garguelles/archpass/ent/transaction"

	"github.com/garguelles/archpass/ent"
	"github.com/garguelles/archpass/internal/adapter/database"
	"github.com/garguelles/archpass/internal/domain/dto"
	"github.com/garguelles/archpass/internal/domain/repository"
)

type TransactionRepository struct {
	client ent.Client
	ctx    *context.Context
}

func NewTransactionRepository(ctx *context.Context) repository.ITransactionRepository {
	return &TransactionRepository{
		client: *database.EntClient(),
		ctx:    ctx,
	}
}

func (t *TransactionRepository) Create(input dto.CreateTransactionInput) (int, error) {
	id, err := t.client.Transaction.
		Create().
		SetWalletAddress(input.WalletAddress).
		SetEventType(input.EventType).
		SetTransactionHash(input.TransactionHash).
		SetBlockNumber(input.BlockNumber).
		OnConflictColumns(
			transaction.FieldEventType,
			transaction.FieldWalletAddress,
			transaction.FieldTransactionHash,
		).
		UpdateNewValues().
		ID(*t.ctx)
	if err != nil {
		return 0, err
	}

	return id, nil
}
