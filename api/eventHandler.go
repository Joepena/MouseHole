package api

import "github.com/joepena/mouse_hole/models"

var eventChan requestChan = make(chan Request, 10)

type Event struct {
	Title   string
	Content string
	Tags    []string
}

func processEventRequest(req Request) {
	writeReq := &models.WriteRequest{
		req.DbName,
		req.RequestType.toString(),
		&Event{
			req.Title,
			req.Content,
			req.Tags,
		},
	}
	models.GetDBInstance().Write(writeReq)
}
