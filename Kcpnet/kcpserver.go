package Kcpnet

// by udp

import (
	"crypto/sha1"
	"fmt"
	"net"

	"github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"

	//cli "gopkg.in/urfave/cli.v2"
	"context"
	"os"
	"sync"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/pprof"
	cli "github.com/urfave/cli"
)

type KcpServer struct {
	sw           sync.WaitGroup
	svrName      string
	pack         IMessagePack
	addr         string
	ppAddr       string
	cancel       context.CancelFunc
	offCh        chan *KcpServerSession
	exCollection *ExternalCollection
}

func NewKcpServer(Name string, addr string, pprofAddr string, exCol *ExternalCollection) *KcpServer {
	return &KcpServer{
		svrName:      Name,
		addr:         addr,
		ppAddr:       pprofAddr,
		offCh:        make(chan *KcpServerSession, 1000),
		exCollection: exCol,
	}
}

func (this *KcpServer) Run() {
	os.Setenv("GOTRACEBACK", "crash")

	ctx, _ := context.WithCancel(context.Background())
	pprof.Run(ctx)

	app := &cli.App{
		Name:    this.svrName,
		Usage:   "a server...",
		Version: "v1.0",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "listen",
				Value: this.addr,
				Usage: "local listen address",
			},
			cli.StringFlag{
				Name:  "key",
				Value: "innergateway",
				Usage: "key",
			},
			cli.StringFlag{
				Name:  "crypt",
				Value: "aes-128",
				Usage: "crypt",
			},
			cli.DurationFlag{
				Name:  "tcp_readDeadline",
				Value: pingPeriod,
				Usage: "tcp_readDeadline",
			},
			cli.DurationFlag{
				Name:  "tcp_writeDeadline",
				Value: pingPeriod,
				Usage: "sockbuf",
			},
			cli.DurationFlag{
				Name:  "udp_readDeadline",
				Value: pingPeriod,
				Usage: "udp_readDeadline",
			},
			cli.DurationFlag{
				Name:  "udp_writeDeadline",
				Value: pingPeriod,
				Usage: "udp_writeDeadline",
			},
			cli.IntFlag{
				Name:  "tcp_sockbuf_w",
				Value: tcpBuffSize,
				Usage: "tcp_sockbuf_w",
			},
			cli.IntFlag{
				Name:  "tcp_sockbuf_r",
				Value: tcpBuffSize,
				Usage: "tcp_sockbuf_r",
			},
			cli.IntFlag{
				Name:  "udp_sockbuf_w",
				Value: udpBuffSize,
				Usage: "udp_sockbuf_w",
			},
			cli.IntFlag{
				Name:  "udp_sockbuf_r",
				Value: udpBuffSize,
				Usage: "udp_sockbuf_r",
			},
			cli.IntFlag{
				Name:  "queuelen",
				Value: queueSize,
				Usage: "queue length",
			},
			cli.IntFlag{
				Name:  "dscp",
				Value: dscp,
				Usage: "dscp",
			},
			cli.IntFlag{
				Name:  "udp-sndwnd",
				Value: 2048,
				Usage: "udp-sndwnd",
			},
			cli.IntFlag{
				Name:  "udp-rcvwnd",
				Value: 2048,
				Usage: "udp-rcvwnd",
			},
			cli.IntFlag{
				Name:  "udp-mtu",
				Value: 1400,
				Usage: "udp-mtu",
			},
			cli.IntFlag{
				Name:  "dataShard",
				Value: 10,
				Usage: "dataShard",
			},
			cli.IntFlag{
				Name:  "parityShards",
				Value: 3,
				Usage: "parityShards",
			},
			cli.IntFlag{
				Name:  "nodelay",
				Value: 1,
				Usage: "nodelay",
			},
			cli.IntFlag{
				Name:  "interval",
				Value: 40,
				Usage: "interval",
			},
			cli.IntFlag{
				Name:  "resend",
				Value: 2,
				Usage: "resend",
			},
			cli.IntFlag{
				Name:  "nc",
				Value: 1,
				Usage: "nc",
			},
		},
		Action: func(c *cli.Context) error {
			akLog.FmtPrintln("action begin...")

			//setup net param
			config := &KcpSvrConfig{
				listen:            c.String("listen"),
				key:               c.String("key"),
				crypt:             c.String("crypt"),
				tcp_readDeadline:  c.Duration("tcp_readDeadline"),
				tcp_writeDeadline: c.Duration("tcp_writeDeadline"),
				udp_readDeadline:  c.Duration("udp_readDeadline"),
				udp_writeDeadline: c.Duration("udp_writeDeadline"),
				tcp_sockbuf_w:     c.Int("tcp_sockbuf_w"),
				tcp_sockbuf_r:     c.Int("tcp_sockbuf_r"),
				udp_sockbuf_w:     c.Int("udp_sockbuf_w"),
				udp_sockbuf_r:     c.Int("udp_sockbuf_r"),
				queuelen:          c.Int("queuelen"),
				dscp:              c.Int("dscp"),
				sndwnd:            c.Int("udp-sndwnd"),
				rcvwnd:            c.Int("udp-rcvwnd"),
				mtu:               c.Int("udp-mtu"),
				dataShard:         c.Int("dataShard"),
				parityShards:      c.Int("parityShards"),
				nodelay:           c.Int("nodelay"),
				interval:          c.Int("interval"),
				resend:            c.Int("resend"),
				nc:                c.Int("nc"),
			}
			// init services
			//startup(c)
			// init timer
			//initTimer(c.Int("rpm-limit"))

			// start udp server...
			this.sw.Add(1)
			go this.kcpAccept(config)
			go this.loopOffline()
			this.sw.Wait()
			return nil
		},
	}

	app.Run(os.Args)
}

func (this *KcpServer) listenEcho(c *KcpSvrConfig) (net.Listener, error) {
	var pass = pbkdf2.Key([]byte(c.key), []byte(c.crypt), 4096, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(pass)
	return kcp.ListenWithOptions(c.listen, block, c.dataShard, c.parityShards)
}

func (this *KcpServer) kcpAccept(c *KcpSvrConfig) {
	l, err := this.listenEcho(c)
	if err != nil {
		panic(err)
	}
	akLog.FmtPrintln("kcp listening on:", l.Addr())
	kcplis := l.(*kcp.Listener)
	if err := kcplis.SetReadBuffer(c.udp_sockbuf_r); err != nil {
		panic(fmt.Errorf("SetReadBuffer, err: %v.", err))
	}
	if err := kcplis.SetWriteBuffer(c.udp_sockbuf_w); err != nil {
		panic(fmt.Errorf("SetWriteBuffer, err: %v.", err))
	}
	if err := kcplis.SetDSCP(c.dscp); err != nil {
		panic(fmt.Errorf("SetDSCP, err: %v.", err))
	}
	// loop accepting
	for {
		conn, err := kcplis.AcceptKCP()
		if err != nil {
			akLog.FmtPrintln("accept failed:", err)
			continue
		}

		// set kcp parameters
		conn.SetWindowSize(c.sndwnd, c.rcvwnd)
		conn.SetNoDelay(c.nodelay, c.interval, c.resend, c.nc)
		conn.SetStreamMode(true)
		conn.SetMtu(c.mtu)

		// start a goroutine for every incoming connection for read and write
		//go handleClient(conn, config)
		sess := NewKcpSvrSession(conn, this.offCh, c, this.exCollection)
		sess.Handler()
	}
}

func (this *KcpServer) loopOffline() {
	for {
		offsession := <-this.offCh
		offsession.Offline()
	}
}
