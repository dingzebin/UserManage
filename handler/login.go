package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/usermanage/model"
)

func Login(context *gin.Context) {
	u := model.User{}
	if err := context.BindJSON(&u); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user := model.GetUser(u.Account, u.Password)
	if user == nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, "The account or password is incorrect")
		return
	}
}
