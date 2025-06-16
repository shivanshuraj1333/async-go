package cache

import (
	"container/list"
	"sync"
)

type entry struct {
	key   string
	value interface{}
}

type Cache struct {
	mu        sync.Mutex
	capacity  int
	items     map[string]*list.Element
	evictList *list.List // List of *entry, most recently used at the front
}

// NewCache initializes an LRU cache with the given capacity
func NewCache(capacity int) *Cache {
	return &Cache{
		capacity:  capacity,
		items:     make(map[string]*list.Element),
		evictList: list.New(),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// If item exists, update and move to front
	if el, found := c.items[key]; found {
		el.Value.(*entry).value = value
		c.evictList.MoveToFront(el)
		return
	}

	// If at capacity, evict least recently used item
	if c.evictList.Len() >= c.capacity {
		backEl := c.evictList.Back()
		if backEl != nil {
			c.evictList.Remove(backEl)
			delete(c.items, backEl.Value.(*entry).key)
		}
	}

	// Insert new item
	en := &entry{key: key, value: value}
	el := c.evictList.PushFront(en)
	c.items[key] = el
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, found := c.items[key]; found {
		c.evictList.MoveToFront(el)
		return el.Value.(*entry).value, true
	}
	return nil, false
}
