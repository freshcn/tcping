package tcping

import (
	"fmt"
	"testing"
	"time"
)

func TestTcping(t *testing.T) {
	var (
		host = "www.baidu.com"
		port = 443
	)

	client := New(Config{
		Host:        host,
		Port:        port,
		ConnTimeout: 0,
	})

	var (
		i = 0
	)
	fmt.Printf("start ping %s:%d\n", host, port)
	for {
		data := client.Ping()
		if data.Err != nil {
			fmt.Printf("%s\n", data.Err)
		} else {
			fmt.Printf("Connected to %s sqe=%d time=%s \n", data.RemoteAddr, i, data.HumanTime())
		}
		i++
		time.Sleep(time.Second * 1)
	}
}
