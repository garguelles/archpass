package repository

import (
	"github.com/garguelles/archpass/internal/domain/dto"
)

type IAttendeeRepository interface {
	Create(input dto.CreateAttendeeInput) (int, error)
}
