package person

import (
	"context"

	"github.com/xm-chentl/goframedemo/internal/dao"
	"github.com/xm-chentl/goframedemo/internal/model/entity"
	"github.com/xm-chentl/goframedemo/internal/service"
)

type sPerson struct{}

func (s *sPerson) Get(ctx context.Context, ID int) (entry entity.Person, err error) {
	err = dao.Person.Ctx(ctx).Where("id = ?", ID).Scan(&entry)
	return
}

func init() {
	service.RegisterPerson(&sPerson{})
}
