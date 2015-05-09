package gows

import (
//"fmt"
)

var VERSION string = "13"

type Opcode byte

const (
	CONTINUE Opcode = iota
	TEXT
	BINARY
	CLOSE Opcode = iota + 5
	PING
	PONG
)

func (opc Opcode) String() string {
	opcodes := []string{
		"CONTINUE",
		"TEXT",
		"BINARY",
		"", "", "", "", "", //RESERVED
		"CLOSE",
		"PING",
		"PONG",
	}
	return opcodes[int(opc)]
}

type State byte

const (
	OPEN State = iota + 1
	CONNECTING
	CLOSING
	CLOSED
)

func (st State) String() string {
	states := []string{
		"OPEN",
		"CONNECTING",
		"CLOSING",
		"CLOSED",
	}
	return states[int(st)-1]

}

type CloseCode uint16

const (
	NORMAL_CLOSURE CloseCode = iota + 1000
	GOING_AWAY
	PROTOCOL_ERROR
	UNSUPPORTED_DATA
	NO_STATUS_RCVD CloseCode = iota + 1001
	ABNORMAL_CLOSURE
	INVALID_FRAME_PAYLOAD_DATA
	POLICY_VIOLATION
	MESSAGE_TOO_BIG
	MANDATORY_EXT
	INTERNAL_SERVER_ERROR
	TLS_HANDSHAKE
)

func (c CloseCode) String() string {
	closeCodes := []string{
		"NORMAL_CLOSURE",
		"GOING_AWAY",
		"PROTOCOL_ERROR",
		"UNSUPPORTED_DATA",
		"", // RESERVED
		"NO_STATUS_RCVD",
		"ABNORMAL_CLOSURE",
		"INVALID_FRAME_PAYLOAD_DATA",
		"POLICY_VIOLATION",
		"MESSAGE_TOO_BIG",
		"MANDATORY_EXT",
		"INTERNAL_SERVER_ERROR",
		"TLS_HANDSHAKE",
	}
	return closeCodes[int(c)-1000]
}

var HandshakeRequest []byte = []byte(`GET /chat HTTP/1.1
Host: server.example.com
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==
Origin: http://example.com
Sec-WebSocket-Protocol: chat, superchat
Sec-WebScoket-Version: 13
`)

var HandshakeResponse []byte = []byte(`HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-Websocket-Accept: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=
`)
