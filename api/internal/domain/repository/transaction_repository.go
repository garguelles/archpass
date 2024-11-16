package repository

import (
	"github.com/garguelles/archpass/internal/domain/dto"
)

type ITransactionRepository interface {
	Create(input dto.CreateTransactionInput) (int, error)
}
