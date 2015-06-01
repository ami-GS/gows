package gows

import (
	"net/url"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	reqUrl := "ws://example.com:8080"
	u, _ := url.Parse(reqUrl)
	actual, _ := NewClient(u.Host)
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
