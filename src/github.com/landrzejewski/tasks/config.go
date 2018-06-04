package tasks

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/gin-contrib/location"
	"gopkg.in/gin-contrib/cors.v1"
)

func initDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "tasks.db")
	if err != nil {
		panic("Failed to connect")
	}
	db.AutoMigrate(&Task{})
	return db
}

func DbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("db", db)
		context.Next()
	}
}

func initialize() (*gin.Engine, *gorm.DB) {
	db := initDb()
	engine := gin.Default()

	/*config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	engine.Use(cors.New(config))*/

	engine.Use(cors.Default(), location.Default(), DbMiddleware(db))
	return engine, db
}

