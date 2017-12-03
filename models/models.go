package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"sync"
	"gopkg.in/mgo.v2/bson"
)

const MONGOD_URL = "localhost:27017"

var (
	once       sync.Once
	dbInstance DB
)

type DB struct {
	session *mgo.Session
	logger  *log.Entry
}

type ReadRequest struct {
	dBName         string
	collectionName string
	logger         *log.Entry
}

type WriteRequest struct {
	DbName         string
	CollectionName string
	Data           interface{}
}

func (db *DB) Write(wR *WriteRequest) {
	collection := db.session.DB(wR.DbName).C(wR.CollectionName) // maybe can cache this later?
	err := collection.Insert(wR.Data)
	if err != nil {
		log.Fatal("failed to insert data into db. dbName: %v, collectionName: %v, data: %v, error: %v", wR.DbName, wR.CollectionName, wR.Data, err)
	}
}

func (db *DB) CreateCappedCollection(dbName string, collectionName string, capacity int) {
	collectionInfo := &mgo.CollectionInfo{
		Capped:   true,
		MaxBytes: capacity,
	}
	err := db.session.DB(dbName).C(collectionName).Create(collectionInfo)
	if err != nil {
		log.WithField("Function", "CreateCappedCollection").Error(err)
	}
}

func (db *DB) GetUser(id string) (User, error){
	collection := db.session.DB("auth").C("users")

	var user User
	err := collection.Find(bson.M{"_id": id}).One(&user)

	return user, err
}

func GetDBInstance() *DB {
	once.Do(func() {
		dbInstance = DB{
			session: establishMongoDBSession(),
			logger:  log.WithField("Component", "DB"),
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

func NewReadRequest() ReadRequest {
	return ReadRequest{
		logger: log.WithField("Component", "readRequest"),
	}
}

func (rR *ReadRequest) SetDB(dBName string) {
	rR.dBName = dBName
}

func (rR *ReadRequest) DBName() string {
	return rR.dBName
}

func (rR *ReadRequest) SetCollection(collectionName string) {
	rR.collectionName = collectionName
}

func (rR *ReadRequest) CollectionName() string {
	return rR.collectionName
}
