package cache

import "log"

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

func New(typ string) Cache {
	var c Cache
	if typ == "rocksdb"{
		c = newRocksdbCache()
	}
	if typ == "memory"{
		c = newMemoryCache()
	}else{
		log.Println("unknow cache type "+typ)
	}
	log.Println(typ, "ready to serve")
	return c
}