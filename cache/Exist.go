package cache

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/allegro/bigcache"
)

type cacheExist struct {
	rwl           sync.RWMutex
	defaultSecond uint
	store         *bigcache.BigCache
}
type cacheExistItem struct {
	Value      any   `json:"value"`
	Expiration int64 `json:"expiration"`
}

func (it cacheExistItem) Bytes() []byte {
	marshal, _ := json.Marshal(&it)
	return marshal
}

func (it cacheExistItem) Item(bs []byte) (cacheExistItem, error) {
	var temp cacheExistItem
	err := json.Unmarshal(bs, &temp)
	return temp, err
}

func (c *cacheExist) Set(key string, value any, second uint) {
	c.rwl.Lock()
	if second <= 0 {
		second = c.defaultSecond
	}
	var it = cacheExistItem{
		Value:      value,
		Expiration: time.Now().Add(time.Second * time.Duration(second)).Unix(),
	}
	_ = c.store.Delete(key)
	_ = c.store.Set(key, it.Bytes())
	c.rwl.Unlock()
}

func (c *cacheExist) Get(key string) (any, error) {
	c.rwl.RLock()
	defer c.rwl.RUnlock()
	if v, err := c.store.Get(key); err != nil {
		return nil, err
	} else {
		temp, err := cacheExistItem{}.Item(v)
		if err != nil {
			_ = c.store.Delete(key)
			return nil, err
		}
		if time.Now().Unix() > temp.Expiration {
			_ = c.store.Delete(key)
			return nil, fmt.Errorf("已过期")
		}
		return temp.Value, nil
	}
}

func (c *cacheExist) Exist(key string) bool {
	c.rwl.RLock()
	defer c.rwl.RUnlock()
	if v, err := c.store.Get(key); err != nil {
		return false
	} else {
		temp, err := cacheExistItem{}.Item(v)
		if err != nil {
			_ = c.store.Delete(key)
			return false
		}
		if time.Now().Unix() > temp.Expiration {
			_ = c.store.Delete(key)
			return false
		}
		return true
	}
}

func (c *cacheExist) Del(key string) {
	c.rwl.Lock()
	_ = c.store.Delete(key)
	c.rwl.Unlock()
}

// ExistDelete 存在即删除
// 返回值，存在返回true，不存在返回false
func (c *cacheExist) ExistDelete(key string) bool {
	c.rwl.Lock()
	defer c.rwl.Unlock()
	if v, err := c.store.Get(key); err != nil {
		return false
	} else {
		temp, err := cacheExistItem{}.Item(v)
		if err != nil {
			_ = c.store.Delete(key)
			return false
		}
		if time.Now().Unix() > temp.Expiration {
			_ = c.store.Delete(key)
			return false
		}
		_ = c.store.Delete(key)
		return true
	}
}

// ExistOrSet 存在则不动，不存在则set
// 返回值，存在返回true，不存在返回false
func (c *cacheExist) ExistOrSet(key string, value any, second uint) bool {
	c.rwl.Lock()
	defer c.rwl.Unlock()
	if second <= 0 {
		second = c.defaultSecond
	}
	var it = cacheExistItem{
		Value:      value,
		Expiration: time.Now().Add(time.Second * time.Duration(second)).Unix(),
	}
	if v, err := c.store.Get(key); err != nil {
		_ = c.store.Set(key, it.Bytes())
		return false
	} else {
		temp, err := cacheExistItem{}.Item(v)
		if err != nil {
			_ = c.store.Set(key, it.Bytes())
			return false
		}
		if time.Now().Unix() > temp.Expiration {
			_ = c.store.Set(key, it.Bytes())
			return false
		}
		return true
	}
}

func (c *cacheExist) Clean() {
	c.rwl.Lock()
	_ = c.store.Reset()
	c.rwl.Unlock()
}
