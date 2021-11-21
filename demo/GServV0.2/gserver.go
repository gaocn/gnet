package main

import "github.com/gaocn/gnet/gserv"

// 基于gnet框架开发的服务端应用程序

func main() {
	// 基于gnet框架实例化服务器并启动服务
	s := gserv.NewServer("gserv-0.2")
	s.Serve()
}
