package main

import (
	"github.com/jamesgotech/go-bank-backend/api"
	"github.com/jamesgotech/go-bank-backend/migrations"
)

func main() {
	migrations.MigrateTransactions()
	api.StartApi()
}
