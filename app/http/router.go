package http

import (
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/http/controllers"
	"github.com/codecrafters-io/http-server-starter-go/app/http/request"
	res "github.com/codecrafters-io/http-server-starter-go/app/http/response"
	"regexp"
	"strings"
)

type RequestResponse interface {
	Handle(r request.Request) string
}

func ProcessRequest(r request.Request) string {

	actionKey, r := getController(r)
	actionHandler := getRouteAction(actionKey)

	return actionHandler.Handle(r, res.Response{})
}

func availableRoutes() map[string]string {
	router := make(map[string]string)
	router["GET+/"] = "HomeController"
	router["GET+/echo/danielhe4rt"] = "MeController"
	router["GET+/echo/{message}"] = "EchoController"
	router["GET+/user-agent"] = "UserAgentController"
	router["GET+/files/{fileName}"] = "FilesController"
	router["POST+/files/{fileName}"] = "FileUploaderController"

	return router
}

func getRouteAction(controller string) controllers.BaseController {
	switch controller {
	case "HomeController":
		return controllers.HomeController{}
	case "MeController":
		return controllers.MeController{}
	case "EchoController":
		return controllers.EchoController{}
	case "UserAgentController":
		return controllers.UserAgentController{}
	case "FilesController":
		return controllers.FilesController{}
	case "FileUploaderController":
		return controllers.FileUploaderController{}
	default:
		return controllers.NotFoundController{}
	}
}

func getController(r request.Request) (string, request.Request) {
	routeList := availableRoutes()

	routerPathCounter := 0

	//fmt.Printf("Current Route: %v %v\n", r.Verb, r.Path)
	currentRequestPathStructure := strings.Split(r.Path, "/")
	for route, response := range routeList {
		availableRoute := strings.Split(route, "+")

		requestVerb := request.VerbType(availableRoute[0])
		if requestVerb != r.Verb {
			continue
		}

		availablePath := availableRoute[1]
		//fmt.Printf("Tested Path: %v\n", availablePath)
		availablePathStructure := strings.Split(availablePath, "/")

		if len(availablePathStructure) != len(currentRequestPathStructure) {
			//fmt.Printf("Route '%v' doesnt match %v != %v\n", availablePath, len(availablePathStructure), len(currentRequestPathStructure))
			continue
		}
		fmt.Printf("Args count: %v\n", routerPathCounter)
		for idx, pathItem := range availablePathStructure {

			if pathItem == currentRequestPathStructure[idx] {
				routerPathCounter++
				continue
			}

			fmt.Println(pathItem)
			//fmt.Printf("Matching router: [%v] with Req: [%v]\n", pathItem, currentRequestPathStructure[idx])

			hasVariable, err := regexp.MatchString("{(.*)}", pathItem)

			if err != nil {
				//fmt.Printf("Regex Error: %v\n", err.Error())
				continue
			}

			if hasVariable {
				//fmt.Printf("Found Regex: %v\n", hasVariable)
				routerPathCounter++
				pathKey := pathItem[1 : len(pathItem)-1]
				fmt.Println(pathKey, currentRequestPathStructure[idx])

				r.Params[pathKey] = currentRequestPathStructure[idx]
			}

			//fmt.Printf("Arguments Counter: %v\n", routerPathCounter)
		}

		if routerPathCounter == len(availablePathStructure) {
			return response, r
		}
		routerPathCounter = 0
		r.Params = make(map[string]string)
	}

	return "NotFoundController", r
}
