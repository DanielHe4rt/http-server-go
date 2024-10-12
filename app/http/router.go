package http

import "github.com/codecrafters-io/http-server-starter-go/app/controllers"

type RequestResponse interface {
	Handle(r controllers.Request) string
}

func ProcessRequest(r controllers.Request) string {

	// Available Routes
	router := make(map[string]string)
	router["GET+/"] = controllers.HomeController{Req: r}.Handle()

	key := r.GetKey()

	value, containsRoute := router[key]

	if containsRoute {
		return value
	}

	return "HTTP/1.1 404 Not Found\r\n\r\n"
}
