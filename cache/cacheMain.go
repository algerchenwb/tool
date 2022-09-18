package cache

import (
	"sync"
	"time"

	"github.com/allegro/bigcache"
)

type Exist interface {
	Set(key string, value any, second uint)             //设置：键，值，过期秒数
	Get(key string) (any, error)                        //获取：键。返回：值，错误（有错表示不存在或已过期）
	Exist(key string) bool                              //是否存在：键。返回：是否存在，true存在，false不存在
	Del(key string)                                     //删除：键
	ExistDelete(key string) bool                        //存在即删除：键；返回：是否存在，true存在，false不存在
	ExistOrSet(key string, value any, second uint) bool //不存在就设置：键，值，过期秒数；返回：是否存在，true存在，false不存在
	Clean()                                             //清空所有
}

// NewExist 创建一个是否存在的记录器
// defaultExpirationSecond 默认过期时间（秒）
func NewExist(defaultExpirationSecond uint) Exist {
	if defaultExpirationSecond <= 0 {
		defaultExpirationSecond = 60
	}
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute * 10))
	return &cacheExist{
		rwl:           sync.RWMutex{},
		defaultSecond: defaultExpirationSecond,
		store:         cache,
	}
}
