package controllers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/http/request"
	"github.com/codecrafters-io/http-server-starter-go/app/http/response"
)

type EchoController struct{}

func (c EchoController) Handle(req request.Request, response response.Response) string {

	if compression, exists := req.Headers["Accept-Encoding"]; exists {
		return response.Compress(compression).Success().Build()
	}

	return response.Text(req.Params["message"]).Success().Build()
}
