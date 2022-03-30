package route

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/usermanage/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	authMiddleware := InitJwt()

	router.POST("/login", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	initRouteWithAuth(router, authMiddleware)
	return router
}

func initRouteWithAuth(router *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	auth := router.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	auth.POST("/user", handler.SaveUser)
	auth.POST("/user/list", handler.UserList)
}
