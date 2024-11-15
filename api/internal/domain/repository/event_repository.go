package repository

import (
	"github.com/garguelles/archpass/ent"
	"github.com/garguelles/archpass/internal/domain/dto"
)

type IEventRepository interface {
	Create(input dto.CreateEventInput, userId int) (ent.Event, error)
	ListByOrganizerId(limit int, offset int, userId int) (ent.Events, error)
	GetByIdAndOrganizerId(eventId int, userId int) (ent.Event, error)
}
