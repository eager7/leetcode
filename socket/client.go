package socket

import (
	"fmt"
	"net"
)

func NewClient(addr string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8989")
	if err != nil {
		return fmt.Errorf("resolve tcp addr:%v", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return fmt.Errorf("dial tcp:%v", err)
	}
	if _, err := conn.Write([]byte("hello world!")); err != nil {
		return fmt.Errorf("write:%v", err)
	}
	return nil
}
