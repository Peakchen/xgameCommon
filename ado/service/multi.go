package service

import (
	"github.com/Peakchen/xgameCommon/ado"
	"github.com/Peakchen/xgameCommon/aktime"
)

var (
	clusterProvider = &TClusterDBProvider{}
)

func StartMultiDBProvider(Server string, rediscfg *ado.TRedisConfig, mgocfg *ado.TMgoConfig) {
	clusterProvider.Start(Server)
	aktime.InitAkTime(clusterProvider.GetRedisConn(), rediscfg, mgocfg)
}
