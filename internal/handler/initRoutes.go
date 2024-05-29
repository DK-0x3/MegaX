package handler

import (
	"MegaX/internal/handler/categoryHandle"
	"MegaX/internal/handler/productHandle"
	"MegaX/internal/handler/userHandle"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(server *gin.Engine, db *sqlx.DB) {
	//? Работа с юзерами
	server.GET("/users", func(c *gin.Context) {
		userHandle.HandleUserId_GET(c, db)
	})
	server.GET("/users/addres", func(c *gin.Context) {
		userHandle.HandleUsersAndAddres_GET(c, db)
	})
	server.POST("/users", func(c *gin.Context) {
		userHandle.HandleUserPOST(c, db)
	})
	server.PUT("/users", func(c *gin.Context) {
		userHandle.HandleUserPUT(c, db)
	})
	server.DELETE("/users", func(c *gin.Context) {
		userHandle.HandleUserDEL(c, db)
	})

	server.GET("/users/:role", func(c *gin.Context) {
		userHandle.HandleUsersIsRoleGET(c, db)
	})
	server.GET("/user/:id", func(c *gin.Context) {
		userHandle.HandleUserId_GET(c, db)
	})
	//? Работа с аддресами
	server.GET("/addresUser", func(c *gin.Context) {
		userHandle.HandleAddresGET(c, db)
	})
	server.POST("/addresUser", func(c *gin.Context) {
		userHandle.HandleAddresPOST(c, db)
	})
	server.DELETE("/addresUser", func(c *gin.Context) {
		userHandle.HandleAddresDEL(c, db)
	})
	server.PUT("/addresUser", func(c *gin.Context) {
		userHandle.HandleAddresPUT(c, db)
	})

	//? Работа с MainCategory

	server.GET("/maincategory", func(c *gin.Context) {
		categoryHandle.HandleMainCategoryGET(c, db)
	})
	server.POST("/maincategory", func(c *gin.Context) {
		categoryHandle.HandleMainCategoryPOST(c, db)
	})
	server.PUT("/maincategory", func(c *gin.Context) {
		categoryHandle.HandleMainCategoryPUT(c, db)
	})
	server.DELETE("/maincategory", func(c *gin.Context) {
		categoryHandle.HandleMainCategoryDEL(c, db)
	})
	server.GET("/maincategory/category", func(c *gin.Context) {
		categoryHandle.HandleMainCategoryAndCategory(c, db)
	})
	server.GET("/maincategory/:id", func(c *gin.Context) {
		categoryHandle.HandleMainCategoryId_GET(c, db)
	})

	server.GET("/maincategory/:id/category", func(c *gin.Context) {
		categoryHandle.HandleMainCategoryAndCategoryId_GET(c, db)
	})

	//? Работа с категориями
	server.GET("/category", func(c *gin.Context) {
		categoryHandle.HandleCategoryGET(c, db)
	})
	server.GET("/categoryId", func(c *gin.Context) {
		categoryHandle.HandleCategoryGETid(c, db)
	})
	server.POST("/category", func(c *gin.Context) {
		categoryHandle.HandleCategoryPOST(c, db)
	})
	server.DELETE("/category", func(c *gin.Context) {
		categoryHandle.HandleCategoryDEL(c, db)
	})
	server.PUT("/category", func(c *gin.Context) {
		categoryHandle.HandleCategoryPUT(c, db)
	})

	//? Работа с продуктом
	server.GET("/product", func(c *gin.Context) {
		productHandle.HandleProductGET(c, db)
	})
	server.GET("/product/price/max", func(c *gin.Context) {
		productHandle.HandleGetProductPriceMax(c, db)
	})
	server.GET("/product/price/min", func(c *gin.Context) {
		productHandle.HandleGetProductPriceMin(c, db)
	})
	server.GET("/product/name/desc", func(c *gin.Context) {
		productHandle.HandleGetProductNameDesc(c, db)
	})
	server.GET("/product/name/asc", func(c *gin.Context) {
		productHandle.HandleGetProductNameAsc(c, db)
	})
	server.GET("/product/category/desc", func(c *gin.Context) {
		productHandle.HandleGetProductCategoryDesc(c, db)
	})
	server.GET("/product/category/asc", func(c *gin.Context) {
		productHandle.HandleGetProductCategoryAsc(c, db)
	})
	server.GET("/product/:id", func(c *gin.Context) {
		productHandle.HandleProductGETid(c, db)
	})
	server.GET("/product/:id/:params", func(c *gin.Context) {
		productHandle.HandleProductGETid(c, db)
	})
	server.GET("/productCategory", func(c *gin.Context) {
		productHandle.HandleProductGETCategory(c, db)
	})
	server.GET("/productMainCategory", func(c *gin.Context) {
		productHandle.HandleProductGETMainCategory(c, db)
	})
	server.POST("/product", func(c *gin.Context) {
		productHandle.HandleProductPOST(c, db)
	})
	server.DELETE("/product", func(c *gin.Context) {
		productHandle.HandleProductDEL(c, db)
	})
	server.PUT("/product", func(c *gin.Context) {
		productHandle.HandleProductPUT(c, db)
	})
	server.POST("/product/params", func(c *gin.Context) {
		productHandle.HandleParamsPOST(c, db)
	})
	server.DELETE("/product/params", func(c *gin.Context) {
		productHandle.HandleParamsDEL(c, db)
	})
	server.PUT("/product/params", func(c *gin.Context) {
		productHandle.HandleParamsPUT(c, db)
	})
}
