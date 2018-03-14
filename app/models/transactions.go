package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TransactionModel struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Status   string        `bson:"status" json:"status"`
	Coins  float64        `bson:"coins" json:"coins"`
	Amount  string        `bson:"amount" json:"amount"`
	Blockchain interface{}        `bson:"blockchain" json:"blockchain"`
	Currency string        `bson:"currency" json:"currency"`
	User TransactionUserModel        `bson:"user" json:"user"`
	CreatedAt   time.Time        `bson:"created_at" json:"created_at"`
}

type TransactionUserModel struct {
	Email  string        `bson:"email" json:"email"`
	Document	string        `bson:"document" json:"document"`
	Bank TransactionBankModel        `bson:"bank" json:"bank"`
}

type TransactionBankModel struct {
	Agency  string        `bson:"agency" json:"agency"`
	Bank   	string        `bson:"bank" json:"bank"`
	Number  string        `bson:"number" json:"number"`
	Type   	string        `bson:"type" json:"type"`
}

