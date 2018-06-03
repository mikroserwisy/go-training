package chat

import (
	"net"
	"bufio"
	"log"
	"strconv"
	"math/rand"
)

type Server struct {

	Address string

}

func (server Server) Start() {
	mainRoom := server.createRoom("main")
	listener, err := net.Listen("tcp", server.Address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		} else {
			client := server.createClient(mainRoom, connection, strconv.Itoa(rand.Int()))
			mainRoom.add(client)
		}
	}
}

func (server Server) createClient(room *room, connection net.Conn, userName string) *client {
	newClient := client{
		name: userName,
		room: room,
		reader: bufio.NewReader(connection),
		writer: bufio.NewWriter(connection),
	}
	go newClient.read()
	return &newClient
}

func (server Server) createRoom(name string) *room {
	newRoom := room{
		name: name,
		channel: make(chan message),
		clients: make([]*client, 0),
	}
	go newRoom.start()
	return &newRoom
}