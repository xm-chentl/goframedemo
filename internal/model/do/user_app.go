// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserApp is the golang structure of table user_app for DAO operations like Where/Data.
type UserApp struct {
	g.Meta         `orm:"table:user_app, do:true"`
	Id             interface{} // 数据标识
	UserId         interface{} // 用户标识
	AppId          interface{} // 应用标识
	AppName        interface{} // 应用名
	Duration       interface{} // 有效时长(天)
	StartTime      *gtime.Time // 开始时间
	PrvContent     interface{} // 私钥
	PubContent     interface{} // 公钥
	LicenseContent interface{} // 许可证
}
