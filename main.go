package main

import (
	App "github.com/velrino/BlockchainPayments/app"
)

func main() {
	App.DatabaseInit()
	App.Routes()
}