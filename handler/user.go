package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usermanage/model"
	"github.com/usermanage/util"
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

/**
 * 保存用户
 */
func UserList(context *gin.Context) {
	p := util.Pager{}
	if err := context.BindJSON(&p); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	context.JSON(http.StatusAccepted, model.UserList(p))
}
