package his

import (
	"sync"

	"github.com/garyburd/redigo/redis"
)

type exist struct {
	Topic string //主题，唯一值，比如：添加好友
	RWL   sync.RWMutex
	RD    *redis.Pool
}

//IsExist 该key是否存在
func (e *exist) IsExist(key string) (bool, error) {
	e.RWL.RLock()
	defer e.RWL.RUnlock()
	conn := e.RD.Get()
	defer conn.Close()
	i, err := redis.Int(conn.Do("SISMEMBER", "Exist_"+e.Topic, key))
	if err != nil {
		return false, err
	}
	if i == 0 {
		return false, nil
	}
	return true, nil
}

//Add 添加key
func (e *exist) Add(key string) {
	e.RWL.Lock()
	defer e.RWL.Unlock()
	conn := e.RD.Get()
	_, _ = conn.Do("SADD", "Exist_"+e.Topic, key)
	conn.Close()
}

//Del 删除key
func (e *exist) Del(key string) {
	e.RWL.Lock()
	defer e.RWL.Unlock()
	conn := e.RD.Get()
	_, _ = conn.Do("SREM", "Exist_"+e.Topic, key)
	conn.Close()
}

//Count 一共有多少个key
func (e *exist) Count() (sum int, err error) {
	e.RWL.RLock()
	defer e.RWL.RUnlock()
	conn := e.RD.Get()
	defer conn.Close()
	i, err := redis.Int(conn.Do("SCARD", "Exist_"+e.Topic))
	if err != nil {
		return 0, err
	}
	return i, nil
}

//GetAllKey 返回所有的key
func (e *exist) GetAllKey() (all []string, err error) {
	e.RWL.RLock()
	defer e.RWL.RUnlock()
	conn := e.RD.Get()
	defer conn.Close()
	is, err := redis.Strings(conn.Do("SMEMBERS", "Exist_"+e.Topic))
	if err != nil {
		return nil, err
	}
	return is, nil
}
