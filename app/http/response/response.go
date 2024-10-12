package response

import (
	"fmt"
	"net/http"
	"strconv"
)

type Response struct {
	body       string
	status     int
	statusText string
	headers    map[string]string
}

func (res *Response) Text(payload string) *Response {
	res.headers = make(map[string]string)
	res.headers["Content-Type"] = "text/plain"
	res.headers["Content-Length"] = strconv.Itoa(len(payload))

	res.body = payload
	return res
}

func (res *Response) Success() *Response {
	res.status = http.StatusOK
	res.statusText = "OK"

	return res
}

func (res *Response) NotFound() *Response {
	res.status = http.StatusNotFound
	res.statusText = "Not Found"

	return res
}

func (res *Response) Build() string {
	var response string

	response += fmt.Sprintf("%v %v %v\r\n", "HTTP/1.1", res.status, res.statusText)

	for header, value := range res.headers {
		response += fmt.Sprintf("%v: %v\r\n", header, value)
	}
	// End of Headers
	response += "\r\n"
	if len(res.body) > 0 {
		response += res.body
	}

	fmt.Printf("Payload: %v\n", response)
	return response
}
