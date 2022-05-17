package lingo

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	base "github.com/lingo-dev/lingo-base"
)

type Response struct {
	output *base.Output
}

func NewResponse(output *base.Output) *Response {
	return &Response{output: output}
}

// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----

// OK success return
func (resp *Response) OK() *Response {
	resp.output.StatusCode = http.StatusOK
	return resp
}

// Error error return
func (resp *Response) Error() *Response {
	resp.output.StatusCode = http.StatusInternalServerError
	return resp
}

// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----

// Status set the status code
func (resp *Response) Status(statusCode int) *Response {
	resp.output.StatusCode = statusCode

	return resp
}

// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----

// AddHeader add header
func (resp *Response) AddHeader(key, value string) *Response {
	resp.output.Header[key] = value

	return resp
}

// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----

// String set text format response body
func (resp *Response) String(data string) *Response {
	resp.output.Body = []byte(data)
	resp.output.Header[contentTypeHeaderKey] = plainContentType

	return resp
}

// Json set json format response body
func (resp *Response) Json(data any) *Response {
	resp.output.Body, resp.output.Err = json.Marshal(data)
	resp.output.Header[contentTypeHeaderKey] = jsonContentType

	return resp
}

// Xml set xml format response body
func (resp *Response) Xml(data any) *Response {
	resp.output.Body, resp.output.Err = xml.Marshal(data)
	resp.output.Header[contentTypeHeaderKey] = xmlContentType

	return resp
}

// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----
// ----- ----- ----- ----- ----- ----- ----- ----- ----- -----

const contentTypeHeaderKey = "Content-Type"

var (
	plainContentType = "text/plain; charset=utf-8"
	jsonContentType  = "application/json; charset=utf-8"
	xmlContentType   = "application/xml; charset=utf-8"
)
