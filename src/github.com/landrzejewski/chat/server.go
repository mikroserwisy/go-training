package chat

import (
	"net"
	"bufio"
	"log"
	"strconv"
	"math/rand"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Server struct {

	Address string

}

func (server Server) Start() {
	mainRoom := server.createRoom("main", server.initDb())
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

func (server Server) createRoom(name string, db *gorm.DB) *room {
	newRoom := room{
		name: name,
		channel: make(chan message),
		clients: make([]*client, 0),
		db: db,
	}
	go newRoom.start()
	return &newRoom
}

func (server Server) initDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "chat.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&chatMessage{})
	return db
}