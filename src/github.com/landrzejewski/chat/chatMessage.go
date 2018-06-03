package chat

import "github.com/jinzhu/gorm"

type chatMessage struct {

	gorm.Model

	Text string

	Sender string

}
