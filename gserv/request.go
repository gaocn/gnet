package gserv

import "github.com/gaocn/gnet/gifac"

type Request struct {
	// 请求关联的链接
	conn gifac.IConnection
	// 请求数据
	data []byte
}

func (r *Request) GetConnection() gifac.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
