package api

import "log"

var badReqChan = make(chan Request, 10)

func processBadRequest(req Request) {
	log.Println("bad request was made. request: %v", req)
}
