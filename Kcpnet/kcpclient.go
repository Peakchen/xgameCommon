package Kcpnet

// by udp

import (
	"github.com/Peakchen/xcommon/pprof"
	"context"
)

type KcpClient struct {
	sw      sync.WaitGroup
	svrName string
	pack    IMessagePack
	Addr    string
	ppAddr  string
	cancel  context.CancelFunc
	sw      sync.WaitGroup
}

func NewKcpClient(addr, pprofAddr string, name string) *KcpClient {
	return &KcpClient{
		svrName: name,
		Addr:    addr,
		ppAddr:  pprofAddr,
	}
}

func (this *KcpClient) Run() {
	os.Setenv("GOTRACEBACK", "crash")

	var ctx context.Context
	ctx, this.cancel = context.WithCancel(context.Background())
	pprof.Run(ctx)

	this.connect(ctx, &this.sw)

}

func (this *KcpClient) connect(ctx context.Context, sw *sync.WaitGroup) {

}
