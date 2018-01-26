// 开发接口
package gserver

import (
	Conversion "github.com/gfandada/gserver/autoconversion"
	DB "github.com/gfandada/gserver/db"
	LOADER "github.com/gfandada/gserver/loader"
	Logger "github.com/gfandada/gserver/logger"
	Module "github.com/gfandada/gserver/module"
	Network "github.com/gfandada/gserver/network"
	Services "github.com/gfandada/gserver/services"
	Discovery "github.com/gfandada/gserver/services/discovery"
	GateService "github.com/gfandada/gserver/services/gateway"
	Service "github.com/gfandada/gserver/services/service"
)

// 运行模块
// @params 模块转载器...
func Run(mods ...Module.Imodule) {
	Module.Run(mods...)
}

// 启动一个websocket网关服务
// @params log:日志配置 path:网关配置文件 discpath:服务发现配置 coder:消息编码器(注意消息需要开发者自行注册)
func RunWSGateway(log, path, discpath string, coder Network.Imessage) {
	Logger.Start(log)
	gate := new(GateService.WsGateway)
	gate.Config = path
	gate.Coder = coder
	Discovery.Init(discpath)
	Run(gate)
}

// 启动一个tcp网关服务
// @params log:日志配置 path:网关配置文件 discpath:服务发现配置 coder:消息编码器(注意消息需要开发者自行注册)
func RunTCPGateway(log, path, discpath string, coder Network.Imessage) {
	Logger.Start(log)
	gate := new(GateService.TcpGateway)
	gate.Config = path
	gate.Coder = coder
	Discovery.Init(discpath)
	Run(gate)
}

// 启动一个通用的service服务
// @params log:配置 path:服务配置 coder:消息编码器(注意消息需要开发者自行注册)
func RunService(log, path string, coder Network.Imessage) {
	Logger.Start(log)
	service := new(Service.Service)
	service.Config = path
	service.Coder = coder
	Run(service)
}

// 注册消息handler
// 本接口注意使用提供全局容器，保存id-handler的映射关系
// 非线程安全
// @params list:消息列表
func RegisterHandler(list []*Services.MsgHandler) {
	for _, v := range list {
		Services.Register(v.MsgId, v.MsgHandler)
	}
}

// 启动mysql数据库
// @params log:日志配置 path:数据库日志
func RunMysqlService(log, path string) {
	Logger.Start(log)
	DB.NewMysql(path)
}

// 启动xlsx加载器
// @params log:日志配置 path:xlsx目录
func RunXlsxService(log, path string) {
	Logger.Start(log)
	LOADER.Init(path)
}

// 解析配置，自动生成类型代码
// @params src : 源目录(yml文件)
// @params dest: 目地目录(go源文件)
func RunConversion(src, dest string) {
	Conversion.Conversion(src, dest)
}
