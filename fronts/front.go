package fronts

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is index"})
}

func UserHomePage(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is user homepage"})
}

func ArticleDetail(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is article detail"})
}