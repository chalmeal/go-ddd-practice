package usecase

import (
	e "go-ddd-practice/pkg/errors"

	"go-ddd-practice/pkg/infrastracture/db"
	"go-ddd-practice/pkg/infrastracture/repository"
	"go-ddd-practice/pkg/model"
	"go-ddd-practice/pkg/usecase/dto"

	"gorm.io/gorm"
)

type MasterService struct {
	db      *gorm.DB
	rep     repository.AccountRepository
	account model.Accounts
}

func NewMasterService() *MasterService {
	return &MasterService{
		db:      db.GetDB(),
		rep:     repository.NewAccountRepository(),
		account: *model.NewAccount(),
	}
}

func (master *MasterService) GetAccount(id string) *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		ac, err := master.rep.Get(tx, id)
		if err != nil {
			return &dto.Dto{
				Error: e.GET_ACCOUNT_ACCOUNT_NOT_FOUND,
			}
		}

		return &dto.Dto{
			Result: &dto.GetAccountDto{
				Account: ac,
			},
			Error: nil,
		}
	})
}

func (master *MasterService) GetAccountAll() *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		ac, _ := master.rep.GetAll(tx)

		return &dto.Dto{
			Result: &dto.GetAccountAllDto{
				Accounts: ac,
			},
			Error: nil,
		}
	})
}

func (master *MasterService) RegisterAccount(param *model.RegAccount) *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		reg := master.account.Create(param)
		if err := master.rep.Save(tx, reg); err != nil {
			return &dto.Dto{
				Error: e.REGISTER_ACCOUNT_BAD_REQUEST,
			}
		}

		return &dto.Dto{
			Result: &dto.RegisterAccountDto{
				Message: e.REGISTER_ACCOUNT_SUCCESS.Error(),
			},
			Error: nil,
		}
	})
}

func (master *MasterService) EditAccount(param *model.EdtAccount) *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		chg := master.account.Edit(param)
		if err := master.rep.Update(tx, param.Id, chg); err != nil {
			return &dto.Dto{
				Error: e.EDIT_ACCOUNT_BAD_REQUEST,
			}
		}

		return &dto.Dto{
			Result: &dto.RegisterAccountDto{
				Message: e.EDIT_ACCOUNT_SUCCESS.Error(),
			},
			Error: nil,
		}
	})
}

func (master *MasterService) DeleteAccount(id string) *dto.Dto {
	return db.Tx(master.db, func(tx *gorm.DB) *dto.Dto {
		ac, err := master.rep.Get(tx, id)
		if err != nil {
			return &dto.Dto{
				Error: e.INTERNAL_SERVER_ERROR,
			}
		}
		err = master.account.VerifyDelete(ac)
		if err != nil {
			return &dto.Dto{
				Error: err,
			}
		}

		if err := master.rep.Delete(tx, id); err != nil {
			return &dto.Dto{
				Error: e.DELETE_ACCOUNT_BAD_REQUEST,
			}
		}

		return &dto.Dto{
			Result: &dto.DeleteAccountDto{
				Message: e.DELETE_ACCOUNT_SUCCESS.Error(),
			},
			Error: nil,
		}
	})
}
