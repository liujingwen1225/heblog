package user

import (
	"context"
	"github.com/jinzhu/copier"
	"heblog/internal/heblog/store"
	"heblog/internal/pkg/errno"
	"heblog/internal/pkg/model"
	v1 "heblog/pkg/api/heblog/v1"
	"regexp"
)

type UserBiz interface {
	Create(cxt context.Context, request *v1.CreateUserRequest) error
}

var _ UserBiz = &userBiz{}

type userBiz struct {
	ds store.IStore
}

func NewUserBiz(ds store.IStore) *userBiz {
	return &userBiz{ds: ds}
}

func (u *userBiz) Create(cxt context.Context, request *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, request)
	//log.C(cxt).Infow(fmt.Sprintf("%s"), userM)
	if err := u.ds.Users().Create(cxt, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}
	return nil
}
