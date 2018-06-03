package chat

import (
	"bufio"
	"log"
	"io"
	"fmt"
)

type client struct {
	 name string
	 room *room
	 reader *bufio.Reader
	 writer *bufio.Writer
}

func (client *client) read() {
	for {
		line, err := client.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				message := message{"Client disconnected\n", client}
				client.room.broadcast(message)
			} else {
				log.Printf("Exception: %v", err)
			}
			return
		}
		client.room.channel <- message{line, client}
	}
}

func (client *client) write(message message) {
	messageText := fmt.Sprintf("%s: %s", message.sender.name, message.text)
	client.writer.WriteString(messageText)
	client.writer.Flush()
}
