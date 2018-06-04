package tasks

import "github.com/jinzhu/gorm"

type Task struct {

	gorm.Model

	Title string `json:"name"`

	Completed bool `json:"done"`

}
