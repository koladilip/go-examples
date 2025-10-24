package utils

type LRUCache[K comparable, T any] struct {
	capacity int
	cache    map[K]*DoublyLinkedListNode[K, T]
	head     *DoublyLinkedListNode[K, T]
	tail     *DoublyLinkedListNode[K, T]
}

type DoublyLinkedListNode[K comparable, T any] struct {
	key   K
	value T
	prev  *DoublyLinkedListNode[K, T]
	next  *DoublyLinkedListNode[K, T]
}

func NewLRUCache[K comparable, T any](capacity int) *LRUCache[K, T] {
	return &LRUCache[K, T]{
		capacity: capacity,
		cache:    make(map[K]*DoublyLinkedListNode[K, T], capacity),
	}
}

func (lru *LRUCache[K, T]) Get(key K) (T, bool) {
	node, ok := lru.cache[key]
	if !ok {
		var zero T
		return zero, false
	}
	lru.moveToHead(node)
	return node.value, true
}

func (lru *LRUCache[K, T]) moveToHead(node *DoublyLinkedListNode[K, T]) {
	if node == lru.head {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if lru.tail == node {
		lru.tail = node.prev
	}
	node.prev = nil
	node.next = lru.head
	lru.head.prev = node
	lru.head = node
}

func (lru *LRUCache[K, T]) Put(key K, value T) {
	node, ok := lru.cache[key]
	if ok {
		node.value = value
		lru.moveToHead(node)
		return
	}
	if len(lru.cache) >= lru.capacity {
		delete(lru.cache, lru.tail.key)
		lru.tail = lru.tail.prev
		if lru.tail != nil {
			lru.tail.next = nil
		}
	}
	node = &DoublyLinkedListNode[K, T]{
		key:   key,
		value: value,
	}
	lru.cache[key] = node
	if lru.head == nil {
		lru.head = node
		lru.tail = node
		return
	}
	lru.moveToHead(node)
}
