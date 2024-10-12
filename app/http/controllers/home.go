package controllers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/http/request"
	"github.com/codecrafters-io/http-server-starter-go/app/http/response"
)

type HomeController struct{}

func (c HomeController) Handle(_ request.Request, response response.Response) string {
	return response.Success().Build()
}
