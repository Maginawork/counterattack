package main

import (
	"counterattack/server/cache"
	"counterattack/server/http"
)

func main(){
	c := cache.New("memory")

	http.New(c).Listen()
}