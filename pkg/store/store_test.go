package store_test

import (
	"context"
	store "example/pkg/store"
	"testing"
)

func TestStore(t *testing.T) {
	t.Run("basic operations", func(t *testing.T) {
		s := store.NewStore()

		// Test Set and Get
		key := store.StoreKey("test-key")
		val := "test-value"
		s.Set(key, val)

		got, exists := s.Get(key)
		if !exists {
			t.Error("expected key to exist")
		}
		if got != val {
			t.Errorf("got %v, want %v", got, val)
		}

		// Test Delete
		s.Delete(key)
		_, exists = s.Get(key)
		if exists {
			t.Error("key should not exist after deletion")
		}

		// Test Clear
		s.Set(key, val)
		s.Clear()
		_, exists = s.Get(key)
		if exists {
			t.Error("store should be empty after clear")
		}
	})

	t.Run("context operations", func(t *testing.T) {
		s := store.NewStore()
		ctx := context.Background()

		// Test WithStore
		ctxWithStore := store.WithStore(ctx, s)

		// Test FromContext
		gotStore, ok := store.FromContext(ctxWithStore)
		if !ok {
			t.Error("expected store to exist in context")
		}
		if gotStore != s {
			t.Error("got different store instance than expected")
		}

		// Test FromContext with empty context
		_, ok = store.FromContext(ctx)
		if ok {
			t.Error("expected no store in empty context")
		}
	})

	t.Run("concurrent operations", func(t *testing.T) {
		s := store.NewStore()
		done := make(chan bool)

		go func() {
			s.Set(store.StoreKey("key1"), "value1")
			done <- true
		}()

		go func() {
			s.Set(store.StoreKey("key2"), "value2")
			done <- true
		}()

		<-done
		<-done

		if _, exists := s.Get(store.StoreKey("key1")); !exists {
			t.Error("key1 should exist")
		}
		if _, exists := s.Get(store.StoreKey("key2")); !exists {
			t.Error("key2 should exist")
		}
	})
}
