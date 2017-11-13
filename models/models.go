package models

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"sync"
	"log"
	"reflect"
)

const MONGOD_URL = "localhost:27017"
var (
	once       sync.Once
	dbInstance DB
)

type DB struct {
	session *mgo.Session
}

type ReadRequest struct {
	DbName         string
	CollectionName string
	//DataType           reflect.Type
}

func (db *DB) Read(rR *ReadRequest) []Event {
	log.Println(rR)
	collection := db.session.DB(rR.DbName).C(rR.CollectionName) // maybe can cache this later?
	var result []Event
	err := collection.Find(nil).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func GetDBInstance() *DB {
	once.Do(func() {
		dbInstance = DB{
			session: establishMongoDBSession(),
		}
	})

	return &dbInstance
}

func establishMongoDBSession() *mgo.Session {
	session, err := mgo.Dial(MONGOD_URL)
	if err != nil {
		panic(fmt.Sprint("failed to establish session to %v. error: ", MONGOD_URL, err.Error()))
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

func retrieveTypeSlice(t reflect.Type, slice interface{}){
	if t.Name() == "Event" {
		slice = make([]Event, 10)
	}
}
