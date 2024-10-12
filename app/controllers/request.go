package controllers

import (
	"bytes"
	"fmt"
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
	headers map[string]string
}

func (r Request) GetKey() string {
	return fmt.Sprintf("%v+%v", r.Verb, r.Path)
}

func NewRequest(payload []byte) Request {

	payloadSlices := bytes.Split(payload, []byte("\r\n"))

	verb, path, version := extractRequestLine(payloadSlices[0])

	req := Request{
		Verb:    VerbType(verb),
		Version: version,
		Path:    path,
		headers: nil,
	}

	fmt.Println(req.Verb)
	fmt.Println(req.Path)
	fmt.Println(req.Version)

	return req
}

func extractRequestLine(requestLine []byte) (string, string, string) {

	splittedRequestLine := bytes.Split(requestLine, []byte(" "))
	verb := string(splittedRequestLine[0])
	path := string(splittedRequestLine[1])
	version := string(splittedRequestLine[2])
	return verb, path, version
}
