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

  //? Работа с юзерами
	server.GET("/users", handle.HandleUsersGET)
	server.POST("/users", handle.HandleUserPOST)
	server.PUT("/users", handle.HandleUserPUT)
	server.DELETE("/users", handle.HandleUserDEL)

	//? Работа с аддресами
	server.GET("/addresUser", handle.HandleAddresGET)
	server.POST("/addresUser", handle.HandleAddresPOST)
	server.DELETE("/addresUser", handle.HandleAddresDEL)
	server.PUT("addresUser", handle.HandleAddresPUT)



	server.Run(":8080")

}
