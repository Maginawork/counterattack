package tcp

import (
	"bufio"
	"net"
	"strconv"
	"strings"
	"fmt"
)

func readLen(r *bufio.Reader)(int, error){
	tmp, e := r.ReadString(' ')
	if e != nil {
		return 0,e
	}

	return strconv.Atoi(strings.TrimSpace(tmp))
}

func sendResponse(value []byte, err error, conn net.Conn) error{
	if err != nil{
		errString := err.Error()
		tmp := fmt.Sprintf("-%d ",len(errString)) + errString
		_,e := conn.Write([]byte(tmp))
		return e
	}
	vlen := fmt.Sprintf("%d ",len(value))
	_,e := conn.Write(append([]byte(vlen),value...))
	return e
}
