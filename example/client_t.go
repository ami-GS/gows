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
	client.Ping(data)
	time.Sleep(time.Second)
	client.Ping(data)
	time.Sleep(time.Second)
	client.Send(data, true)
	time.Sleep(time.Second)
}
