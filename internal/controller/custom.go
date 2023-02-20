package controller

import (
	v1 "github.com/xm-chentl/goframedemo/api/v1"

	"github.com/gogf/gf/v2/net/ghttp"
)

type CustomAPI struct{}

func (c CustomAPI) Call(req *ghttp.Request) (res *v1.DefaultRes, err error) {
	return
}
