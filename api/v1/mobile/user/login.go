package user

import (
	"context"

	"github.com/xm-chentl/goframedemo/internal/contract"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginAPI struct {
	g.Meta `method:"POST"`

	Person contract.IPersonService `inject:""`

	Name string `v:"name @required|length:6,30#请输入用户名称|用户名称长度非法"`
}

func (s LoginAPI) Call(ctx context.Context) (res interface{}, err error) {

	res = map[string]interface{}{
		"name":       s.Name,
		"person.say": s.Person.Say(),
	}
	return
}
