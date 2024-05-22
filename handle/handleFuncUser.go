package handle

import (
	"MegaX/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

<<<<<<< HEAD
const connectionString = "host=127.0.0.1 port=5432 user=postgres password=akeceqm dbname=mega_xxx sslmode=disable"
=======

const connectionString = "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=mega_xxx sslmode=disable"
>>>>>>> 0315d67fc7032d38c4d2a701b0c0fb1659fef4c0

func InitDB() {
	db, err = sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("failed connection db")
	}
}

// ? Вызов всех user из таблицы
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

// ? Вызов всех  адресов
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
}
