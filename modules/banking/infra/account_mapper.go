package infra

import (
	"github.com/google/uuid"
	"github.com/truepay/go-boilerplate/modules/banking/domain/entities"
)

type AccountMapper struct{}

func NewAccountMapper() *AccountMapper {
	return &AccountMapper{}
}

func (_ *AccountMapper) toData(entity *entities.Account) *Account {
	return &Account{
		ID:        uuid.MustParse(entity.ID),
		UserID:    uuid.MustParse(entity.UserID),
		Balance:   entity.Balance,
		Limit:     entity.Limit,
		Status:    entity.Status,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func (_ *AccountMapper) toEntity(data *Account) *entities.Account {
	return &entities.Account{
		ID:        data.ID.String(),
		UserID:    data.UserID.String(),
		Balance:   data.Balance,
		Limit:     data.Limit,
		Status:    data.Status,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
