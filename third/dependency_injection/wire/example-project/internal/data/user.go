package data

import (
	"context"
	"fmt"
	"go-app/third/dependency_injection/wire/example-project/internal/biz"
	"go-app/third/dependency_injection/wire/example-project/pkg/jsonx"
	"time"

	verr "github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	userCachePrefix     = "userCache"
	userCacheExpiration = time.Minute
)

func GetUserCacheKey(id int64) string {
	return fmt.Sprintf("%s:%d", userCachePrefix, id)
}

type UserRepo struct {
	data *Data
}

func NewUserRepo(logger *zap.SugaredLogger, data *Data) *UserRepo {
	return &UserRepo{data: data}
}

func (u *UserRepo) CreateUser(ctx context.Context, user *biz.User) error {
	if err := u.data.DB.WithContext(ctx).Create(user).Error; err != nil {
		return verr.WithStack(err)
	}
	return nil
}

func (u *UserRepo) DeleteUserByID(ctx context.Context, id int64) error {
	if err := u.data.DB.WithContext(ctx).Where("id = ?", id).Delete(&biz.User{}).Error; err != nil {
		return verr.WithStack(err)
	}
	return u.removeUserCache(ctx, id)
}

func (u *UserRepo) UpdateUserByID(ctx context.Context, id int64, update map[string]interface{}) error {
	if err := u.data.DB.WithContext(ctx).Where("id = ?", id).Updates(update).Error; err != nil {
		return verr.WithStack(err)
	}
	return u.removeUserCache(ctx, id)
}

func (u *UserRepo) GetUserByID(ctx context.Context, id int64) (*biz.User, error) {
	var user biz.User
	if err := u.data.DB.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, verr.WithStack(err)
	}
	if err := u.addUserCache(ctx, &user); err != nil {
		return &user, verr.WithStack(err)
	}
	return &user, nil
}

func (u *UserRepo) removeUserCache(ctx context.Context, id int64) error {
	key := GetUserCacheKey(id)
	if _, err := u.data.RC.Del(ctx, key).Result(); err != nil {
		return verr.WithStack(err)
	}
	return nil
}

func (u *UserRepo) addUserCache(ctx context.Context, user *biz.User) error {
	key := GetUserCacheKey(user.ID)
	val := jsonx.ToStr(user)
	if _, err := u.data.RC.Set(ctx, key, val, userCacheExpiration).Result(); err != nil {
		return verr.WithStack(err)
	}
	return nil
}
