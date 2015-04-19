package gows

type Opcode byte

const (
	CONTINUE Opcode = iota
	TEXT
	BINARY
	CLOSE Opcode = iota + 8
	PING
	PONG
)

type State byte

const (
	LISTENING State = iota
	OPEN
	CONNECTING
	CLOSING
	CLOSED
)

type CloseCode uint16

const (
	NORMAL_CLOSURE CloseCode = iota + 1000
	GOING_AWAY
	PROTOCOL_ERROR
	UNSUPPORTED_DATA
	NO_STATUS_RCVD CloseCode = iota + 1005
	ABNORMAL_CLOSURE
	INVALID_FRAME_PAYLOAD_DATA
	POLICY_VIOLATION
	MESSAGE_TOO_BIG
	MANDATORY_EXT
	INTERNAL_SERVER_ERROR
	TLS_HANDSHAKE
)

var HandshakeRequest []byte = []byte(`GET /chat HTTP/1.1
Host: server.example.com
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==
Origin: http://example.com
Sec-WebSocket-Protocol: chat, superchat
Sec-WebScoket-Version: 13`)

var HandshakeResponse []byte = []byte(`HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-Websocket-Accept: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=`)
