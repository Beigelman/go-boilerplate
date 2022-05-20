package application

import (
	"github.com/truepay/go-boilerplate/modules/banking/domain/repositories"
)

type DepositMoneyDTO struct {
	accountId string
	amount    int
}

func MakeDepositMoney(repository repositories.AccountRepository) func(props DepositMoneyDTO) error {
	return func(props DepositMoneyDTO) error {
		account, err := repository.FindByID(props.accountId)
		if err != nil {
			return err
		}

		err = account.Deposit(props.amount)
		if err != nil {
			return err
		}

		err = repository.Store(account)
		if err != nil {
			return err
		}

		return nil
	}
}
