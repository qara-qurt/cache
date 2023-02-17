package cache

import (
	"errors"
	"fmt"
)

type Cache struct {
	data map[string]interface{}
}

func New() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, error) {
	val, ok := c.data[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Value with %s doesn't exist", key))
	}
	return val, nil
}

func (c *Cache) Delete(key string) error {
	if _, ok := c.data[key]; !ok {
		return errors.New(fmt.Sprintf("Value with %s doesn't exist", key))
	}
	delete(c.data, key)
	return nil
}
