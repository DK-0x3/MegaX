package handle

import (
	//"MegaX/database"
	//"MegaX/middlewares"
	"MegaX/database"
	"MegaX/middlewares"
	"log"
	"net/http"

	//"net/http"

	//"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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

//? Запросы для Users

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

func HandleUserPOST(c *gin.Context) {
	var user database.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(`INSERT INTO users (phone, password, name, surname) VALUES ($1,$2,$3,$4) RETURNING id`, user.Phone, middlewares.PasswordHash(user.Password), user.Name, user.Surname)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}

func HandleUserDEL(c *gin.Context) {
	var user database.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
	}

	_, err := db.Exec(`DELETE FROM users WHERE id = $1`, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
	}

	c.JSON(http.StatusOK, &user)
}

func HandleUserPUT(c *gin.Context) {
	var user database.User
	var userDB database.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
	}
	
	err := db.Get(&userDB, `SELECT * FROM users WHERE id = $1`, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
	}

	if user.Phone != "" {
		userDB.Phone = user.Phone
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("")); err != nil {
		userDB.Password = string(middlewares.PasswordHash(user.Phone))
	}
	if user.Name != "" {
		userDB.Name = user.Name
	}
	if user.Surname != "" {
		userDB.Surname = user.Surname
	}
	if user.Id_Addr.Int32 != 0 {
		userDB.Id_Addr = user.Id_Addr
	}
	if user.Role != "" {
		userDB.Role = user.Role
	}

	_, err = db.Exec(`UPDATE users SET phone = $1, password = $2, name = $3, surname = $4, id_addr = $5, role = $6 WHERE id = $7`, userDB.Phone, userDB.Password, userDB.Name, userDB.Surname, userDB.Id_Addr, userDB.Role, userDB.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
	}

	c.JSON(http.StatusOK, &userDB)
}
