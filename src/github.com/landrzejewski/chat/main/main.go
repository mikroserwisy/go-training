package main

import "github.com/landrzejewski/chat"

func main() {
	chat.Server{Address:":6000"}.Start()
}
