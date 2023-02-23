package license

import (
	"context"
	"encoding/hex"
	"encoding/json"

	"github.com/farmerx/gorsa"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/xm-chentl/goframedemo/internal/service"
	"github.com/xm-chentl/goframedemo/utils"
)

type CheckAPI struct {
	g.Meta `method:"post"`

	UserAppID   int    `v:"user_app_id @required#请输入应用标识"`
	LicenseCode string `v:"license_code @required#请输入许可证"`
}

func (s CheckAPI) Call(ctx context.Context) (res interface{}, err error) {
	// todo: 返回有效期(开始、结束时间)
	entry, err := service.UserApp().Get(ctx, s.UserAppID)
	if err != nil {
		err = utils.NewCustomError(602, "获取信息失败")
		return
	}
	if entry.Id == 0 {
		err = utils.NewCustomError(604, "未找到对应有效应用")
		return
	}

	gorsa.RSA.SetPrivateKey(entry.PrvContent)
	gorsa.RSA.SetPublicKey(entry.PubContent)
	licenseCode, _ := hex.DecodeString(s.LicenseCode)
	licenseCodeBytes, err := gorsa.RSA.PubKeyDECRYPT(licenseCode)
	if err != nil {
		err = utils.NewCustomError(607, "无效的许可证")
		return
	}

	var licenseDataResp licenseData
	if err = json.Unmarshal(licenseCodeBytes, &licenseDataResp); err != nil {
		err = utils.NewCustomError(607, "无效的许可证")
		return
	}
	res = licenseDataResp

	return
}
