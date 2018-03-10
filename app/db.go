package app

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	Config "github.com/velrino/BlockchainPayments/app/config"
)

func DatabaseInit() {
	Migrations()
}

func Migrations() {

	db := Config.Database()

	defer db.Close()
}

