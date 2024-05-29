package main

import (
	"MegaX/internal/database"
	"MegaX/internal/handler"
	"MegaX/internal/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

var server *gin.Engine

func main() {
	var err error
	db, err = database.InitDB(database.ConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	handler.InitRoutes(server, db)
	server = gin.Default()
	middlewares.LogFile(server)

	server.Run(":8080")

}
