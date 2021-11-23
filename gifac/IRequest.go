package gifac

type IRequest interface {
	// 获取请求绑定的链接
	GetConnection() IConnection
	// 获取请求的数据
	GetData() []byte
}
