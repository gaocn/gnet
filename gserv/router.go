package gserv

import "github.com/gaocn/gnet/gifac"

// 实现router时，继承该基类以便根据需求重写
type BaseRouter struct {
}

// 仅仅是简单实现，因为不知道具体的业务处理逻辑
func (r *BaseRouter) PreHandle(request gifac.IRequest) {}

func (r *BaseRouter) Handle(request gifac.IRequest) {}

func (r *BaseRouter) PostHandle(request gifac.IRequest) {}
