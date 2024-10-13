package controllers

import (
	"fmt"
	args2 "github.com/codecrafters-io/http-server-starter-go/app/http/args"
	"github.com/codecrafters-io/http-server-starter-go/app/http/request"
	"github.com/codecrafters-io/http-server-starter-go/app/http/response"
	"os"
)

type FileUploaderController struct{}

func (c FileUploaderController) Handle(req request.Request, response response.Response) string {

	args := args2.GetArgs()

	fileName := req.Params["fileName"]
	absolutePath := fmt.Sprintf("%v%v", args.Directory, fileName)
	fmt.Sprintf("%v", req.Body)

	err := os.WriteFile(absolutePath, []byte(req.Body), 0644)

	if err != nil {
		return response.InternalServerError().Text(err.Error()).Build()
	}

	return response.Created().Build()
}
