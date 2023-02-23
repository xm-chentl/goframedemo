package api

import (
	"github.com/xm-chentl/goframedemo/api/app"
	"github.com/xm-chentl/goframedemo/api/license"
	"github.com/xm-chentl/goframedemo/api/mobile/user"
	"github.com/xm-chentl/goframedemo/utils/apicontainer"
)

func init() {
	apicontainer.Register(map[string]apicontainer.APIHandler{
		"app/buy":           &app.BuyAPI{},
		"license/app-gen":   &license.AppGenAPI{},
		"license/check":     &license.CheckAPI{},
		"mobile/user/info":  &user.InfoAPI{},
		"mobile/user/login": &user.LoginAPI{},
		"mobile/user/panic": &user.PanicAPI{},
	})
}
