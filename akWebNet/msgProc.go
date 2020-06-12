package akWebNet

/*
	初级消息处理
*/

import (
	"fmt"
	"reflect"

	"github.com/Peakchen/xgameCommon/ado/dbStatistics"
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/gorilla/websocket"
)

type receiveMsgProc func(sess *WebSession, msg *wsMessage)

var (
	_procMsgs = map[int]receiveMsgProc{
		websocket.TextMessage:   TextMessageFunc,
		websocket.BinaryMessage: BinaryMessageFunc,
		websocket.CloseMessage:  CloseMessageFunc,
		websocket.PingMessage:   PingMessageFunc,
		websocket.PongMessage:   PongMessageFunc,
	}
)

func GetMessageHandler(id int) receiveMsgProc {
	return _procMsgs[id]
}

func TextMessageFunc(sess *WebSession, msg *wsMessage) {
	fmt.Println("read TextMessage data: ", string(msg.data))
	sess.Write(websocket.TextMessage, []byte("hello,too!"))
}

func BinaryMessageFunc(sess *WebSession, src *wsMessage) {
	protocolPack := sess.GetProtoPack()
	msg, cb, unpackerr, exist := protocolPack.UnPackData()
	if unpackerr != nil || !exist {
		akLog.FmtPrintf("unpack data, err: %v.", unpackerr)
		return
	}

	// record db operation stack log.
	mainid, subid := protocolPack.GetMessageID()
	// sessionobj.RefreshHeartBeat(mainid, subid)
	identify := protocolPack.GetIdentify()
	dbStatistics.DBMsgStatistics(identify, mainid, subid)

	//callback define: func (sess *WebSession, proto Message)(bool,error)
	params := []reflect.Value{
		reflect.ValueOf(sess),
		reflect.ValueOf(msg),
	}
	ret := cb.Call(params)
	succ := ret[0].Interface().(bool)
	reterr := ret[1].Interface()
	if reterr != nil || !succ {
		akLog.FmtPrintln("[client] message return err: ", reterr.(error).Error())
	}
}

func CloseMessageFunc(sess *WebSession, msg *wsMessage) {
	fmt.Println("read CloseMessage.")
	sess.offch <- sess
}

func PingMessageFunc(sess *WebSession, msg *wsMessage) {
	fmt.Println("read PingMessage.")
}

func PongMessageFunc(sess *WebSession, msg *wsMessage) {
	fmt.Println("read PongMessage.")
}
