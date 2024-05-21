package main

import (
	handle "MegaX/handle"
	"MegaX/middlewares"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var server *gin.Engine

func main() {
	server = gin.Default()
	middlewares.LogFile(server)
	handle.InitDB()

	server.GET("/users", handle.HandleUsersGET)
	//server.POST("/users", handle.HandleUserPOST)
	//server.POST("/authorization", handle.HandleAuthorizationAcc)

	server.Run(":8080")

}
