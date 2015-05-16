package gows

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	clients map[string]*Connection
	serv    *net.Listener
	addr    *Addr // binding addr
	version string
}

func NewServer(addr string) *Server {
	if !strings.Contains(addr, ":") {
		addr += ":80"
	}
	ad := strings.Split(addr, ":")
	serv, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	port, _ := strconv.ParseUint(ad[1], 10, 16)
	server := &Server{
		map[string]*Connection{"": &Connection{}},
		&serv, &Addr{ad[0], uint16(port)},
		VERSION}

	return server
}

// TODO: gather validate function to 1 (client and server)
func (self *Server) ValidateRequest(buffer []byte, addr string) (validate bool) {
	af := strings.Split(string(buffer), "\n")
	validate = true
	for i, v := range af {
		if i == 1 {
			continue
		}

		header := strings.Split(v, ": ")
		fmt.Printf("%v ", header)
		if header[0] == "Host" {
			validate = true // authority ?
		} else if header[0] == "Upgrade" && header[1] != "websocket" {
			validate = false
		} else if header[0] == "Connection" && header[1] != "Upgrade" {
			validate = false
		} else if header[0] == "Sec-Websocket-Key" && header[1] != "dGhlIHNhbXBsZSBub25jZQ==" {
			validate = false
		} else if header[0] == "Sec-WebScoket-Version" && !strings.Contains(header[1], self.version) {
			validate = false
		} else if header[0] == "Origin" {
			self.clients[addr].IsBrowser = true
		} else if header[0] == "Sec-WebSocket-Protocol" {
			self.clients[addr].SubProto = header[1]
		} else if header[0] == "Sec-WebSocket-Extensions" {
			self.clients[addr].Extention = header[1]
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

		connection := NewConnection(&conn, conn.LocalAddr().String(), false)
		addr := conn.LocalAddr().String()
		self.clients[addr] = connection
		buffer, _ := connection.Read(1024)
		if !self.ValidateRequest(buffer, addr) {
			connection.Close()
		} else {
			connection.ResponseOpeningHandshake()
			go connection.ReceiveLoop()
		}
	}
}
