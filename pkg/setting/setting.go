package pkg_setting

import (
	"log"
	"github.com/go-ini/ini"
	"flag"
	"time"
)

var Setting = SettingConf{}

type SettingConf struct {
	Database DatabaseIni
	Server   ServerIni
	Redis    RedisIni
	Captcha  TencentCaptchaIni
	AliSms   AliSmsIni
	Log      GinLogIni
}

type ServerIni struct {
	Port string
	Mode string
}

type DatabaseIni struct {
	Type        string
	User        string
	Password    string
	Host        string
	Db          string
	TablePrefix string
	MaxIdle     int
	MaxOpen     int
}

type RedisIni struct {
	Host        string
	Password    string
	DB          int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type TencentCaptchaIni struct {
	Aid          string
	AppSecretKey string
}

type GinLogIni struct {
	Enable int
	Path   string
}

type AliSmsIni struct {
	AccessId            string
	SecretKey           string
	SignName            string
	DefaultTemplateCode string
}

var confPath *string = flag.String("conf", "app.ini", "Use -conf <conf path>")
var logPath *string = flag.String("log", "", "Use -log <log output path>")

func init() {
	flag.Parse()
	cfg, err := ini.Load(*confPath)
	if err != nil {
		log.Fatalf("Fail to parse 'app.ini': %v", err)
	}
	loadMysqlConfig(cfg)
	loadServerIni(cfg)
	loadRedisCacheIni(cfg)
	loadTencentCaptchaIni(cfg)
	loadAliSmsIni(cfg)
	loadGinLogIni(cfg)
}

//载入数据库配置
func loadMysqlConfig(cfg *ini.File) {
	sec, err := cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}
	var iniData DatabaseIni
	if sec.MapTo(&iniData) != nil {
		log.Fatalf("Fail to Struct database config")
	}
	Setting.Database = iniData
}

//载入服务器配置
func loadServerIni(cfg *ini.File) {
	sec, err := cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	var iniData ServerIni
	if sec.MapTo(&iniData) != nil {
		log.Fatalf("Fail to Struct server config")
	}
	Setting.Server = iniData
}

//载入Redis缓存配置
func loadRedisCacheIni(cfg *ini.File) {
	sec, err := cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get section 'redis_cache': %v", err)
	}
	var iniData RedisIni
	if sec.MapTo(&iniData) != nil {
		log.Fatalf("Fail to Struct redis_cache config")
	}
	Setting.Redis = iniData
}

//载入腾讯验证码配置
func loadTencentCaptchaIni(cfg *ini.File) {
	sec, err := cfg.GetSection("tencent_captcha")
	if err != nil {
		log.Fatalf("Fail to get section 'tencent_captcha': %v", err)
	}
	var iniData TencentCaptchaIni
	if sec.MapTo(&iniData) != nil {
		log.Fatalf("Fail to Struct tencent_captcha config")
	}
	Setting.Captcha = iniData
}

//载入阿里云短信配置
func loadAliSmsIni(cfg *ini.File) {
	sec, err := cfg.GetSection("ali_sms")
	if err != nil {
		log.Fatalf("Fail to get section 'ali_sms': %v", err)
	}
	var iniData AliSmsIni
	if sec.MapTo(&iniData) != nil {
		log.Fatalf("Fail to Struct ali_sms config")
	}
	Setting.AliSms = iniData
}

//载入日志配置
func loadGinLogIni(cfg *ini.File) {
	sec, err := cfg.GetSection("log")
	if err != nil {
		log.Fatalf("Fail to get section 'log': %v", err)
	}
	var iniData GinLogIni
	if sec.MapTo(&iniData) != nil {
		log.Fatalf("Fail to Struct log config")
	}
	Setting.Log = iniData
	//如果通过命令启动设置过日志路径的话，直接优先使用命令路径
	if *logPath != "" {
		Setting.Log.Path = *logPath
	}
}
