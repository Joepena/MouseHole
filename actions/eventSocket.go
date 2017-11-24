package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/joepena/mouse_hole/models"
	"github.com/sirupsen/logrus"
)

const EVENTS_COLLECTION = "events"

func eventSocketHandler(c buffalo.Context) error {
	// get readReq from auth
	// place this as an auth middleware. So req -> auth -> app DBName attached to the Context object
	wsConn, err := c.Websocket()
	if err != nil {
		logrus.WithField("Function", "eventSocketHandler").Error(err)
		return err
	}
	readReq := c.Data()[READ_REQUEST].(models.ReadRequest)
	readReq.SetCollection(EVENTS_COLLECTION)
	go readReq.SubscribeToIterator(wsConn) //add read request

	return nil
}
