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
	IPerson interface {
		Get(ctx context.Context, ID int) (entry entity.Person, err error)
	}
)

var (
	localPerson IPerson
)

func Person() IPerson {
	if localPerson == nil {
		panic("implement not found for interface IPerson, forgot register?")
	}
	return localPerson
}

func RegisterPerson(i IPerson) {
	localPerson = i
}
