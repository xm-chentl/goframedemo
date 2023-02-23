package license

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/farmerx/gorsa"
	"github.com/gogf/gf/v2/frame/g"
	userApp "github.com/xm-chentl/goframedemo/internal/model/views/user_app"
	"github.com/xm-chentl/goframedemo/internal/service"
	"github.com/xm-chentl/goframedemo/utils"
)

const (
	LayoutDateTime = "2006-01-02 15:04:05"
)

type licenseData struct {
	MachineCode string `json:"machine_code"`
	ExpireTime  string `json:"expire_time"`
}

type AppGenAPI struct {
	g.Meta `method:"post"`

	UserAppID   int    `v:"user_app_id @required#请输入应用标识"`
	MachineCode string `v:"machine_code @required#请输入机器码"`
}

func (s AppGenAPI) Call(ctx context.Context) (res interface{}, err error) {
	// todo: 暂时用统计的用户信息
	entry, err := service.UserApp().Get(ctx, s.UserAppID)
	if err != nil {
		err = utils.NewCustomError(602, "获取信息失败")
		return
	}
	if entry.Id == 0 {
		err = utils.NewCustomError(604, "未找到对应有效应用")
		return
	}

	prvBytes, pubBytes, err := utils.GenRsaKey()
	if err != nil {
		err = utils.NewCustomError(603, "生成密钥失败")
		return
	}

	licenseDataResp := licenseData{
		MachineCode: s.MachineCode,
		ExpireTime:  time.Now().Add(time.Duration(entry.Duration*24) * time.Hour).Format(LayoutDateTime),
	}
	licenseCodeByte, err := json.Marshal(licenseDataResp)
	if err != nil {
		err = utils.NewCustomError(605, "生成许可证失败")
		return
	}

	gorsa.RSA.SetPrivateKey(string(prvBytes))
	gorsa.RSA.SetPublicKey(string(pubBytes))
	newLicenseCodeBytes, err := gorsa.RSA.PriKeyENCTYPT(licenseCodeByte)
	if err != nil {
		panic(err)
	}

	// 私钥加密
	resp := userApp.LicenseAppGenReq{
		UserAppID:   s.UserAppID,
		Pub:         string(pubBytes),
		Prv:         string(prvBytes),
		LicenseCode: hex.EncodeToString(newLicenseCodeBytes),
	}
	err = service.UserApp().UpdateLicense(ctx, resp)
	if err != nil {
		err = utils.NewCustomError(601, "保存许可信息失败: %v", err)
		return
	}

	// lf, err := os.Create(fmt.Sprintf("license_app_%d.pri", s.UserAppID))
	// defer lf.Close()

	// _, err = lf.WriteString(hex.EncodeToString(newLicenseCodeBytes))
	res = map[string]interface{}{
		"license_code": resp.LicenseCode,
		"pub":          resp.Pub,
	}

	return
}
