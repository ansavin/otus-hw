package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity      int
	queue         List
	items         map[Key]*ListItem
	reversedItems map[*ListItem]Key
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity:      capacity,
		queue:         NewList(),
		items:         make(map[Key]*ListItem, capacity),
		reversedItems: make(map[*ListItem]Key, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	if item, ok := c.items[key]; ok {
		if item.Value == value {
			c.queue.MoveToFront(item)
		} else {
			c.queue.Remove(item)
			c.queue.PushFront(value)
		}
		c.items[key] = c.queue.Front()
		return true
	}
	c.queue.PushFront(value)
	c.items[key] = c.queue.Front()
	c.reversedItems[c.queue.Front()] = key
	if c.queue.Len() > c.capacity {
		delete(c.items, c.reversedItems[c.queue.Back()])
		delete(c.reversedItems, c.queue.Back())
		c.queue.Remove(c.queue.Back())
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item, ok := c.items[key]
	if !ok {
		return nil, false
	}
	c.queue.MoveToFront(item)
	res := item.Value
	return res, true
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
	c.reversedItems = make(map[*ListItem]Key, c.capacity)
}
