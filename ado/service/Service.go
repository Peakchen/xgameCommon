package service

import (
	"Config/serverConfig"
	"Log"
	"MgoConn"
	"RedisConn"
	"ado"
	"public"
	"context"
	"sync"
)

type TDBProvider struct {
	rconn  *RedisConn.TAokoRedis
	mconn  *MgoConn.TAokoMgo
	Server string
	ctx    context.Context
	cancle context.CancelFunc
	wg     sync.WaitGroup
}

func (this *TDBProvider) StartDBService(Server string, upcb public.UpdateDBCacheCallBack) {
	this.Server = Server
	rediscfg := serverConfig.GRedisconfigConfig.Get()
	this.rconn = RedisConn.NewRedisConn(rediscfg.Connaddr, rediscfg.DBIndex, rediscfg.Passwd, upcb)

	mgocfg := serverConfig.GMgoconfigConfig.Get()
	this.mconn = MgoConn.NewMgoConn(Server, mgocfg.Username, mgocfg.Passwd, mgocfg.Host)
}

func (this *TDBProvider) GetAkRedis() *RedisConn.TAokoRedis {
	return this.rconn
}

func (this *TDBProvider) RediSave(identify string, rediskey string, data []byte, Oper ado.EDBOperType) (err error) {
	err, _ = this.rconn.SaveEx(identify, rediskey, data, Oper)
	if err != nil {
		Log.ErrorIDCard(identify, "update redis fail, rediskey: ", rediskey, ", err: ", err)
	}
	return
}

func (this *TDBProvider) GetMogoConn() *MgoConn.TAokoMgo {
	return this.mconn
}
