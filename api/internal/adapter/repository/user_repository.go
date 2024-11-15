package repository

import (
	"context"

	"github.com/garguelles/archpass/ent"
	"github.com/garguelles/archpass/ent/user"
	databasse "github.com/garguelles/archpass/internal/adapter/database"
	"github.com/garguelles/archpass/internal/domain/dto"
	"github.com/garguelles/archpass/internal/domain/repository"
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

func (u *UserRepository) FindByWalletAddress(walletAddress string) (ent.User, error) {
	user, err := u.client.
		User.Query().
		Where(user.WalletAddress(walletAddress)).
		Only(*u.ctx)

	if err != nil {
		return ent.User{}, err
	}

	return *user, nil
}
