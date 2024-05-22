package handle

import (
	"MegaX/database"
	"MegaX/middlewares"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"

)

var db *sqlx.DB
var err error

const connectionString = "host=127.0.0.1 port=5432 user=postgres password=akeceqm dbname=mega_xxx sslmode=disable"


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

// ?Запросы для Addres
func HandleAddresGET(c *gin.Context) {
	var addres []database.Addres_User

	err := db.Select(&addres, `Select * FROM addres_user`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, addres)
}

func HandleAddresPOST(c *gin.Context) {
	var addres database.Addres_User

	if err := c.ShouldBindJSON(&addres); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(`INSERT INTO addres_user(city,street,house,flat,entrance) VALUES($1,$2,$3,$4,$5) RETURNING id`, addres.City, addres.Street, addres.House, addres.Flat, addres.Entrance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &addres)
}


func HandleAddresDEL(c *gin.Context) {
	var addres database.Addres_User

	if err := c.ShouldBindJSON(&addres); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(`DELETE FROM users WHERE id_addr = $1`, addres.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec(`DELETE FROM addres_user WHERE id = $1;`, addres.Id) // Added semicolon
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		return
	}

	c.JSON(http.StatusOK, &addres)

}

func HandleAddresPUT(c *gin.Context) {
	var addres database.Addres_User
	var addresDB database.Addres_User
	if err := c.ShouldBindJSON(&addres); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := db.Get(&addresDB, `SELECT * FROM addres_user WHERE id = $1`, addres.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if addres.City != "" {
		addresDB.City = addres.City
	}

	if addres.Street != "" {
		addresDB.Street = addres.Street
	}

	if addres.House != "" {
		addresDB.House = addres.House
	}

	if addres.Flat != "" {
		addresDB.Flat = addres.Flat
	}

	if addres.Entrance != "" {
		addresDB.Entrance = addres.Entrance
	}

	_, err = db.Exec(`UPDATE addres_user SET City=$1, Street=$2, House=$3, Flat=$4, Entrance=$5 WHERE id =$6`, addresDB.City, addresDB.Street, addresDB.House, addresDB.Flat, addresDB.Entrance, addresDB.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &addresDB)
	_, err := db.Exec(`INSERT INTO users (phone, password, name, surname) VALUES ($1,$2,$3,$4) RETURNING id`, user.Phone, middlewares.PasswordHash(user.Password), user.Name, user.Surname)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}


