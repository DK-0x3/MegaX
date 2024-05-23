package handle

import (
	"MegaX/database"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// ? Запросы для MainCategory

func HandleMainCategoryGET(c *gin.Context, db *sqlx.DB) {
	var MainCat []database.Main_Category

	err := db.Select(&MainCat, `SELECT * FROM main_category`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}
	c.JSON(http.StatusOK, &MainCat)
}

func HandleMainCategoryPOST(c *gin.Context, db *sqlx.DB) {
	var MainCat database.Main_Category

	if err := c.ShouldBindJSON(&MainCat); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error: ": err.Error()})
		return
	}

	_, err := db.Exec(`INSERT INTO main_category (name) VALUES ($1) RETURNING id`, MainCat.Name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &MainCat)
}

func HandleMainCategoryPUT(c *gin.Context, db *sqlx.DB) {
	var MainCat database.Main_Category
	var MainCatDB database.Main_Category

	if err := c.ShouldBindJSON(&MainCat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := db.Get(&MainCatDB, `SELECT * FROM main_category WHERE id = $1`, MainCat.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if MainCat.Name != "" {
		MainCatDB.Name = MainCat.Name
	}

	_, err = db.Exec(`UPDATE main_category SET name = $1 WHERE id = $2`, MainCatDB.Name, MainCatDB.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &MainCatDB)

}

func HandleMainCategoryDEL(c *gin.Context, db *sqlx.DB) {
	var MainCat database.Main_Category

	if err := c.ShouldBindJSON(&MainCat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	var deletedCategory database.Main_Category

	err := db.QueryRow(`DELETE FROM main_category WHERE id = $1 RETURNING id, name`, MainCat.Id).Scan(&deletedCategory.Id, &deletedCategory.Name)
	if err != nil {
    	if err == sql.ErrNoRows {
       		c.JSON(http.StatusNotFound, gin.H{"Error": "No rows found"})
    	} else {
    	    c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
    	}
    return
	}
    
	c.JSON(http.StatusOK, gin.H{"Deleted_Maincategory": deletedCategory})
}

func HandleMainCategoryAndCategory(c *gin.Context, db *sqlx.DB) {
	var MainCat []database.Main_Category
	var Cat []database.Category
	var MainCatAndCat []database.Main_CategoryAndCategory

	err := db.Select(&MainCat, `SELECT * FROM main_category`) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	err = db.Select(&Cat, `SELECT * FROM category`) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	for _, mainVal := range MainCat {
		var CategoryIsMainCat []database.Category

		for _, Val := range Cat {
			if Val.Main_Category == mainVal.Id {
				CategoryIsMainCat = append(CategoryIsMainCat, Val)
			}
		}
		MainCatAndCat = append(MainCatAndCat, database.Main_CategoryAndCategory{
			Id: mainVal.Id,
			Name: mainVal.Name,
			Categories: CategoryIsMainCat,
		})
	}

	c.JSON(http.StatusOK, &MainCatAndCat)
}

func HandleMainCategoryId_GET(c *gin.Context, db *sqlx.DB) {
	var MainCat []database.Main_Category

	err := db.Select(&MainCat, `SELECT * FROM main_category`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, val := range MainCat {
		if val.Id == id {
			c.JSON(http.StatusOK, &val)
			return
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("There is NO such MainCategory with id = %d", id)})
}

func HandleMainCategoryAndCategoryId_GET(c *gin.Context, db *sqlx.DB) {
	var MainCat []database.Main_Category
	var Cat []database.Category

	err := db.Select(&MainCat, `SELECT * FROM main_category`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err = db.Select(&Cat, `SELECT * FROM category`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}
	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, mainVal := range MainCat {
		if mainVal.Id == id {
			var cateforyis []database.Category

			for _, val := range Cat {
				if val.Main_Category == mainVal.Id {
					cateforyis = append(cateforyis, val)
				}
			}
			var result = database.Main_CategoryAndCategory{
				Id: mainVal.Id,
				Name: mainVal.Name,
				Categories: cateforyis,
			}
			c.JSON(http.StatusOK, &result)
			return
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("There is NO such MainCategory with id = %d", id)})
}