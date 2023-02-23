package userapp

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/xm-chentl/goframedemo/internal/dao"
	"github.com/xm-chentl/goframedemo/internal/model/entity"
	userApp "github.com/xm-chentl/goframedemo/internal/model/views/user_app"
	"github.com/xm-chentl/goframedemo/internal/service"
)

type sUserApp struct{}

func (s *sUserApp) Get(ctx context.Context, id int) (entry entity.UserApp, err error) {
	err = dao.UserApp.Ctx(ctx).Where("id = ?", id).Scan(&entry)
	return
}

func (s *sUserApp) Insert(ctx context.Context, req userApp.BuyAppReq) (id int64, err error) {
	entry := entity.UserApp{
		UserId:    req.UserID,
		AppId:     req.AppID,
		AppName:   req.AppName,
		Duration:  int(req.Duration),
		StartTime: gtime.Now(),
	}
	result, err := dao.UserApp.Ctx(ctx).Insert(entry)
	if err != nil {
		return
	}
	id, err = result.LastInsertId()

	return
}

func (s *sUserApp) UpdateLicense(ctx context.Context, req userApp.LicenseAppGenReq) (err error) {
	_, err = dao.UserApp.Ctx(ctx).Data(g.Map{
		"pub_content":     req.Pub,
		"prv_content":     req.Prv,
		"license_content": req.LicenseCode,
	}).Where("id = ?", req.UserAppID).Update()

	return
}

func init() {
	service.RegisterUserApp(&sUserApp{})
}
