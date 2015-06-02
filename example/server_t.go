package main

import (
	"github.com/ami-GS/gows"
)

func main() {
	server := gows.NewServer("127.0.0.1:8080")
	server.WaitClient()
}
