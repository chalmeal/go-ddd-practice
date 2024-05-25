package repository

import (
	"go-ddd-practice/pkg/model"
	"log"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Get(tx *gorm.DB, id string) (*model.Accounts, error)
	GetAll(tx *gorm.DB) (*[]model.Accounts, error)
	Save(tx *gorm.DB, param *model.Accounts) error
	Update(tx *gorm.DB, id string, param *model.Accounts) error
	Delete(tx *gorm.DB, id string) error
}

type AccountRepositoryImpl struct{}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{}
}

func (rep *AccountRepositoryImpl) Get(tx *gorm.DB, id string) (*model.Accounts, error) {
	var account model.Accounts
	e := tx.Where("id=?", id).First(&account)
	if e.Error != nil {
		return &model.Accounts{}, e.Error
	}
	return &account, nil
}

func (rep *AccountRepositoryImpl) GetAll(tx *gorm.DB) (*[]model.Accounts, error) {
	var accounts []model.Accounts
	e := tx.Find(&accounts)
	if e.Error != nil {
		log.Println(e.Error)
		return &[]model.Accounts{}, e.Error
	}
	return &accounts, nil
}

func (rep *AccountRepositoryImpl) Save(tx *gorm.DB, param *model.Accounts) error {
	if err := tx.Create(param).Error; err != nil {
		return err
	}
	return nil
}

func (rep *AccountRepositoryImpl) Update(tx *gorm.DB, id string, param *model.Accounts) error {
	var account model.Accounts
	if err := tx.Model(&account).Where("id=?", id).Updates(param).Error; err != nil {
		return err
	}
	return nil
}

func (rep *AccountRepositoryImpl) Delete(tx *gorm.DB, id string) error {
	var account model.Accounts
	if err := tx.Where("id=?", id).Delete(&account).Error; err != nil {
		return err
	}
	return nil
}
