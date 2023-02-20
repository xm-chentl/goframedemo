package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type InfoAPI struct {
	g.Meta `method:"get"`

	Mode string `h:"mode"`

	ID int `p:"sid"`
}

func (s InfoAPI) Call(ctx context.Context) (res interface{}, err error) {
	panic("测试报错")
	res = map[string]interface{}{
		"id":    s.ID,
		"debug": s.Mode,
	}
	return
}
