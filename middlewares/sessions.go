package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("7f4b8a6c2d8b72e5b59c8a61f0d9a04f5b47f5b2587a9c8a8b07bfb32f5d10c5"))

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, "mysession")
		c.Set("session", session)
		c.Next()
	}
}
