package controllers

import (
	"github.com/codecrafters-io/http-server-starter-go/app/http/request"
	"github.com/codecrafters-io/http-server-starter-go/app/http/response"
)

type MeController struct{}

func (c MeController) Handle(_ request.Request, response response.Response) string {
	return response.Text("danielzin apenas sit down be humble").NotFound().Build()
}
