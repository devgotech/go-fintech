package transactions

import (
	"github.com/jamesgotech/go-bank-backend/helpers"
	"github.com/jamesgotech/go-bank-backend/interfaces"
)

func CreateTransaction(From uint, To uint, Amount int) {
	db := helpers.ConnectDB()
	transaction := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	db.Create(&transaction)

	defer db.Close()
}
