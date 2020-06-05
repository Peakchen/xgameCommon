package akWebNet

//
//from https://www.godoc.org/github.com/gorilla/websocket
//

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/Peakchen/xgameCommon/pprof"
	"github.com/gorilla/websocket"
	//"strings"
	//"strconv"
)

type WebSocketSvr struct {
	Addr      string
	pprofAddr string
	offch     chan *WebSession //离线通道
	cancel    context.CancelFunc
}

func NewWebsocketSvr(addr string, pprofAddr string) *WebSocketSvr {
	return &WebSocketSvr{
		Addr:      addr,
		offch:     make(chan *WebSession, 1024),
		pprofAddr: pprofAddr,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (this *WebSocketSvr) wsSvrHandler(resp http.ResponseWriter, req *http.Request) {

	wsSocket, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		fmt.Println("upgrader websocket fail, err: ", err.Error())
		return
	}

	sess := NewWebSession(wsSocket, this.offch)
	sess.Handle()
	fmt.Println("connect ws socket: ", sess.RemoteAddr, aktime.Now().Unix())
	GwebSessionMgr.AddSession(sess)
}

func (this *WebSocketSvr) disconnloop(ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
		this.exit()
	}()

	for {
		select {
		case sess := <-this.offch:
			mosterid := sess.GetId()
			fmt.Println("exit ws socket: ", sess.RemoteAddr, mosterid, aktime.Now().Unix())
			GwebSessionMgr.RemoveSession(sess.RemoteAddr)
			if proc := GetGameLogicProcMsg(MID_logout); proc != nil {
				proc(sess, []uint32{mosterid})
			}
		case <-ctx.Done():
			return
		}
	}
}

func (this *WebSocketSvr) Run() {
	http.HandleFunc("/ws", this.wsSvrHandler)
	var ctx context.Context
	ctx, this.cancel = context.WithCancel(context.Background())
	pprof.Run(ctx)
	var sw sync.WaitGroup
	sw.Add(1)
	go this.disconnloop(ctx, &sw)
	go this.loopSignalCheck(ctx, &sw)
	go func() {
		akLog.FmtPrintln("[client] run http server, host: ", this.pprofAddr)
		http.ListenAndServe(this.pprofAddr, nil)
	}()
	sw.Wait()
}

func (this *WebSocketSvr) exit() {
	this.cancel()
}

func (this *WebSocketSvr) loopSignalCheck(ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
		this.exit()
	}()

	chsignal := make(chan os.Signal, 1)
	signal.Notify(chsignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		select {
		case <-ctx.Done():
			return
		case s := <-chsignal:
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				akLog.FmtPrintln("signal exit:", s)
				return
			default:
				akLog.FmtPrintln("other signal:", s)
			}
		}
	}
}
