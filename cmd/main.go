package main

import (
	handle "MegaX/handle"
	"MegaX/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

var server *gin.Engine

const connectionString = "host=127.0.0.1 port=5432 user=postgres password=123456 dbname=mega_xxx sslmode=disable"

func main() {
	var err error
	db, err = handle.InitDB(connectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	server = gin.Default()
	middlewares.LogFile(server)

	//? Работа с юзерами
	server.GET("/users", func(c *gin.Context) {
		handle.HandleUsersGET(c, db)
	})
	server.GET("/users/addres", func(c *gin.Context) {
		handle.HandleUsersAndAddres_GET(c, db)
	})
	server.POST("/users", func(c *gin.Context) {
		handle.HandleUserPOST(c, db)
	})
	server.PUT("/users", func(c *gin.Context) {
		handle.HandleUserPUT(c, db)
	})
	server.DELETE("/users", func(c *gin.Context) {
		handle.HandleUserDEL(c, db)
	})

	server.GET("/users/:role", func(c *gin.Context) {
		handle.HandleUsersIsRoleGET(c, db)
	})

	server.GET("/user/:id", func (c *gin.Context) {
		handle.HandleUserId_GET(c, db)
	})


	//? Работа с аддресами
	server.GET("/addresUser", func(c *gin.Context) {
		handle.HandleAddresGET(c, db)
	})
	server.POST("/addresUser", func(c *gin.Context) {
		handle.HandleAddresPOST(c, db)
	})
	server.DELETE("/addresUser", func(c *gin.Context) {
		handle.HandleAddresDEL(c, db)
	})
	server.PUT("/addresUser", func(c *gin.Context) {
		handle.HandleAddresPUT(c, db)
	})

	//? Работа с MainCategory

	server.GET("/maincategory", func(c *gin.Context) {
		handle.HandleMainCategoryGET(c, db)
	})
	server.POST("/maincategory", func(c *gin.Context) {
		handle.HandleMainCategoryPOST(c, db)
	})
	server.PUT("/maincategory", func(c *gin.Context) {
		handle.HandleMainCategoryPUT(c, db)
	})
	server.DELETE("/maincategory", func(c *gin.Context) {
		handle.HandleMainCategoryDEL(c, db)
	})
	server.GET("/maincategory/category", func(c *gin.Context) {
		handle.HandleMainCategoryAndCategory(c, db)
	})
	server.GET("/maincategory/:id", func(c *gin.Context) {
		handle.HandleMainCategoryId_GET(c, db)
	})
	server.GET("/maincategory/:id/category", func (c *gin.Context) {

		handle.HandleMainCategoryAndCategoryId_GET(c, db)
	})

	//? Работа с категориями
	server.GET("/category", func(c *gin.Context) {
		handle.HandleCategoryGET(c, db)
	})
	server.GET("/categoryId", func(c *gin.Context) {
		handle.HandleCategoryGETid(c, db)
	})
	server.POST("/category", func(c *gin.Context) {
		handle.HandleCategoryPOST(c, db)
	})
	server.DELETE("/category", func(c *gin.Context) {
		handle.HandleCategoryDEL(c, db)
	})
	server.PUT("/category", func(c *gin.Context) {
		handle.HandleCategoryPUT(c, db)
	})

	//? Работа с продуктом
	server.GET("/product", func(c *gin.Context) {
		handle.HandleProductGET(c, db)
	})
	server.GET("/product/price/max", func(c *gin.Context) {
		handle.HandleGetProductPriceMax(c, db)
	})
	server.GET("/product/price/min", func(c *gin.Context) {
		handle.HandleGetProductPriceMin(c, db)
	})
	server.GET("/product/name/desc", func(c *gin.Context) {
		handle.HandleGetProductNameDesc(c, db)
	})
	server.GET("/product/name/asc", func(c *gin.Context) {
		handle.HandleGetProductNameAsc(c, db)
	})
	server.GET("/product/category/desc", func(c *gin.Context) {
		handle.HandleGetProductCategoryDesc(c, db)
	})
	server.GET("/product/category/asc", func(c *gin.Context) {
		handle.HandleGetProductCategoryAsc(c, db)
	})
	server.GET("/product/:id", func(c *gin.Context) {
		handle.HandleProductGETid(c, db)
	})
	server.GET("/product/:id/:params", func(c *gin.Context) {
		handle.HandleProductGETid(c, db)
	})
	server.GET("/productCategory", func(c *gin.Context) {
		handle.HandleProductGETCategory(c, db)
	})
	server.GET("/productMainCategory", func(c *gin.Context) {
		handle.HandleProductGETMainCategory(c, db)
	})
	server.POST("/product", func(c *gin.Context) {
		handle.HandleProductPOST(c, db)
	})
	server.DELETE("/product", func(c *gin.Context) {
		handle.HandleProductDEL(c, db)
	})
	server.PUT("/product", func(c *gin.Context) {
		handle.HandleProductPUT(c, db)
	})
	server.POST("/product/params", func(c *gin.Context) {
		handle.HandleParamsPOST(c, db)
	})
	server.DELETE("/product/params", func(c *gin.Context) {
		handle.HandleParamsDEL(c, db)
	})
	server.PUT("/product/params", func(c *gin.Context) {
		handle.HandleParamsPUT(c, db)
	})

	server.Run(":8080")

}
