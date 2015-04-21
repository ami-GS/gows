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

func (self *Client) Send(data []byte, opc Opcode) {
	self.connections[0].Send(data, opc)
}

func (self *Client) Connect(addr string) {
	for _, con := range self.connections {
		if addr == con.addr {
			// refuse connect
		}
	}
	conn, _ := net.Dial("tcp", addr)
	connection := NewConnection(addr)
	connection.conn = &conn
	connection.ValidateHandshake()
	self.connections = append(self.connections, *connection)
	go connection.ReceiveLoop()
}
