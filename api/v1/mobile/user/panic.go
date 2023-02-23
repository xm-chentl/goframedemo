package user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type PanicAPI struct {
	g.Meta `method:"get"`
}

func (s PanicAPI) Call(ctx context.Context) (res interface{}, err error) {
	panic("异常报错")
}
