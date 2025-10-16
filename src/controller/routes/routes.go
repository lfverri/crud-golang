package routes

import (
	"crud-golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	// Initialize your routes here
	r.GET("/ping", controller.Ping)
	r.GET("/users/", controller.FindUsers)
	r.GET("/users/:userId", controller.FindUserByID)
	r.GET("/users/:userEmail", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:userId", controller.UpdateUser)
	r.DELETE("/users/:userId", controller.DeleteUser)
}