package models

import (
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	Id      bson.ObjectId `bson:"_id"`
	Title   string
	Content string
	Tags    []string
}

func (rR *ReadRequest) SubscribeToIterator(conn *websocket.Conn) {
	db := GetDBInstance()
	collection := db.session.DB(rR.dBName).C(rR.collectionName)
	iter := collection.Find(nil).Sort("$natural").Tail(-1)
	var result Event
	var lastId bson.ObjectId
	for {
		for iter.Next(&result) {
			conn.WriteJSON(result)
			lastId = result.Id
		}
		if iter.Err() != nil {
			rR.logger.WithField("Function", "SubscribeToIterator").Error(iter.Close())
		}
		if iter.Timeout() {
			continue
		}
		query := collection.Find(bson.M{"_id": bson.M{"$gt": lastId}})
		iter = query.Sort("$natural").Tail(-1)
	}
	iter.Close()
}
