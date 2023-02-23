package utils

import (
	"fmt"
	"net"
	"time"
)

func CheckTcpConnection(host string, port string) error {
	timeout := time.Second
	address := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return fmt.Errorf("unable to connect to %s", address)
	}
	if conn != nil {
		defer conn.Close()
	}
	return nil
}
