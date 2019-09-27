package http

import (
	"counterattack/server/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}


func (s *Server) cacheHandler() http.Handler{
	return &cacheHandler{s}
}

func (s *Server) statusHandler() http.Handler{
	return &statusHandler{s}
}

func New(c cache.Cache) *Server{
	return &Server{c}
}

func (s *Server) Listen(){
	http.Handle("/cache/",s.cacheHandler())
	http.Handle("/status",s.statusHandler())
	http.ListenAndServe(":55555",nil)
}