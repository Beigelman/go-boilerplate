package application

import (
	"github.com/truepay/go-boilerplate/modules/banking/domain/entities"
	"github.com/truepay/go-boilerplate/modules/banking/domain/repositories"
)

type CreateAccountDTO struct {
	UserID string
}

type CreateAccount func(props CreateAccountDTO) (*entities.Account, error)

func MakeCreateAccount(repository repositories.AccountRepository) CreateAccount {
	return func(props CreateAccountDTO) (*entities.Account, error) {
		id := repository.GetNextID()
		account := entities.NewAccount(entities.AccountProps{
			ID:     id,
			UserID: props.UserID,
		})

		err := repository.Store(account)
		if err != nil {
			return nil, err
		}

		return account, nil
	}
}
