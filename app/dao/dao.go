package dao

import (
	. "github.com/velrino/BlockchainPayments/app/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DataBaseParams struct {
}

var db *mgo.Database
// Establish a connection to database
func (m *DataBaseParams) Connect() {
	session, err := mgo.Dial("mongodb://172.17.0.1:27018")
	if err != nil {
		panic("failed to connect database")
	}
	db = session.DB("blockchain")
}

func init() {
	var dao = DataBaseParams{}
	dao.Connect()
}
// Insert a database
func (m *DataBaseParams) Insert(collection string, model interface{}) (interface{}, error) {
	err := db.C(collection).Insert(&model)
	return model, err
}

// Find a transaction by its id
func (m *DataBaseParams) FindById(collection string, model interface{}, id string) (interface{}, error) {
	err := db.C(collection).FindId(bson.ObjectIdHex(id)).One(&model)
	return model, err
}

// Delete an existing transaction
func (m *DataBaseParams) Delete(collection string,transaction TransactionModel) error {
	err := db.C(collection).Remove(&transaction)
	return err
}

// Update an existing transaction
func (m *DataBaseParams) Update(collection string,transaction TransactionModel) error {
	err := db.C(collection).UpdateId(transaction.ID, &transaction)
	return err
}