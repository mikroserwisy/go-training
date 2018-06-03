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
		if client.name != message.sender.name {
			client.write(message)
		}
	}
}

func (room *room) add(client *client) {
	room.clients = append(room.clients, client)
}

func (room *room) start() {
	for {
		message := <-room.channel
		room.broadcast(message)
		/*select {
		case message := <-room.channel:
			room.broadcast(message)
		}*/
	}
}
