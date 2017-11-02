package fronts

import (
	"github.com/gin-gonic/gin"
	. "backend/commons"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.tmpl", gin.H{"static": STATIC_URL})

}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is register"})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is login"})
}

func Logout(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is logout"})
}

func UserHomePage(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is user homepage"})
}

func ArticleDetail(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "this is article detail"})
}
