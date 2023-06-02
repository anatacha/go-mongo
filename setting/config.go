package setting

import(
	"github.com/patcharp/golib/v2/util"
	"github.com/patcharp/golib/v2/database"
)
var (
	// 1x getEnv
	getEnv       = util.GetEnv
)

type Cfg struct {
	Db            database.Database

}