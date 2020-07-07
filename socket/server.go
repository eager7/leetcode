package socket

import (
	"bufio"
	"context"
	"fmt"
	"net"
)

func NewServer() error {
	ctx := context.Background()
	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8989")
	if err != nil {
		return fmt.Errorf("resolve tcp addr:%v", err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return fmt.Errorf("listen tcp:%v", err)
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			return fmt.Errorf("accept tcp:%v", err)
		}
		go Handler(ctx, conn)
	}
}

func Handler(ctx context.Context, conn *net.TCPConn) {
	fmt.Println("start:", conn.RemoteAddr().String())
	if err := conn.SetKeepAlive(true); err != nil {
		fmt.Println("set keepalive:", err)
		return
	}
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("close conn err:", err)
		}
	}()
	reader := bufio.NewReader(conn)
	for {
		select {
		default:
			ret, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("read err:", err)
				return
			}
			fmt.Println("read msg:", ret)
			n, err := conn.Write([]byte(ret))
			if err != nil {
				fmt.Println("write err:", err)
				return
			}
			fmt.Println("write success:", n)
		case <-ctx.Done():
			fmt.Println("exit:", conn.RemoteAddr().String())
			return
		}
		break
	}
}
