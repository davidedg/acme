package interaction

import "fmt"

type responder struct{}

var responses = map[string]*Response{}

// Auto-responder.
var Responder Interactor = responder{}

func (responder) Status(c *StatusInfo) (StatusSink, error) {
	return nil, fmt.Errorf("not supported")
}

func (responder) Prompt(c *Challenge) (*Response, error) {
	if c.UniqueID == "" {
		return nil, fmt.Errorf("cannot auto-respond to a challenge without a unique ID")
	}

	res := responses[c.UniqueID]
	if res == nil {
		return nil, fmt.Errorf("unknown unique ID, cannot respond: %#v", c.UniqueID)
	}

	return res, nil
}

func SetResponse(uniqueID string, res *Response) {
	responses[uniqueID] = res
}

// © 2015 Hugo Landau <hlandau@devever.net>    MIT License
