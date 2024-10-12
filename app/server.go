package main

import (
	"bytes"
	"fmt"
	"github.com/codecrafters-io/http-server-starter-go/app/controllers"
	"github.com/codecrafters-io/http-server-starter-go/app/http"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	buffer := make([]byte, 1024)

	bytes_readed, err := conn.Read(buffer)

	payload := make([]byte, bytes_readed)
	copy(payload, buffer)

	content := bytes.Split(payload, []byte("\r\n"))

	req := controllers.NewRequest(content)

	res := http.ProcessRequest(req)

	fmt.Printf("Verb: %v\n", req.Verb)
	fmt.Printf("Version: %v\n", req.Version)
	fmt.Printf("Path: %v\n", req.Path)

	fmt.Printf("Content: %v", content)

	fmt.Printf("Bytes received: %v\n", bytes_readed)

	bytes_sent, err := conn.Write([]byte(res))

	if err != nil {
		fmt.Println("Error responding client: ", err.Error())
	}

	fmt.Printf("Bytes sent: %v", bytes_sent)
}
