package main

import (
	"context"

	"github.com/xm-chentl/goframedemo/internal/cmd"
	"github.com/xm-chentl/goframedemo/internal/contract"
	_ "github.com/xm-chentl/goframedemo/internal/packed"
	"github.com/xm-chentl/goframedemo/internal/services"
	"github.com/xm-chentl/goresource"
	"github.com/xm-chentl/goresource/mongoex"
	"github.com/xm-chentl/goresource/mysqlex"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/xm-chentl/gocore/iocex"
)

func main() {
	// todo: 初始配置
	cfg := g.Cfg()
	mysqlLink, _ := cfg.Get(context.TODO(), "database.link")
	mongoLink, _ := cfg.Get(context.TODO(), "mongo.link")
	// 注入对象
	iocex.Set(new(contract.IPersonService), services.NewPersonService())
	iocex.SetMap(new(goresource.IResource), map[string]interface{}{
		"mysql": mysqlex.New(mysqlLink.String()),
		"mongo": mongoex.New("mytest", mongoLink.String()),
	})

	cmd.Main.Run(gctx.New())


	go grpc.run()
}
