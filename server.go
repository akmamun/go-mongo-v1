package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"github.com/go-mongo/controllers"


)

var todosCollection *mgo.Collection
var session *mgo.Session

var router *gin.Engine
func init() {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	// get a Collection of todo
	todosCollection = session.DB("test-todo").C("todo")
}
func main() {
	router := gin.Default()
 
	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", controllers.CreateTodo)
		v1.GET("/", controllers.FetchAllTodo)
		v1.GET("/:id", controllers.FetchSingleTodo)
		v1.PUT("/:id", controllers.UpdateTodo)
		v1.DELETE("/:id", controllers.DeleteTodo)
	}

	// Start serving the application
	router.Run(":8000")

}
 


 