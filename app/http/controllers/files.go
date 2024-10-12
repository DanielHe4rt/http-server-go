package controllers

import (
	"fmt"
	args2 "github.com/codecrafters-io/http-server-starter-go/app/http/args"
	"github.com/codecrafters-io/http-server-starter-go/app/http/request"
	"github.com/codecrafters-io/http-server-starter-go/app/http/response"
)

type FilesController struct{}

func (c FilesController) Handle(req request.Request, response response.Response) string {

	args := args2.GetArgs()
	fmt.Println("PQP AQUI:" + args.Directory)

	fileName := req.Params["fileName"]
	return response.Download(args.Directory, fileName).Build()
}
