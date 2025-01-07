package store

import (
	"context"
	"example/pkg/store"
)

const (
	ExampleKey store.StoreKey = "example"
	// Add more keys here
)

func Get[T any](ctx context.Context, key store.StoreKey) (T, bool) {
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

	return val.(T), true
}

func Store[T any](ctx context.Context, key store.StoreKey, val T) bool {
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
