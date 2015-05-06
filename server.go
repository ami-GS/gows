package gows

import (
	"fmt"
	"net"
	"strings"
)

type Server struct {
	clients map[string]*Connection
	serv    *net.Listener
	addr    string
	version string
}

func NewServer(addr string) *Server {
	serv, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	server := &Server{
		map[string]*Connection{"": &Connection{}},
		&serv, addr, "13"}

	return server
}

// TODO: gather validate function to 1 (client and server)
func (self *Server) ValidateRequest(buffer []byte) (validate bool) {
	af := strings.Split(string(buffer), "\n")
	validate = true
	for _, v := range af {
		if strings.Contains(v, "HTTP/1.1") {
			continue
		}
		header := strings.Split(v, ": ")
		fmt.Printf("%v ", header)
		if header[0] == "Upgrade" && header[1] != "websocket" {
			validate = false
		} else if header[0] == "Connection" && header[1] != "Upgrade" {
			validate = false
		} else if header[0] == "Sec-Websocket-Key" && header[1] != "dGhlIHNhbXBsZSBub25jZQ==" {
			validate = false
		} else if header[0] == "Sec-WebSocket-Protocol" && !strings.Contains(header[1], "chat") {
			validate = false
		} else if header[0] == "Sec-WebScoket-Version" && !strings.Contains(header[1], self.version) {
			validate = false
		}

	}
	return
}

func (self *Server) WaitClient() {
	for {
		conn, err := (*self.serv).Accept()
		if err != nil {
			panic(err)
		}

		connection := NewConnection(&conn, conn.LocalAddr().String())
		self.clients[conn.LocalAddr().String()] = connection
		buffer, _ := connection.Read(1024)
		if !self.ValidateRequest(buffer) {
			connection.Close()
		} else {
			connection.ResponseOpeningHandshake()
			go connection.ReceiveLoop()
		}
	}
}
