package store

import (
	"time"

	expirable "github.com/hashicorp/golang-lru/v2/expirable"
)

type InmemStore struct {
	cache *expirable.LRU[string, string]
}

func NewInmemStore() *InmemStore {
	return &InmemStore{
		cache: expirable.NewLRU[string, string](
			1023,
			nil,
			24*time.Hour,
		),
	}
}

func (s *InmemStore) AddUser(id string, token string) User {
	s.cache.Add(id, token)
	return User{
		ID:    id,
		Token: token,
	}
}

func (s *InmemStore) DeleteUser(id string) {
	s.cache.Remove(id)
}

func (s *InmemStore) GetUser(id string) *User {
	token, ok := s.cache.Get(id)
	if !ok {
		return nil
	}
	return &User{
		ID:    id,
		Token: token,
	}
}
