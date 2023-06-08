package setting

import (


	"github.com/patcharp/golib/v2/util"
	// "github.com/sirupsen/logrus"
)

var (
	AppName     string
	Version     string
	BuildTime   string
	BuildCommit string

	AllowOrigin  = "*"
	AllowHeaders = "Origin, Content-Type, Accept"

	// 1x getEnv
	getEnv = util.GetEnv
)

type Cfg struct {
	AppName      string
	Version      string
	BuildTime    string
	BuildCommit  string
	Production   bool
	Debug        bool
	AllowOrigin  string
	AllowHeaders string

	Db Config

	// Server
	// Server server.Server
}

type Config struct {
	Host     string
	Port     string
	Name     string
	Filename string
}

var theCfg *Cfg

func NewCfg() *Cfg {
	return &Cfg{
		AppName:      AppName,
		Version:      Version,
		BuildTime:    BuildTime,
		BuildCommit:  BuildCommit,
		AllowOrigin:  AllowOrigin,
		AllowHeaders: AllowHeaders,
	}
}

func GetCfg() *Cfg {
	if theCfg == nil {
		theCfg = NewCfg()
	}
	return theCfg
}
func (cfg *Cfg) Load() error {

	cfg.Db = Config{
		Host: getEnv("DB_HOST", "127.0.0.1"),
		Port: getEnv("DB_PORT", "27017"),
		Name: getEnv("DB_NAME", "golang-test"),
	}
	return nil
}

