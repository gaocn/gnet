package main

import (
	"log"

	"github.com/gaocn/gnet/gifac"
	"github.com/gaocn/gnet/gserv"
)

// 基于gnet框架开发的服务端应用程序

// 定义服务端路由用于处理客户端请求
type PingRouter struct {
	gserv.BaseRouter
}

func (r *PingRouter) PreHandle(request gifac.IRequest) {
	log.Println("prehandle called")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("before ping..."))
	if err != nil {
		log.Println("prehandle error: ", err)
		return
	}
}

func (r *PingRouter) Handle(request gifac.IRequest) {
	log.Println("hadle called")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("ping...ping...ping..."))
	if err != nil {
		log.Println("handle error: ", err)
		return
	}
}

func (r *PingRouter) PostHandle(request gifac.IRequest) {
	log.Println("posthandle called")
	_, err := request.GetConnection().GetTcpConnection().Write([]byte("post ping..."))
	if err != nil {
		log.Println("posthandle error: ", err)
		return
	}
}

func main() {
	// 基于gnet框架实例化服务器并启动服务
	s := gserv.NewServer("gserv-0.3")

	s.AddRouter(&PingRouter{})
	s.Serve()
}
