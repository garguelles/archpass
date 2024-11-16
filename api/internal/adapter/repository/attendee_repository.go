package repository

import (
	"context"
	"github.com/garguelles/archpass/ent"
	"github.com/garguelles/archpass/ent/attendee"
	"github.com/garguelles/archpass/internal/adapter/database"
	"github.com/garguelles/archpass/internal/domain/dto"
	"github.com/garguelles/archpass/internal/domain/repository"
)

type AttendeeRepository struct {
	client ent.Client
	ctx    *context.Context
}

func NewAttendeeRepository(ctx *context.Context) repository.IAttendeeRepository {
	return &AttendeeRepository{
		client: *database.EntClient(),
		ctx:    ctx,
	}
}

func (a *AttendeeRepository) Create(input dto.CreateAttendeeInput) (int, error) {
	id, err := a.client.Attendee.
		Create().
		SetEventID(input.EventId).
		SetTicketID(input.TicketId).
		SetUserID(input.UserId).
		SetTokenID(input.TokenId).
		SetTransactionHash(input.TransactionHash).
		SetBlockNumber(input.BlockNumber).
		OnConflictColumns(
			attendee.FieldTicketID,
			attendee.FieldTokenID,
			attendee.FieldUserID,
		).
		UpdateNewValues().
		ID(*a.ctx)

	if err != nil {
		return 0, err
	}

	return id, nil
}
