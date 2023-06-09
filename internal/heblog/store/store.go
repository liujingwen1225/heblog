package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	S    *datastore
)

type IStore interface {
	Users() UserStore
}

var _ IStore = &datastore{}

type datastore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		S = &datastore{db: db}
	})
	return S
}

func (ds *datastore) Users() UserStore {
	return newUserStore(ds.db)
}
