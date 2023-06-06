package setting

import (
	"fmt"

	"github.com/patcharp/golib/v2/database"
	"github.com/patcharp/golib/v2/server"
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
	DbHost       string
	DbPort       string
	DbName       string
	DbFilename   string
	Db           database.Database

	// Server
	Server server.Server
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

	cfg.DbHost = getEnv("DB_HOST", "127.0.0.1")
	cfg.DbPort = getEnv("DB_PORT", "27017")
	cfg.DbName = getEnv("DB_NAME", "golang-test")
	cfg.DbFilename = constructMongoDBConnectionString(cfg.DbHost, cfg.DbPort)
	cfg.Server = server.New(server.Config{
		Host: getEnv("SERVER_HOST", "0.0.0.0"),
		Port: getEnv("SERVER_PORT", "3000"),
		// Config: &fiberCfg,
	})
	return nil
}

// Function to construct MongoDB connection string
func constructMongoDBConnectionString(host, port string) string {
	return fmt.Sprintf("mongodb://%s:%s", host, port)

}
