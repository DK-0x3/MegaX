package handle

import (
	//"MegaX/database"
	//"MegaX/middlewares"
	"MegaX/database"
	"log"
	"net/http"

	//"net/http"

	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	//"golang.org/x/crypto/bcrypt"
)

var db *sqlx.DB
var err error
const connectionString = "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=mega_xxx sslmode=disable"


func InitDB() {
	db, err = sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("failed connection db")
	}
}

func HandleUsersGET(c *gin.Context) {
	var users []database.User

	//users = make([]database.User, 0)

	err := db.Select(&users, `SELECT * FROM users`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}


