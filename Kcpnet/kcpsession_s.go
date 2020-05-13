package Kcpnet

import (
	//"aktime"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/stacktrace"
	"fmt"
	"github.com/xtaci/kcp-go"
	"time"
	"github.com/Peakchen/xgameCommon/aktime"
	"sync"
)

type KcpServerSession struct {
	conn *kcp.UDPSession

	readCh  	chan bool
	writeCh 	chan []byte

	RemoteAddr 	string
	pack 		IMessagePack
	offCh 		chan *KcpServerSession
	isAlive		bool
	closeOnce   sync.Once
}

func NewKcpSvrSession(c *kcp.UDPSession, offCh chan *KcpServerSession) *KcpServerSession {
	return &KcpServerSession{
		conn:       c,
		readCh:     make(chan bool, 1000),
		writeCh:    make(chan []byte, 1000),
		RemoteAddr: this.conn.RemoteAddr().String(),
		pack:       &KcpServerProtocol{},
		offCh:		offCh,
	}
}

func (this *KcpServerSession) Handler() {
	this.isAlive = true
	go this.readloop()
	go this.writeloop()
}

func (this *KcpServerSession) close() {
	closeOnce.Do(func(){
		this.isAlive = false
		this.offCh <-this
		this.conn.Close()
	})
}

func (this *KcpServerSession) Offline() {
	// notify some one server...
}

func (this *KcpServerSession) heartBeatloop(sw *sync.WaitGroup) {

	ticker := time.NewTicker(time.Duration(cstCheckHeartBeatMonitorSec) * time.Second)

	defer func() {
		ticker.Stop()
		sw.Done()
	}()

	for {
		select {
		// case <-this.ctx.Done():
		// 	return
		case <-ticker.C:
			if this.heartBeatDeadline == 0 {
				continue
			}

			var disconnectionSec int
			if this.RegPoint == 0 {
				disconnectionSec = cstClientDisconnectionSec
			} else {
				disconnectionSec = cstSvrDisconnectionSec
			}

			if aktime.Now().Unix()-this.heartBeatDeadline >= int64(disconnectionSec) {
				//close connection...
				this.close()
				this.heartBeatDeadline = 0
			}
		}
	}
}

func (this *KcpServerSession) readloop() {

	defer func() {
		this.close()
	}()

	for {
		select {
		case rspcliented := <-this.readCh:
			this.dispatch(rspcliented)
		default:
			this.conn.SetReadDeadline(time.Now().Add(config.readDeadline))
			//是否加个消息队列处理 ?
			this.read(data)
		}
	}
}

func (this *KcpServerSession) read(data []byte) (succ bool) {

	defer func() {
		stacktrace.Catchcrash()
	}()

	var responseCliented bool
	if this.RegPoint == 0 {
		succ = UnPackExternalMsg(this.conn, this.pack)
		if !succ {
			return
		}
		this.pack.SetRemoteAddr(this.RemoteAddr)
	} else {
		succ = UnPackInnerMsg(this.conn, this.pack)
		if !succ {
			return
		}
		this.StrIdentify = this.pack.GetIdentify()
		if this.SvrType == Define.ERouteId_ER_ESG {
			responseCliented = true
		}
	}

	this.readCh <- responseCliented
}

func (this *KcpServerSession) dispatch(responseCliented bool) (succ bool) {
	defer func() {
		stacktrace.Catchcrash()
	}()

	var route Define.ERouteId
	mainID, SubID := this.pack.GetMessageID()
	Log.FmtPrintf("recv message: mainID: %v, subID: %v.", mainID, SubID)
	if mainID == uint16(MSG_MainModule.MAINMSG_SERVER) &&
		this.SvrType == Define.ERouteId_ER_ESG {
		route = Define.ERouteId_ER_ISG
		this.RegPoint = Define.ERouteId_ER_ISG
		this.Push(Define.ERouteId_ER_ISG) //外网关加入内网关session
		RegisterMessageRet(this)
		succ = true
		return
	}

	if (mainID == uint16(MSG_MainModule.MAINMSG_SERVER) ||
		mainID == uint16(MSG_MainModule.MAINMSG_LOGIN)) && len(this.StrIdentify) == 0 {
		this.StrIdentify = this.RemoteAddr
	}

	if len(this.pack.GetIdentify()) == 0 {
		this.pack.SetIdentify(this.StrIdentify)
	}

	if mainID == uint16(MSG_MainModule.MAINMSG_LOGIN) {
		route = Define.ERouteId_ER_Login
	} else if mainID >= uint16(MSG_MainModule.MAINMSG_PLAYER) {
		route = Define.ERouteId_ER_Game
	}

	if mainID != uint16(MSG_MainModule.MAINMSG_SERVER) && mainID != uint16(MSG_MainModule.MAINMSG_HEARTBEAT) &&
		(this.SvrType == Define.ERouteId_ER_ESG || this.SvrType == Define.ERouteId_ER_ISG) {
		if this.SvrType == Define.ERouteId_ER_ESG {
			succ = externalRouteAct(route, this, responseCliented)
		} else {
			succ = innerMsgRouteAct(ESessionType_Server, route, mainID, this.pack.GetSrcMsg())
		}
	} else {
		succ = msgCallBack(this) //路由消息回调处理
	}
}

