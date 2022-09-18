package cache

import (
	"time"

	"gitee.com/ruige_fun/util/rlog"

	"github.com/allegro/bigcache"
)

// NewCache 创建一个缓存管理，Expired 缓存过期时间。
// Expired 不能小于2秒，小于2秒则不设置过期时间。
func NewCache(Expired time.Duration) *bigcache.BigCache {
	config := bigcache.DefaultConfig(Expired)
	if Expired >= time.Second*2 {
		config.CleanWindow = time.Second * 2
	}
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		rlog.ErrorP(rlog.ModelNoPath, "创建缓存失败：", err)
	}
	return cache
}
