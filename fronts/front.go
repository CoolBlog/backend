package fronts

import (
	"github.com/gin-gonic/gin"
	. "backend/global"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{"static": STATIC_URL})

}

func UserHomePage(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is user homepage"})
}

func ArticleDetail(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is article detail"})
}
