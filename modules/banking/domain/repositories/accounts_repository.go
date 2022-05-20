package repositories

import (
	"github.com/truepay/go-boilerplate/modules/banking/domain/entities"
)

type AccountRepository interface {
	GetNextID() string
	Store(account *entities.Account) error
	FindByID(id string) (*entities.Account, error)
}
