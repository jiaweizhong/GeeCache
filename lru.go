package lru

import (
	"container/list"
)

// LRU cache
type Cache struct {
	maxBytes   int64                         // max memory
	curBytes   int64                         // current used memory
	doubleList *list.List                    // double linkedlist
	cache      map[string]*list.Element      // string : node in double linkedlist
	OnEvicted  func(key string, value Value) // evication callback
}

// entry for cache
type entry struct {
	key   string
	value Value
}

// generic Value as interface
type Value interface {
	Len() int // return memory usage
}

// cache constructor
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {

	return &Cache{
		maxBytes:   maxBytes,
		doubleList: list.New(),
		cache:      make(map[string]*list.Element),
		OnEvicted:  onEvicted,
	}
}
