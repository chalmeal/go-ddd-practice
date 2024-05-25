package model

import (
	e "go-ddd-practice/pkg/errors"

	"go-ddd-practice/pkg/infrastracture/security"

	"gorm.io/gorm"
)

var (
	sec = security.NewSecurity()
)

type Accounts struct {
	gorm.Model
	Password    string `gorm:"type:varchar(500)" json:"password"`
	AccessToken string `gorm:"type:varchar(500)" json:"access_token"`
	Name        string `gorm:"type:varchar(50)" json:"name"`
	Email       string `gorm:"type:varchar(50)" json:"email"`
}

func NewAccount() *Accounts {
	return &Accounts{}
}

type VerLogin struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (account *Accounts) VerifyLogin(a *Accounts, p *VerLogin) bool {
	if a.Name == "" || a.Password != sec.Hash(p.Password) {
		return false
	}
	return true
}

type UpdAccessToken struct {
	AccessToken string `json:"access_token" binding:"required"`
}

func (account *Accounts) UpdateAccessToken(token string) *Accounts {
	return &Accounts{
		AccessToken: sec.Hash(token),
	}
}

func (account *Accounts) VerifyDelete(ac *Accounts) error {
	if ac.DeletedAt.Valid == false {
		return e.DELETE_ACCOUNT_DELETED
	}
	return nil
}

type RegAccount struct {
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func (account *Accounts) Create(p *RegAccount) *Accounts {
	return &Accounts{
		Password: sec.Hash(p.Password),
		Name:     p.Name,
		Email:    p.Email,
	}
}

type EdtAccount struct {
	Id    string `json:"id" binding:"required"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (account *Accounts) Edit(p *EdtAccount) *Accounts {
	return &Accounts{
		Name:  p.Name,
		Email: p.Email,
	}
}
