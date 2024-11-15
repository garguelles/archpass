package repository

import (
	"context"

	"github.com/pragma-collective/0xStarter-api/ent"
	"github.com/pragma-collective/0xStarter-api/internal/adapter/database"
	"github.com/pragma-collective/0xStarter-api/internal/domain/dto"
	"github.com/pragma-collective/0xStarter-api/internal/domain/repository"
)

type UserRepository struct {
	client ent.Client
	ctx    *context.Context
}

func NewUserRepository(ctx *context.Context) repository.IUserRepository {
	return &UserRepository{
		client: *databasse.EntClient(),
		ctx:    ctx,
	}
}

func (u *UserRepository) Create(input dto.CreateUserInput) (ent.User, error) {
	result, err := u.client.User.
		Create().
		SetWalletAddress(input.WalletAddress).
		SetBio(input.Bio).
		Save(*u.ctx)

	if err != nil {
		return ent.User{}, err
	}

	return *result, nil
}

func (u *UserRepository) List() (ent.Users, error) {
	return u.client.User.Query().All(*u.ctx)
}
