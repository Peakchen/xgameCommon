package akWebNet

type TActor struct {
	ActorType ACTOR_TYPE
	Route     *MsgRoute
}

func (this *TActor) GetMsgRoute() *MsgRoute {
	return this.Route
}

func (this *TActor) GetActorType() ACTOR_TYPE {
	return this.ActorType
}
