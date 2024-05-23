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

const connectionString = "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=mega_xxx sslmode=disable"

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
	server.GET("/users/addres", func(c *gin.Context) {
		handle.HandleUsersAndAddres_GET(c, db)
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
	server.GET("/user/:id", func (c *gin.Context) {
		handle.HandleUserId_GET(c, db)
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

	//? Работа с MainCategory
	server.GET("/maincategory", func (c *gin.Context)  {
		handle.HandleMainCategoryGET(c, db)
	})
	server.POST("/maincategory", func (c *gin.Context)  {
		handle.HandleMainCategoryPOST(c, db)
	})
	server.PUT("/maincategory", func (c *gin.Context) {
		handle.HandleMainCategoryPUT(c, db)
	})
	server.DELETE("/maincategory", func (c *gin.Context) {
		handle.HandleMainCategoryDEL(c, db)
	})
	server.GET("/maincategory/category", func (c *gin.Context) {
		handle.HandleMainCategoryAndCategory(c, db)
	})
	server.GET("/maincategory/:id", func (c *gin.Context) {
		handle.HandleMainCategoryId_GET(c, db)
	})
	server.GET("/maincategory/:id/category", func (c *gin.Context) {
		handle.HandleMainCategoryAndCategoryId_GET(c, db)
	})
	

	server.Run(":8080")

}
