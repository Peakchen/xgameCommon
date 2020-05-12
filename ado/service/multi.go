package service

import (
	"github.com/Peakchen/xgameCommon/aktime"
)

var (
	clusterProvider *TClusterDBProvider
)

func init() {
	clusterProvider = &TClusterDBProvider{}
}

func StartMultiDBProvider(Server string) {
	clusterProvider.Start(Server)
	aktime.InitAkTime(clusterProvider.GetRedisConn())
}
