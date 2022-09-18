package his

import (
	"fmt"
	"sync"

	"github.com/garyburd/redigo/redis"
)

type myMap struct {
	Topic  string
	rwlMax uint
	RWL    []sync.RWMutex
	RD     *redis.Pool
}

func (m *myMap) key(id uint) string {
	return fmt.Sprint("Map_", m.Topic, "_", id)
}

func (m *myMap) SaddOne(id uint, value string) {
	m.RWL[id%m.rwlMax].Lock()
	defer m.RWL[id%m.rwlMax].Unlock()
	conn := m.RD.Get()
	_, _ = conn.Do("SADD", m.key(id), value)
	conn.Close()
}

func (m *myMap) Sadd(id uint, values []string) int {
	m.RWL[id%m.rwlMax].Lock()
	defer m.RWL[id%m.rwlMax].Unlock()
	conn := m.RD.Get()

	startCount, _ := redis.Int(conn.Do("SCARD", m.key(id)))
	var temp = make([]any, 0, len(values)+1)
	temp = append(temp, m.key(id))
	for _, d := range values {
		temp = append(temp, d)
	}
	_, _ = conn.Do("SADD", temp...)
	endCount, _ := redis.Int(conn.Do("SCARD", m.key(id)))
	conn.Close()
	return endCount - startCount
}

func (m *myMap) SpopOne(id uint) (string, error) {
	m.RWL[id%m.rwlMax].Lock()
	defer m.RWL[id%m.rwlMax].Unlock()
	conn := m.RD.Get()
	s, err := redis.String(conn.Do("SPOP", m.key(id)))
	conn.Close()
	return s, err
}

func (m *myMap) Spop(id uint, count int) ([]string, error) {
	m.RWL[id%m.rwlMax].Lock()
	defer m.RWL[id%m.rwlMax].Unlock()
	conn := m.RD.Get()
	var rts = make([]string, 0, count)
	for i := 0; i < count; i++ {
		s, err := redis.String(conn.Do("SPOP", m.key(id)))
		if err == nil {
			rts = append(rts, s)
		}
	}
	conn.Close()
	if len(rts) <= 0 {
		return nil, fmt.Errorf("没有数据了")
	}
	return rts, nil
}

func (m *myMap) SpopAll(id uint) ([]string, error) {
	m.RWL[id%m.rwlMax].Lock()
	defer m.RWL[id%m.rwlMax].Unlock()
	conn := m.RD.Get()
	rts, err := redis.Strings(conn.Do("SMEMBERS", m.key(id)))
	if err == nil {
		_, _ = conn.Do("DEL", m.key(id))
	}
	conn.Close()
	if len(rts) <= 0 {
		return nil, fmt.Errorf("没有数据了")
	}
	return rts, err
}

func (m *myMap) RangeAll(id uint) ([]string, error) {
	m.RWL[id%m.rwlMax].Lock()
	defer m.RWL[id%m.rwlMax].Unlock()
	conn := m.RD.Get()
	rts, err := redis.Strings(conn.Do("SMEMBERS", m.key(id)))
	conn.Close()
	if len(rts) <= 0 {
		return nil, fmt.Errorf("没有数据了")
	}
	return rts, err
}

func (m *myMap) DeleteKey(id uint) {
	m.RWL[id%m.rwlMax].Lock()
	defer m.RWL[id%m.rwlMax].Unlock()
	conn := m.RD.Get()
	_, _ = conn.Do("DEL", m.key(id))
	conn.Close()
}

func (m *myMap) Length(id uint) int {
	m.RWL[id%m.rwlMax].RLock()
	defer m.RWL[id%m.rwlMax].RUnlock()
	conn := m.RD.Get()
	size, err := redis.Int(conn.Do("SCARD", m.key(id)))
	conn.Close()
	if err != nil {
		return 0
	}
	return size
}
