package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"fmt"
	"net/http"
	"github.com/gin-contrib/location"
	"strconv"
)

func getLocationUrl(context *gin.Context, id uint) string {
	return fmt.Sprintf("%v%v/%v", location.Get(context), context.Request.URL, id)
}

func createTask(context *gin.Context)  {
	var task Task
	context.Bind(&task)
	db := context.MustGet("db").(*gorm.DB)
	db.Save(&task)
	context.Header("Location", getLocationUrl(context, task.ID))
	context.JSON(http.StatusCreated, gin.H{"resourceId": task.ID})
}

func getTasks(context *gin.Context) {
	var tasks []Task
	db := context.MustGet("db").(*gorm.DB)
	db.Find(&tasks)
	context.JSON(http.StatusOK, tasks)
}

func getTask(context *gin.Context) {
	context.JSON(http.StatusOK, context.MustGet("task").(Task))
}

func getTaskMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var task Task
		db.First(&task, context.Param("id"))
		if task.ID == 0 {
			context.AbortWithStatus(http.StatusNotFound)
		} else {
			context.Set("task", task)
			context.Next()
		}
	}
}

func updateTask(context *gin.Context)  {
	paramId := context.Param("id")
	db := context.MustGet("db").(gorm.DB)
	var updatedTodo Task
	context.Bind(&updatedTodo)
	id, _ := strconv.Atoi(paramId)
	updatedTodo.ID = uint(id)
	db.Model(&updatedTodo).Update(updatedTodo)
	context.Status(http.StatusNoContent)
}

func deleteTask(context *gin.Context)  {
	todo, _ := context.MustGet("task").(Task)
	db := context.MustGet("db").(gorm.DB)
	db.Delete(&todo)
	context.Status(http.StatusNoContent)
}

func Router() *gin.Engine {
	engine, db := initialize()
	api := engine.Group("/api/v1/")
	api.POST("tasks", createTask)
	api.GET("tasks", getTasks)
	api.GET("tasks/:id", getTaskMiddleware(db), getTask)
	api.PUT("tasks/:id", getTaskMiddleware(db), updateTask)
	api.DELETE("tasks/:id", getTaskMiddleware(db), deleteTask)
	return engine
}
