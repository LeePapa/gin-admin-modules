package admin_service

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/pkg/errors"
	"crypto/md5"
	"encoding/hex"
	"gin-modules/modules/admin/model"
	"gin-modules/middleware"
)

//登录操作
func Login(username string, password string, ip string) (string, error) {
	adminInfo, has := admin_model.GetAdminByUsername(username, "*")
	if !has {
		return "", errors.New("账号不存在")
	}
	if bcrypt.CompareHashAndPassword([]byte(adminInfo.Password), []byte(password)) != nil {
		return "", errors.New("账号或者密码错误")
	}
	if adminInfo.Status != 0 {
		return "", errors.New("该账号异常，无法正常登录")
	}
	//添加管理员登录记录
	admin_model.AddLoginRecord(admin_model.AdminLoginRecord{
		AdminID: adminInfo.ID,
		IP:      ip,
		Address: "",
	})
	//生成token
	token, err := createToken(adminInfo, ip)
	if err != nil {
		return "", err
	}
	return token, nil
}

//通过短信验证码登录操作
func LoginByCode(username string, code string, ip string) (string, error) {
	adminInfo, has := admin_model.GetAdminByUsername(username, "*")
	if !has {
		return "", errors.New("账号不存在")
	}
	if err := CheckSmsCode(2, username, code); err != nil {
		return "", err
	}
	if adminInfo.Status != 0 {
		return "", errors.New("该账号异常，无法正常登录")
	}
	//添加管理员登录记录
	admin_model.AddLoginRecord(admin_model.AdminLoginRecord{
		Type:    2,
		AdminID: adminInfo.ID,
		IP:      ip,
		Address: "",
	})
	//生成token
	token, err := createToken(adminInfo, ip)
	if err != nil {
		return "", err
	}
	return token, nil
}

//用户注册业务
func Register(u admin_model.AdminInfo, code string) error {
	_, has := admin_model.GetAdminByUsername(u.Username, "id")
	if has {
		return errors.New("该账号已存在")
	}
	err := CheckSmsCode(1, u.Username, code)
	if err != nil {
		return err
	}
	if admin_model.AddAdmin(u) == 0 {
		return errors.New("用户信息注册失败")
	}
	return nil
}

//创建登录token
func createToken(adminInfo admin_model.AdminInfo, ip string) (string, error) {
	h := md5.New()
	h.Write([]byte(ip))
	token, err := middleware.NewJWT().CreateToken(middleware.CustomClaims{
		ID:       adminInfo.ID,
		Nickname: adminInfo.Nickname,
		HeadImg:  adminInfo.HeadImg,
		HashIP:   hex.EncodeToString(h.Sum(nil)),
		RoleId:   adminInfo.RoleID,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}
