package main

import "goRedis/server"

func main() {
	var se server.Server = server.NewTCPServer(5535, nil)
	se.StartListen()
}
