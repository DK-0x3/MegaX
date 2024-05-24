package handle

import (
	"MegaX/database"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func HandleParamsPOST(c *gin.Context, db *sqlx.DB) {
	var param database.Parameters

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(`INSERT INTO parameters (name, value, id_product) VALUES ($1,$2,$3) RETURNING id`, param.Name, param.Value, param.Id_Product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &param)
}

func HandleParamsDEL(c *gin.Context, db *sqlx.DB) {
	var param database.Parameters

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	var deletedParam database.Parameters

	err := db.QueryRow(`DELETE FROM parameters WHERE id = $1`, param.Id).Scan(&deletedParam.Id, &deletedParam.Name, &deletedParam.Value, &deletedParam.Id_Product)
	if err != nil {
    	if err == sql.ErrNoRows {
       		c.JSON(http.StatusNotFound, gin.H{"Error": "No rows found"})
    	} else {
    	    c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
    	}
    return
	}
    
	c.JSON(http.StatusOK, gin.H{"Deleted_Maincategory": deletedParam})
}

func HandleParamsPUT(c *gin.Context, db*sqlx.DB) {
	var param database.Parameters
	var paramDB database.Parameters

	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := db.Get(&paramDB, `SELECT * FROM parameters WHERE id = $1`, param.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if param.Name != "" {
		paramDB.Name = param.Name
	}
	if param.Value != "" {
		paramDB.Value = param.Value
	}
	if param.Id_Product != 0 {
		paramDB.Id_Product = param.Id_Product
	}

	_, err = db.Exec(`UPDATE parameters SET name = $1, value = $2, id_product = $3 WHERE id = $4`, paramDB.Name, paramDB.Value, paramDB.Id_Product, paramDB.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &paramDB)
}