func (this *KcpServerSession) writeloop() {

	defer func() {
		this.close()
	}()

	for {
		select {
		case data := <-this.writeCh:
			n, err := this.conn.Write(data)
			if err != nil {
				Log.Error("send reply data fail, size: %v, err: %v.", n, err)
				return
			}
		}
	}
}

func (this *KcpServerSession) Alive() bool {
	return this.isAlive
}

func (this *KcpServerSession) SetSendCache(data []byte) {
	this.writeCh <- data
}

func (this *KcpServerSession) GetPack() (obj IMessagePack) {
	return this.pack
}

func (this *KcpServerSession) GetRemoteAddr() string {
	return this.RemoteAddr
}

func (this *KcpServerSession) Push(RegPoint Define.ERouteId) {
	this.RegPoint = RegPoint
	GServer2ServerSession.AddSession(this.RemoteAddr, this)
}

func (this *KcpServerSession) SetIdentify(StrIdentify string) {
	session := GServer2ServerSession.GetSessionByIdentify(this.StrIdentify)
	if session != nil {
		GServer2ServerSession.RemoveSession(this.StrIdentify)
		this.StrIdentify = StrIdentify
		GServer2ServerSession.AddSession(StrIdentify, session)
	}
}

func (this *KcpServerSession) Offline() {

}

func (this *KcpServerSession) SendSvrClientMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[server] send msg session disconnection, mainid: %v, subid: %v.", mainid, subid)
		Log.FmtPrintln("send msg err: ", err)
		return false, err
	}

	data, err := this.pack.PackClientMsg(mainid, subid, msg)
	if err != nil {
		return succ, err
	}
	this.SetSendCache(data)
	return true, nil
}

func (this *KcpServerSession) SendInnerSvrMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[server] send svr session disconnection, mainid: %v, subid: %v.", mainid, subid)
		Log.FmtPrintln("send msg err: ", err)
		return false, err
	}

	data, err := this.pack.PackInnerMsg(mainid, subid, msg)
	if err != nil {
		return succ, err
	}
	this.SetSendCache(data)
	return true, nil
}

func (this *KcpServerSession) SendInnerClientMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[server] session disconnection, mainid: %v, subid: %v.", mainid, subid)
		Log.FmtPrintln("send msg err: ", err)
		return false, err
	}

	if len(this.GetIdentify()) > 0 {
		this.pack.SetIdentify(this.GetIdentify())
	}

	this.pack.SetPostType(MsgPostType_Single)

	data, err := this.pack.PackInnerMsg(mainid, subid, msg)
	if err != nil {
		return succ, err
	}
	this.SetSendCache(data)
	return true, nil
}

func (this *KcpServerSession) SendInnerBroadcastMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error) {
	if !this.isAlive {
		err = fmt.Errorf("[server] session disconnection, mainid: %v, subid: %v.", mainid, subid)
		Log.FmtPrintln("send msg err: ", err)
		return false, err
	}

	if len(this.GetIdentify()) > 0 {
		this.pack.SetIdentify(this.GetIdentify())
	}

	this.pack.SetPostType(MsgPostType_Broadcast)

	data, err := this.pack.PackInnerMsg(mainid, subid, msg)
	if err != nil {
		return succ, err
	}
	this.SetSendCache(data)
	return true, nil
}

func (this *KcpServerSession) GetIdentify() string {
	return this.StrIdentify
}

func (this *KcpServerSession) GetRegPoint() (RegPoint Define.ERouteId) {
	return this.RegPoint
}
