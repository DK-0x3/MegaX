package handle

import (
	"MegaX/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// ? Запросы для product

func HandleProductGETid(c *gin.Context, db *sqlx.DB) {
	var Products []database.Product

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = db.Select(&Products, `SELECT * FROM product WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	params := c.Param("params")

	if params == "true" {
		var parameters []database.Parameters
		err = db.Select(&parameters, `SELECT * FROM parameters WHERE id_product = $1`, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
			return
		}
		resultProduct := database.Product_Parameters{
			Id:          Products[0].Id,
			Name:        Products[0].Name,
			Price:       Products[0].Price,
			Description: Products[0].Description,
			Category:    Products[0].Category,
			Parameters:  parameters,
		}
		c.JSON(http.StatusOK, &resultProduct)
		return
	} else {
		if len(Products) > 0 {
			c.JSON(http.StatusOK, &Products)
			return
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("There is NO such Product with id = %d", id)})
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

func HandleGetProductPriceMax(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Select(&product, `SELECT * FROM product 
								ORDER BY product.price DESC`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &product)
}

func HandleGetProductPriceMin(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Select(&product, `SELECT * FROM product 
								ORDER BY product.price ASC`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &product)
}

func HandleGetProductNameDesc(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Select(&product, `SELECT * FROM product 
								ORDER BY product.name DESC`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &product)
}

func HandleGetProductNameAsc(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Select(&product, `SELECT * FROM product 
								ORDER BY product.name ASC`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &product)
}

func HandleGetProductCategoryDesc(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Select(&product, `SELECT * FROM product 
								ORDER BY product.category DESC`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &product)
}

func HandleGetProductCategoryAsc(c *gin.Context, db *sqlx.DB) {
	var product database.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Select(&product, `SELECT * FROM product 
								ORDER BY product.category ASC`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &product)
}
