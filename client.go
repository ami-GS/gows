package gows

import (
	"fmt"
	"net"
	"net/url"
	"strings"
)

type Client struct {
	connections []Connection
}

func NewClient(reqUrl string) (*Client, error) {
	u, _ := url.Parse(reqUrl)
	client := &Client{[]Connection{}}
	err := client.Connect(u.Host)
	if err != nil {
		return nil, err
	}
	return client, nil
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

// TODO: gather validate function to 1 (client and server)
func (self *Client) ValidateResponse(buffer []byte) (validate bool) {
	af := strings.Split(string(buffer), "\n")
	validate = true
	for i, v := range af {
		if i == 1 {
			continue
		}

		header := strings.Split(v, ": ")
		fmt.Printf("%v ", header)
		if header[0] == "Upgrade" && header[1] != "websocket" {
			validate = false
		} else if header[0] == "Connection" && header[1] != "Upgrade" {
			validate = false
		} else if header[0] == "Sec-Websocket-Accept" && header[1] != "s3pPLMBiTxaQ9kYGzzhZRbK+xOo=" {
			validate = false
		}
	}
	return
}

func (self *Client) Connect(addr string) (err error) {
	for _, con := range self.connections {
		if addr == con.addr {
			// refuse connect
		}
	}
	conn, err := net.Dial("tcp", addr)
	connection := NewConnection(&conn, addr, true)
	connection.SendHandshake()
	buffer, err := connection.Read(256)
	if !self.ValidateResponse(buffer) {
		connection.Close()
	} else {
		self.connections = append(self.connections, *connection)
		go connection.ReceiveLoop()
	}
	return
}
