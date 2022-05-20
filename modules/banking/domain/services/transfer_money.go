package service

import "github.com/truepay/go-boilerplate/modules/banking/domain/entities"

type TransferMoneyDTO struct {
	FromAccount *entities.Account
	ToAccount   *entities.Account
	Amount      int
}

func TranferMoney(props TransferMoneyDTO) error {
	err := props.FromAccount.Withdraw(props.Amount)
	if err != nil {
		return err
	}

	err = props.ToAccount.Deposit(props.Amount)
	if err != nil {
		return err
	}

	return nil
}
