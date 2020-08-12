package akWebNet

import (
	"context"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/pprof"
	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	Addr      string
	pprofAddr string
	offch     chan *WebSession //离线通道
	cancel    context.CancelFunc
	session   *WebSession
}

func NewWebsocketClient(addr string, pprofAddr string) *WebSocketClient {
	return &WebSocketClient{
		Addr:      addr,
		offch:     make(chan *WebSession, 1024),
		pprofAddr: pprofAddr,
	}
}

func (this *WebSocketClient) Run() {
	var ctx context.Context
	ctx, this.cancel = context.WithCancel(context.Background())
	pprof.Run(ctx)
	this.newDail()
	var sw sync.WaitGroup
	sw.Add(2)
	go this.checkReconnect(ctx, &sw)
	go loopSignalCheck(ctx, &sw)
	sw.Wait()
	this.exit()
}

func (this *WebSocketClient) newDail() {
	url := url.URL{Scheme: "ws", Host: this.Addr, Path: "/echo"}
	wsDialer := &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: 45 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
	}
	c, _, err := wsDialer.Dial(url.String(), nil)
	if err != nil {
		akLog.Error("dail fail, err: ", err)
		return
	}
	this.session = NewWebSession(c, this.offch)
	this.session.Handle()
	this.sendRegisterMsgs()
}

func (this *WebSocketClient) sendRegisterMsgs() {
	req := &MSG_Server.CS_ServerRegister_Req{}
	req.ServerType = int32(this.SvrType)
	req.Msgs = akNet.GetAllMessageIDs()
	akLog.FmtPrintln("register context: ", req.Msgs)
	msg, err := PackMsgOp(uint16(MSG_MainModule.MAINMSG_SERVER),
		uint16(MSG_Server.SUBMSG_CS_ServerRegister),
		req, PACK_PROTO)
	if msg == nil || err != nil {
		akLog.Error("pack msg fail: ", mainId, subId, err)
		return
	}
	this.session.Write(websocket.BinaryMessage, msg)
}

func (this *WebSocketClient) checkReconnect(ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case <-this.offch:
			this.newDail()
		}
	}
}

func (this *WebSocketClient) exit() {
	close(this.offch)
	this.cancel()
}
