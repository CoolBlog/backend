package main

import (
	"github.com/gin-gonic/gin"
	. "backend/apis"
	. "backend/fronts"
	. "backend/commons"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Use(AuthMiddleWare())

	router.GET("/", Index)
	router.POST("/register", Register)
	router.POST("/login", Login)
	router.POST("/logout", Logout)
	router.GET("/u/:id", UserHomePage)
	router.GET("/a/:id", ArticleDetail)

	router.GET("/articles", GetArticlesAPI)
	router.POST("/articles", AddArticleAPI)
	router.GET("/articles/:id", GetArticleAPI)
	router.PUT("/articles/:id", ModArticleAPI)

	return router
}
