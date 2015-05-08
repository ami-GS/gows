package gows

import (
	"fmt"
	"net"
)

type Connection struct {
	state     State
	conn      *net.Conn
	WaitPong  bool
	addr      string
	IsBrowser bool
	SubProto  string
	Extention string
}

func NewConnection(conn *net.Conn, addr string) (connection *Connection) {
	connection = &Connection{OPEN, conn, false, addr, false, "", ""}
	return

}

func (self *Connection) SendHandshake() {
	(*self.conn).Write(HandshakeRequest)
	self.state = CONNECTING
}

func (self *Connection) ResponseOpeningHandshake() {
	(*self.conn).Write(HandshakeResponse)
}

func (self *Connection) Close() {
	(*self.conn).Close()
	self.state = CLOSED
}

func (self *Connection) Send(data []byte, opc Opcode) {
	if opc == CLOSE {
		self.state = CLOSING
	} else if opc == PING {
		self.WaitPong = true
	}
	fmt.Printf("Send\n\tOpcode=%s, Data=%s\n", opc.String(), data)
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
		fmt.Printf("Recv\n\tOpcode=%s, Data=%s\n", frame.opc.String(), frame.Payload)
		if frame.opc == CLOSE {
			self.state = CLOSING
			self.Send([]byte(""), CLOSE)
			break
		} else if frame.opc == PING {
			self.Send(frame.Payload, PONG)
		} else if frame.opc == PONG {
			self.WaitPong = false
		}
	}

	self.Close()
}
