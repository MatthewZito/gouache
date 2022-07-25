package cache

import (
	"sync"
	"time"
)

type cacheRecord struct {
	value   interface{}
	expires int64
}

type cache struct {
	state map[string]*cacheRecord
	l     sync.Mutex
}

func NewCache() *cache {

	c := &cache{
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

func (c *cache) Get(key string) interface{} {
	c.l.Lock()
	defer c.l.Unlock()

	if v, ok := c.state[key]; ok {
		return v.value
	}

	return nil
}

func (c *cache) Put(key string, value interface{}, expires int64) {
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

func (c *cache) Delete(key string) bool {
	c.l.Lock()
	defer c.l.Unlock()

	if _, ok := c.state[key]; ok {
		delete(c.state, key)
		return true
	}

	return false
}
