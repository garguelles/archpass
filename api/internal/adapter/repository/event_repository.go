package repository

import (
	"context"
	"fmt"

	"github.com/garguelles/archpass/ent"
	"github.com/garguelles/archpass/ent/event"
	"github.com/garguelles/archpass/internal/adapter/database"
	"github.com/garguelles/archpass/internal/application/utils"
	"github.com/garguelles/archpass/internal/domain/dto"
	"github.com/garguelles/archpass/internal/domain/repository"
)

type EventRepository struct {
	client ent.Client
	ctx    *context.Context
}

func NewEventRepository(ctx *context.Context) repository.IEventRepository {
	return &EventRepository{
		client: *database.EntClient(),
		ctx:    ctx,
	}
}

func (e *EventRepository) generateUniqueSlug(name string) (string, error) {
	baseSlug := utils.Slugify(name)
	slug := baseSlug

	exists, err := e.client.Event.Query().
		Where(event.EventSlugEQ(slug)).
		Exist(*e.ctx)
	if err != nil {
		return "", err
	}

	// If the slug already exists, append a numeric suffix until it's unique
	i := 1
	for exists {
		slug = fmt.Sprintf("%s-%d", baseSlug, i)
		exists, err = e.client.Event.Query().
			Where(event.EventSlugEQ(slug)).
			Exist(*e.ctx)
		if err != nil {
			return "", err
		}
		i++
	}

	return slug, nil
}

func (e *EventRepository) Create(input dto.CreateEventInput, userId int) (ent.Event, error) {
	slug, err := e.generateUniqueSlug(input.Name)
	if err != nil {
		return ent.Event{}, nil
	}

	event, err := e.client.Event.
		Create().
		SetName(input.Name).
		SetEventSlug(slug).
		SetDate(input.Date).
		SetContractAddress(input.ContractAddress).
		// SetStartDate(input.StartDate).
		// SetEndDate(input.EndDate).
		SetDescription(input.Description).
		SetImageURL(*input.ImageUrl).
		Save(*e.ctx)
	if err != nil {
		return ent.Event{}, err
	}

	return *event, nil
}

func (e *EventRepository) ListByOrganizerId(limit int, offset int, userId int) (ent.Events, error) {
	events, err := e.client.Event.
		Query().
		Where(
			event.UserIDEQ(userId),
		).
		Limit(limit).
		Offset(offset).
		All(*e.ctx)

	if err != nil {
		return []*ent.Event{}, err
	}

	return events, nil
}

func (e *EventRepository) GetByIdAndOrganizerId(eventId int, userId int) (ent.Event, error) {
	event, err := e.client.Event.
		Query().
		Where(
			event.IDEQ(eventId),
			event.UserIDEQ(userId),
		).
		Only(*e.ctx)
	if err != nil {
		return ent.Event{}, err
	}

	return *event, nil
}

func (e *EventRepository) GetBySlug(eventSlug string) (ent.Event, error) {
	event, err := e.client.Event.
		Query().
		Where(event.EventSlugEQ(eventSlug)).
		WithTickets().
		WithUser().
		Only(*e.ctx)

	if err != nil {
		return ent.Event{}, err
	}

	return *event, nil
}
