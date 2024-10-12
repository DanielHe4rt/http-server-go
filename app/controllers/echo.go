package controllers

type EchoController struct {
	Req Request
}

func (c EchoController) Handle() string {
	return "HTTP/1.1 200 OK\r\n\r\n"
}
