package akWebNet

import (
	"sync"
)

type wsClientSession struct {
	sessMap sync.Map
}

var (
	GwebSessionMgr *wsClientSession
)

func (this *wsClientSession) AddSession(sess *WebSession){
	this.sessMap.Store(sess.RemoteAddr, sess)
}

func (this *wsClientSession) GetSession(addr string) (sess *WebSession){
	val, exist := this.sessMap.Load(addr)
	if exist {
		sess = val.(*WebSession)
	}
	return 
}

func (this *wsClientSession) RemoveSession(addr string) {
	this.sessMap.Delete(addr)
}

func (this *wsClientSession) GetSessions()sync.Map{
	return this.sessMap
}

func init(){
	GwebSessionMgr = &wsClientSession{}
}