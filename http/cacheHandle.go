package http

import (
	"net/http"
	"strings"
	"io/ioutil"
	"log"
	"fmt"
)

type cacheHandler struct {
	*Server
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter,r *http.Request){

	key := strings.Split(r.URL.EscapedPath(),"/")[2]
	fmt.Println("key is ", key)
	if len(key) == 0{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := r.Method
	if m == http.MethodPut{
		b,_ := ioutil.ReadAll(r.Body)
		if len(b) != 0 {
			if e :=h.Set(key,b);e != nil{
				log.Println(e)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		return
	}

	if m == http.MethodGet{
		b,err := h.Get(key)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(b) == 0{
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(b)
		return
	}

	if m == http.MethodDelete{
		err := h.Del(key)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}


