package Kcpnet

// by udp

import (
	"github.com/Peakchen/xgameCommon/pprof"
	"context"
	"github.com/xtaci/kcp-go"
)

type KcpClient struct {
	sw      sync.WaitGroup
	svrName string
	pack    IMessagePack
	Addr    string
	ppAddr  string
	cancel  context.CancelFunc
	sw      sync.WaitGroup
	sesson  *KcpClientSession
	offCh 	chan *KcpClientSession
}

func NewKcpClient(addr, pprofAddr string, name string) *KcpClient {
	return &KcpClient{
		svrName: name,
		Addr:    addr,
		ppAddr:  pprofAddr,
		offCh: 	 make(chan *KcpClientSession, 1000),
	}
}

func (this *KcpClient) Run() {
	os.Setenv("GOTRACEBACK", "crash")

	var ctx context.Context
	ctx, this.cancel = context.WithCancel(context.Background())
	pprof.Run(ctx)

	this.connect(ctx, &this.sw)
	go this.loopconnect(ctx, &this.sw)
	go this.loopOffline(ctx, &this.sw)
}

func (this *KcpClient) connect(ctx context.Context, sw *sync.WaitGroup) {
	conn, err := kcp.Dial("127.0.0.1:10086")
    if err != nil {
        fmt.Println(err)
        return
	}
	this.sesson = NewKcpClientSession(conn, this.offCh)
	this.sesson.Handler()
}

func (this *KcpClient) loopconnect(ctx context.Context, sw *sync.WaitGroup){
	tick := time.NewTicker(time.Duration(5)*time.Second)
	for {
		<-tick.C
		if this.sesson == nil || !this.sesson.Alive() {
			this.connect(ctx, sw)
		}
	}
}

func (this *KcpClient) loopOffline(ctx context.Context, sw *sync.WaitGroup){
	for {
		offsession := <-this.offCh
		offsession.Offline()
	}
}