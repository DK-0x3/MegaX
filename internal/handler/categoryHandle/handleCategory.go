package categoryHandle

import (
	"MegaX/internal/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// ? Запросы для category

func HandleCategoryGETid(c *gin.Context, db *sqlx.DB) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := db.Get(&category, `SELECT id FROM category WHERE name = $1`, category.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func HandleCategoryGET(c *gin.Context, db *sqlx.DB) {
	var category []models.Category

	err := db.Select(&category, `SELECT * FROM category`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func HandleCategoryPOST(c *gin.Context, db *sqlx.DB) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(`INSERT INTO category (name, main_category) VALUES ($1,$2) RETURNING id`, category.Name, category.Main_Category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &category)
}

func HandleCategoryDEL(c *gin.Context, db *sqlx.DB) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	_, err := db.Exec(`DELETE FROM category WHERE id = $1`, category.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &category)
}

func HandleCategoryPUT(c *gin.Context, db *sqlx.DB) {
	var category models.Category
	var categoryDB models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := db.Get(&categoryDB, `SELECT * FROM category WHERE id = $1`, category.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if category.Name != "" {
		categoryDB.Name = category.Name
	}

	if category.Main_Category != 0 {
		categoryDB.Main_Category = category.Main_Category
	}

	_, err = db.Exec(`UPDATE category SET name = $1,main_category=$2  WHERE id = $3`, categoryDB.Name, categoryDB.Main_Category, categoryDB.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}
	_, err = db.Exec(`UPDATE category SET name = $1,main_category=$2  WHERE id = $3`, categoryDB.Name, categoryDB.Main_Category, category.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}
	_, err = db.Exec(`UPDATE category SET name = $1,main_category=$2  WHERE id = $3`, categoryDB.Name, categoryDB.Main_Category, category.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &categoryDB)

}
