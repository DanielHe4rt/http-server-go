package controllers

type MeController struct {
	Req Request
}

func (c MeController) Handle() string {
	return "HTTP/1.1 201 OK\r\n\r\n"
}
