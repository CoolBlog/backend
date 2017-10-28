package apis

import "github.com/gin-gonic/gin"

func GetArticlesAPI(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
}

func AddArticleAPI(c *gin.Context) {

}

func GetArticleAPI(c *gin.Context) {

}

func ModArticleAPI(c *gin.Context) {

}
