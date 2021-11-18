package gserv

import (
	"fmt"
	"log"
	"net"

	"github.com/gaocn/gnet/gifac"
)

// IServer接口的实现类

type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的IP版本
	IPVersion string
	// 服务器绑定的IP
	IP string
	// 服务器监听的端口
	Port int
}

// 启动服务器
// 非阻塞方法，只用于启动服务器
func (s *Server) Start() {
	go func() {
		log.Printf("Starting Server listening at %s:%d\n", s.IP, s.Port)
		// 1. 获取一个TCP的地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Println("Error starting server:", err)
			return
		}
		// 2. 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Println("Error listen tpc: ", err)
			return
		}
		log.Println("Start gnet server successfully, ", s.Name, "is listening...")
		// 3. 阻塞等待客户端连接，处理客户端业务（读写）
		for {
			// 若有连接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("Error accept connection: ", err)
				continue
			}
			go func() {
				for {
					buf := make([]byte, 512)
					count, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Error read connection: ", err)
						return
					}

					//直接写回
					if _, err := conn.Write(buf[:count]); err != nil {
						fmt.Println("Error write connection:", err)
						return
					}
				}
			}()
		}
	}()

}

// 停止服务器
func (s *Server) Stop() {
	// 回收服务器资源后停止服务器

}

// 启动服务
func (s *Server) Serve() {
	s.Start()

	// TODO: 可以做一些服务器启动后的额外业务，便于扩展

	// 阻塞
	select {}
}

// Server工厂方法，用于创建Server实例
func NewServer(name string) gifac.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
