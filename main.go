package main

import (
	"counterattack/server/cache"
	"counterattack/server/http"
	"counterattack/server/tcp"
)

func main(){
	c := cache.New("memory")
	go tcp.New(c).Listen()
	http.New(c).Listen()
}