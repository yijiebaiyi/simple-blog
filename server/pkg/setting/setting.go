package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize     int
	JwtSecret    string
	Md5Salt      string
	SessionName  string
	CookieSecret string
)

func init() {
	var err error
	var RunCfg *ini.File

	RunCfg, err = ini.Load("conf/run.ini")
	if err != nil {
		log.Fatalf("Failed to parse run.ini: %v", err)
	}

	RunMode = RunCfg.Section("").Key("RUN_MODE").MustString("release")
	if RunMode == "release" {
		Cfg, err = ini.Load("conf/app_release.ini")
	} else {
		Cfg, err = ini.Load("conf/app_debug.ini")
	}

	if err != nil {
		log.Fatalf("Failed to parse app.ini: %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Failed to get section 'server': %v", err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Failed to get section 'app': %v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	Md5Salt = sec.Key("MD5_SALT").MustString("")
	CookieSecret = sec.Key("COOKIE_SECRET").MustString("cookie_secret")
	SessionName = sec.Key("SESSION_NAME").MustString("session_name")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
