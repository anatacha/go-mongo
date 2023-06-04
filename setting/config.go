package setting

import (
	"github.com/patcharp/golib/v2/database"
	"github.com/patcharp/golib/v2/util"
	"github.com/sirupsen/logrus"
)

var (
	// 1x getEnv
	getEnv = util.GetEnv
)

type Cfg struct {
	Db database.Database 
	Debug	bool
}

func (cfg *Cfg) Load() error {
	cfg.Debug = getEnv("DEBUG", "false") == "true"
	if cfg.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
		// Resource
		cfg.Db = database.NewWithConfig(
			database.Config{
				Host:      getEnv("DB_HOST", "127.0.0.1"),
				Port:      getEnv("DB_PORT", "3000"),
				Username:  getEnv("DB_USERNAME", ""),
				Password:  getEnv("DB_PASSWORD", ""),
				Name:      getEnv("DB_NAME", ""),
				DebugMode: cfg.Debug,
			},
			database.DriverMySQL,
		)
	return nil
}
