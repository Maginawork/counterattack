package http

import (
	"net/http"
	"encoding/json"
	"log"
)

type statusHandler struct {
	*Server
}


func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	b,e := json.Marshal(h.GetStat())
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(b)
}