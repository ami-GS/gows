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

func (self *Client) Send(data string, isBin bool) {
	if isBin {
		self.connections[0].Send([]byte(data), BINARY)
	} else {
		self.connections[0].Send([]byte(data), TEXT)
	}
}

func (self *Client) Ping(data string) {
	self.connections[0].Send([]byte(data), PING)
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
