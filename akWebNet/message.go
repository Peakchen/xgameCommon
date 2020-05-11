package akWebNet

import (
	"encoding/binary"
	"github.com/gorilla/websocket"
	"fmt"
)

type WsCallback func (ws *WebSession, data []uint32) (error, bool)

var (
	msgs = map[int]WsCallback{}
)

func MsgRegister(id int, cb WsCallback) {
	msgs[id] = cb
}

func GetGameLogicProcMsg(id int) WsCallback{
	return msgs[id]
}

func packBroadCastMsg(msgid int, msgparams []uint32)(msg []byte){
	if len(msgparams) == 0 {
		panic("invalid  broadcast msg content.")
	}

	msg = make([]byte, 4*(len(msgparams)+2))
	//消息ID
	var pos int
	binary.LittleEndian.PutUint32(msg[pos:], uint32(msgid))
	pos+=4
	
	//消息长度
	binary.LittleEndian.PutUint32(msg[pos:], uint32(len(msgparams)))
	pos+=4

	//消息内容
	for i:=0; i<len(msgparams); i++{
		binary.LittleEndian.PutUint32(msg[pos:], msgparams[i])
		pos+=4
	}
	return
}

//广播消息
/*
	@param 1: 自身会话
	@param 2：是否不给自己广播
	@param 3：消息ID
	@param 4：消息参数 
*/
func BroadCastMsgExceptSession(selfsess *WebSession, bMsg2Me bool, msgid int, msgparams []uint32) {
	msg := packBroadCastMsg(msgid, msgparams)
	sesses := GwebSessionMgr.GetSessions()
	sesses.Range(func (k, v interface{}) bool{
		if v != nil {
			sess := v.(*WebSession)
			if !bMsg2Me && sess.RemoteAddr == selfsess.RemoteAddr{
				return true
			}
			sess.Write(websocket.BinaryMessage, msg)
		}
		
		return true
	})
}

func BroadCastMsgExceptID(msgid int, msgparams []uint32) {
	msg := packBroadCastMsg(msgid, msgparams)
	sesses := GwebSessionMgr.GetSessions()
	sesses.Range(func (k, v interface{}) bool{
		if v != nil {
			sess := v.(*WebSession)
			sess.Write(websocket.BinaryMessage, msg)
		}
		
		return true
	})
}

func SendMsg(sess *WebSession, msgid int, msgparams []uint32) {
	if len(msgparams) == 0 {
		panic("invalid  broadcast msg content.")
	}

	var (
		msg = make([]byte, 4*(len(msgparams)+2))
	)

	//消息ID
	var pos int
	binary.LittleEndian.PutUint32(msg[pos:], uint32(msgid))
	pos+=4
	
	//消息长度
	binary.LittleEndian.PutUint32(msg[pos:], uint32(len(msgparams)))
	pos+=4

	//消息内容
	for i:=0; i<len(msgparams); i++{
		binary.LittleEndian.PutUint32(msg[pos:], msgparams[i])
		pos+=4
	}

	sess.Write(websocket.BinaryMessage, msg)
}

func HeartBeat(sess *WebSession, data []uint32) (error, bool) {
	fmt.Println("proc HeartBeat message ... ")
	var (
		rspmsg = []uint32{}
	)
	rspmsg = append(rspmsg, 0)
	SendMsg(sess, MID_HeartBeat, rspmsg)
	return nil, true
}

func init(){
	MsgRegister(MID_HeartBeat, HeartBeat)
}