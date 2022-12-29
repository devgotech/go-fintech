package useraccounts

import (
	"fmt"

	"github.com/jamesgotech/go-bank-backend/helpers"
	"github.com/jamesgotech/go-bank-backend/interfaces"
	"github.com/jamesgotech/go-bank-backend/transactions"
)

func updateAccount(id uint, amount int) interfaces.ResponseAccount {
	db := helpers.ConnectDB()
	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	db.Where("id = ? ", id).First(&account)
	account.Balance = uint(amount)
	db.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = int(account.Balance)
	defer db.Close()
	return responseAcc
}

func getAccount(id uint) *interfaces.Account {
	db := helpers.ConnectDB()
	account := &interfaces.Account{}
	if db.Where("id = ?", id).First(&account).RecordNotFound() {
		return nil
	}
	defer db.Close()
	return account
}

func Transaction(userId uint, from uint, to uint, amount int, jwt string) map[string]interface{} {
	userIdString := fmt.Sprint(userId)
	isValid := helpers.ValidateToken(userIdString, jwt)

	if isValid {
		fromAccount := getAccount(from)
		toAccount := getAccount(to)

		if fromAccount == nil || toAccount == nil {
			return map[string]interface{}{"message": "Account not found"}
		} else if fromAccount.UserID != userId {
			return map[string]interface{}{"message": "You are not owner of the account"}
		} else if int(fromAccount.Balance) < amount {
			return map[string]interface{}{"message": "Account balance is too small"}
		}
		updateAcct := updateAccount(from, int(fromAccount.Balance)-amount)
		updateAccount(to, int(fromAccount.Balance)+amount)

		transactions.CreateTransaction(from, to, amount)

		var response = map[string]interface{}{"message": "all is fine"}
		response["data"] = updateAcct
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
