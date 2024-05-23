package main

import (
	handle "MegaX/handle"
	"MegaX/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

var server *gin.Engine

const connectionString = "host=127.0.0.1 port=5432 user=postgres password=akeceqm dbname=mega_xxx sslmode=disable"

func main() {
	var err error
	db, err = handle.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	server = gin.Default()
	middlewares.LogFile(server)

	//? Работа с юзерами
	server.GET("/users", func(c *gin.Context) {
		handle.HandleUsersGET(c, db)
	})
	server.POST("/users", func(c *gin.Context) {
		handle.HandleUserPOST(c, db)
	})
	server.PUT("/users", func(c *gin.Context) {
		handle.HandleUserPUT(c, db)
	})
	server.DELETE("/users", func(c *gin.Context) {
		handle.HandleUserDEL(c, db)
	})

	//? Работа с аддресами
	server.GET("/addresUser", func(c *gin.Context) {
		handle.HandleAddresGET(c, db)
	})
	server.POST("/addresUser", func(c *gin.Context) {
		handle.HandleAddresPOST(c, db)
	})
	server.DELETE("/addresUser", func(c *gin.Context) {
		handle.HandleAddresDEL(c, db)
	})
	server.PUT("/addresUser", func(c *gin.Context) {
		handle.HandleAddresPUT(c, db)
	})

	//? Работа с категориями
	server.GET("/category", func(c *gin.Context) {
		handle.HandleCategoryGET(c, db)
	})
	server.POST("/category", func(c *gin.Context) {
		handle.HandleCategoryPOST(c, db)
	})
	server.DELETE("/category", func(c *gin.Context) {
		handle.HandleCategoryDEL(c, db)
	})
	server.PUT("/category", func(c *gin.Context) {
		handle.HandleCategoryPUT(c, db)
	})

	server.Run(":8080")

}
