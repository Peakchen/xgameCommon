package serverConfig

import "github.com/Peakchen/xgameCommon/akLog"

var (
	SvrPath string
)

func LoadSvrAllConfig(CfgPath string) {
	akLog.FmtPrintln("load config path: ", CfgPath)
	SvrPath = CfgPath
	loadExternalgwConfig()
	loadGameConfig()
	loadInnergwConfig()
	loadLoginConfig()
	loadMgoConfig()
	loadNetFilterConfig()
	loadRedisConfig()
	loadCenterconfig()
	loadServerglobalconfig()
}
