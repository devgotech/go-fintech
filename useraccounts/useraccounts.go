package useraccounts

import (
	"github.com/jamesgotech/go-bank-backend/helpers"
	"github.com/jamesgotech/go-bank-backend/interfaces"
)

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)
	defer db.Close()
}
