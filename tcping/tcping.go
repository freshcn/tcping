package tcping

import (
	"fmt"
	"net"
	"time"
)

var (
	// DefaultPort 默认端口
	DefaultPort = 80
	// DefaultConnTimeout 默认连接超时时间
	DefaultConnTimeout = 5 * time.Second
)

// Client 连接的客户端
type Client struct {
	config Config
}

// Config 配置
type Config struct {
	// Host 主机地址
	Host string
	// Port 端口号
	Port int
	// ConnTimeout 连接超时时间
	ConnTimeout time.Duration
}

// Response 请求结果返回
type Response struct {
	// StartTime 请求开始的时间
	StartTime int64
	// EndTime 请求结束的时间
	EndTime int64
	// Err 错误消息
	Err error
	// RemoteAddr 远程地址信息
	RemoteAddr net.Addr
	// LocalAddr 本地地址信息
	LocalAddr net.Addr
}

// New 创建一个新 Client 的对象
// host设置要请求的服务器地址,port为要请求的服务端端口
func New(config Config) *Client {

	if config.Port == 0 {
		config.Port = DefaultPort
	}

	if config.ConnTimeout == 0 {
		config.ConnTimeout = DefaultConnTimeout
	}
	return &Client{
		config,
	}
}

// Ping 测试一个地址
func (t *Client) Ping() (r Response) {
	r = Response{
		StartTime: time.Now().UnixNano(),
	}
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", t.config.Host, t.config.Port), t.config.ConnTimeout)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	if conn != nil {
		r.LocalAddr = conn.LocalAddr()
		r.RemoteAddr = conn.RemoteAddr()
	}
	r.EndTime = time.Now().UnixNano()

	if err != nil {
		r.Err = err
	}
	return
}

// HumanTime 返回一个人类能看懂的时间
func (r *Response) HumanTime() (s string) {
	var times = float64(r.EndTime) - float64(r.StartTime)

	if tmp := times / 1000000000; tmp > 1 {
		s = fmt.Sprintf("%.2fs", tmp)
		return
	}

	if tmp := times / 1000000; tmp > 1 {
		s = fmt.Sprintf("%.2fms", tmp)
		return
	}

	if tmp := times / 1000; tmp > 1 {
		s = fmt.Sprintf("%.2fμs", tmp)
		return
	}

	return fmt.Sprintf("%fns", times)
}
