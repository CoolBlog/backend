package main

import (
	"github.com/gin-gonic/gin"
	. "backend/apis"
	. "backend/front"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", Index)
	router.GET("/u/:id", UserHomePage)
	router.GET("/a/:id", ArticleDetail)

	router.GET("/articles", GetArticlesAPI)
	router.POST("/articles", AddArticleAPI)
	router.GET("/articles/:id", GetArticleAPI)
	router.PUT("/articles/:id", ModArticleAPI)

	return router
}