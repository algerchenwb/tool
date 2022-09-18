package his

import (
	"sync"

	"github.com/garyburd/redigo/redis"
)

//limit 用于上限记录，数字累计。
type limit struct {
	Topic string //主题，唯一值，比如：添加好友
	RWL   sync.RWMutex
	RD    *redis.Pool
}

//Add 记录自增
func (l *limit) Add(key string, delta int) {
	l.RWL.Lock()
	defer l.RWL.Unlock()
	conn := l.RD.Get()
	_, _ = conn.Do("HINCRBY", "Limit_"+l.Topic, key, delta)
	conn.Close()
}

//Sub 记录自减
func (l *limit) Sub(key string, delta int) {
	l.RWL.Lock()
	defer l.RWL.Unlock()
	conn := l.RD.Get()
	_, _ = conn.Do("HINCRBY", "Limit_"+l.Topic, key, -delta)
	conn.Close()
}

//Get 获取当前量
func (l *limit) Get(key string) (cur int, e error) {
	conn := l.RD.Get()
	defer conn.Close()
	return redis.Int(conn.Do("HGET", "Limit_"+l.Topic, key))
}

//Reset 重置为0
func (l *limit) Reset(key string) (cur int, e error) {
	conn := l.RD.Get()
	defer conn.Close()
	return redis.Int(conn.Do("HSET", "Limit_"+l.Topic, key, 0))
}

//IsToLimit 是否已达指定上限
//返回：是否达上限，当前量，错误
func (l *limit) IsToLimit(key string, max int) (b bool, cur int, e error) {
	l.RWL.RLock()
	defer l.RWL.RUnlock()
	conn := l.RD.Get()
	defer conn.Close()
	i, err := redis.Int(conn.Do("HGET", "Limit_"+l.Topic, key))
	if err != nil {
		return false, 0, err
	}
	if i >= max {
		return true, i, nil
	}
	return false, i, nil
}
