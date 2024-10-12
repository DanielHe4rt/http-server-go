package controllers

type HomeController struct {
	Req Request
}

func (c HomeController) Handle() string {
	return "HTTP/1.1 200 OK\r\n\r\n"
}
