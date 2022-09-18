package his

import (
	"sync"

	"github.com/garyburd/redigo/redis"
)

//record 用于消息记录，累计
type keyValue struct {
	Topic string //主题，唯一值，比如：添加好友
	RWL   sync.RWMutex
	RD    *redis.Pool
}

func (r *keyValue) Set(key string, value string) {
	r.RWL.Lock()
	defer r.RWL.Unlock()
	conn := r.RD.Get()
	_, _ = conn.Do("HSET", "KeyValue_"+r.Topic, key, value)
	conn.Close()
}

func (r *keyValue) Get(key string) (value string, err error) {
	r.RWL.RLock()
	defer r.RWL.RUnlock()
	conn := r.RD.Get()
	defer conn.Close()
	return redis.String(conn.Do("HGET", "KeyValue_"+r.Topic, key))
}

func (r *keyValue) Del(key string) {
	r.RWL.Lock()
	defer r.RWL.Unlock()
	conn := r.RD.Get()
	_, _ = conn.Do("HDEL", "KeyValue_"+r.Topic, key)
	conn.Close()
}
