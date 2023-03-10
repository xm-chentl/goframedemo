// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/xm-chentl/goframedemo/internal/model/entity"
)

type (
	IUserBuyHistory interface {
		Get(ctx context.Context, id int) (entry entity.UserBuyHistory, err error)
	}
)

var (
	localUserBuyHistory IUserBuyHistory
)

func UserBuyHistory() IUserBuyHistory {
	if localUserBuyHistory == nil {
		panic("implement not found for interface IUserBuyHistory, forgot register?")
	}
	return localUserBuyHistory
}

func RegisterUserBuyHistory(i IUserBuyHistory) {
	localUserBuyHistory = i
}
