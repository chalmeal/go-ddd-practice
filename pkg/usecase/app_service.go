package usecase

import (
	e "go-ddd-practice/pkg/errors"
	"go-ddd-practice/pkg/infrastracture/db"
	"go-ddd-practice/pkg/infrastracture/repository"
	"go-ddd-practice/pkg/model"
	"go-ddd-practice/pkg/usecase/dto"

	"gorm.io/gorm"
)

type AppService struct {
	db      *gorm.DB
	rep     repository.AccountRepository
	account model.Accounts
}

func NewAppService() *AppService {
	return &AppService{
		db:      db.GetDB(),
		rep:     repository.NewAccountRepository(),
		account: *model.NewAccount(),
	}
}

func (app *AppService) Login(param *model.VerLogin) *dto.Dto {
	return db.Tx(app.db, func(tx *gorm.DB) *dto.Dto {
		a, err := app.rep.Get(tx, param.Id)
		if err != nil {
			return &dto.Dto{
				Error: err,
			}
		}
		if !app.account.VerifyLogin(a, param) {
			return &dto.Dto{
				Error: e.LOGIN_FAILURE,
			}
		}

		// 適当なアクセストークンをとりあえず付与
		token := "header.payload.signature"

		p := app.account.UpdateAccessToken(token)
		if err := app.rep.Update(tx, param.Id, p); err != nil {
			return &dto.Dto{
				Error: e.LOGIN_FAILURE_UPDATE_TOKEN,
			}
		}

		return &dto.Dto{
			Result: &dto.LoginDto{
				AccessToken: token,
			},
			Error: nil,
		}
	})
}
