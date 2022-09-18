package his

import (
	"fmt"
	"strings"
	"sync"

	"github.com/garyburd/redigo/redis"
	"tools/rlog"
)

func NewQueue(rd *redis.Pool, minLength int, autoPushFunc func(LetterId string, gender uint, countryCode string) []int32) *queue {
	return &queue{
		RWL:          sync.RWMutex{},
		RD:           rd,
		MinLength:    minLength,
		AutoPushFunc: autoPushFunc,
	}
}

type queue struct {
	RWL          sync.RWMutex
	RD           *redis.Pool
	MinLength    int
	AutoPushFunc func(LetterId string, gender uint, countryCode string) []int32
}

// Lpop 获取并移除最左元素
func (q *queue) Lpop(LetterId string, gender uint, countryCode string) (int32, error) {
	q.RWL.Lock()
	defer q.RWL.Unlock()
	conn := q.RD.Get()
	defer conn.Close()
	key := fmt.Sprint("LetterID_", LetterId, "_", gender, "_", countryCode)
	key = strings.ToLower(key) //转换为小写
	length, err := redis.Int(conn.Do("LLEN", key))
	if err != nil {
		rlog.Error("获取Redis长度失败：", key, err)
	} else {
		if length <= q.MinLength {
			nodes := q.AutoPushFunc(LetterId, gender, countryCode)
			if len(nodes) <= 0 {
				return 0, fmt.Errorf("没有数据了")
			}
			var ok = 0
			var no = 0
			for _, jg := range nodes {
				_, err = conn.Do("RPUSH", key, jg) //如果插入失败
				if err == nil {
					ok++
				} else {
					no++
				}
			}
			length2, _ := redis.Int(conn.Do("LLEN", key))
			rlog.Info("本次插入到Redis数量：[ 成功:", ok, "] [ 失败:", no, "] [ 总量：", len(nodes), "]", " [ Redis现在量：", length2, "] ：key =", key)
		}
	}
	uid, err := redis.Int64(conn.Do("LPOP", key))
	if err != nil {
		rlog.Info("LPOP失败：查询长度：", length, key, err)
	}
	return int32(uid), err
}

////Rpop 获取并移除最右元素
//func (q *queue) Rpop(LetterId string, sex uint, country_code string) (int32, error) {
//	q.RWL.Lock()
//	defer q.RWL.Unlock()
//	conn := q.RD.Get()
//	uid, err := redis.Int64(conn.Do("RPOP", fmt.Sprint("LetterID_", LetterId, "_", sex, "_", country_code)))
//	conn.Close()
//	return int32(uid), err
//}

////Len 获取当前长度
//func (q *queue) Len(LetterId string, sex uint, country_code string) (int, error) {
//	q.RWL.RLock()
//	defer q.RWL.RUnlock()
//	conn := q.RD.Get()
//	length, err := redis.Int(conn.Do("LLEN", fmt.Sprint("LetterID_", LetterId, "_", sex, "_", country_code)))
//	conn.Close()
//	return length, err
//}

// LPush 最左压入一个元素
func (q *queue) LPush(LetterId string, sex uint, countryCode string, uid int32) error {
	q.RWL.Lock()
	defer q.RWL.Unlock()
	conn := q.RD.Get()
	_, err := conn.Do("LPUSH", fmt.Sprint("LetterID_", LetterId, "_", sex, "_", countryCode), uid)
	conn.Close()
	return err
}

////RPush 最右压入一个元素
//func (q *queue) RPush(LetterId string, sex uint, country_code string, uid int32) error {
//	q.RWL.Lock()
//	defer q.RWL.Unlock()
//	conn := q.RD.Get()
//	_, err := conn.Do("RPUSH", fmt.Sprint("LetterID_", LetterId, "_", sex, "_", country_code), uid)
//	conn.Close()
//	return err
//}

////RPush 最右压入多个元素
//func (q *queue) RPushs(LetterId string, sex uint, country_code string, uids []int32) error {
//	q.RWL.Lock()
//	defer q.RWL.Unlock()
//	conn := q.RD.Get()
//	var v []any
//	v = append(v, fmt.Sprint("LetterID_", LetterId, "_", sex, "_", country_code))
//	for _, d := range uids {
//		v = append(v, d)
//	}
//	_, err := conn.Do("RPUSH", v...)
//	conn.Close()
//	return err
//}
