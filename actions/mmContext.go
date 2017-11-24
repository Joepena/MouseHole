package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type MMContext struct {
	buffalo.Context
}

func (c MMContext) Websocket() (*websocket.Conn, error) {
	// upgrade the conn to a socket
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.WithField("Function", "MMContext/Websocket").Error(err)
		return nil, err
	}

	return conn, nil
}
