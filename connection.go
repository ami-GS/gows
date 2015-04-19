package gows

import (
	"net"
)

type Connection struct {
	state State
	conn  *net.Conn
}

func (self *Connection) OpeningHandshake() {
	(*self.conn).Write(HandshakeRequest)
}

func (self *Connection) ResponseOpeningHandshake() {
	(*self.conn).Write(HandshakeResponse)
}
