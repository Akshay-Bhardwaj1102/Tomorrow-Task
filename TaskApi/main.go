package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"task.com/config"
	"task.com/controller"
)

func main() {

	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.POST("/Add", controller.CreateLeave)
	r.GET("/GetLeaves/:id", controller.GetLeaves)
	r.GET("/GetAll", controller.GetAllData)
	defer r.Run()

}
