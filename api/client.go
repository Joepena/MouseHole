package api

import (
	"time"

	"encoding/json"
	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/joepena/mouse_hole/models"
)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn
	// dbName
	dbName string
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		var req Request

		err = json.Unmarshal(message, &req)
		if err != nil {
			log.Println(err)
		}
		log.Infof("ADDED DBNAME %v", c.dbName)
		req.DbName = c.dbName
		c.hub.requestQueue <- req
	}
}

// serveWs handles websocket requests from the peer.
func RegisterClient(hub *Hub, c buffalo.Context) {
	wsConn, err := c.Websocket()
	if err != nil {
		log.WithField("Function", "RegisterClient").Error(err)
	}
	uDbName := c.Data()["User"].(models.User).DBName
	client := &Client{hub: hub, conn: wsConn, dbName: uDbName}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.readPump()
}
