package controller

import (
	"context"
	"go-app/third/dependency_injection/wire/example-project/internal/biz"
	"go-app/third/dependency_injection/wire/example-project/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(NewUserController)

type UserController struct {
	userBiz *biz.UserBiz
	logger  *zap.SugaredLogger
}

func NewUserController(logger *zap.SugaredLogger, userBiz *biz.UserBiz) *UserController {
	return &UserController{
		userBiz: userBiz,
		logger:  logger,
	}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var req model.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "user format invalid")
		return
	}

	res, err := u.userBiz.CreateUser(context.Background(), &req)
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
