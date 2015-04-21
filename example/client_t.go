package main

import (
	"../../gows"
	//"fmt"
	"time"
)

func main() {
	client := gows.NewClient()
	client.Connect("127.0.0.1:8080")
	data := []byte("aiueo")
	client.Send(data, gows.PING)
	time.Sleep(time.Second)
	client.Send(data, gows.PING)
	time.Sleep(time.Second)
}
