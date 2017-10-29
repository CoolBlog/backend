package commons

// AuthMiddleWare
/*
session_id = cookie.get('session_id')
if session_id is not None:
	session_time = session_cache.get(session_id)
	if time.now - session_time > one_week:
		context.user = anonymous_user
	else:
		context.user = user_info_cache.get(session_id)
else:
	context.user = anonymous_user

context.Next()
return
*/


/*

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}
router.GET("/auth/signin", func (c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "123",
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)
	c.String(http.StatusOK, "Login successful")
})

router.GET("/home", AuthMiddleWare(), func (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "home"})
})
*/
