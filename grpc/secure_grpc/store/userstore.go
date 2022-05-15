// Package store contains ...
package store

import (
	"fmt"
	"github.com/connect2naga/go-examples/grpc/secure_grpc/authz"
	"sync"
)

/*
Author : Nagarjuna S
Date : 15-05-2022 23:41
Project : secure_grpc
File : userstore.go
*/

type UserStore interface {
	Save(user *authz.User) error
	Find(username string) (*authz.User, error)
}

type InMemoryUserStore struct {
	mutex sync.RWMutex
	users map[string]*authz.User
}

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{
		users: make(map[string]*authz.User),
	}
}

func (store *InMemoryUserStore) Save(user *authz.User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return fmt.Errorf("User Already exist..")
	}

	store.users[user.Username] = user.Clone()
	return nil
}

func (store *InMemoryUserStore) Find(username string) (*authz.User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.users[username]
	if user == nil {
		return nil, nil
	}

	return user.Clone(), nil
}
