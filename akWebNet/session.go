package akWebNet

import (
	"time"
	"github.com/gorilla/websocket"
	"fmt"
	//"encoding/binary"
	//"io"
	//"strings"
	"utls"
	"sync"
)

type wsMessage struct {
    messageType int //消息类型 TextMessage/BinaryMessage/CloseMessage/PingMessage/PongMessage
    data        []byte //消息内容
}

type WebSession struct {
	wsconn *websocket.Conn  
	offch 	chan *WebSession	//离线通道
	RemoteAddr string
	writeCh chan *wsMessage //写通道
	readCh  chan *wsMessage //读通道
	IdCh    *uint32
	stopWrite bool
	stopRead  bool

	one 	sync.Once
}

func NewWebSession(conn *websocket.Conn, off chan *WebSession) *WebSession{
	return &WebSession{
		wsconn: conn,
		offch: off,
		RemoteAddr: conn.RemoteAddr().String(),
		writeCh: make(chan *wsMessage, maxWriteMsgSize),
		readCh: make(chan *wsMessage, maxWriteMsgSize),
		IdCh: new(uint32),
	}
}

func (this *WebSession) SetId(id uint32){
	*this.IdCh = id
}

func (this *WebSession) GetId() uint32{
	return *(this.IdCh)
}

func (this *WebSession) Handle(){
	go this.readloop()
	go this.writeloop()
}

func (this *WebSession) offline(){
	mosterid := this.GetId()
	fmt.Println("exit ws socket: ", this.RemoteAddr, mosterid, time.Now().Unix())
	GwebSessionMgr.RemoveSession(this.RemoteAddr)
	if proc := GetGameLogicProcMsg(MID_logout); proc != nil {
		proc(this, []uint32{mosterid})
	}
}

func (this *WebSession) exit(){
	//conn close 只执行一次
	this.one.Do(func(){
		this.offline()
		//this.offch <-this	

		if _, noclosed := <-this.writeCh; noclosed {
			this.stopWrite = true
			close(this.writeCh)
		} 
		
		if _, noclosed := <-this.readCh; noclosed {
			this.stopRead = true
			close(this.readCh)
		}
	
		this.wsconn.Close()
	})
}

func (this *WebSession) readloop(){

	defer func(){
		this.exit()
	}()

	this.wsconn.SetReadLimit(maxMessageSize)
	for {
		this.wsconn.SetReadDeadline(time.Now().Add(pongWait))
        msgType, data, err := this.wsconn.ReadMessage()
        if err != nil {
            websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			fmt.Println("msg read fail, err: ", err.Error(), time.Now().Unix())
            return
		}

		if this.stopRead {
			return
		}
		
        this.readCh <- &wsMessage{
            messageType: msgType,
            data: data,
		}
		
		go this.read()
	}
}

func (this *WebSession) read(){
	msg := <- this.readCh
	fmt.Println("read messageType: ", msg.messageType, len(msg.data), time.Now().Unix())
	if handler := GetMessageHandler(msg.messageType); handler != nil {
		handler(this, msg)
	}else{
		panic(fmt.Errorf("invalid message type: %v.", msg.messageType))
	}
}

func (this *WebSession) writeloop(){
	ticker := time.NewTicker(pingPeriod)
	deadline := time.Duration(pingPeriod/2)
    defer func() {
		ticker.Stop()
		this.exit()
	}()
	
	for {
		select {
			case msg := <-this.writeCh:
				this.wsconn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := this.wsconn.WriteMessage(msg.messageType, msg.data); err != nil {
					fmt.Println("send msg fail, err: ", err.Error(), time.Now().Unix())
					return
				}
			case <-ticker.C:
				if err := this.wsconn.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(deadline)); err != nil {
					fmt.Println("send msg over time, err: ", err.Error(), time.Now().Unix())
					return
				}
			}
	}
}

func (this *WebSession) sendOffline(){
	/* send message to close connetion... */
	if err := this.wsconn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "now closing..."), time.Now().Add(time.Second)); err != nil {
		fmt.Println("send close fail, err: ", err)
		return
	}
}

func (this *WebSession) Write(msgtype int, data []byte) {
	if this.stopWrite {
		return
	}

	fmt.Println("session writed channel data len: ", len(this.writeCh), utls.SizeVal(this.writeCh), time.Now().Unix())
	this.writeCh <- &wsMessage{
		messageType: msgtype,
		data: data,
	}
}

func (this *WebSession) broadcast(msgtype int, data []byte) {
	sesses := GwebSessionMgr.GetSessions()
	sesses.Range(func (k, v interface{}) bool{
		if v != nil {
			sess := v.(*WebSession)
			sess.Write(msgtype, data)
		}
		
		return true
	})
}

