package repository

import (
	"github.com/garguelles/archpass/ent"
	"github.com/garguelles/archpass/internal/domain/dto"
)

type ITicketRepository interface {
	Create(input dto.CreateTicketInput, userId int) (ent.Ticket, error)
	ListByEventId(eventId int, userId int) (ent.Tickets, error)
	GetBySlugAndEvent(eventSlug string, ticketSlug string) (ent.Ticket, error)
	GetByAttendee(walletAddress string) (ent.Attendees, error)
}
