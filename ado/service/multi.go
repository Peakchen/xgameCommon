package service

import (
	"github.com/Peakchen/xgameCommon/aktime"
)

var (
	clusterProvider = &TClusterDBProvider{}
)

func StartMultiDBProvider(Server string, rediscfg *TRedisConfig, mgocfg *TMgoConfig) {
	clusterProvider.Start(Server)
	aktime.InitAkTime(clusterProvider.GetRedisConn(), rediscfg, mgocfg)
}
