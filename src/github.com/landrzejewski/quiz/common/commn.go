package common

import (
	"io/ioutil"
	"github.com/jinzhu/gorm"
	"github.com/landrzejewski/bank"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func PanicIfError(err error)  {
	if err != nil {
		panic(err)
	}
}

func LoadData(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	PanicIfError(err)
	return bytes
}

func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "quiz.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&bank.Account{})
	return db
}