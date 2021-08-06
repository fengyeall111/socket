package main

import (
	"./socket_client"
	"./socket_server"
)

func main() {
	var ip = "127.0.0.1:8080"
	go socket_server.NewServer(ip)
	go socket_client.NewClient(ip)
	for {
	}
}
