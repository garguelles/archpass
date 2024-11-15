package repository

import (
	"context"
	"fmt"

	"github.com/garguelles/archpass/ent"
	"github.com/garguelles/archpass/ent/attendee"
	"github.com/garguelles/archpass/ent/event"
	"github.com/garguelles/archpass/ent/ticket"
	"github.com/garguelles/archpass/ent/user"
	"github.com/garguelles/archpass/internal/adapter/database"
	"github.com/garguelles/archpass/internal/application/utils"
	"github.com/garguelles/archpass/internal/domain/dto"
	"github.com/garguelles/archpass/internal/domain/repository"
)

type TicketRepository struct {
	client ent.Client
	ctx    *context.Context
}

func NewTicketRepository(ctx *context.Context) repository.ITicketRepository {
	return &TicketRepository{
		client: *database.EntClient(),
		ctx:    ctx,
	}
}

func (t *TicketRepository) generateUniqueTicketSlug(name string) (string, error) {
	baseSlug := utils.Slugify(name)
	slug := baseSlug

	exists, err := t.client.Ticket.Query().
		Where(ticket.TicketSlugEQ(slug)).
		Exist(*t.ctx)
	if err != nil {
		return "", err
	}

	// If the slug already exists, append a numeric suffix until it's unique
	i := 1
	for exists {
		slug = fmt.Sprintf("%s-%d", baseSlug, i)
		exists, err = t.client.Ticket.Query().
			Where(ticket.TicketSlugEQ(slug)).
			Exist(*t.ctx)
		if err != nil {
			return "", err
		}
		i++
	}

	return slug, nil
}

func (t *TicketRepository) Create(input dto.CreateTicketInput, userId int) (ent.Ticket, error) {
	exists, err := t.client.Event.
		Query().
		Where(
			event.IDEQ(input.EventId),
			event.UserIDEQ(userId),
		).
		Exist(*t.ctx)
	if err != nil {
		return ent.Ticket{}, err
	}

	if !exists {
		return ent.Ticket{}, fmt.Errorf("Event does not belong to the user.")
	}

	slug, err := t.generateUniqueTicketSlug(input.Name)

	if err != nil {
		return ent.Ticket{}, nil
	}

	ticket, err := t.client.Ticket.
		Create().
		SetName(input.Name).
		SetTicketSlug(slug).
		SetDescription(input.Description).
		SetEventID(input.EventId).
		SetQuantity(input.Quantity).
		SetMintPrice(input.MintPrice).
		Save(*t.ctx)
	if err != nil {
		return ent.Ticket{}, err
	}

	return *ticket, nil
}

func (t *TicketRepository) ListByEventId(eventId int, userId int) (ent.Tickets, error) {
	exists, err := t.client.Event.
		Query().
		Where(
			event.IDEQ(eventId),
			event.UserIDEQ(userId),
		).
		Exist(*t.ctx)
	if err != nil {
		return []*ent.Ticket{}, err
	}

	if !exists {
		return []*ent.Ticket{}, fmt.Errorf("Tickets does not belong to the user.")
	}

	tickets, err := t.client.Ticket.
		Query().
		Where(
			ticket.EventIDEQ(eventId),
		).
		All(*t.ctx)

	if err != nil {
		return []*ent.Ticket{}, err
	}

	return tickets, nil
}

func (t *TicketRepository) GetBySlugAndEvent(eventSlug string, ticketSlug string) (ent.Ticket, error) {
	exists, err := t.client.Event.
		Query().
		Where(
			event.EventSlugEQ(eventSlug),
		).
		Exist(*t.ctx)
	if err != nil {
		return ent.Ticket{}, err
	}

	if !exists {
		return ent.Ticket{}, fmt.Errorf("Ticket does not belong to the event.")
	}

	ticket, err := t.client.Ticket.
		Query().
		Where(ticket.TicketSlugEQ(ticketSlug)).
		WithEvent().
		Only(*t.ctx)
	if err != nil {
		return ent.Ticket{}, err
	}

	return *ticket, nil
}

func (t *TicketRepository) GetByAttendee(walletAddress string) (ent.Attendees, error) {
	// Find the user by wallet address
	user, err := t.client.User.
		Query().
		Where(user.WalletAddressEQ(walletAddress)).
		Only(*t.ctx)
	if err != nil {
		return nil, err
	}

	attendees, err := t.client.Attendee.
		Query().
		Where(attendee.UserIDEQ(user.ID)).
		WithUser().
		WithTicket(func(tq *ent.TicketQuery) {
			tq.WithEvent()
		}).
		All(*t.ctx)
	if err != nil {
		return nil, err
	}

	return attendees, nil
}
