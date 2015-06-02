package gows

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	server := NewServer("127.0.0.1:8080")
	go server.WaitClient()
	reqUrl := "ws://127.0.0.1:8080"
	actual, _ := NewClient(reqUrl)
	u, _ := url.Parse(reqUrl)
	expect := &Client{[]Connection{}}
	_ = expect.Connect(u.Host)
	if !reflect.DeepEqual(actual.connections, expect.connections) {
		t.Errorf("got %v\nwant %v", actual.connections, expect.connections)
	}
}

func TestValidateResponse(t *testing.T) {
	cli, _ := NewClient("ws://example.com:8080")
	buffer := []byte(HandshakeResponse)
	actual := cli.ValidateResponse(buffer)
	expect := true
	if actual != expect {
		t.Errorf("got %v\nwant %v", actual, expect)
	}
}
