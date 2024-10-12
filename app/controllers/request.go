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

func NewRequest(payload [][]byte) Request {

	verb, path, version := extractRequestLine(payload[0])
	for i := range payload {
		fmt.Printf("Content Parsed: %v \n", string(payload[i]))
	}
	req := Request{
		Verb:    VerbType(verb),
		Version: version,
		Path:    path,
		headers: nil,
	}

	return req
}

func extractRequestLine(requestLine []byte) (string, string, string) {

	splittedRequestLine := bytes.Split(requestLine, []byte(" "))
	verb := string(splittedRequestLine[0])
	path := string(splittedRequestLine[1])
	version := string(splittedRequestLine[2])
	return verb, path, version
}
