package infra

import (
	"github.com/google/uuid"
	"github.com/truepay/go-boilerplate/modules/banking/domain/entities"
	"gorm.io/gorm"
)

type PostgresAccountRepository struct {
	db *gorm.DB
}

func NewPostgresAccountRepository(db *gorm.DB) *PostgresAccountRepository {
	return &PostgresAccountRepository{
		db: db,
	}
}

func (repository *PostgresAccountRepository) GetNextID() string {
	return uuid.NewString()
}

func (repository *PostgresAccountRepository) Store(account *entities.Account) error {
	mapper := NewAccountMapper()
	dbAccount := mapper.toData(account)

	var count int64
	err := repository.db.Find(&dbAccount, account.ID).Count(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return repository.db.Save(&dbAccount).Error
	}

	err = repository.db.Create(&dbAccount).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *PostgresAccountRepository) FindByID(id string) (*entities.Account, error) {
	mapper := NewAccountMapper()

	dbAccount := Account{}

	err := repository.db.First(&dbAccount, id).Error
	if err != nil {
		return nil, err
	}

	return mapper.toEntity(&dbAccount), nil
}
