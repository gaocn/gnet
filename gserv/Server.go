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

	// 目前一个服务器只能绑定一个Router，即所有处理逻辑相同
	Router gifac.IRouter
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
		var cid uint32
		cid = 0
		// 3. 阻塞等待客户端连接，处理客户端业务（读写）
		for {
			// 若有连接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				log.Println("Error accept connection: ", err)
				continue
			}

			// 创建链接对象，并绑定业务处理方法，然后启动链接业务处理逻辑
			dealConn := NewConnection(conn, cid, EchoToClient)
			cid++
			go dealConn.Start()
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

func (s *Server) AddRouter(router gifac.IRouter) {
	s.Router = router
	log.Println("Add Router: ", s.Router)
}

// Server工厂方法，用于创建Server实例
func NewServer(name string) gifac.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}
