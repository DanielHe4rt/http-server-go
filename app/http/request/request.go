package request

import (
	"bytes"
	"fmt"
	"strings"
)

type VerbType string

const (
	VerbGet  VerbType = "GET"
	VerbPost VerbType = "POST"
)

type Request struct {
	Verb    VerbType
	Version string
	Path    string
	Headers map[string]string
	Params  map[string]string
	Body    string
}

func (r Request) GetKey() string {
	return fmt.Sprintf("%v+%v", r.Verb, r.Path)
}

func (r Request) GetPathSlices() []string {
	return strings.Split(r.Path, "/")
}

func NewRequest(payload []byte) Request {

	payloadSlices := bytes.Split(payload, []byte("\r\n"))

	verb, path, version := extractRequestLine(payloadSlices[0])
	headers := extractHeaders(payloadSlices[1 : len(payloadSlices)-2])
	body := payloadSlices[len(payloadSlices)-1]
	req := Request{
		Verb:    VerbType(verb),
		Version: version,
		Path:    path,
		Headers: headers,
		Params:  map[string]string{},
		Body:    string(body),
	}

	fmt.Println(req.Verb)
	fmt.Println(req.Path)
	fmt.Println(req.Version)

	return req
}

func extractHeaders(rawHeaders [][]byte) map[string]string {
	headers := make(map[string]string)

	for _, headerBytes := range rawHeaders {
		fmt.Printf("%v\n", string(headerBytes))
		data := bytes.SplitN(headerBytes, []byte(": "), 2)

		key, value := string(data[0]), string(data[1]) // Accept (?)
		headers[key] = value
	}

	return headers
}

func extractRequestLine(requestLine []byte) (string, string, string) {

	splittedRequestLine := bytes.Split(requestLine, []byte(" "))
	verb := string(splittedRequestLine[0])
	path := string(splittedRequestLine[1])
	version := string(splittedRequestLine[2])
	return verb, path, version
}
