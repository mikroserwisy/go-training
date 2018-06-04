package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"fmt"
	"net/http"
	"github.com/gin-contrib/location"
)

func getLocationUrl(context *gin.Context, id uint) string {
	return fmt.Sprintf("%v%v/%v", location.Get(context), context.Request.URL, id)
}

func createTask(context *gin.Context)  {
	var task Task
	context.Bind(&task)
	db, _ := context.MustGet("db").(*gorm.DB)
	db.Save(&task)
	context.Header("Location", getLocationUrl(context, task.ID))
	context.JSON(http.StatusCreated, gin.H{"resourceId": task.ID})
}

func Router() *gin.Engine {
	engine := getEngine()
	api := engine.Group("/api/v1/")
	api.POST("tasks", createTask)
	return engine
}
