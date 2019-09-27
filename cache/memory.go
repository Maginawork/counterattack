package cache

import (
	"sync"
	"errors"
)

type MemoryCache struct{
	c map[string][]byte
	mutex sync.RWMutex
	Stat
}

func (mc *MemoryCache) Set(k string, v []byte) error{
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	tmp, exist := mc.c[k]
	if exist {
		mc.del(k, tmp)
	}
	mc.c[k] = v
	mc.add(k,v)

	return nil
}

func (mc *MemoryCache) Get(k string) ([]byte,error){
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	v,isexist := mc.c[k]

	if isexist{
		return v,nil
	}else{
		return v,errors.New("there is no key: " + k)
	}

}

func (mc *MemoryCache)Del(k string) error{
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	v,isexist := mc.c[k]
	if isexist{
		mc.Stat.del(k,v)
		delete(mc.c,k)
	}

	return nil

}

func (mc *MemoryCache)GetStat() Stat{
	return mc.Stat
}

func newMemoryCache() *MemoryCache{
	var mc MemoryCache
	mc.c = make(map[string][]byte)

	return &mc
}