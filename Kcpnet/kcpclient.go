package Kcpnet

// by udp

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Peakchen/xgameCommon/pprof"
	"github.com/xtaci/kcp-go"
)

type KcpClient struct {
	sw      sync.WaitGroup
	svrName string
	pack    IMessagePack
	Addr    string
	ppAddr  string
	cancel  context.CancelFunc
	sesson  *KcpClientSession
	offCh   chan *KcpClientSession
}

func NewKcpClient(addr, pprofAddr string, name string) *KcpClient {
	return &KcpClient{
		svrName: name,
		Addr:    addr,
		ppAddr:  pprofAddr,
		offCh:   make(chan *KcpClientSession, 1000),
	}
}

func (this *KcpClient) Run() {
	os.Setenv("GOTRACEBACK", "crash")

	var ctx context.Context
	ctx, this.cancel = context.WithCancel(context.Background())
	pprof.Run(ctx)

	this.connect(ctx, &this.sw)
	this.sw.Add(2)
	go this.loopconnect(ctx, &this.sw)
	go this.loopOffline(ctx, &this.sw)
	this.sw.Wait()
}

func (this *KcpClient) connect(ctx context.Context, sw *sync.WaitGroup) {
	conn, err := kcp.Dial(this.Addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	this.sesson = NewKcpClientSession(conn, this.offCh)
	this.sesson.Handler()
}

func (this *KcpClient) loopconnect(ctx context.Context, sw *sync.WaitGroup) {
	defer sw.Done()

	tick := time.NewTicker(time.Duration(5) * time.Second)
	for {
		select {
		case <-tick.C:
			if this.sesson == nil || !this.sesson.Alive() {
				this.connect(ctx, sw)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (this *KcpClient) loopOffline(ctx context.Context, sw *sync.WaitGroup) {
	defer sw.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case offsession := <-this.offCh:
			offsession.Offline()
		}
	}
}
