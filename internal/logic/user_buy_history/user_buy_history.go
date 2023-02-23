package userbuyhistory

import (
	"context"

	"github.com/xm-chentl/goframedemo/internal/model/entity"
	"github.com/xm-chentl/goframedemo/internal/service"
)

type sUserBuyHistory struct{}

func (s *sUserBuyHistory) Get(ctx context.Context, id int) (entry entity.UserBuyHistory, err error) {
	return
}

func init() {
	service.RegisterUserBuyHistory(&sUserBuyHistory{})
}
