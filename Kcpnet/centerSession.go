package Kcpnet

import "sync"

/*
	server session, key: remoteAddr, value: session obj
	relationship of player and server.
	[
		map1 key: player(addr or identify), value: server session, route server session by player key.
		map2 key: server session, value: arrary for players (addr or identify), if servers disconnect, will delete all value.
	]
*/

type CenterSessionMgr struct {
	playerSessions sync.Map //key,value:server addr
	serverSessions sync.Map //key,value:[]string
}

func (this *CenterSessionMgr) AddPlayerSession(pkey string, session TcpSession) {
	this.playerSessions.Store(pkey, session)
	this.AppendSvrSession(session, pkey)
}

func (this *CenterSessionMgr) AppendSvrSession(session TcpSession, vKey string) {
	data, exist := this.serverSessions.Load(session.GetRemoteAddr())
	if !exist {
		data = []string{}
	}
	data = append(data, vKey)
	this.serverSessions.Store(session.GetRemoteAddr(), data)
}

func (this *CenterSessionMgr) GetPlayerSession(pkey string) (session TcpSession) {
	session, _ = this.playerSessions.Load(pkey)
	return
}

func (this *CenterSessionMgr) ClearSvrSession(session TcpSession) {
	pkeys, exist := this.serverSessions.Load(session.GetRemoteAddr())
	if exist {
		for _, pkey := range pkeys {
			this.ClearPlayerSession(pkey)
		}
	}
	this.serverSessions.Delete(session.GetRemoteAddr())
}

func (this *CenterSessionMgr) ClearPlayerSession(pkey string) {
	this.playerSessions.Delete(pkey)
}
