package gows

import (
	"net"
	"strconv"
	"testing"
)

func TestNewConnection(t *testing.T) {
	host := "127.0.0.1"
	port := "8080"
	portInt, _ := strconv.ParseUint(port, 10, 16)
	addr := host + ":" + port
	conn, _ := net.Dial("tcp", addr)
	actual := NewConnection(&conn, addr, true)
	expect := &Connection{OPEN, &conn, true, false,
		&Addr{host, uint16(portInt)}, false, 0, "", ""}
	if actual.state != expect.state {
		t.Errorf("got %v\nwant %v", actual.state, expect.state)
	}
	if actual.conn != expect.conn {
		t.Errorf("got %v\nwant %v", actual.conn, expect.conn)
	}
	if actual.IsClient != expect.IsClient {
		t.Errorf("got %v\nwant %v", actual.IsClient, expect.IsClient)
	}
	if actual.WaitPong != expect.WaitPong {
		t.Errorf("got %v\nwant %v", actual.WaitPong, expect.WaitPong)
	}
	if actual.addr != expect.addr {
		t.Errorf("got %v\nwant %v", actual.addr, expect.addr)
	}
	if actual.IsBrowser != expect.IsBrowser {
		t.Errorf("got %v\nwant %v", actual.IsBrowser, expect.IsBrowser)
	}
	if actual.RSV != expect.RSV {
		t.Errorf("got %v\nwant %v", actual.IsBrowser, expect.IsBrowser)
	}
	if actual.SubProto != expect.SubProto {
		t.Errorf("got %v\nwant %v", actual.SubProto, expect.SubProto)
	}
	if actual.Extention != expect.Extention {
		t.Errorf("got %v\nwant %v", actual.Extention, expect.Extention)
	}
}
