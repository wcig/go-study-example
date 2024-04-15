package biz

import (
	"context"
	"go-app/third/dependency_injection/wire/example-project/internal/model"
	"time"

	verr "github.com/pkg/errors"

	"github.com/google/wire"

	"go.uber.org/zap"
)

var ProviderSet = wire.NewSet(NewUserBiz)

type User struct {
	ID       int64  `gorm:"primaryKey;column:id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	CreateAt int64  `gorm:"column:create_at" json:"create_at"`
	UpdateAt int64  `gorm:"column:update_at" json:"update_at"`
}

func (u *User) TableName() string {
	return "user"
}

func CreateUserReqToUser(req *model.CreateUserReq) *User {
	return &User{
		Name:     req.Name,
		CreateAt: time.Now().Unix(),
		UpdateAt: time.Now().Unix(),
	}
}

func UserToCreateUserRes(user *User) *model.CreateUserRes {
	return &model.CreateUserRes{
		ID:       user.ID,
		Name:     user.Name,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
}

type UserBizRepo interface {
	CreateUser(ctx context.Context, user *User) error
	DeleteUserByID(ctx context.Context, id int64) error
	UpdateUserByID(ctx context.Context, id int64, update map[string]interface{}) error
	GetUserByID(ctx context.Context, id int64) (*User, error)
}

type UserBiz struct {
	userRepo UserBizRepo
	logger   *zap.SugaredLogger
}

func NewUserBiz(logger *zap.SugaredLogger, userRepo UserBizRepo) *UserBiz {
	return &UserBiz{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (u *UserBiz) CreateUser(ctx context.Context, req *model.CreateUserReq) (res *model.CreateUserRes, err error) {
	user := CreateUserReqToUser(req)
	if err = u.userRepo.CreateUser(ctx, user); err != nil {
		return nil, verr.WithStack(err)
	}
	res = UserToCreateUserRes(user)
	return res, nil
}
