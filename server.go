package gows

import (
	"fmt"
	"net"
)

type Server struct {
	clients []Connection
	addr    string
}

func NewServer(addr string) *Server {
	server := &Server{[]Connection{}, addr}
	return server
}

func (self *Server) WaitClient() {
	serv, err := net.Listen("tcp", self.addr)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := serv.Accept()
		if err != nil {
			panic(err)
		}

		connection := Connection{OPEN, &conn, false, conn.LocalAddr().String()}
		self.clients = append(self.clients, connection)
		buffer, err := connection.Read(1024)
		if err != nil {
			//
		}
		fmt.Printf("%s\n", buffer)
		connection.ResponseOpeningHandshake()
		go connection.ReceiveLoop()
	}
}
