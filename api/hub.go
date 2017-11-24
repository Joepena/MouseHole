package api

type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	requestQueue chan Request

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		requestQueue: make(chan Request, 10),
		register:     make(chan *Client, 10),
		unregister:   make(chan *Client, 10),
		clients:      make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	go listenToChannels()
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
			}
		case req := <-h.requestQueue:
			go handleRequest(req)
		}
	}
}

// listens to all channels and routes to proper request handler
func listenToChannels() {
	for {
		select {
		case eventReq := <-eventChan:
			go processEventRequest(eventReq)
		case badReq := <-badReqChan:
			go processBadRequest(badReq)
		}
	}
}

// Handles incoming requests.
func handleRequest(req Request) {
	if value, ok := requestRouter.get(req.RequestType); ok {
		value <- req
	} else {
		badReqChan <- req
	}
}
