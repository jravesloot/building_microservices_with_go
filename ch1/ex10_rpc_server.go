// ex10.go implements a hello world rpc server

package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/building-microservices-with-go/chapter1/rpc/contract"
)

const port = 8080

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.name
	return nil
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", error))
	}

	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
