package gows

import (
	//"fmt"
	"net"
)

type Client struct {
	connections []Connection
}

func NewClient() *Client {
	client := &Client{[]Connection{}}
	return client
}

func (self *Client) Send(data []byte, isBin bool) {
	if isBin {
		self.connections[0].Send(data, BINARY)
	} else {
		self.connections[0].Send(data, TEXT)
	}
}

func (self *Client) Ping(data []byte) {
	self.connections[0].Send(data, PING)
}

func (self *Client) Connect(addr string) {
	for _, con := range self.connections {
		if addr == con.addr {
			// refuse connect
		}
	}
	conn, _ := net.Dial("tcp", addr)
	connection := NewConnection(&conn, addr)
	connection.ValidateHandshake()
	self.connections = append(self.connections, *connection)
	go connection.ReceiveLoop()
}
