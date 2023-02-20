package api

import (
	"github.com/xm-chentl/goframedemo/api/mobile/user"
	"github.com/xm-chentl/goframedemo/utility/apicontainer"
)

func init() {
	apicontainer.Register(map[string]apicontainer.APIHandler{
		"mobile/user/info":  &user.InfoAPI{},
		"mobile/user/login": &user.LoginAPI{},
	})
}
