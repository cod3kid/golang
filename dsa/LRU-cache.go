package main

import (
	"container/list"
	"fmt"
)

type entry struct {
	key   string
	value int
}
// O(1) time complexity
type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
}

// capacity 5
// "f" <-> "e" <-> "d" <-> "c" <->"a"  <-> nil
func (lru *LRUCache) Get(key string) (int, bool) {
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return 0, false
}

func (lru *LRUCache) Put(key string, value int) {
	if elem, ok := lru.cache[key]; ok {
		elem.Value.(*entry).value = value
		lru.list.MoveToFront(elem)
		return
	}

	if lru.list.Len() >= lru.capacity {
		lruElem := lru.list.Back()
		if lruElem != nil {
			lru.list.Remove(lruElem)
			delete(lru.cache, lruElem.Value.(*entry).key)
		}
	}

	elem := lru.list.PushFront(&entry{key, value})
	lru.cache[key] = elem
}

func main() {
	lruCache := &LRUCache{
		capacity: 2,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}

	lruCache.Put("youtube", 7)
	lruCache.Put("instagram", 34)
	lruCache.Put("tiktok", 567)
	lruCache.Put("facebook", 56)

	fmt.Println(lruCache.Get("youtube"))
}
