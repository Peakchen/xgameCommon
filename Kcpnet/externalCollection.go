package Kcpnet

/*
	external struct collection for server or client for special using.
*/

type ExternalCollection struct {
	centerSession *CenterSessionMgr
	// ...
}

func (this *ExternalCollection) GetCenterSession() *CenterSessionMgr {
	return this.centerSession
}
