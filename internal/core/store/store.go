package store

import (
	"context"
	"example/pkg/store"
)

const (
	ExampleKey store.Key = "example"
	// Add your store key
)

func Get[T any](ctx context.Context, key store.Key) (T, bool) {
	var defaultValue T

	if ctx == nil || key == "" {
		return defaultValue, false
	}

	storeInstance, ok := store.FromContext(ctx)
	if !ok {
		return defaultValue, false
	}

	val, ok := storeInstance.Get(key)
	if !ok {
		return defaultValue, false
	}

	castVal, okType := val.(T)
	if !okType {
		return defaultValue, false
	}

	return castVal, true
}

func Set[T any](ctx context.Context, key store.Key, val T) bool {
	if ctx == nil || key == "" {
		return false
	}

	storeInstance, ok := store.FromContext(ctx)
	if !ok {
		return false
	}

	storeInstance.Set(key, val)
	_, ok = storeInstance.Get(key)
	return ok
}
