package user

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/xm-chentl/goresource"
)

type InfoAPI struct {
	g.Meta `method:"get"`

	Mongo goresource.IResource `inject:"mongo"`
	MySql goresource.IResource `inject:"mysql"`

	Mode string `h:"mode"`

	ID int `p:"sid"`
}

func (s InfoAPI) Call(ctx context.Context) (res interface{}, err error) {
	fmt.Println(s.Mongo, s.MySql)
	res = map[string]interface{}{
		"id":    s.ID,
		"debug": s.Mode,
		"m":     s.Mongo,
		"s":     s.MySql,
	}
	return
}
