package lingo

import base "github.com/lingo-dev/lingo-base"

type Request struct {
	input *base.Input
}

func NewRequest(input *base.Input) *Request {
	return &Request{input: input}
}

// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----

// Query get query parameter of the key in request
func (req *Request) Query(key string) string {
	return req.input.Query[key]
}

// QueryDefault get query parameter of the key with default value for backoff in request
func (req *Request) QueryDefault(key string, defVal string) string {
	if v, ok := req.input.Query[key]; ok {
		return v
	}

	return defVal
}

// Header get query header value of the key in request
func (req *Request) Header(key string) string {
	return req.input.Header[key]
}

// HeaderDefault get header value of the key with default value for backoff in request
func (req *Request) HeaderDefault(key string, defVal string) string {
	if v, ok := req.input.Header[key]; ok {
		return v
	}

	return defVal
}

// Body get request body
func (req *Request) Body() []byte {
	return req.input.Body
}
