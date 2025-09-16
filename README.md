# ginroute
自动生成gin路由注册文件

package main

import "github.com/qinxl/ginroute"

func main() {
	cfg := &ginroute.GenCfg{
		Path: "internal/routes", // 默认routes
	}
	ginroute.Generate(cfg)
}
