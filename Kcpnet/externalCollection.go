package Kcpnet

/*
	external struct collection for server or client for special using.
*/

type ExternalCollection struct {
	centerSession *CenterSessionMgr
	client        *KcpClient
	// ...
}

func (this *ExternalCollection) GetCenterSession() *CenterSessionMgr {
	return this.centerSession
}

func (this *ExternalCollection) GetClient() *KcpClient {
	return this.client
}
