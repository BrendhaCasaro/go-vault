package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLRUCache(t *testing.T) {
	lru := NewLRUCache(2)
	require.NotNil(t, lru)
	assert.Equal(t, 2, lru.capacity)
	assert.NotNil(t, lru.items)
	assert.NotNil(t, lru.left)
	assert.NotNil(t, lru.right)
	assert.Equal(t, lru.left.next, lru.right)
	assert.Equal(t, lru.right.prev, lru.left)
}

func TestPutAndGet(t *testing.T) {
	lru := NewLRUCache(2)

	// Inserção simples
	lru.Put("a", "1")
	val, ok := lru.Get("a")
	require.True(t, ok)
	assert.Equal(t, "1", val)

	// Substituição de valor existente
	lru.Put("a", "2")
	val, ok = lru.Get("a")
	require.True(t, ok)
	assert.Equal(t, "2", val)
}

func TestEviction(t *testing.T) {
	lru := NewLRUCache(2)

	lru.Put("a", "1")
	lru.Put("b", "2")
	lru.Put("c", "3") // deve remover o "a"

	_, ok := lru.Get("a")
	assert.False(t, ok, "expected 'a' to be evicted")

	val, ok := lru.Get("b")
	require.True(t, ok)
	assert.Equal(t, "2", val)

	val, ok = lru.Get("c")
	require.True(t, ok)
	assert.Equal(t, "3", val)
}

func TestGetUpdatesOrder(t *testing.T) {
	lru := NewLRUCache(2)

	lru.Put("a", "1")
	lru.Put("b", "2")

	// acesso recente a "a" deve movê-lo pro fim
	lru.Get("a")
	lru.Put("c", "3") // agora "b" deve ser removido, não "a"

	_, ok := lru.Get("b")
	assert.False(t, ok, "expected 'b' to be evicted")

	val, ok := lru.Get("a")
	require.True(t, ok)
	assert.Equal(t, "1", val)

	val, ok = lru.Get("c")
	require.True(t, ok)
	assert.Equal(t, "3", val)
}

func TestDeleteExistingKey(t *testing.T) {
	lru := NewLRUCache(2)

	lru.Put("x", "100")
	err := lru.Delete("x")
	require.NoError(t, err)

	_, ok := lru.Get("x")
	assert.False(t, ok, "expected deleted key to not be found")
}

func TestDeleteNonExistentKey(t *testing.T) {
	lru := NewLRUCache(2)

	err := lru.Delete("missing")
	require.Error(t, err)
	assert.Equal(t, "Error: provided key does not exist", err.Error())
}

func TestMultipleInsertionsAndEvictions(t *testing.T) {
	lru := NewLRUCache(3)

	lru.Put("a", "1")
	lru.Put("b", "2")
	lru.Put("c", "3")

	lru.Get("a")      // a vira o mais recente
	lru.Put("d", "4") // remove "b"

	_, ok := lru.Get("b")
	assert.False(t, ok, "expected 'b' to be evicted")

	val, ok := lru.Get("a")
	require.True(t, ok)
	assert.Equal(t, "1", val)

	val, ok = lru.Get("d")
	require.True(t, ok)
	assert.Equal(t, "4", val)
}
