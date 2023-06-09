package store

import (
	"context"
	"gorm.io/gorm"
	"heblog/internal/pkg/model"
)

type UserStore interface {
	Create(cxt context.Context, user *model.UserM) error
}

var _ UserStore = &users{}

type users struct {
	db *gorm.DB
}

func newUserStore(db *gorm.DB) *users {
	return &users{db: db}
}

func (u *users) Create(cxt context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}
