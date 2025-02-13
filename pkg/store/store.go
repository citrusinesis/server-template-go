package store

import (
	"context"
	"sync"
)

type Key string

type Store struct {
	mu    sync.RWMutex
	store map[Key]any
}

func NewStore() *Store {
	return &Store{
		store: make(map[Key]any),
	}
}

func (s *Store) Get(key Key) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.store[key]
	return val, ok
}

func (s *Store) Set(key Key, val any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = val
}

func (s *Store) Delete(key Key) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
}

func (s *Store) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store = make(map[Key]any)
}

type storeContextKey struct{}

func WithStore(ctx context.Context, store *Store) context.Context {
	return context.WithValue(ctx, storeContextKey{}, store)
}

func FromContext(ctx context.Context) (*Store, bool) {
	store, ok := ctx.Value(storeContextKey{}).(*Store)
	return store, ok
}
