package akWebNet

import "github.com/Peakchen/xgameCommon/define"

type TActor struct {
	ActorType define.ERouteId
	Route     *MsgRoute
}

func (this *TActor) GetMsgRoute() *MsgRoute {
	return this.Route
}

func (this *TActor) GetActorType() define.ERouteId {
	return this.ActorType
}
