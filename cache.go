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

func (c *Cache) Set(key string, value interface{}) error {
	if key := validate(key); key != nil {
		return errors.New("key is empty")
	}

	c.data[key] = value

	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	if key := validate(key); key != nil {
		return nil, errors.New("key is empty")
	}

	val, ok := c.data[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Value with %s doesn't exist", key))
	}
	return val, nil
}

func (c *Cache) Delete(key string) error {
	if key := validate(key); key != nil {
		return errors.New("key is empty")
	}

	if _, ok := c.data[key]; !ok {
		return errors.New(fmt.Sprintf("Value with %s doesn't exist", key))
	}
	delete(c.data, key)
	return nil
}

func validate(val string) error {
	if val == "" {
		return errors.New("key is empty")
	}

	return nil
}
