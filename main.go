package main

import (
	"github.com/xm-chentl/goframedemo/internal/cmd"
	"github.com/xm-chentl/goframedemo/internal/contract"
	_ "github.com/xm-chentl/goframedemo/internal/packed"
	"github.com/xm-chentl/goframedemo/internal/services"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/xm-chentl/gocore/iocex"
)

func main() {
	// 初始配置
	// 注入对象
	iocex.Set(new(contract.IPersonService), services.NewPersonService())

	cmd.Main.Run(gctx.New())
}
