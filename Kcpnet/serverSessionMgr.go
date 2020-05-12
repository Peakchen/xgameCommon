package Kcpnet

import (
	"sync"
)

type SessionMgr struct {
	sessionMap sync.Map
}

var (
	GServer2ServerSession *SessionMgr
	GClient2ServerSession *SessionMgr
)

func (this *SessionMgr) AddSession(key interface{}, sess TcpSession) {
	this.sessionMap.Store(key, sess)
}

func (this *SessionMgr) GetSession(key interface{}) (sess TcpSession) {
	v, exist := this.sessionMap.Load(key)
	if exist {
		sess = v.(*TcpSession)
	}
	return
}

func (this *SessionMgr) RemoveSession(key interface{}) (succ bool) {
	_, exist := this.sessionMap.Load(key)
	if exist {
		this.sessionMap.Delete(key)
	}
}

func (this *SessionMgr) GetSessionByIdentify(key interface{}) (sess TcpSession) {
	val, exist := this.sessionMap.Load(key)
	if exist {
		sess = val.(TcpSession)
	}
	return
}

func (this *SessionMgr) GetAllSession() (sess sync.Map) {
	return this.sessionMap
}

func init() {
	GServer2ServerSession = &SessionMgr{}
	GClient2ServerSession = &SessionMgr{}
}
