package cache

import (
	"sync"
	"time"

	"github.com/MatthewZito/gouache/models"
)

type cacheRecord struct {
	value   interface{}
	expires int64
}

type Cache struct {
	state map[string]*cacheRecord
	l     sync.Mutex
}

func NewCache() *Cache {
	c := &Cache{
		state: make(map[string]*cacheRecord),
	}

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				c.l.Lock()
				for k, v := range c.state {
					if v.Expired(time.Now().UnixNano()) {
						delete(c.state, k)
					}
				}
				c.l.Unlock()
			}
		}
	}()

	return c
}

func (cr *cacheRecord) Expired(time int64) bool {
	if cr.expires == 0 {
		return true
	}
	return time > cr.expires
}

func (c *Cache) Get(key string) interface{} {
	c.l.Lock()
	defer c.l.Unlock()

	if v, ok := c.state[key]; ok {
		return v.value
	}

	return nil
}

func (c *Cache) Put(key string, value interface{}, expires int64) {
	c.l.Lock()
	defer c.l.Unlock()

	if v, ok := c.state[key]; ok {
		v.expires = expires
		v.value = value
	} else {
		c.state[key] = &cacheRecord{
			value:   value,
			expires: expires,
		}
	}
}

func (c *Cache) Delete(key string) bool {
	c.l.Lock()
	defer c.l.Unlock()

	if _, ok := c.state[key]; ok {
		delete(c.state, key)
		return true
	}

	return false
}

func (c *Cache) All() []models.Resource {
	resources := []models.Resource{}

	for k, r := range c.state {
		resources = append(resources, models.Resource{
			Key:     k,
			Value:   r.value,
			Expires: r.expires,
		})
	}

	return resources
}
