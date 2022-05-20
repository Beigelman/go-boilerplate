package application

import (
	"github.com/truepay/go-boilerplate/modules/banking/domain/repositories"
)

type WithdrawDTO struct {
	accountId string
	amount    int
}

func MakeWithdrawMoney(repository repositories.AccountRepository) func(props WithdrawDTO) error {
	return func(props WithdrawDTO) error {
		account, err := repository.FindByID(props.accountId)
		if err != nil {
			return err
		}

		err = account.Withdraw(props.amount)
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
