package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TransactionModel struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Email   string        `bson:"email" json:"email"`
	Status   string        `bson:"status" json:"status"`
	Amount  string        `bson:"amount" json:"amount"`
	Address string        `bson:"address" json:"address"`
	Blockchain interface{}        `bson:"blockchain" json:"blockchain"`
	Document string        `bson:"document" json:"document"`
	Currency string        `bson:"currency" json:"currency"`
	Bank TransactionBankModel        `bson:"bank" json:"bank"`
	CreatedAt   time.Time        `bson:"created_at" json:"created_at"`
}

type TransactionBankModel struct {
	Agency  string        `bson:"agency" json:"agency"`
	Bank   	string        `bson:"bank" json:"bank"`
	Number  string        `bson:"number" json:"number"`
	Type   	string        `bson:"type" json:"type"`
}

