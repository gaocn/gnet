package gserv

import (
	"log"
	"net"

	"github.com/gaocn/gnet/gifac"
)

// 链接模块的实现
type Connection struct {
	// 当前链接的socket
	Conn *net.TCPConn
	//链接Id
	ConnId uint32
	// 当前链接的状态
	IsClosed bool
	// 当前链接的业务处理方法
	HandleAPI gifac.HandleFunc

	// 告知链接是否需要终止的通道
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connId uint32, callbackApi gifac.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnId:    connId,
		HandleAPI: callbackApi,
		IsClosed:  false,
		// buffered channel
		ExitChan: make(chan bool, 1),
	}
	return c
}

// 链接的读业务方法，
func (c *Connection) StartReader() {
	log.Printf("Reader Goroutine is running...\n")
	defer log.Printf("ConnId: %d, Reader Goroutine exits, remote addr: %s\n", c.ConnId, c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			log.Printf("Error reading connection: %s\n", err)
			continue
		}

		// 调用链接绑定的HandlAPI处理接收到的数据
		if err := c.HandleAPI(c.Conn, buf, cnt); err != nil {
			log.Printf("ConnId:%d, handle read data error:%v\n", c.ConnId, err)
			break
		}
	}

}

// 启动链接
func (c *Connection) Start() {
	log.Printf("Conn Starting.... ConnId: %d\n", c.ConnId)
	// 启动从当前链接读数据的业务
	go c.StartReader()

	// TODO 启动从当前链接写数据的业务
}

// 终止链接
func (c *Connection) Stop() {
	log.Printf("Conn Stopped.... ConnId: %d\n", c.ConnId)
	if c.IsClosed {
		return
	}
	c.IsClosed = true
	// 关闭链接
	c.Conn.Close()
	// 关闭管道，回收资源
	close(c.ExitChan)
}

// 获取
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnectionId() uint32 {
	return c.ConnId
}
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send() error {
	return nil
}
