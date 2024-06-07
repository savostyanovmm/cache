package cache

import (
	"errors"
)

type Cache struct {
	items map[string]Item
}

type Item struct {
	Value interface{}
	Key   string
}

func New() *Cache {
	items := make(map[string]Item)
	cache := Cache{
		items: items,
	}
	return &cache
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = Item{
		Value: value,
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	return item.Value, true
}

func (c *Cache) Delete(key string) error {
	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}
	delete(c.items, key)
	return nil
}
