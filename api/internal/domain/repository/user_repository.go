package repository

import (
	"github.com/garguelles/archpass/ent"
	"github.com/garguelles/archpass/internal/domain/dto"
)

type IUserRepository interface {
	Create(input dto.CreateUserInput) (ent.User, error)
	List() (ent.Users, error)
	FindByWalletAddress(walletAddress string) (ent.User, error)
}
