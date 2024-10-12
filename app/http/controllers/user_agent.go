package controllers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/http/request"
	"github.com/codecrafters-io/http-server-starter-go/app/http/response"
)

type UserAgentController struct{}

func (c UserAgentController) Handle(req request.Request, response response.Response) string {
	return response.Text(req.Headers["User-Agent"]).Success().Build()
}
