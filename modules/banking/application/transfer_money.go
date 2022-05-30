package application

import (
	"github.com/truepay/go-boilerplate/modules/banking/domain/repositories"
	service "github.com/truepay/go-boilerplate/modules/banking/domain/services"
)

type TransferMoneyDTO struct {
	fromAccountId string
	toAccountId   string
	amount        int
}

func MakeTransferMoney(repository repositories.AccountRepository) func(props TransferMoneyDTO) error {
	return func(props TransferMoneyDTO) error {
		fromAccount, err := repository.FindByID(props.fromAccountId)
		if err != nil {
			return err
		}

		toAccount, err := repository.FindByID(props.toAccountId)
		if err != nil {
			return err
		}

		err = service.TransferMoney(service.TransferMoneyDTO{
			FromAccount: fromAccount,
			ToAccount:   toAccount,
			Amount:      props.amount,
		})
		if err != nil {
			return err
		}

		err = repository.Store(fromAccount)
		if err != nil {
			return err
		}

		err = repository.Store(toAccount)
		if err != nil {
			return err
		}

		return nil
	}
}
