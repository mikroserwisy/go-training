package chat

import (
	"net"
)

type room struct {
	name     string
	channel  chan message
	clients  []*client
	listener net.Listener
}

func (room *room) broadcast(message message) {
	for _, client := range room.clients {
		if client.connection != message.sender.connection {
			client.write(message)
		}
	}
}

func (room *room) add(client *client) {
	room.clients = append(room.clients, client)
}

func (room *room) init() {
	for {
		select {
		case message := <-room.channel:
			room.broadcast(message)
		}
	}
}
