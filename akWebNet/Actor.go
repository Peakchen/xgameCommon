package akWebNet

type TActor struct {
	Route *MsgRoute
}

func (this *TActor) GetMsgRoute() *MsgRoute {
	return this.Route
}
