package gows

import (
	"fmt"
	"net"
)

type Connection struct {
	state    State
	conn     *net.Conn
	WaitPong bool
	addr     string
}

func NewConnection(addr string) (conn *Connection) {
	conn = &Connection{OPEN, nil, false, addr}
	return

}

func (self *Connection) ValidateHandshake() {
	self.openingHandshake()
	for {
		buffer, _ := self.Read(1024)
		fmt.Printf("%s\n", buffer)
		break
		// validate handshake response from peer
		// finally break
	}
}

func (self *Connection) openingHandshake() {
	(*self.conn).Write(HandshakeRequest)
	self.state = CONNECTING
}

func (self *Connection) ResponseOpeningHandshake() {
	(*self.conn).Write(HandshakeResponse)
}

func (self *Connection) Send(data []byte, opc Opcode) {
	if opc == CLOSE {
		self.state = CLOSING
	} else if opc == PING {
		self.WaitPong = true
	}
	(*self.conn).Write(Pack(data, opc))
}

func (self *Connection) Read(length uint32) (buffer []byte, err error) {
	buffer = make([]byte, length)
	_, err = (*self.conn).Read(buffer)
	if err != nil {
		//panic(err)
	}
	return
}

func (self *Connection) ReceiveLoop() {
	for {
		frame, err := Parse(self)
		if err != nil {
			break
		}
		fmt.Printf("Opcode=%s, %s\n", frame.opc.String(), frame.Payload)
		if frame.opc == CLOSE {
			self.state = CLOSING
			self.Send([]byte(""), CLOSE)
			break
		} else if frame.opc == PING {
			fmt.Printf("%s\n", frame.Payload)
			self.Send(frame.Payload, PONG)
		} else if frame.opc == PONG {
			self.WaitPong = false
		}
	}

	(*self.conn).Close()
	self.state = CLOSED
}
