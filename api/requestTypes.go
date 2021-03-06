package api

const (
	EventType requestType = "event"
)

var requestRouter = requestRouterType{
	typeToRequestChan: map[requestType](requestChan){
		EventType: eventChan,
	},
}

type requestType string

type requestChan chan Request

type requestRouterType struct {
	typeToRequestChan map[requestType](requestChan)
}

type Request struct {
	ApplicationName string      `json:"ApplicationName"`
	Title           string      `json:"Title"`
	Content         string      `json:"Content"`
	Tags            []string    `json:"Tags"`
	RequestType     requestType `json:"ResponseType"`
}

func (r *requestRouterType) get(key requestType) (requestChan, bool) {
	val, ok := r.typeToRequestChan[key]
	return val, ok
}

func (rT requestType) toString() string {
	return string(rT)
}
