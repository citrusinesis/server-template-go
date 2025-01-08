package store_test

import (
	"context"
	"testing"

	store "example/internal/core/store"
	pkgstore "example/pkg/store"
)

func TestGetAndSet(t *testing.T) {
	ctx := context.Background()

	// Create a new store and attach it to the context.
	s := pkgstore.NewStore()
	ctx = pkgstore.WithStore(ctx, s)

	// Set a string value.
	success := store.Set[string](ctx, store.ExampleKey, "hello world")
	if !success {
		t.Fatalf("Set should have succeeded")
	}

	// Get the string value.
	val, ok := store.Get[string](ctx, store.ExampleKey)
	if !ok {
		t.Fatalf("Get should have returned true for the existing key")
	}
	if val != "hello world" {
		t.Errorf("expected %q, got %q", "hello world", val)
	}

	// Test retrieving a non-existent key.
	_, ok = store.Get[string](ctx, "non_existent_key")
	if ok {
		t.Errorf("Get should have returned false for a non-existent key")
	}
}

func TestTypeMismatch(t *testing.T) {
	ctx := context.Background()

	s := pkgstore.NewStore()
	ctx = pkgstore.WithStore(ctx, s)

	// Set an integer value.
	store.Set[int](ctx, store.ExampleKey, 12345)

	// Attempt to retrieve it as a string.
	val, ok := store.Get[string](ctx, store.ExampleKey)
	if ok {
		t.Fatalf("Get should have returned false when there's a type mismatch, but got val=%v", val)
	}
}

func TestNilContextOrKey(t *testing.T) {
	//lint:ignore SA1012 for test when input is nil
	success := store.Set(nil, store.ExampleKey, "value")
	if success {
		t.Error("Set should fail with nil context")
	}

	//lint:ignore SA1012 for test when input is nil
	val, ok := store.Get[string](nil, store.ExampleKey)
	if ok {
		t.Error("Get should fail with nil context")
	}
	if val != "" {
		t.Errorf("expected empty string, got %q", val)
	}

	// empty key
	ctx := context.Background()
	ctx = pkgstore.WithStore(ctx, pkgstore.NewStore())

	success = store.Set[string](ctx, "", "value")
	if success {
		t.Error("Set should fail with empty key")
	}
	val, ok = store.Get[string](ctx, "")
	if ok {
		t.Error("Get should fail with empty key")
	}
	if val != "" {
		t.Errorf("expected empty string, got %q", val)
	}
}
