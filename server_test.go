package gows

import (
	"net"
	"reflect"
	"strconv"
	"testing"
)

func TestNewServer(t *testing.T) {
	host := "127.0.0.1"
	port := "8080"
	portInt, _ := strconv.ParseUint(port, 10, 16)
	addr := host + ":" + port
	serv, _ := net.Listen("tcp", addr)
	actual := NewServer(addr)
	expect := &Server{map[string]*Connection{"": &Connection{}},
		&serv, &Addr{host, uint16(portInt)}, VERSION}

	if reflect.DeepEqual(actual.clients, expect.clients) {
		t.Errorf("got %v\nwant %v", actual.clients, expect.clients)
	}
	if actual.serv != expect.serv {
		t.Errorf("got %v\nwant %v", actual.serv, expect.serv)
	}
	if actual.addr != expect.addr {
		t.Errorf("got %v\nwant %v", actual.addr, expect.addr)
	}
	if actual.version != expect.version {
		t.Errorf("got %v\nwant %v", actual.version, expect.version)
	}
}

func TestValidateRequest(t *testing.T) {
	server := NewServer("127.0.0.1:8080")
	actual := server.ValidateRequest([]byte(HandshakeRequest), "172.168.0.10")
	expect := true
	if actual != expect {
		t.Errorf("got %v\nwant %v", actual, expect)
	}
}
