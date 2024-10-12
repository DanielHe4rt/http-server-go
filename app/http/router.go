package http

import (
	"errors"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/controllers"
	"regexp"
	"strings"
)

type RequestResponse interface {
	Handle(r controllers.Request) string
}

func ProcessRequest(r controllers.Request) string {

	// Available Routes

	response, err := getController(r)

	if err != nil {
		return "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	return response
}

func availableRoutes(r controllers.Request) map[string]string {
	router := make(map[string]string)
	router["GET+/"] = controllers.HomeController{Req: r}.Handle()
	router["GET+/echo/danielhe4rt"] = controllers.MeController{Req: r}.Handle()
	router["GET+/echo/{message}"] = controllers.EchoController{Req: r}.Handle()

	return router
}

func getController(r controllers.Request) (string, error) {
	routeList := availableRoutes(r)

	routerPathCounter := 0

	currentRequestPathStructure := strings.Split(r.Path, "/")
	for route, response := range routeList {

		fmt.Println("------------------")
		availableRoute := strings.Split(route, "+")
		availablePath := availableRoute[1]
		fmt.Printf("Route: %v\n", availablePath)
		availablePathStructure := strings.Split(availablePath, "/")

		if len(availablePathStructure) != len(currentRequestPathStructure) {
			fmt.Printf("Route '%v' doesnt match %v != %v\n", availablePath, len(availablePathStructure), len(currentRequestPathStructure))
			continue
		}

		for idx, pathItem := range availablePathStructure {

			if pathItem == currentRequestPathStructure[idx] {
				routerPathCounter++
			}

			fmt.Println(pathItem)
			fmt.Printf("Matching router: [%v] with Req: [%v]\n", pathItem, currentRequestPathStructure[idx])

			hasVariable, err := regexp.MatchString("{(.*)}", pathItem)

			if err != nil {
				fmt.Printf("Regex Error: %v\n", err.Error())
				continue
			}

			if hasVariable {
				fmt.Printf("Found Regex: %v\n", hasVariable)
				routerPathCounter++
			}

			fmt.Printf("Arguments Counter: %v\n", routerPathCounter)
		}

		if routerPathCounter == len(availablePathStructure) {
			return response, nil
		}
		routerPathCounter = 0

	}

	return "", errors.New("Not Found :x")
}
