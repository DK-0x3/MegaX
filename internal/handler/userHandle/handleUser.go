package userHandle

import (
	"MegaX/internal/database/models"
	"MegaX/internal/middlewares"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// ? Запросы для Users
func HandleUsersGET(c *gin.Context, db *sqlx.DB) {
	var users []models.User
	//users = make([]database.User, 0)

	err := db.Select(&users, `SELECT * FROM users`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func HandleUsersAndAddres_GET(c *gin.Context, db *sqlx.DB) {
	var users []models.User
	var Address []models.Addres_User
	var UserAdres []models.User_Addr

	err := db.Select(&users, `SELECT * FROM users`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	err = db.Select(&Address, `SELECT * FROM addres_user`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	for _, user := range users {
		for _, addr := range Address {
			if addr.Id == int(user.Id_Addr.Int32) {
				UserAdres = append(UserAdres, models.User_Addr{
					Id:       user.Id,
					Phone:    user.Phone,
					Password: user.Password,
					Name:     user.Name,
					Surname:  user.Surname,
					Role:     user.Role,
					IpAddres: user.IpAddres,
					Addres:   addr,
				})
			}
		}
	}

	c.JSON(http.StatusOK, &UserAdres)
}

func HandleUsersIsRoleGET(c *gin.Context, db *sqlx.DB) {
	var users []models.User

	role := c.Param("role")
	err := db.Select(&users, `SELECT * FROM users WHERE role = $1`, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func HandleUserPOST(c *gin.Context, db *sqlx.DB) {
	var user models.User

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

func HandleUserDEL(c *gin.Context, db *sqlx.DB) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	var deletedUser models.User

	err := db.QueryRow(`DELETE FROM users WHERE id = $1`, user.Id).Scan(&deletedUser.Id, &deletedUser.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"Error": "No rows found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"Deleted_Maincategory": deletedUser})
}

func HandleUserPUT(c *gin.Context, db *sqlx.DB) {
	var user models.User
	var userDB models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := db.Get(&userDB, `SELECT * FROM users WHERE id = $1`, user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
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
		return
	}

	c.JSON(http.StatusOK, &userDB)

}

func HandleUserId_GET(c *gin.Context, db *sqlx.DB) {
	var Users []models.User_Addr

	err := db.Select(&Users, `SELECT * FROM users`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, val := range Users {
		if val.Id == id {
			c.JSON(http.StatusOK, &val)
			return
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("There is NO such User with id = %d", id)})

}
