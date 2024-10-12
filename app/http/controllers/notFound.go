package controllers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/http/request"
	"github.com/codecrafters-io/http-server-starter-go/app/http/response"
)

type NotFoundController struct{}

func (c NotFoundController) Handle(_ request.Request, response response.Response) string {
	return response.Text("sai daqui porra").NotFound().Build()
}
