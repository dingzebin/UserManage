package route

import (
	"github.com/gin-gonic/gin"
	"github.com/usermanage/handler"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/user", handler.SaveUser)
	router.POST("/userList", handler.UserList)
	return router
}
