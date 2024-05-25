package dto

import "go-ddd-practice/pkg/model"

type GetAccountDto struct {
	Account *model.Accounts
}

type GetAccountAllDto struct {
	Accounts *[]model.Accounts
}

type RegisterAccountDto struct {
	Message string
}

type EditAccountDto struct {
	Message string
}

type DeleteAccountDto struct {
	Message string
}
