// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserApp is the golang structure for table user_app.
type UserApp struct {
	Id             int         `json:"id"              ` // 数据标识
	UserId         int         `json:"user_id"         ` // 用户标识
	AppId          int         `json:"app_id"          ` // 应用标识
	AppName        string      `json:"app_name"        ` // 应用名
	Duration       int         `json:"duration"        ` // 有效时长(天)
	StartTime      *gtime.Time `json:"start_time"      ` // 开始时间
	PrvContent     string      `json:"prv_content"     ` // 私钥
	PubContent     string      `json:"pub_content"     ` // 公钥
	LicenseContent string      `json:"license_content" ` // 许可证
}