package admin_service

import (
	"time"
	"github.com/fastgoo/alisms-go"
	"fmt"
	"math/rand"
	"errors"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gin-modules/modules/admin/model"
	"gin-modules/pkg/setting"
)

const (
	SmsHost = "http://dysmsapi.aliyuncs.com"
)

type CheckTicketRes struct {
	Response  string `json:"response"`
	EvilLevel string `json:"evil_level"`
	ErrMsg    string `json:"err_msg"`
}

//发送短信验证码
func SendSmsCode(codeType int, mobile string, ip string) error {
	_, has := admin_model.GetAdminByUsername(mobile, "id")
	switch codeType {
	case 1:
		if has {
			return errors.New("该账号已被注册，不可重新注册")
		}
		break
	case 2:
		if !has {
			return errors.New("该账号不存在，请先注册")
		}
		break
	}
	code := createCaptcha()
	expireTime := 5
	ret := admin_model.AddSmsCode(admin_model.SmsCode{
		Type:       codeType,
		IP:         ip,
		Mobile:     mobile,
		Code:       code,
		ExpireTime: int(time.Now().Unix()) + expireTime*60,
	})
	if ret == 0 {
		return errors.New("短信验证码记录生成失败")
	}
	err := alisms.InitConfig(SmsHost, pkg_setting.Setting.AliSms.AccessId, pkg_setting.Setting.AliSms.SecretKey, pkg_setting.Setting.AliSms.SignName).Send(mobile, `{"code":"`+code+`"}`, pkg_setting.Setting.AliSms.DefaultTemplateCode)
	if err != nil {
		print(err.Error())
		switch err.Error() {
		case "isv.BUSINESS_LIMIT_CONTROL":
			return errors.New("短信验证码发送频繁，请稍后再试")
		default:
			return errors.New("短信验证码发送失败")
		}
	}
	return nil
}

//验证短信验证码
func CheckSmsCode(codeType int, mobile string, code string) error {
	codeInfo, has := admin_model.GetSmsByMobile(codeType, mobile, "*")
	if !has {
		return errors.New("验证码不存在")
	}
	if codeInfo.Code != code {
		return errors.New("验证码错误")
	}
	if codeInfo.Status != 0 || codeInfo.ExpireTime <= int(time.Now().Unix()) {
		return errors.New("短信验证码已过期")
	}
	admin_model.UpdateSmsCode(map[string]interface{}{"status": 1}, codeInfo.ID)
	return nil
}

//验证腾讯验证码的ticket
func CheckTicket(ticket string, randStr string, userIP string) error {
	resp, err := http.Get(fmt.Sprintf("https://ssl.captcha.qq.com/ticket/verify?aid=%s&AppSecretKey=%s&Ticket=%s&Randstr=%s&UserIP=%s", pkg_setting.Setting.Captcha.Aid, pkg_setting.Setting.Captcha.AppSecretKey, ticket, randStr, userIP))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var ret CheckTicketRes
	json.Unmarshal(body, &ret)
	if ret.Response == "1" {
		return nil
	}
	return errors.New("人机验证失败：" + ret.ErrMsg)
}

//创建6位验证码
func createCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
