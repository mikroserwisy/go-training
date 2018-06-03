package chat

import (
	"net"
	"github.com/jinzhu/gorm"
)

type room struct {

	name     string

	channel  chan message

	clients  []*client

	listener net.Listener

	db *gorm.DB

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
		room.db.Create(&chatMessage{
			Text:message.text,
			Sender: message.sender.name,
		})
		/*select {
		case message := <-room.channel:
			room.broadcast(message)
		}*/
	}
}
