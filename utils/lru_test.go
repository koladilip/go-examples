package utils_test

import (
	"testing"

	"github.com/koladilip/go-examples/utils"
	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	lru := utils.NewLRUCache[int, int](2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	val, ok := lru.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 1, val)
	val, ok = lru.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 2, val)
	lru.Put(3, 3)
	val, ok = lru.Get(3)
	assert.True(t, ok)
	assert.Equal(t, 3, val)
	_, ok = lru.Get(1)
	assert.False(t, ok)
}
