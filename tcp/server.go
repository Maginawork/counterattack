package tcp

import (
	"counterattack/server/cache"
	"net"
	"bufio"
	"io"
	"log"
)



type Server struct {
	cache.Cache
}


func (s *Server) Listen(){
	l, e := net.Listen("tcp",":12346")
	if e != nil{
		panic(e)
	}

	for {
		c, e := l.Accept()
		if e != nil {
			panic(e)
		}
		go s.process(c)
	}
}



func New(c cache.Cache) *Server{
	return  &Server{c}
}

func (s *Server) get(conn net.Conn,r *bufio.Reader) error{
	k,e := s.readKey(r)
	if e != nil{
		return e
	}
	v,e := s.Get(k)
	return sendResponse(v,e,conn)
}

func (s *Server) set(conn net.Conn, r *bufio.Reader) error{
	k,v,e:= s.readKeyAndValue(r)
	if e !=nil {
		return e
	}
	return sendResponse(nil,s.Set(k,v),conn)
}

func (s *Server) del(conn net.Conn, r *bufio.Reader) error{
	k, e := s.readKey(r)
	if e != nil {
		return e
	}

	return  sendResponse(nil,s.Del(k),conn)
}

func (s *Server) process(conn net.Conn){
	defer conn.Close()
	r := bufio.NewReader(conn)

	for {
		if op, e := r.ReadByte();e != nil{
			if e != io.EOF{
				log.Println("close connection due to error : ",e)
			}
			return
		}else{
			switch op {
			case 'S':
				e = s.set(conn,r)
			case 'G':
				e = s.get(conn,r)
			case 'D':
				e = s.del(conn,r)
			default:
				log.Println("close connetion due to invalid operation: ", op)
				return
			}
			if e != nil {
				log.Println("close connection due to error : ",e)
				return
			}
		}

	}
}

func (s *Server) readKey(r *bufio.Reader) (string, error){
	klen,e := readLen(r)
	if e != nil{
		return "",e
	}

	k :=make([]byte,klen)
	if _,err := io.ReadFull(r,k);err !=nil{
		return "",err
	}

	return string(k),nil

}

func (s *Server) readKeyAndValue(r *bufio.Reader)(string,[]byte,error){
	klen, e := readLen(r)
	if e != nil{
		return "",nil,e
	}

	vlen, e := readLen(r)
	if e != nil{
		return "", nil, e
	}

	k := make([]byte,klen)
	_,e = io.ReadFull(r,k)
	if e != nil{
		return "", nil, e
	}

	v := make([]byte,vlen)
	_,e = io.ReadFull(r,v)

	if e != nil{
		return "",nil,e
	}
	return string(k),v,nil
}