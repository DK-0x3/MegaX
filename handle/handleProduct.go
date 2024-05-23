package handle

import (
	"MegaX/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// ? Запросы для product

func HandleProductGETid(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := db.Get(&product, `SELECT id FROM product WHERE name = $1`, product.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func HandleProductGETCategory(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := db.Get(&product, `SELECT DISTINCT category.name FROM category
							INNER JOIN product ON product.category = category.id
							WHERE product.category = $1`, product.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func HandleProductGETMainCategory(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := db.Get(&product, `SELECT DISTINCT main_category.name FROM main_category
							INNER JOIN category ON category.main_category = main_category.id
							INNER JOIN product ON product.category = category.id
							WHERE product.category = $1`, product.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func HandleProductGET(c *gin.Context, db *sqlx.DB) {
	var product []database.Product

	err := db.Select(&product, `SELECT * FROM product`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func HandleProductPOST(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверить, существует ли категория
	var categoryExists bool
	err := db.Get(&categoryExists, `SELECT EXISTS(SELECT 1 FROM category WHERE id=$1)`, product.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось проверить существование категории", "details": err.Error()})
		return
	}

	if !categoryExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Категория не существует"})
		return
	}

	_, err = db.Exec(`INSERT INTO product (name, price, description, category) VALUES ($1, $2, $3, $4) RETURNING id`,
		product.Name, product.Price, product.Description, product.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &product)
}

func HandleProductDEL(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	_, err := db.Exec(`DELETE FROM product WHERE id = $1`, product.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &product)
}

func HandleProductPUT(c *gin.Context, db *sqlx.DB) {
	var product database.Product
	var productDB database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	err := db.Get(&productDB, `SELECT * FROM product WHERE id = $1`, product.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	if product.Name != "" {
		productDB.Name = product.Name
	}

	if product.Price != 0 {
		productDB.Price = product.Price
	}

	if product.Description != "" {
		productDB.Description = product.Description
	}

	if product.Category != 0 {
		productDB.Category = product.Category
	}

	_, err = db.Exec(`UPDATE product SET name = $1, price=$2, description=$3, category=$4 WHERE id = $5`, productDB.Name, productDB.Price, productDB.Description, productDB.Category, productDB.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &productDB)

}
