package app

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	userapp "github.com/xm-chentl/goframedemo/internal/model/views/user_app"
	"github.com/xm-chentl/goframedemo/internal/service"
	"github.com/xm-chentl/goframedemo/mock"
	"github.com/xm-chentl/goframedemo/utility"
)

type BuyAPI struct {
	g.Meta `method:"post"`

	AppID    int    `v:"app_id @required#请输入应用标识"`
	AppName  string `p:"app_name"`
	Duration int64  `v:"duration @required#请输入应用购买时长"`
}

func (s BuyAPI) Call(ctx context.Context) (res interface{}, err error) {
	// todo: 用户信息(暂时先mock)
	user := mock.CurrentUser
	id, err := service.UserApp().Insert(ctx, userapp.BuyAppReq{
		AppID:    s.AppID,
		AppName:  s.AppName,
		Duration: s.Duration,
		UserID:   user.ID,
		UserName: user.Name,
	})
	if err != nil {
		err = utility.NewCustomError(601, "购买失败")
		return
	}
	res = id

	return
}
