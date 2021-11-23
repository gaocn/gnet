package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// gnet框架相关的配置文件
// 配置文件默认路径是在客户端的conf/gnet.json文件中
//
type GlobalConfig struct {
	// 服务器名称
	Name string `json:"name"`
	// 服务器监听地址
	Host string `json:"host"`
	// 服务器监听端口
	Port int `json:"port"`
	// 版本号
	Version string `json:"version"`

	//最大链接数
	MaxConnections uint32 `json:"maxConnections"`
	// 最大报文长度
	MaxPacketSize uint32 `json:"maxPacketSize"`

	// 关联服务器实例
}

// 定义全局配置
var Conf *GlobalConfig

// 配置对象初始化，若未配置使用缺省值
func init() {
	Conf = &GlobalConfig{
		Name:           "Gnet Server",
		Host:           "0.0.0.0",
		Port:           8999,
		Version:        "Gnet-V1.0",
		MaxConnections: 100,
		MaxPacketSize:  2048,
	}

	// 读取配置文件
	Conf.Reload()
}

func (c *GlobalConfig) Reload() {
	data, err := ioutil.ReadFile("conf/gnet.json")
	if err != nil {
		log.Println("conf file read error: ", err)
		return
	}
	err = json.Unmarshal(data, Conf)
	if err != nil {
		log.Println("conf json unmarshal error: ", err)
	}
}

func (c *GlobalConfig) String() string {
	data, _ := json.Marshal(c)
	return string(data)
}
