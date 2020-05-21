package Kcpnet

import "time"

type KcpSvrConfig struct {
	listen                        string
	readDeadline                  time.Duration
	sockbuf                       int
	udp_sockbuf                   int
	txqueuelen                    int
	dscp                          int
	sndwnd                        int
	rcvwnd                        int
	mtu                           int
	nodelay, interval, resend, nc int
}
