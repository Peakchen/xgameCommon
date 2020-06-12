package akWebNet

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

//广播消息
/*
	@param 1: 自身会话
	@param 2：是否不给自己广播
	@param 3：消息ID
	@param 4：消息参数
*/
func BroadCastMsgExceptSession(sess *WebSession, bMsg2Me bool, mainId, subId uint16, data proto.Message) {
	protocolPack := sess.GetProtoPack()
	msg, err := protocolPack.PackClientMsg(mainId, subId, data)
	if err != nil {
		akLog.Error("pack msg fail: ", err)
		return
	}
	sesses := GwebSessionMgr.GetSessions()
	sesses.Range(func(k, v interface{}) bool {
		if v != nil {
			sess := v.(*WebSession)
			if !bMsg2Me && sess.RemoteAddr == sess.RemoteAddr {
				return true
			}
			sess.Write(websocket.BinaryMessage, msg)
		}

		return true
	})
}

func BroadCastMsgExceptID(mainId, subId uint16, data proto.Message) {
	sesses := GwebSessionMgr.GetSessions()
	sesses.Range(func(k, v interface{}) bool {
		if v != nil {
			sess := v.(*WebSession)
			protocolPack := sess.GetProtoPack()
			msg, err := protocolPack.PackClientMsg(mainId, subId, data)
			if err != nil {
				akLog.Error("pack msg fail: ", err)
				return false
			}
			sess.Write(websocket.BinaryMessage, msg)
		}

		return true
	})
}

func SendMsg(sess *WebSession, mainId, subId uint16, data proto.Message) {
	protocolPack := sess.GetProtoPack()
	msg, err := protocolPack.PackClientMsg(mainId, subId, data)
	if err != nil {
		return
	}
	sess.Write(websocket.BinaryMessage, msg)
}
