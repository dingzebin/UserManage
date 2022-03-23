package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/usermanage/model"
	"net/http"
)

/**
 * 保存用户
 */
func SaveUser(context *gin.Context) {
	u := model.User{}
	if err := context.BindJSON(&u); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	fmt.Println(u)
	model.SaveUser(&u)
	context.JSON(http.StatusAccepted, u.Id)
}
