package repository

import (
	"github.com/pragma-collective/0xStarter-api/ent"
	"github.com/pragma-collective/0xStarter-api/internal/domain/dto"
)

type IUserRepository interface {
	Create(input dto.CreateUserInput) (ent.User, error)
	List() (ent.Users, error)
}
