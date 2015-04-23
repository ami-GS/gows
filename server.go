package gows

import (
	"fmt"
	"net"
)

type Server struct {
	clients map[string]*Connection
	serv    *net.Listener
	addr    string
}

func NewServer(addr string) *Server {
	serv, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	server := &Server{
		map[string]*Connection{"": &Connection{}},
		&serv, addr}

	return server
}

func (self *Server) WaitClient() {
	for {
		conn, err := (*self.serv).Accept()
		if err != nil {
			panic(err)
		}

		connection := NewConnection(&conn, conn.LocalAddr().String())
		self.clients[conn.LocalAddr().String()] = connection
		buffer, err := connection.Read(1024)
		if err != nil {
			//
		}
		fmt.Printf("%s\n", buffer)
		connection.ResponseOpeningHandshake()
		go connection.ReceiveLoop()
	}
}
